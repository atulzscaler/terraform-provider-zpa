terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

// CrowdStrike_ZTA_Score_Policy
resource "zpa_policyset_rule" "crwd_zpa_pre_zta" {
  name                          = "CrowdStrike_ZPA_Pre-ZTA"
  description                   = "CrowdStrike_ZPA_Pre-ZTA"
  action                        = "DENY"
  rule_order                    = 1
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
  conditions {
    negated = false
    operator = "OR"
    operands {
      object_type = "APP_GROUP"
      lhs = "id"
      rhs = zpa_segment_group.sg_sgio_intranet_web_apps.id
    }
  }
  conditions {
    negated = false
    operator = "OR"
    operands {
      object_type = "APP"
      lhs = "id"
      rhs = zpa_application_segment.as_intranet_web_apps.id
    }
  }
  conditions {
    negated = false
    operator = "OR"
    operands {
      object_type = "POSTURE"
      lhs = data.zpa_posture_profile.crwd_zpa_pre_zta.posture_udid
      rhs = false
    }
  }
  conditions {
     negated = false
     operator = "OR"
    operands {
      object_type = "SAML"
      lhs = data.zpa_saml_attribute.email_sgio_user_sso.id
      rhs = "alison.abbas@securitygeek.io"
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}


resource "zpa_server_group" "sgio_intranet_web_apps" {
  name = "SGIO Intranet Web Apps"
  description = "SGIO Intranet Web Apps"
  enabled = true
  dynamic_discovery = false
  servers {
    id = [
      zpa_application_server.intranet.id
    ]
  }
  app_connector_groups {
    id = [data.zpa_app_connector_group.sgio-vancouver.id]
  }
}

resource "zpa_application_segment" "as_intranet_web_apps" {
    name = "SGIO Intranet Web Apps"
    description = "SGIO Intranet Web Apps"
    enabled = true
    health_reporting = "ON_ACCESS"
    bypass_type = "NEVER"
    tcp_port_ranges = ["80", "80"]
    domain_names = ["intranet.securitygeek.io", "qa.securitygeek.io"]
    segment_group_id = zpa_segment_group.sg_sgio_intranet_web_apps.id
    server_groups {
        id = [zpa_server_group.sgio_intranet_web_apps.id]
    }
}

   resource "zpa_segment_group" "sg_sgio_intranet_web_apps" {
   name = "SGIO Intranet Web Apps"
   description = "SGIO Intranet Web Apps"
   enabled = true
   policy_migrated = true
 }

 resource "zpa_application_server" "intranet" {
  name                          = "intranet.securitygeek.io"
  description                   = "intranet.securitygeek.io"
  address                       = "intranet.securitygeek.io"
  enabled                       = true
}

data "zpa_policy_set_global" "all" {}

data "zpa_idp_controller" "sgio_user_okta" {
 name = "SGIO-User-Okta"
}

data "zpa_saml_attribute" "email_sgio_user_sso" {
    name = "Email_SGIO-User-Okta"
}

data "zpa_posture_profile" "crwd_zpa_pre_zta" {
 name = "CrowdStrike_ZPA_Pre-ZTA"
}

data "zpa_app_connector_group" "sgio-vancouver" {
  name = "SGIO-Vancouver"
}
