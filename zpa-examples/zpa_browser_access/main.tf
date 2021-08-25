terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

resource "zpa_browser_access" "example" {
    name = "example"
    description = "example"
    enabled = true
    healthreporting = "ON_ACCESS"
    ipanchored = true
    doubleencrypt = false
    bypasstype = "NEVER"
    iscnameenabled = false
    tcpportranges = ["443", "443"]
    domainnames = ["example.securitygeek.io"]
    segmentgroupid = 216196257331282481
    clientlessapps {
        name = "example.securitygeek.io"
        allowoptions = false
        applicationprotocol = "HTTPS"
        applicationport = "443"
        certificateid = "216196257331282104"
        trustuntrustedcert = true
        enabled = true
        domain = "example.securitygeek.io"
    }
    servergroups {
        id = 216196257331282097
    }
}

output "all_application_segment" {
  value = zpa_browser_access.example
}
