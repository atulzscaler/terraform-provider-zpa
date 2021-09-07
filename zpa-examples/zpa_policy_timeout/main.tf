terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}


resource "zpa_policy_timeout" "all_other_services" {
  name                          = "All Other Services"
  description                   = "All Other Services"
  action                        = "RE_AUTH"
  reauth_idle_timeout = "600"
  reauth_timeout = "172800"
  rule_order                     = 1
  operator = "AND"
  policy_set_id = data.zpa_policy_timeout.all.id
}

output "all_zpa_policy_timeout" {
  value = zpa_policy_timeout.all_other_services
}

data "zpa_policy_timeout" "all" {
}

// output "all_zpa_policyset_rule" {
//   value = data.zpa_policy_timeout.all
// }

/*


data "zpa_application_segment" "all_other_services"{
  id = 216196257331283285
}

data "zpa_idp_controller" "sgio_user_okta" {
 id = 216196257331281933
}

data "zpa_scim_groups" "engineering" {
 id = 255066
}
*/