terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

data "zpa_scim_attribute_header" "department" {
    name = "department"
    idp_name = "SGIO-User-Okta"
}

output "scim_groups_department" {
    value = data.zpa_scim_attribute_header.department.id
}