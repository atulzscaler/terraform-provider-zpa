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
    id = 255069
    // idp_id = "216196257331281933"
}

output "scim_groups" {
    value = data.zpa_scim_groups.all
}