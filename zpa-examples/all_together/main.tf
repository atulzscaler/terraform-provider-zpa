terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

resource "zpa_application_server" "sales" {
  name                          = "sales.securitygeek.io"
  description                   = "sales.securitygeek.io"
  address                       = "sales.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "intranet" {
  name                          = "intranet.securitygeek.io"
  description                   = "intranet.securitygeek.io"
  address                       = "intranet.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "qa" {
  name                          = "qa.securitygeek.io"
  description                   = "qa.securitygeek.io"
  address                       = "qa.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "jenkins" {
  name                          = "jenkins.securitygeek.io"
  description                   = "jenkins.securitygeek.io"
  address                       = "jenkins.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "pan220" {
  name                          = "pan220.securitygeek.io"
  description                   = "pan220.securitygeek.io"
  address                       = "pan220.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "vcenter" {
  name                          = "vcenter.securitygeek.io"
  description                   = "vcenter.securitygeek.io"
  address                       = "vcenter.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "cahlesx01" {
  name                          = "cahlesx01.securitygeek.io"
  description                   = "cahlesx01.securitygeek.io"
  address                       = "cahlesx01.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "cahlesx02" {
  name                          = "cahlesx02.securitygeek.io"
  description                   = "cahlesx02.securitygeek.io"
  address                       = "cahlesx02.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "trafficgen" {
  name                          = "trafficgen.securitygeek.io"
  description                   = "trafficgen.securitygeek.io"
  address                       = "trafficgen.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "zpa131" {
  name                          = "zpa131.securitygeek.io"
  description                   = "zpa131.securitygeek.io"
  address                       = "zpa131.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "splunk" {
  name                          = "splunk.securitygeek.io"
  description                   = "splunk.securitygeek.io"
  address                       = "splunk.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "nss128" {
  name                          = "nss128.securitygeek.io"
  description                   = "nss128.securitygeek.io"
  address                       = "nss128.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "rdp125" {
  name                          = "rdp125.securitygeek.io"
  description                   = "rdp125.securitygeek.io"
  address                       = "rdp125.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "rdp126" {
  name                          = "rdp126.securitygeek.io"
  description                   = "rdp126.securitygeek.io"
  address                       = "rdp126.securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

resource "zpa_application_server" "vcd125-ad01" {
  name                          = "vcd125-ad01.securitygeek.io"
  description                   = "vcd125-ad01.securitygeek.io"
  address                       = "vcd125-ad01securitygeek.io"
  enabled                       = true
  appservergroupids             = [ for a in each.value.app_server_appconnectorgroups : metanetworks_group.this[a].id ]
}

data "zpa_app_connector_group" "sgio-vancouver" {
  name = "SGIO-Vancouver"
}

// Server Groups
resource "zpa_server_group" "all_other_services" {
  name = "all_other_services"
  description = "All Other Services"
  enabled = true
  dynamicdiscovery = true
  appconnectorgroups {
    id = [data.zpa_app_connector_group.sgio-vancouver.id]
  }
}

resource "zpa_server_group" "sgio_devops_servers" {
  name = "sgio_devops_servers"
  description = "SGIO DevOps Servers"
  enabled = true
  dynamicdiscovery = false
    servers {
      id = [
        zpa_application_server.jenkins.id
      ]
  }
    appconnectorgroups {
    id = [
      data.zpa_app_connector_group.sgio-vancouver.id
    ]
  }
}

resource "zpa_server_group" "sgio_intranet_web_apps" {
  name = "sgio_intranet_web_apps"
  description = "SGIO Intranet Web Apps"
  enabled = true
  dynamicdiscovery = false
    servers {
      id = [
        zpa_application_server.intranet.id
        zpa_application_server.qa.id
        zpa_application_server.sales.id
      ]
    }
    appconnectorgroups {
    id = [
      data.zpa_app_connector_group.sgio-vancouver.id
    ]
  }
}

resource "zpa_server_group" "sgio_monitoring_tools" {
  name = "sgio_monitoring_tools"
  description = "SGIO Monitoring Tools"
  enabled = true
  dynamicdiscovery = false
    servers {
      id = [
        zpa_application_server.splunk.id
      ]
    }
    appconnectorgroups {
    id = [
      data.zpa_app_connector_group.sgio-vancouver.id
    ]
  }
}

resource "zpa_server_group" "sgio_rdp_servers" {
  name = "sgio_rdp_servers"
  description = "SGIO RDP Services Group"
  enabled = true
  dynamicdiscovery = false
    servers {
      id = [
        zpa_application_server.rdp125.id
        zpa_application_server.rdp126.id
      ]
    }
    appconnectorgroups {
    id = [
      data.zpa_app_connector_group.sgio-vancouver.id
    ]
  }
}

resource "zpa_server_group" "sgio_rdp_servers" {
  name = "sgio_rdp_servers"
  description = "SGIO RDP Services Group"
  enabled = true
  dynamicdiscovery = false
    servers {
      id = [
        zpa_application_server.rdp125.id
        zpa_application_server.rdp126.id
      ]
    }
    appconnectorgroups {
    id = [
      data.zpa_app_connector_group.sgio-vancouver.id
    ]
  }
}

resource "zpa_server_group" "sgio_ssh_access" {
  name = "sgio_ssh_access"
  description = "SGIO SSH Access"
  enabled = true
  dynamicdiscovery = false
    servers {
      id = [
        zpa_application_server.splunk.id
        zpa_application_server.pan220.id
        zpa_application_server.trafficgen.id
        zpa_application_server.nss128.id
      ]
    }
    appconnectorgroups {
    id = [
      data.zpa_app_connector_group.sgio-vancouver.id
    ]
  }
}

resource "zpa_server_group" "sgio_vcenter_server" {
  name = "sgio_vcenter_server"
  description = "SGIO vCenter Server"
  enabled = true
  dynamicdiscovery = false
    servers {
      id = [
        zpa_application_server.vcenter.id
        zpa_application_server.cahlesx01.id
        zpa_application_server.cahlesx02.id
      ]
    }
    appconnectorgroups {
    id = [
      data.zpa_app_connector_group.sgio-vancouver.id
    ]
  }
}

// Creating Segment Groups
resource "zpa_segment_group" "all_other_services" {
  name = "all_other_services"
  description = "All Other Services"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.all_other_services.id
    ]
  }
}

resource "zpa_segment_group" "sgio_browser_access_apps" {
  name = "sgio_browser_access_apps"
  description = "SGIO Browser Access Apps"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.sgio_browser_access_apps.id
    ]
  }
}

resource "zpa_segment_group" "sgio_crm_apps_group" {
  name = "sgio_crm_apps_group"
  description = "SGIO CRM App Group"
  enabled = true
  policymigrated = true,
  // applications {
  //   id = [
  //     zpa_application_segment.sgio_browser_access_apps.id
  //   ]
  }
}

resource "zpa_segment_group" "sgio_devops_servers" {
  name = "sgio_devops_servers"
  description = "SGIO DevOps Servers"
  enabled = true
  policymigrated = true,
  // applications {
  //   id = [
  //     zpa_application_segment.sgio_browser_access_apps.id
  //   ]
  }
}

resource "zpa_segment_group" "sgio_domain_controllers" {
  name = "sgio_domain_controllers"
  description = "SGIO Domain Controllers"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.sgio_domain_controllers.id
    ]
  }
}

resource "zpa_segment_group" "sgio_intranet_web_apps" {
  name = "sgio_intranet_web_apps"
  description = "SGIO Intranet Web Apps"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.sgio_intranet_web_apps.id
    ]
  }
}

resource "zpa_segment_group" "sgio_monitoring_tools" {
  name = "sgio_monitoring_tools"
  description = "SGIO Monitoring Tools"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.sgio_monitoring_tools.id
    ]
  }
}

resource "zpa_segment_group" "sgio_rdp_services_group" {
  name = "sgio_rdp_services_group"
  description = "SGIO RDP Services Group"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.sgio_rdp_services_group.id
    ]
  }
}

resource "zpa_segment_group" "sgio_ssh_access" {
  name = "sgio_ssh_access"
  description = "SGIO SSH Access"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.sgio_ssh_access.id
    ]
  }
}

resource "zpa_segment_group" "sgio_ssh_access" {
  name = "sgio_ssh_access"
  description = "SGIO SSH Access"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.sgio_ssh_access.id
    ]
  }
}

resource "zpa_segment_group" "sgio_vcenter_server" {
  name = "sgio_vcenter_server"
  description = "SGIO vCenter Server"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.sgio_vcenter_server.id
    ]
  }
}

resource "zpa_segment_group" "sgio_vcenter_server" {
  name = "sgio_vcenter_server"
  description = "SGIO vCenter Server"
  enabled = true
  policymigrated = true,
  applications {
    id = [
      zpa_application_segment.sgio_vcenter_server.id
    ]
  }
}

output "zpa_segment_group" {
  value = zpa_segment_group.zpa
}

// Testing Data Source Segment Group
data "zpa_segment_group" "example" {
  id = 216196257331282105
}

output "all_segment_group" {
  value = data.zpa_segment_group.example
}

// Testing Data Source App Connector Group
data "zpa_app_connector_group" "example" {
  id = 216196257331281931
}

output "all_app_connector_group" {
  value = data.zpa_app_connector_group.example
}

// Testing Data Source Server Group
data "zpa_server_group" "example" {
  id = 216196257331282100
}

output "all_server_group" {
  value = data.zpa_server_group.example
}

// Testing Data Source Posture Profile
data "zpa_posture_profile" "example" {
}

output "all_posture_profile" {
  value = data.zpa_posture_profile.example
}

// Testing Data Source Trusted Network 
data "zpa_trusted_network" "example" {
}

output "all_trusted_network" {
  value = data.zpa_trusted_network.example
}

// Testing Data Source Machine Group
data "zpa_machine_group" "all" {
    // totalpages = 20
 // id = 216196257331282185
}

output "all_machine_group" {
  value = data.zpa_machine_group.all
}

// Testing Data Source IdP Controller
data "zpa_idp_controller" "all" {
  id = 216196257331282178
}

output "idp_controller" {
    value = data.zpa_idp_controller.all
}

// Testing Data Source Saml Attributes
data "zpa_saml_attribute" "example" {
}

output "all_saml_attribute" {
  value = data.zpa_saml_attribute.example
}

// Testing Data Source Application Management
data "zpa_application_segment" "all" {
  id = 216196257331282101
}

output "application_segment" {
    value = data.zpa_application_segment.all
}