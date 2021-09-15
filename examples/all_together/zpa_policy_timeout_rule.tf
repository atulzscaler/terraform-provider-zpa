resource "zpa_policy_timeout" "temp_contractors_browser_access" {
  name                          = "Temp Contractors Browser Access"
  description                   = "Temp Contractors Browser Access"
  action                        = "RE_AUTH"
  reauth_idle_timeout           = "600"
  reauth_timeout                = "172800"
  operator                      = "AND"
  policy_set_id                 = data.zpa_policy_timeout.all.id

  conditions {
    negated = false
    operator = "OR"
    operands {
      name =  "zpn_client_type_exporter"
      object_type = "CLIENT_TYPE"
      lhs = "id"
      rhs = "zpn_client_type_exporter"
    }
  }
  conditions {
     negated = false
     operator = "OR"
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.contractors.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}