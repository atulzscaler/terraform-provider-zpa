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

// data "zscaler_segment_group" "all" {
//     id = "216196257331282481"
// }

// output "all_segment_group" {
//     value = data.zscaler_segment_group.all
// }

data "zscaler_application_segment" "all" { 
  id = 216196257331282483
}

output "application_segment" {
    value = data.zscaler_application_segment.all.id
}

 resource "zscaler_segment_group" "example" {
   name = "example"
   description = "example"
   enabled = true
   policymigrated = false
    applications {
        name = [data.zscaler_application_segment.all.name]
        //id = [data.zscaler_application_segment.all.id]
    }
 }