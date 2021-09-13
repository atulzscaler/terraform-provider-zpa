// Access to vCenter Server Rule
resource "zpa_policyset_rule" "as_vcenter_servers" {
  name                          = "SGIO vCenter Servers"
  description                   = "SGIO vCenter Servers"
  action                        = "ALLOW"
  rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
  conditions {
    negated = false
    operator = "OR"
    operands {
      name =  "SGIO vCenter Servers"
      object_type = "APP"
      lhs = "id"
      rhs = zpa_application_segment.as_vcenter_servers.id
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

// Access to Intranet Web Apps
resource "zpa_policyset_rule" "as_intranet_web_apps" {
  name                          = "SGIO Intranet Web Apps"
  description                   = "SGIO Intranet Web Apps"
  action                        = "ALLOW"
  rule_order                     = 3
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
  conditions {
    negated = false
    operator = "OR"
    operands {
      name =  "SGIO Intranet Web Apps"
      object_type = "APP"
      lhs = "id"
      rhs = zpa_application_segment.as_intranet_web_apps.id
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
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.sales.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.finance.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.executives.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}

// Browser Access Rule
resource "zpa_policyset_rule" "browser_access_apps" {
  name                          = "Browser Access Apps"
  description                   = "Browser Access Apps"
  action                        = "ALLOW"
  rule_order                     = 5
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
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
      rhs = data.zpa_scim_groups.sales.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.finance.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.executives.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}

// Access to all other Apps
resource "zpa_policyset_rule" "all_other_services" {
  name                          = "All Other Services"
  description                   = "All Other Services"
  action                        = "ALLOW"
 rule_order                     = 4
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
  conditions {
    negated = false
    operator = "OR"
    operands {
      name =  "All Other Services"
      object_type = "APP"
      lhs = "id"
      rhs = zpa_application_segment.all_other_services.id
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
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.sales.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.finance.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.executives.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}
