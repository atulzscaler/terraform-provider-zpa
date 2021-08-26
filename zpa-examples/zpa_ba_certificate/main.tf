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
    id = 216196257331282104

}

output "all_zpa_ba_certificate" {
  value = data.zpa_ba_certificate.all
}