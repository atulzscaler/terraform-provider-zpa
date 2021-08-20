terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

data "zpa_scim_groups" "all" {
    list {
        idpid = 216196257331282118
    }
}

output "scim_groups" {
    value = data.zpa_scim_groups.all
}