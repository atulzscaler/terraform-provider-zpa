terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

resource "zpa_server_group" "example" {
  name = "example"
  description = "example"
  enabled = true
  dynamic_discovery = false
  applications {
    id = ["216196257331283686", "216196257331283691"]
  }
  servers {
    id = ["216196257331283699", "216196257331283697"]
  }
  app_connector_groups {
    id = ["216196257331281931", "216196257331282724"]
  }
}

output "all_zpa_server_group" {
  value = zpa_server_group.example
}