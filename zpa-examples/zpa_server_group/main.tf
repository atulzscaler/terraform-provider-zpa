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
  //id = 216196257331282435
  id = 216196257331282097
  //name = "SGIO-CORP-Server-Group"
}

output "all_server_group" {
  value = data.zpa_server_group.example
}

// data "zpa_application_server" "example" {
//   id = 216196257331282449

// }

// // output "all_application_server" {
// //   value = data.zpa_application_server.example
// // }

// data "zpa_app_connector_group" "example" {
//   id = 216196257331281931
//   //name = "SGIO-Vancouver"
// }

// // output "all_app_connector_group" {
// //   value = data.zpa_app_connector_group.example.name
// // }

// resource "zpa_server_group" "example" {
//   name = "example"
//   description = "example"
//   enabled = true
//   dynamicdiscovery = true
//   // servers {
//   //   id = [data.zpa_application_server.example.id]
//   // }
//   appconnectorgroups {
//     id = [data.zpa_app_connector_group.example.id]
//   }
// }