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
//   id = 216196257331282475
// }

// output "segment_group" {
//     value = data.zpa_segment_group.all
// }


resource "zpa_application_segment" "all_other_services" {
    name = "All Other Services"
    description = "All Other Services"
    enabled = true
    healthreporting = "ON_ACCESS"
    bypasstype = "NEVER"
    tcpportranges = ["1", "52", "54", "65535"]
    domainnames = ["*.securitygeek.io"]
    segmentgroupid = zpa_segment_group.sg_all_other_services.id
    // servergroups {
    //     id = 216196257331282438
    // }
}

 resource "zpa_segment_group" "sg_all_other_services" {
   name = "All Other Services"
   description = "All Other Services"
   enabled = true
   policymigrated = true
    // applications  {
    //     id = zpa_application_segment.all_other_services.id
    // }
 }