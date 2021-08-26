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

// output "all_policy_set_global" {
//   value = data.zpa_policy_set_global.all
// }



resource "zpa_policyset_rule" "example" {
  //policysetid = data.zpa_policy_set_global.all.id
  name                          = "example1"
  description                   = "example1"
  action                        = "ALLOW"
  ruleorder                     = 1
  // conditions {
  //   operands {
  //       name = "Example"
  //       objecttype = "APP"
  //       operator = "AND"
  //   }
  //   operands {
  //       name = "SGIO-User-Okta"
  //       objecttype = ["SCIM_GROUP"]
  //       operator = "AND"  
  //   }
  // }
}