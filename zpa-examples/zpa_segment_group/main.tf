terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

// data "zpa_segment_group" "all" {
//     id = "216196257331282481"
// }

// output "all_segment_group" {
//     value = data.zpa_segment_group.all
// }

data "zpa_application_segment" "all" { 
  id = 216196257331282483
}

output "application_segment" {
    value = data.zpa_application_segment.all.id
}

 resource "zpa_segment_group" "example" {
   name = "example"
   description = "example"
   enabled = true
   policymigrated = false
    applications {
        name = [data.zpa_application_segment.all.name]
        //id = [data.zpa_application_segment.all.id]
    }
 }