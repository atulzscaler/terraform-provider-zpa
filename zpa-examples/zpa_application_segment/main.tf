terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

// data "zpa_application_segment" "all" { 
//   id = 216196257331282452
// }

// output "application_segment" {
//     value = data.zpa_application_segment.all
// }


// data "zscaler_server_group" "example" {
//   id = 216196257331282482
//   //name = "SGIO-CORP-Server-Group"
// }

resource "zpa_application_segment" "example" {
    name = "example"
    description = "example"
    enabled = false
    healthreporting = "ON_ACCESS"
    ipanchored = false
    doubleencrypt = false
    bypasstype = "NEVER"
    iscnameenabled = true
    tcpportranges = ["8080", "8080"]
    domainnames = ["acme.com"]
    clientlessApps {
        appid = 
    }
    segmentgroupid = "216196257331282481"
    // servergroups {
    //     id = 216196257331282097
    // }
}