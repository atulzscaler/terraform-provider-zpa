terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

// data "zpa_server_group" "example" {
//   id = 216196257331282482
//   //name = "SGIO-CORP-Server-Group"
// }

// data "zpa_server_group" "example1" {
//   id = 216196257331282476
//   //name = "SGIO-CORP-Server-Group"
// }

// output "all_server_group" {
//   value = data.zpa_server_group.example.name
// }

resource "zpa_application_server" "example" {
  name                          = "example"
  description                   = "example"
  address                       = "1.1.1.1"
  enabled                       = true
  appservergroupids             = [ 
    216196257331282482,
    216196257331282476
    ]
}

output "all_application_server" {
  value = zpa_application_server.example
}