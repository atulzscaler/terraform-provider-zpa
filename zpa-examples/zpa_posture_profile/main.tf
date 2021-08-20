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


// Testing Data Source Posture Profile
data "zscaler_posture_profile" "example" {
    id = 216196257331282068
//  name = "sgio-windows-domain-joined"
}

output "all_posture_profile" {
  value = data.zscaler_posture_profile.example.name
}