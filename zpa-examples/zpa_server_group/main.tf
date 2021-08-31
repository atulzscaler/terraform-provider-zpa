terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

resource "zpa_application_server" "intranet" {
  name                          = "intranet.securitygeek.io"
  description                   = "intranet.securitygeek.io"
  address                       = "intranet.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ zpa_server_group.example1.id ]
}

resource "zpa_server_group" "example1" {
  name = "example1"
  description = "example1"
  enabled = true
  dynamicdiscovery = false
  // applications {
  //   id = 216196257331282730
  // }
  // servers {
  //   id = zpa_application_server.intranet.id
  // }
  appconnectorgroups {
    id = 216196257331281931
  }
}