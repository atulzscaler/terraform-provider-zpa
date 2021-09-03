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

resource "zpa_policyset_rule" "example2" {

  name                          = "example4"
  description                   = "example4"
  action                        = "ALLOW"
  rule_order                     = 1
  policy_set_id = data.zpa_policy_set_global.all.id
  policy_type = 1
}

output "all_zpa_policyset_rule" {
  value = zpa_policyset_rule.example2
}
