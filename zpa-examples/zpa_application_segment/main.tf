terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

// data "zpa_app_connector_group" "example" {
//   id = 216196257331281931
//   //name = "SGIO-Vancouver"
// }

// output "all_app_connector_group" {
//   value = data.zpa_app_connector_group.example
// }

resource "zpa_server_group" "example" {
  name = "example"
  description = "example"
  enabled = true
  dynamicdiscovery = true
//   appconnectorgroups {
//     id = [data.zpa_app_connector_group.example.id]
//   }
}
/*
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
    servergroups {
        id = [zpa_server_group.example.id]
    }
}
*/