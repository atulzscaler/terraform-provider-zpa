terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

// resource "zpa_application_server" "intranet" {
//   name                          = "intranet.securitygeek.io"
//   description                   = "intranet.securitygeek.io"
//   address                       = "intranet.securitygeek.io"
//   enabled                       = true
//   appservergroupids             = [ zpa_server_group.example1.id ]
// }

// data "zpa_app_connector_group" "example1" {
//   id = 216196257331281931
// }

// data "zpa_app_connector_group" "example2" {
//   id = 216196257331282724
// }

resource "zpa_server_group" "example1" {
  name = "example1"
  description = "example1"
  enabled = true
  dynamic_discovery = true
  // applications {
  //   id = 216196257331282730
  // }
  // servers {
  //   id = zpa_application_server.intranet.id
  // }
  app_connector_groups {
    id = "216196257331281931"
  }
}

data "zpa_server_group" "example1" {
  id = zpa_server_group.example1.id
}

output "all_zpa_server_group" {
  value = data.zpa_server_group.example1
}