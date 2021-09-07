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
  conditions {
     negated = false
     operator = "OR"
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.engineering.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
}

output "all_zpa_policy_timeout" {
  value = zpa_policy_timeout.all_other_services
}

data "zpa_policy_timeout" "all" {
}