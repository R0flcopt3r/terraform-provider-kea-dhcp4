package main

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func hostLease() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "subnet id",
				ForceNew:    true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "lease_name in lowercase",
				ForceNew:     true,
				ValidateFunc: validateName,
			},
			"mac_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "mac address in lower-case",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ip address to assign for mac",
			},
			"next_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "0.0.0.0",
				Description: "(Optional) TFTP server to boot from",
			},
			"server_hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "(Optional) TFTP server hostname",
			},
			"boot_file_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "(Optional) Path to boot file on TFTP server",
			},
			"client_classes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "(Optional) List of client classes to apply",
			},
			"option_data": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"always_send": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "(Optional) Force sending this DHCP option",
						},
						"code": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "(Optional) DHCP option code",
						},
						"csv_format": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "(Optional) Set to true if data is specified in hex format",
						},
						"data": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "DHCP option data",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "(Optional) DHCP option name",
						},
						"space": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "dhcp4",
							Description: "(Optional) DHCP space identifier, default is dhcp4",
						},
					},
				},
				Description: "(Optional) List of custom DHCP options for host",
			},
		},
		Create: resourceCreateLease,
		Read:   resourceReadLease,
		Update: resourceUpdateLease,
		Delete: resourceDeleteLease,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}

func resourceCreateLease(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*Client)
	apiClient.lock.Lock()
	defer apiClient.lock.Unlock()

	iClientClasses := d.Get("client_classes").([]interface{})
	ClientClasses := make([]string, len(iClientClasses))
	for i, v := range iClientClasses {
		ClientClasses[i] = fmt.Sprint(v)
	}

	iOptionData := d.Get("option_data").([]interface{})
	OptionData := make([]OptionData, len(iOptionData))
	for i, v := range iOptionData {
		option := v.(map[string]interface{})
		OptionData[i].AlwaysSend = option["always_send"].(bool)
		OptionData[i].CSVFormat = option["csv_format"].(bool)
		if code := option["code"].(interface{}); code != nil {
			OptionData[i].Code = code.(int)
		}
		OptionData[i].Data = option["data"].(string)
		if name := option["name"].(interface{}); name != nil {
			OptionData[i].Name = name.(string)
		}
		OptionData[i].Space = option["space"].(string)
	}

	lease := Reservations{
		Hostname:      d.Get("name").(string),
		HWAddress:     d.Get("mac_address").(string),
		IPAddress:     d.Get("ip_address").(string),
		ClientClasses: ClientClasses,
		OptionData:    OptionData,
		NextServer:    d.Get("next_server").(string),
	}

	err := apiClient.NewLease(lease, d.Get("subnet_id").(int))

	if err != nil {
		return err
	}
	d.SetId(lease.Hostname)
	return nil
}

func resourceUpdateLease(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*Client)
	apiClient.lock.Lock()
	defer apiClient.lock.Unlock()

	iClientClasses := d.Get("client_classes").([]interface{})
	ClientClasses := make([]string, len(iClientClasses))
	for i, v := range iClientClasses {
		ClientClasses[i] = fmt.Sprint(v)
	}

	iOptionData := d.Get("option_data").([]interface{})
	OptionData := make([]OptionData, len(iOptionData))
	for i, v := range iOptionData {
		option := v.(map[string]interface{})
		OptionData[i].AlwaysSend = option["always_send"].(bool)
		OptionData[i].CSVFormat = option["csv_format"].(bool)
		OptionData[i].Code = option["code"].(int)
		OptionData[i].Data = option["data"].(string)
		OptionData[i].Name = option["name"].(string)
		OptionData[i].Space = option["space"].(string)
	}

	lease := Reservations{
		Hostname:      d.Get("name").(string),
		HWAddress:     d.Get("mac_address").(string),
		IPAddress:     d.Get("ip_address").(string),
		ClientClasses: ClientClasses,
		OptionData:    OptionData,
		NextServer:    d.Get("next_server").(string),
	}

	err := apiClient.UpdateLease(lease, d.Get("subnet_id").(int))
	if err != nil {
		return err
	}
	return nil
}
func resourceReadLease(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*Client)
	apiClient.lock.Lock()
	defer apiClient.lock.Unlock()

	iClientClasses := d.Get("client_classes").([]interface{})
	ClientClasses := make([]string, len(iClientClasses))
	for i, v := range iClientClasses {
		ClientClasses[i] = fmt.Sprint(v)
	}

	iOptionData := d.Get("option_data").([]interface{})
	OptionData := make([]OptionData, len(iOptionData))
	for i, v := range iOptionData {
		option := v.(map[string]interface{})
		OptionData[i].AlwaysSend = option["always_send"].(bool)
		OptionData[i].CSVFormat = option["csv_format"].(bool)
		OptionData[i].Code = option["code"].(int)
		OptionData[i].Data = option["data"].(string)
		OptionData[i].Name = option["name"].(string)
		OptionData[i].Space = option["space"].(string)
	}

	lease := Reservations{
		Hostname:      d.Get("name").(string),
		HWAddress:     d.Get("mac_address").(string),
		IPAddress:     d.Get("ip_address").(string),
		ClientClasses: ClientClasses,
		OptionData:    OptionData,
		NextServer:    d.Get("next_server").(string),
	}

	ok := apiClient.ReadLease(lease, d.Get("subnet_id").(int))
	if !ok {
		d.SetId("")
	}
	return nil
}
func resourceDeleteLease(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*Client)
	apiClient.lock.Lock()
	defer apiClient.lock.Unlock()

	iClientClasses := d.Get("client_classes").([]interface{})
	ClientClasses := make([]string, len(iClientClasses))
	for i, v := range iClientClasses {
		ClientClasses[i] = fmt.Sprint(v)
	}

	iOptionData := d.Get("option_data").([]interface{})
	OptionData := make([]OptionData, len(iOptionData))
	for i, v := range iOptionData {
		option := v.(map[string]interface{})
		OptionData[i].AlwaysSend = option["always_send"].(bool)
		OptionData[i].CSVFormat = option["csv_format"].(bool)
		OptionData[i].Code = option["code"].(int)
		OptionData[i].Data = option["data"].(string)
		OptionData[i].Name = option["name"].(string)
		OptionData[i].Space = option["space"].(string)
	}

	lease := Reservations{
		Hostname:      d.Get("name").(string),
		HWAddress:     d.Get("mac_address").(string),
		IPAddress:     d.Get("ip_address").(string),
		ClientClasses: ClientClasses,
		OptionData:    OptionData,
		NextServer:    d.Get("next_server").(string),
	}

	err := apiClient.DeleteLease(lease, d.Get("subnet_id").(int))
	if err != nil {
		return err
	}
	d.SetId("")

	return nil
}
