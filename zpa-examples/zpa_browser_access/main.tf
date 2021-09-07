terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

resource "zpa_server_group" "example" {
  name = "example"
  description = "example"
  enabled = true
  dynamic_discovery = true
  app_connector_groups {
    id = 216196257331281931
  }
}

 resource "zpa_segment_group" "example" {
   name = "example"
   description = "example"
   enabled = true
   policy_migrated = true
    // applications  {
    //     //name = [data.zpa_application_segment.application_segment.name]
    //     id = 216196257331282544
    // }
 }

// DevOps Browser Access
data "zpa_ba_certificate" "jenkins_ba" {
    id = 216196257331282582
}


resource "zpa_browser_access" "jenkins_app" {
    name = "jenkins_app"
    description = "jenkins_app"
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
    app_server_groups {
        id = [zpa_server_group.example.id]
    }
}