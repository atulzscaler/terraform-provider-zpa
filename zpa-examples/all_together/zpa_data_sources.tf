// Application Connector Groups
data "zpa_app_connector_group" "sgio-vancouver" {
  id = 216196257331281931
}

// Sales Browser Certificate
data "zpa_ba_certificate" "sales_ba" {
    id = 216196257331282584
}

data "zpa_policy_set_global" "all" {
}

data "zpa_idp_controller" "sgio_user_okta" {
 id = 216196257331281933
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
