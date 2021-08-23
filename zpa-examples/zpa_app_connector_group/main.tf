
terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {
}

data "zpa_app_connector_group" "example" {
  id = 216196257331281931
  //name = "SGIO-Vancouver"
}

output "all_app_connector_group" {
  value = data.zpa_app_connector_group.example.name
}