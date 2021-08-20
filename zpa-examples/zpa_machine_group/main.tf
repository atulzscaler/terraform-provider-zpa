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


data "zscaler_machine_group" "all" {
    totalpages = 20
 // id = 216196257331282185
}

output "all_machine_group" {
  value = data.zscaler_machine_group.all
}