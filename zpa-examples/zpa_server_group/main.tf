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
  //id = 216196257331282482
  name = "SGIO-CORP-Server-Group"
}

output "all_server_group" {
  value = data.zpa_server_group.example.id
}

/*
resource "zpa_server_group" "example" {
  name = "example"
  description = "example"
  enabled = true
  ipanchored = false
  dynamicdiscovery = true
  appconnectorgroups {
    id = [data.zpa_app_connector_group.example.name]
  }
}
*/