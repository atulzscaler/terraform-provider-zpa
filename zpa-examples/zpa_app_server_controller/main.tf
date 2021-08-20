terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

data "zpa_server_group" "example" {
  id = 216196257331282482
  //name = "SGIO-CORP-Server-Group"
}

output "all_server_group" {
  value = data.zscaler_server_group.example.name
}

resource "zpa_application_server" "example1" {
  name                          = "example1"
  description                   = "example1"
  address                       = "1.1.1.2"
  enabled                       = true
  appservergroupids             = [ data.zscaler_server_group.example.id ]
}

/*
resource "zscaler_application_server" "example2" {
  name                          = "example2"
  description                   = "example2"
  address                       = "2.2.2.2"
  enabled                       = true
  appservergroupids             = [ "216196257331282106" ]
}

output "all_application_server2" {
  value = zscaler_application_server.example2
}

//    "${data.zscaler_server_group.example.id}"
// data "zscaler_server_group" "example" {
//   name = "SGIO-CORP-Server-Group"
// }

// output "all_server_group" {
//   value = "${data.zscaler_server_group.example.id}"
// }

// data "zscaler_application_server" "example" {
// }

// output "all_application_server" {
//   value = data.zscaler_application_server.example
// }

// resource "zscaler_application_server" "zpa" {
//   for_each = { for appserver in var.application_server : appserver.app_server_name => appserver }
//   name                          = each.value.app_server_name
//   description                   = each.value.app_server_description
//   address                       = each.value.app_server_address
//   enabled                       = each.value.app_server_enabled
// //  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
// }
*/