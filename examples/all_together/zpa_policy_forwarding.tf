// Access to DevOps Servers
resource "zpa_policy_forwarding_rule" "sgio_devops_bypass" {
  name                          = "SGIO DevOps Servers Bypass"
  description                   = "SGIO DevOps Servers Bypass"
  action                        = "BYPASS"
  operator = "AND"
  policy_set_id = data.zpa_global_policy_forwarding.policyset.id

  conditions {
    negated = false
    operator = "OR"
    operands {
      name =  "SGIO DevOps Servers"
      object_type = "APP"
      lhs = "id"
      rhs_list = [zpa_application_segment.as_sgio_devops.id]
    }
  }
  conditions {
     negated = false
     operator = "OR"
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs_list = [data.zpa_scim_groups.engineering.id]
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}