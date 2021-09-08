// Application Connector Groups
data "zpa_app_connector_group" "sgio-vancouver" {
  name = "SGIO-Vancouver"
}

// Sales Browser Certificate
data "zpa_ba_certificate" "sales_ba" {
    name = "sales.securitygeek.io"
}

data "zpa_policy_set_global" "all" {
}

data "zpa_idp_controller" "sgio_user_okta" {
 id = "SGIO-User-Okta"
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
