terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

// data "zpa_application_server" "example" {
//   name = "rdp126.securitygeek.io"
// }

// output "all_application_server" {
//   value = data.zpa_application_server.example
// }


resource "zpa_application_server" "example10" {
  name                          = "example10.securitygeek.io"
  description                   = "example10.securitygeek.io"
  address                       = "example10.securitygeek.io"
  enabled                       = true
}