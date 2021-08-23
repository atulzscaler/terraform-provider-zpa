terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

// resource "zpa_application_server" "example" {
//   name                          = "example"
//   description                   = "example"
//   address                       = "1.1.1.1"
//   enabled                       = "true"
// //  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
// }



// data "zpa_server_group" "example" {
//   id = 216196257331282482
//   //name = "SGIO-CORP-Server-Group"
// }


data "zpa_app_connector_group" "example" {
name = "SGIO-Vancouver"
//id = 216196257331281931
}


resource "zpa_server_group" "example" {
  name = "example"
  description = "example"
  enabled = true
  ipanchored = false
  dynamicdiscovery = true
  appconnectorgroups {
    id = [data.zpa_app_connector_group.example.id]
  }
}