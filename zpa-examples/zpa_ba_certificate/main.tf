terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

data "zpa_ba_certificate" "all" {
    id = 216196257331282582
    // name = "jenkins.securitygeek.io"

}

output "all_zpa_ba_certificate" {
  value = data.zpa_ba_certificate.all
}