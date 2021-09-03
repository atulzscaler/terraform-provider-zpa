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

output "all_policy_set_global" {
  value = data.zpa_policy_set_global.all
}



// resource "zpa_policyset_rule" "example" {
//   name                          = "example1"
//   description                   = "example1"
//   action                        = "ALLOW"
//   ruleorder                     = 1
//   policysetid = data.zpa_policy_set_global.all.id
//   policytype = 1
// }

resource "zpa_policyset_rule" "example" {
  name                          = "example1"
  description                   = "example1"
  action                        = "ALLOW"
  ruleorder                     = 1
  policy_set_id = data.zpa_policy_set_global.all.id
  conditions {
      negated = false
    operands {
        object_type = "APP"
        lhs = "id"
        rhs = data.zpa_application_segment.all_other_services.id
        operator = "AND"  
    }
    operands {
        name = "SGIO-User-Okta"
        object_type = ["IDP"]
        operator = "AND"  
    }
    operands {
        name = "SGIO-User-Okta"
        object_type = ["IDP"]
        operator = "AND"  
    }
  }
}