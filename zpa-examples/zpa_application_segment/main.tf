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
  name = "SGIO-Vancouver"
}

data "zpa_policy_set_global" "all" {
}

resource "zpa_server_group" "example" {
  name = "example"
  description = "example"
  enabled = true
  dynamic_discovery = false
  applications {
    id = [zpa_application_segment.example.id]
  }
  app_connector_groups {
    id = [data.zpa_app_connector_group.example.id]
  }
}

resource "zpa_segment_group" "example" {
  name = "example"
  description = "example"
  enabled = true
  policy_migrated = true
}

resource "zpa_application_segment" "example" {
    name = "example"
    description = "example"
    enabled = true
    health_reporting = "ON_ACCESS"
    bypass_type = "NEVER"
    is_cname_enabled = true
    tcp_port_ranges = ["8080", "8080"]
    domain_names = ["acme.com"]
    segment_group_id = zpa_segment_group.example.id
    // server_groups {
    //     id = [ zpa_server_group.example.id]
    // }
}

resource "zpa_policyset_rule" "example" {
  name                          = "example"
  description                   = "example"
  action                        = "ALLOW"
  rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id

  conditions {
    negated = false
    operator = "OR"
    operands {
      name =  "example"
      object_type = "APP"
      lhs = "id"
      rhs = zpa_application_segment.example.id
    }
  }
}

