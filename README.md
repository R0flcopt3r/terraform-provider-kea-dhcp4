# go-terraform-kea-dhcp4

Kea - Dhcp4 plugin for Terraform
=================================


This is a plugin for https://www.isc.org/kea/ DHCP daemon to update its configuration, 
mostly it can be used if you need to control DHCP leases for any kind of Terraform infrastructure as a code.



Installation:
------------

1. "git clone" this repo
2. "go build" inside of it
3. Copy binary file into '~/.terraform.d/plugins/registry.terraform.io/feliksas/kea-dhcp4/<VERSION>/linux_amd64/terraform-provider-kea-dhcp4'
4. update `~/.terraformrc`:

```
provider_installation {
  filesystem_mirror {
    path    = "/home/<USER>/.terraform.d/plugins"
  }
  direct {
    exclude = [
    "registry.terraform.io/feliksas/*",
    ]
  }
}
```


Configuration and Usage:
------------------------

1. The plugin has been written for connection to Kea API with basic auth, so please set up your kea REST API with basic authorisation,
    (it can be done if you have hidden the API behind a HTTP server, like Apache/Nginx/Traefik etc.), its strongly recommended using https protocol to keep credentials safety.
2. To set credentials to Kea API you can use export variables:

```
    export TF_VAR_kea_user=User
    export TF_VAR_kea_pass=Password
```
3. Now to use kea provider put something like this in your .tf files:
```
    # define kea-dhcp4 provider
    provider "kea-dhcp4" {
        kea_server_address    = "https://your.kea.server/kea"
        kea_server_username   = "${var.kea_user}"
        kea_server_password   = "${var.kea_pass}"
        kea_server_configfile = "/etc/kea/kea-dhcp4.conf"
    }
    # create resource lease in dhcp conf.
    resource "kea-dhcp4_host_lease" "exampleVM" {
        subnet_id = 1
        name = "exampleVM"
        mac_address = "aa:bb:cc:00:11:22"
        ip_address = "192.168.10.10"
    }
    option_data  {
        data = "10.0.0.1"
        name = "log-servers"
    }
    option_data  {
        data = "10.0.0.1"
        name = "ntp-servers"
    }
```

Notice: you need to have already defined subnets in your kea server, 
this plugin allows you to define host lease only for host defined in terraform any kind of infrastructure which use Kea DHCPd. 


author michal.lis.1987@gmail.com

GNU Public License see LICENSE file.
