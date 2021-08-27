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


//  resource "zpa_segment_group" "example1" {
//    name = "example"
//    description = "example"
//    enabled = true
//    policymigrated = true
//     // applications  {
//     //     //name = [data.zpa_application_segment.application_segment.name]
//     //     id = 216196257331282544
//     // }
//  }

resource "zpa_policyset_rule" "example" {
  name                          = "example1"
  description                   = "example1"
  action                        = "ALLOW"
  ruleorder                     = 1
  // conditions {
  //   operands {
  //       name = "Example"
  //       objecttype = "APP"
  //       operator = "AND"
  //   }
  //   operands {
  //       name = "SGIO-User-Okta"
  //       objecttype = ["SCIM_GROUP"]
  //       operator = "AND"  
  //   }
  // }
}