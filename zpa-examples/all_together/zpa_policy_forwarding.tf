// Access to DevOps Servers
resource "zpa_policy_forwarding" "sgio_devops_bypass" {
  name                          = "SGIO DevOps Servers Bypass"
  description                   = "SGIO DevOps Servers Bypass"
  action                        = "BYPASS"
  operator = "AND"
  policy_set_id = data.zpa_policy_forwarding.all.id

  conditions {
    negated = false
    operator = "OR"
    operands {
      name =  "SGIO DevOps Servers"
      object_type = "APP"
      lhs = "id"
      rhs = zpa_application_segment.as_sgio_devops.id
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
}

data "zpa_policy_forwarding" "all" {
}