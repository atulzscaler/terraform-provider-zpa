terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

data "zpa_app_connector_group" "example" {
  name = "SGIO-Vancouver"
}

resource "zpa_server_group" "example1" {
  name = "example1"
  description = "example1"
  enabled = true
  dynamic_discovery = true
  app_connector_groups {
    id = [data.zpa_app_connector_group.example.id]
  }
}

resource "zpa_server_group" "example2" {
  name = "example2"
  description = "example2"
  enabled = true
  dynamic_discovery = true
  app_connector_groups {
    id = [data.zpa_app_connector_group.example.id]
  }
}

 resource "zpa_segment_group" "example" {
   name = "example"
   description = "example"
   enabled = true
   policy_migrated = true
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
    domain_names = ["jenkins.securitygeek.io"]
    segment_group_id = zpa_segment_group.example.id

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
        id = [
            zpa_server_group.example1.id,
            zpa_server_group.example2.id
        ]
    }
}