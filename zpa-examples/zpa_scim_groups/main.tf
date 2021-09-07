terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

data "zpa_scim_groups" "engineering" {
    name = "Engineering"
    idp_name = "SGIO-User-Okta"
}

output "scim_groups_engineering" {
    value = data.zpa_scim_groups.engineering.id
}

data "zpa_scim_groups" "sales" {
    name = "Sales"
    idp_name = "SGIO-User-Okta"
}

output "scim_groups_sales" {
    value = data.zpa_scim_groups.sales.id
}

data "zpa_scim_groups" "executives" {
    name = "Executives"
    idp_name = "SGIO-User-Okta"
}

output "scim_groups_executives" {
    value = data.zpa_scim_groups.executives.id
}

data "zpa_scim_groups" "finance" {
    name = "Finance"
    idp_name = "SGIO-User-Okta"
}

output "scim_groups_finance" {
    value = data.zpa_scim_groups.finance.id
}