resource "kea-dhcp4_host_lease" "host3" {
  subnet_id   = 1
  name        = "host3.local"
  mac_address = "00:50:56:ba:2d:e6"
  ip_address  = "10.0.151.3"
}

resource "kea-dhcp4_host_lease" "host4" {
  subnet_id   = 1
  name        = "host4.local"
  mac_address = "00:50:56:ba:4f:39"
  ip_address  = "10.0.151.4"
}

resource "kea-dhcp4_host_lease" "host5" {
  subnet_id   = 1
  name        = "host5.local"
  mac_address = "00:50:56:ba:47:e4"
  ip_address  = "10.0.151.5"
}

resource "kea-dhcp4_host_lease" "host6" {
  subnet_id   = 1
  name        = "host6.local"
  mac_address = "00:50:56:ba:22:5d"
  ip_address  = "10.0.151.6"
}

resource "kea-dhcp4_host_lease" "host7" {
  subnet_id   = 1
  name        = "host7.local"
  mac_address = "00:50:56:ba:b6:40"
  ip_address  = "10.0.151.7"
}

resource "kea-dhcp4_host_lease" "host8" {
  subnet_id   = 1
  name        = "host8.local"
  mac_address = "00:50:56:ba:b6:1b"
  ip_address  = "10.0.151.8"
}

resource "kea-dhcp4_host_lease" "host9" {
  subnet_id   = 1
  name        = "host9.local"
  mac_address = "00:50:56:ba:de:5b"
  ip_address  = "10.0.151.9"
}

resource "kea-dhcp4_host_lease" "host10" {
  subnet_id   = 1
  name        = "host10.local"
  mac_address = "00:50:56:ba:9d:56"
  ip_address  = "10.0.151.10"
}

resource "kea-dhcp4_host_lease" "host11" {
  subnet_id   = 1
  name        = "host11.local"
  mac_address = "00:50:56:ba:c9:af"
  ip_address  = "10.0.151.11"
}

resource "kea-dhcp4_host_lease" "host12" {
  subnet_id   = 1
  name        = "host12.local"
  mac_address = "00:50:56:ba:90:ee"
  ip_address  = "10.0.151.12"
}
