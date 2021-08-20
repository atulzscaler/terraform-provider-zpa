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



data "zscaler_cloud_connector_group" "all" {
}

output "all_cloud_connector_group" {
  value = data.zscaler_cloud_connector_group.all
}
