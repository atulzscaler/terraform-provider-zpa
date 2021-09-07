terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}


resource "zpa_policy_forwarding" "example" {
  name                          = "example"
  description                   = "example"
  action                        = "INTERCEPT"
  operator = "AND"
  policy_set_id = data.zpa_policy_forwarding.all.id
}

output "all_zpa_policy_timeout" {
  value = zpa_policy_forwarding.example
}

data "zpa_policy_forwarding" "all" {
}