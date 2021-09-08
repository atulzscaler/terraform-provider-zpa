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
}

output "all_app_connector_group" {
  value = data.zpa_app_connector_group.example
}

resource "zpa_server_group" "example" {
  name = "example"
  description = "example"
  enabled = true
  dynamic_discovery = true
  app_connector_groups {
    id = data.zpa_app_connector_group.example.id
  }
}

 resource "zpa_segment_group" "example" {
   name = "example"
   description = "example"
   enabled = true
   policy_migrated = false
 }

resource "zpa_application_segment" "example" {
    name = "example"
    description = "example"
    enabled = true
    health_reporting = "ON_ACCESS"
    ip_anchored = true
    double_encrypt = false
    bypass_type = "NEVER"
    is_cname_enabled = false
    icmp_access_type = true
    tcp_port_ranges = ["8080", "8080"]
    domain_names = ["acme.com"]
    segment_group_id = zpa_segment_group.example.id
    server_groups {
        id = zpa_server_group.example.id
    }
}

output "all_application_segment" {
  value = zpa_application_segment.example
}
