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
