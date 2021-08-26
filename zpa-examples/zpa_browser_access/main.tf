terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

 resource "zpa_segment_group" "sg_sgio_browser_access" {
   name = "SGIO Browser Access Apps"
   description = "SGIO Browser Access Apps"
   enabled = true
   policymigrated = true
    // applications  {
    //     //name = [data.zpa_application_segment.application_segment.name]
    //     id = 216196257331282544
    // }
 }


resource "zpa_browser_access" "sg_sgio_browser_access_apps" {
    name = "SGIO Browser Access Apps"
    description = "SGIO Browser Access Apps"
    enabled = true
    healthreporting = "ON_ACCESS"
    ipanchored = false
    doubleencrypt = false
    bypasstype = "NEVER"
    iscnameenabled = false
    tcpportranges = ["80", "80"]
    domainnames = ["sales.securitygeek.io"]
    segmentgroupid = zpa_segment_group.sg_sgio_browser_access.id

    clientlessapps {
        name = "sales.securitygeek.io"
        applicationprotocol = "HTTPS"
        applicationport = "80"
        certificateid = "216196257331282104"
        trustuntrustedcert = true
        enabled = true
        domain = "sales.securitygeek.io"
    }

    servergroups {
        id = 216196257331282476
    }
}

output "all_browser_access" {
  value = zpa_browser_access.sg_sgio_browser_access_apps
}

//216196257331282481