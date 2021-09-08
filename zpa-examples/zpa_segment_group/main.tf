terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

data "zpa_segment_group" "all" { 
  name = "Browser Access Apps"
}

output "segment_group" {
    value = data.zpa_segment_group.all.id
}

/*

resource "zpa_application_segment" "ap_example" {
    name = "ap_example"
    description = "ap_example"
    enabled = true
    health_reporting = "ON_ACCESS"
    bypass_type = "NEVER"
    tcp_port_ranges = ["1", "52", "54", "65535"]
    domain_names = ["*.acme.com"]
    // segmentgroupid = zpa_segment_group.sg_all_other_services.id
    server_groups {
        id = 216196257331283727
    }
}

 resource "zpa_segment_group" "sg_example" {
   name = "sg_example"
   description = "sg_example"
   enabled = true
   policy_migrated = true
    applications  {
        id = zpa_application_segment.ap_example.id
    }
 }
 */