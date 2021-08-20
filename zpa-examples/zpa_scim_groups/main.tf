terraform {
    required_providers {
        zscaler = {
            version = "1.0.0"
            source = "zscaler.com/zscaler/zscaler"
        }
    }
}

provider "zscaler" {
}

data "zscaler_scim_groups" "all" {
    list {
        idpid = 216196257331282118
    }
}

output "scim_groups" {
    value = data.zscaler_scim_groups.all
}