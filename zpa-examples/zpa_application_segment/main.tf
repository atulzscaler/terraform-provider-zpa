terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

/*
data "zpa_application_segment" "all" { 
  //id = 216196257331282452
  id = 216196257331282477
}

output "application_segment" {
    value = data.zpa_application_segment.all
}
*/

 resource "zpa_segment_group" "example" {
   name = "example"
   description = "example"
   enabled = true
   policymigrated = false
 }

resource "zpa_application_segment" "example" {
    name = "example"
    description = "example"
    enabled = true
    healthreporting = "ON_ACCESS"
    ipanchored = true
    doubleencrypt = false
    bypasstype = "NEVER"
    iscnameenabled = false
    tcpportranges = ["8080", "8080"]
    domainnames = ["acme.com"]
    segmentgroupid = zpa_segment_group.example.id
    // servergroups {
    //     id = [216196257331282482]
    // }
}