// Creating Server Groups

resource "zpa_server_group" "all_other_services" {
  name = "All Other Services"
  description = "All Other Services"
  enabled = true
  dynamic_discovery = true
  app_connector_groups {
    id = [data.zpa_app_connector_group.sgio-vancouver.id]
  }
}

resource "zpa_server_group" "browser_access_apps" {
  name = "Browser Access Apps"
  description = "Browser Access Apps"
  enabled = true
  dynamic_discovery = false
  servers {
    id = [zpa_application_server.sales.id]
  }
  app_connector_groups {
    id = [data.zpa_app_connector_group.sgio-vancouver.id]
  }
}

resource "zpa_server_group" "sgio_devops_servers" {
  name = "SGIO DevOps Servers"
  description = "SGIO DevOps Servers"
  enabled = true
  dynamic_discovery = false
    servers {
    id = [
      zpa_application_server.jenkins.id,
      zpa_application_server.pan220.id,
      zpa_application_server.trafficgen.id,
      zpa_application_server.zpa131.id,
      zpa_application_server.splunk.id,
      zpa_application_server.nss128.id,
      zpa_application_server.rdp125.id,
      zpa_application_server.rdp126.id,
      zpa_application_server.vcd125-ad01.id,
    ]
  }
  app_connector_groups {
    id = [data.zpa_app_connector_group.sgio-vancouver.id]
  }
}

resource "zpa_server_group" "sgio_intranet_web_apps" {
  name = "SGIO Intranet Web Apps"
  description = "SGIO Intranet Web Apps"
  enabled = true
  dynamic_discovery = false
  servers {
    id = [
      zpa_application_server.intranet.id,
      zpa_application_server.qa.id
    ]
  }
  app_connector_groups {
    id = [data.zpa_app_connector_group.sgio-vancouver.id]
  }
}

resource "zpa_server_group" "sgio_vcenter_servers" {
  name = "SGIO vCenter Servers"
  description = "SGIO vCenter Servers"
  enabled = true
  dynamic_discovery = false
  servers {
    id = [
      zpa_application_server.vcenter.id,
      zpa_application_server.cahlesx01.id,
      zpa_application_server.cahlesx02.id
      ]
  }
  app_connector_groups {
    id = [data.zpa_app_connector_group.sgio-vancouver.id]
  }
}