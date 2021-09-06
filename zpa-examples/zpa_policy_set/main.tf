terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

// data "zpa_policy_set_global" "all" {
// }

// output "all_zpa_policy_set_global" {
//   value = data.zpa_policy_set_global.all
// }


resource "zpa_policyset_rule" "all_other_services" {
  name                          = "All Other Services"
  description                   = "All Other Services"
  action                        = "ALLOW"
  rule_order                     = 2
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
  app_connector_groups {
    id = ["216196257331281931", "216196257331282724"]
  }

  app_server_groups {
    id = ["216196257331283675", "216196257331283670"]
  }

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

output "all_zpa_policyset_rule" {
  value = zpa_policyset_rule.all_other_services
}

data "zpa_policy_set_global" "all" {
}

data "zpa_application_segment" "all_other_services"{
  //id = 216196257331283691
  name = "All Other Services"
}

data "zpa_idp_controller" "sgio_user_okta" {
 id = 216196257331281933
}

data "zpa_scim_groups" "engineering" {
 id = 255066
}

data "zpa_scim_groups" "sales" {
 id = 255067
}

data "zpa_scim_groups" "finance" {
 id = 255068
}

data "zpa_scim_groups" "executives" {
 id = 255069
}


