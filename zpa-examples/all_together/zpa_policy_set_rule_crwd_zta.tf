// CrowdStrike_ZTA_Score_Policy
resource "zpa_policyset_rule" "crwd_zta_score_40" {
  name                          = "CrowdStrike_ZTA_Score_40"
  description                   = "CrowdStrike_ZTA_Score_40"
  action                        = "DENY"
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
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
      lhs = data.zpa_posture_profile.crwd_zta_score_40.posture_udid
      rhs = false
    }
  }
  conditions {
     negated = false
     operator = "OR"
    operands {
      object_type = "SAML"
      lhs = data.zpa_saml_attribute.email_user_sso.id
      rhs = "alison.abbas@securitygeek.io"
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}

resource "zpa_policyset_rule" "crwd_zta_score_80" {
  name                          = "CrowdStrike_ZTA_Score_80"
  description                   = "CrowdStrike_ZTA_Score_80"
  action                        = "DENY"
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
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
      lhs = data.zpa_posture_profile.crwd_zta_score_80.posture_udid
      rhs = false
    }
  }
  conditions {
     negated = false
     operator = "OR"
    operands {
      object_type = "SAML"
      lhs = data.zpa_saml_attribute.email_user_sso.id
      rhs = "alison.abbas@securitygeek.io"
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}

resource "zpa_policyset_rule" "crwd_zpa_pre_zta" {
  name                          = "CrowdStrike_ZPA_Pre-ZTA"
  description                   = "CrowdStrike_ZPA_Pre-ZTA"
  action                        = "DENY"
  operator = "AND"
  policy_set_id = data.zpa_policy_set_global.all.id
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
      lhs = data.zpa_saml_attribute.email_user_sso.id
      rhs = "alison.abbas@securitygeek.io"
      idp_id = data.zpa_idp_controller.sgio_user_okta.id
    }
  }
}

