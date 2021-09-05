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

/*
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
  conditions {
     negated = false
     operator = "OR"
    operands {
      object_type = "SCIM"
      // lhs = data.zpa_idp_controller.sgio_user_okta.id
      // rhs = data.zpa_scim_groups.engineering.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.engineering.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}

output "all_zpa_policyset_rule" {
  value = zpa_policyset_rule.all_other_services
}

data "zpa_policy_set_global" "all" {
}

data "zpa_application_segment" "all_other_services"{
  id = 216196257331283285
}

data "zpa_idp_controller" "sgio_user_okta" {
 id = "216196257331281956"
}

data "zpa_scim_groups" "engineering" {
 id = "255066"
}
*/