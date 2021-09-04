terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}


data "zpa_policy_set_global" "all" {
}

output "all_zpa_policy_set_global" {
  value = data.zpa_policy_set_global.all
}

data "zpa_application_segment" "all_other_services"{
  id = 216196257331283285
}

resource "zpa_policyset_rule" "all_other_services" {
  name                          = "All Other Services"
  description                   = "All Other Services"
  action                        = "ALLOW"
  rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
  conditions {
     negated = false
     operator = "OR"
    operands {
      name =  "All Other Services"
      object_type = "APP"
      lhs = "id"
      rhs = data.zpa_application_segment.all_other_services.id 
    }
  }
}



// resource "zpa_policyset_rule" "example" {
//   name                          = "example1"
//   description                   = "example1"
//   action                        = "ALLOW"
//   ruleorder                     = 1
//   policysetid = data.zpa_policy_set_global.all.id
//   policytype = 1
// }
