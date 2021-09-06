// Access to DevOps Servers
resource "zpa_policyset_rule" "as_sgio_devops" {
  name                          = "SGIO DevOps Servers"
  description                   = "SGIO DevOps Servers"
  action                        = "ALLOW"
//   rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
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
      object_type = "IDP"
      lhs = "id"
      rhs = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.engineering.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}

resource "zpa_policyset_rule" "block_as_sgio_devops" {
  name                          = "Block SGIO DevOps Servers"
  description                   = "Block SGIO DevOps Servers"
  action                        = "DENY"
//   rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
}

// Access to vCenter Server Rule
resource "zpa_policyset_rule" "as_vcenter_servers" {
  name                          = "SGIO vCenter Servers"
  description                   = "SGIO vCenter Servers"
  action                        = "ALLOW"
//   rule_order                     = 2
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
      object_type = "IDP"
      lhs = "id"
      rhs = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.engineering.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}

resource "zpa_policyset_rule" "block_as_vcenter_servers" {
  name                          = "Block SGIO vCenter Servers"
  description                   = "Block SGIO vCenter Servers"
  action                        = "DENY"
//   rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
}

// Access to Intranet Web Apps
resource "zpa_policyset_rule" "as_intranet_web_apps" {
  name                          = "Block SGIO Intranet Web Apps"
  description                   = "Block SGIO Intranet Web Apps"
  action                        = "ALLOW"
//   rule_order                     = 2
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
      object_type = "IDP"
      lhs = "id"
      rhs = data.zpa_idp_controller.sgio_user_okta.id
    }
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

resource "zpa_policyset_rule" "block_as_intranet_web_apps" {
  name                          = "Block SGIO Intranet Web Apps"
  description                   = "Block SGIO Intranet Web Apps"
  action                        = "DENY"
//   rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
}

/*
// Browser Access Rule
resource "zpa_policyset_rule" "browser_access_apps" {
  name                          = "Browser Access Apps"
  description                   = "Browser Access Apps"
  action                        = "ALLOW"
  rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
  conditions {
    negated = false
    operator = "OR"
    operands {
      name =  "Browser Access Apps"
      object_type = "APP"
      lhs = "id"
      rhs = zpa_browser_access.browser_access_apps.id
    }
  }
  conditions {
     negated = false
     operator = "OR"
    operands {
      object_type = "IDP"
      lhs = "id"
      rhs = data.zpa_idp_controller.sgio_user_okta.id
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

resource "zpa_policyset_rule" "block_browser_access_apps" {
  name                          = "Block Browser Access Apps"
  description                   = "Block Browser Access Apps"
  action                        = "DENY"
//   rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
}
*/
// Access to all other Apps
resource "zpa_policyset_rule" "all_other_services" {
  name                          = "All Other Services"
  description                   = "All Other Services"
  action                        = "ALLOW"
//   rule_order                     = 2
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
  conditions {
     negated = false
     operator = "OR"
    operands {
      object_type = "IDP"
      lhs = "id"
      rhs = data.zpa_idp_controller.sgio_user_okta.id
    }
    operands {
      object_type = "SCIM_GROUP"
      lhs = data.zpa_idp_controller.sgio_user_okta.id
      rhs = data.zpa_scim_groups.engineering.id
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}

resource "zpa_policyset_rule" "block_all_other_services" {
  name                          = "Block Browser Access Apps"
  description                   = "Block Browser Access Apps"
  action                        = "DENY"
//   rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
}

output "all_zpa_policyset_rule" {
  value = zpa_policyset_rule.all_other_services
}

data "zpa_policy_set_global" "all" {
}

data "zpa_idp_controller" "sgio_user_okta" {
 id = 216196257331281933
}

data "zpa_scim_groups" "engineering" {
 id = "255066"
}

data "zpa_scim_groups" "sales" {
 id = "255067"
}

data "zpa_scim_groups" "finance" {
 id = "255068"
}

data "zpa_scim_groups" "executives" {
 id = "255069"
}