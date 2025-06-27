terraform {
  required_providers {
    kea-dhcp4 = {
      source  = "terraform.local/r0flcopt3r/kea-dhcp4"
      version = "1.0.0"
    }
  }
}

provider "kea-dhcp4" {
  kea_server_address    = "http://localhost:8080"
  kea_server_username   = "test"
  kea_server_password   = "1234"
  kea_server_configfile = "/etc/kea/kea-dhcp4.conf"
}
