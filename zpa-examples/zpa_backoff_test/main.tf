terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}


 resource "zpa_segment_group" "sg_all_other_services" {
   count = "5"
   name = "Terraform-${count.index + 1}"
   description = "All Other Services"
   enabled = true
   policy_migrated = true
 }