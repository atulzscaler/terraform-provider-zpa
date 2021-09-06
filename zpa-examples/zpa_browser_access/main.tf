terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

resource "zpa_server_group" "sg_sgio_browser_access" {
  name = "SGIO Browser Access Apps"
  description = "SGIO Browser Access Apps"
  enabled = true
  dynamic_discovery = true
  app_connector_groups {
    id = 216196257331281931
  }
}

 resource "zpa_segment_group" "sg_sgio_browser_access" {
   name = "SGIO Browser Access Apps"
   description = "SGIO Browser Access Apps"
   enabled = true
   policy_migrated = true
    // applications  {
    //     //name = [data.zpa_application_segment.application_segment.name]
    //     id = 216196257331282544
    // }
 }


// Sales Portal Browser Access
data "zpa_ba_certificate" "sales_ba" {
    id = 216196257331282584
}

// QA Browser Access
data "zpa_ba_certificate" "qa_ba" {
    id = 216196257331282583
}

// DevOps Browser Access
data "zpa_ba_certificate" "jenkins_ba" {
    id = 216196257331282582
}


resource "zpa_browser_access" "browser_access_apps" {
    name = "Browser Access Apps"
    description = "Browser Access Apps"
    enabled = true
    health_reporting = "ON_ACCESS"
    bypass_type = "NEVER"
    tcp_port_ranges = ["80", "80", "8080", "8080"]
    domain_names = ["sales.securitygeek.io", "qa.securitygeek.io", "jenkins.securitygeek.io"]
    segment_group_id = zpa_segment_group.sg_sgio_browser_access.id

    clientless_apps {
        name = "sales.securitygeek.io"
        application_protocol = "HTTP"
        application_port = "80"
        certificate_id = data.zpa_ba_certificate.sales_ba.id
        trust_untrusted_cert = true
        enabled = true
        domain = "sales.securitygeek.io"
    }
        clientless_apps {
        name = "qa.securitygeek.io"
        application_protocol = "HTTP"
        application_port = "80"
        certificate_id = data.zpa_ba_certificate.qa_ba.id
        trust_untrusted_cert = true
        enabled = true
        domain = "qa.securitygeek.io"
    }

    clientless_apps {
        name = "jenkins.securitygeek.io"
        application_protocol = "HTTP"
        application_port = "8080"
        certificate_id = data.zpa_ba_certificate.jenkins_ba.id
        trust_untrusted_cert = true
        enabled = true
        domain = "jenkins.securitygeek.io"
    }
    server_groups {
        id = zpa_server_group.sg_sgio_browser_access.id
    }
}




/*
resource "zpa_browser_access" "browser_access_apps" {
    name = "Browser Access Apps"
    description = "Browser Access Apps"
    enabled = true
    healthreporting = "ON_ACCESS"
    bypasstype = "NEVER"
    tcpportranges = ["80", "80", "8080", "8080"]
    domainnames = ["sales.securitygeek.io", "qa.securitygeek.io", "jenkins.securitygeek.io"]
    segmentgroupid = zpa_segment_group.sg_sgio_browser_access.id

    clientlessapps {
        name = "sales.securitygeek.io"
        applicationprotocol = "HTTP"
        applicationport = "80"
        certificateid = data.zpa_ba_certificate.sales_ba.id
        trustuntrustedcert = true
        enabled = true
        domain = "sales.securitygeek.io"
    }
    servergroups {
        id = 216196257331282476
    }
}

 resource "zpa_segment_group" "sg_all_other_services" {
   name = "All Other Services"
   description = "All Other Services"
   enabled = true
   policymigrated = true
    // applications  {
    //     //name = [data.zpa_application_segment.application_segment.name]
    //     id = 216196257331282544
    // }
 }

resource "zpa_application_segment" "all_other_services" {
    name = "All Other Services"
    description = "All Other Services"
    enabled = true
    healthreporting = "ON_ACCESS"
    bypasstype = "NEVER"
    tcpportranges = ["1", "52", "54", "65535"]
    domainnames = ["*.securitygeek.io"]
    segmentgroupid = zpa_segment_group.sg_all_other_services.id
    // servergroups {
    //     id = 216196257331282438
    // }
}
*/