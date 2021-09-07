terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}


// Testing Data Source Posture Profile
data "zpa_trusted_network" "example" {
    // id = "216196257331283524"
 name = "SGIO-Trusted-Networks"
}

output "all_trusted_network" {
  value = data.zpa_trusted_network.example.id
}