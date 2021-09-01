// Creating Server Groups

resource "zpa_server_group" "all_other_services" {
  name = "All Other Services"
  description = "All Other Services"
  enabled = true
  dynamic_discovery = true
  app_connector_groups {
    id = data.zpa_app_connector_group.sgio-vancouver.id
  }
}

resource "zpa_server_group" "browser_access_apps" {
  name = "Browser Access Apps"
  description = "Browser Access Apps"
  enabled = true
  dynamic_discovery = false
  app_connector_groups {
    id = data.zpa_app_connector_group.sgio-vancouver.id
  }
}

resource "zpa_server_group" "sgio_devops_servers" {
  name = "SGIO DevOps Servers"
  description = "SGIO DevOps Servers"
  enabled = true
  dynamic_discovery = false
  app_connector_groups {
    id = data.zpa_app_connector_group.sgio-vancouver.id
  }
}

resource "zpa_server_group" "sgio_intranet_web_apps" {
  name = "SGIO Intranet Web Apps"
  description = "SGIO Intranet Web Apps"
  enabled = true
  dynamic_discovery = false
  app_connector_groups {
    id = data.zpa_app_connector_group.sgio-vancouver.id
  }
}

resource "zpa_server_group" "sgio_vcenter_servers" {
  name = "SGIO vCenter Servers"
  description = "SGIO vCenter Servers"
  enabled = true
  dynamic_discovery = false
  app_connector_groups {
    id = data.zpa_app_connector_group.sgio-vancouver.id
  }
}