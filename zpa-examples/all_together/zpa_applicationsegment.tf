// Creating Application Segments
resource "zpa_application_segment" "all_other_services" {
    name = "All Other Services"
    description = "All Other Services"
    enabled = true
    health_reporting = "ON_ACCESS"
    bypass_type = "NEVER"
    tcp_port_ranges = ["1", "52", "54", "65535"]
    domain_names = ["*.securitygeek.io"]
    segment_group_id = zpa_segment_group.sg_all_other_services.id
    server_groups {
        id = zpa_server_group.all_other_services.id
    }
}

resource "zpa_application_segment" "as_sgio_devops" {
    name = "SGIO DevOps Servers"
    description = "SGIO DevOps Servers"
    enabled = true
    health_reporting = "ON_ACCESS"
    bypass_type = "NEVER"
    tcp_port_ranges = ["8080", "8080"]
    domain_names = ["jenkins.securitygeek.io"]
    segment_group_id = zpa_segment_group.sg_sgio_devops.id
    server_groups {
        id = zpa_server_group.sgio_devops_servers.id
    }
}

resource "zpa_application_segment" "as_vcenter_servers" {
    name = "SGIO vCenter Servers"
    description = "SGIO vCenter Servers"
    enabled = true
    health_reporting = "ON_ACCESS"
    bypass_type = "NEVER"
    tcp_port_ranges = ["8080", "8080"]
    domain_names = ["vcenter.securitygeek.io", "cahlesx01.securitygeek.io", "cahlesx02.securitygeek.io"]
    segment_group_id = zpa_segment_group.sg_sgio_vcenter_servers.id
    server_groups {
        id = zpa_server_group.sgio_vcenter_servers.id
    }
}

resource "zpa_application_segment" "as_intranet_web_apps" {
    name = "SGIO Intranet Web Apps"
    description = "SGIO Intranet Web Apps"
    enabled = true
    health_reporting = "ON_ACCESS"
    bypass_type = "NEVER"
    tcp_port_ranges = ["80", "80"]
    domain_names = ["intranet.securitygeek.io", "qa.securitygeek.io"]
    segment_group_id = zpa_segment_group.sg_sgio_intranet_web_apps.id
    server_groups {
        id = zpa_server_group.sgio_intranet_web_apps.id
    }
}