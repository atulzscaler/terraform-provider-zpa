// Application Connector Groups
data "zpa_app_connector_group" "sgio-vancouver" {
  name = "SGIO-Vancouver"
}


// Sales Browser Certificate
data "zpa_ba_certificate" "sales_ba" {
    name = "sales.securitygeek.io"
}

// Posture Profiles
data "zpa_posture_profile" "crwd_zta_score_40" {
 name = "CrowdStrike_ZPA_ZTA_40"
}

data "zpa_posture_profile" "crwd_zta_score_80" {
 name = "CrowdStrike_ZPA_ZTA_80"
}

data "zpa_posture_profile" "crwd_zpa_pre_zta" {
 name = "CrowdStrike_ZPA_Pre-ZTA"
}

data "zpa_global_access_policy" "policyset" {}

data "zpa_global_policy_timeout" "policyset" {}

data "zpa_global_policy_forwarding" "policyset" {}

data "zpa_idp_controller" "sgio_user_okta" {
 name = "SGIO-User-Okta"
}

data "zpa_saml_attribute" "email_sgio_user_sso" {
    name = "Email_SGIO-User-Okta"
}

// Okta IDP SCIM Groups
data "zpa_scim_groups" "engineering" {
  name = "Engineering"
  idp_name = "SGIO-User-Okta"
}

data "zpa_scim_groups" "sales" {
  name = "Sales"
  idp_name = "SGIO-User-Okta"
}

data "zpa_scim_groups" "executives" {
  name = "Executives"
  idp_name = "SGIO-User-Okta"
}

data "zpa_scim_groups" "finance" {
  name = "Finance"
  idp_name = "SGIO-User-Okta"
}

data "zpa_scim_groups" "contractors" {
  name = "Contractors"
  idp_name = "SGIO-User-Okta"
}