
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

data "zscaler_app_connector_group" "example" {
  //id = 216196257331281931
  name = "SGIO-Vancouver"
}

output "all_app_connector_group" {
  value = data.zscaler_app_connector_group.example.id
}