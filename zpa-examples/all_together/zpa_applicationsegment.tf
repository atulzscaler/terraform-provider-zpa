// Creating Application Segments
resource "zpa_application_segment" "all_other_services" {
    name = "All Other Services"
    description = "All Other Services"
    enabled = true
    healthreporting = "ON_ACCESS"
    bypasstype = "NEVER"
    tcpportranges = ["1", "52", "54", "65535"]
    domainnames = ["*.securitygeek.io"]
    segmentgroupid = zpa_segment_group.sg_all_other_services.id
    servergroups {
        id = zpa_server_group.all_other_services.id
    }
}

resource "zpa_application_segment" "as_sgio_devops" {
    name = "SGIO DevOps Servers"
    description = "SGIO DevOps Servers"
    enabled = true
    healthreporting = "ON_ACCESS"
    bypasstype = "NEVER"
    tcpportranges = ["8080", "8080"]
    domainnames = ["jenkins.securitygeek.io"]
    segmentgroupid = zpa_segment_group.sg_sgio_devops.id
    servergroups {
        id = zpa_server_group.sgio_devops_servers.id
    }
}

resource "zpa_application_segment" "as_vcenter_servers" {
    name = "SGIO DevOps Servers"
    description = "SGIO DevOps Servers"
    enabled = true
    healthreporting = "ON_ACCESS"
    bypasstype = "NEVER"
    tcpportranges = ["8080", "8080"]
    domainnames = ["vcenter.securitygeek.io", "cahlesx01.securitygeek.io", "cahlesx02.securitygeek.io"]
    segmentgroupid = zpa_segment_group.sg_sgio_vcenter_servers.id
    servergroups {
        id = zpa_server_group.sgio_vcenter_servers.id
    }
}

resource "zpa_application_segment" "as_intranet_web_apps" {
    name = "SGIO Intranet Web Apps"
    description = "SGIO Intranet Web Apps"
    enabled = true
    healthreporting = "ON_ACCESS"
    bypasstype = "NEVER"
    tcpportranges = ["80", "80"]
    domainnames = ["intranet.securitygeek.io", "qa.securitygeek.io"]
    segmentgroupid = zpa_segment_group.sg_sgio_intranet_web_apps.id
    servergroups {
        id = zpa_server_group.sgio_intranet_web_apps.id
    }
}