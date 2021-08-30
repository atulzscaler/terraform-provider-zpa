terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}


data "zpa_app_connector_group" "example" {
  id = 216196257331281931
  //name = "SGIO-Vancouver"
}

resource "zpa_server_group" "example1" {
  name = "example1"
  description = "example1"
  enabled = true
  dynamicdiscovery = false
  appconnectorgroups {
    id = data.zpa_app_connector_group.example.id
  }
}

resource "zpa_server_group" "example2" {
  name = "example2"
  description = "example2"
  enabled = true
  dynamicdiscovery = false
  appconnectorgroups {
    id = data.zpa_app_connector_group.example.id
  }
}

resource "zpa_application_server" "example" {
  name                          = "example"
  description                   = "example"
  address                       = "1.1.1.1"
  enabled                       = true
  appservergroupids             = [ 
      zpa_server_group.example1.id,
      zpa_server_group.example2.id,
    ]
}

output "all_application_server" {
  value = zpa_application_server.example
}