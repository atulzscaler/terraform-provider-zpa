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

data "zscaler_application_segment" "all" { 
  id = 216196257331282483
}

output "application_segment" {
    value = data.zscaler_application_segment.all.name
}


// data "zscaler_server_group" "example" {
//   id = 216196257331282482
//   //name = "SGIO-CORP-Server-Group"
// }

// resource "zscaler_application_segment" "example" {
//     name = "example"
//     description = "example"
//     enabled = false
//     healthreporting = "ON_ACCESS"
//     ipanchored = false
//     doubleencrypt = false
//     bypasstype = "NEVER"
//     iscnameenabled = true
//     tcpportranges = ["8080", "8080"]
//     domainnames = ["acme.com"]
//     segmentgroupid = "216196257331282481"
//     // servergroups {
//     //     name = "SGIO RDP Services Group"
//     // }
// }