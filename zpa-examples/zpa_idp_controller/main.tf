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

data "zscaler_idp_controller" "all" {
    name = "SGIO-Admin-Azure"
//  id = 216196257331282178
}

output "idp_controller" {
    value = data.zscaler_idp_controller.all.id
}