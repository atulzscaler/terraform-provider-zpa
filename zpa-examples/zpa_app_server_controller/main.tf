terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

data "zpa_application_server" "example" {
  name = "rdp126.securitygeek.io"
}

output "all_application_server" {
  value = data.zpa_application_server.example
}
