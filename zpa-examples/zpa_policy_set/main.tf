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



resource "zpa_policyset_rule" "example2" {

  name                          = "example2"
  description                   = "example2"
  action                        = "ALLOW"
  rule_order                     = 1
  policy_set_id = data.zpa_policy_set_global.all.id
  policy_type = 1
}


