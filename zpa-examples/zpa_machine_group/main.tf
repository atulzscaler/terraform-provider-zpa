terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}


data "zpa_machine_group" "all" {
  id = 216196257331282185
}

output "all_machine_group" {
  value = data.zpa_machine_group.all
}