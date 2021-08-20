terraform {
    required_providers {
        zscaler = {
            version = "1.0.0"
            source = "zscaler.com/zscaler/zscaler"
        }
    }
}

provider "zscaler" {
}



data "zscaler_policy_set_global" "all" {
}

output "all_policy_set_global" {
  value = data.zscaler_policy_set_global.all
}


/*
resource "zscaler_policyset_rule" "test" {
  policysetid                   = data.zscaler_policy_set_global.all.id
  name                          = "example1"
  description                   = "example1"
  action                        = "ALLOW"
  ruleorder                     = 1
  conditions {
    operands {
        name = "SGIO Domain Controllers"
        objecttype = "APP"
        operator = "AND"
    }
    operands {
        name = "SGIO-User-Okta"
        objecttype = "SCIM_GROUP"
        operator = "AND"  
    }
  }
}
*/