// Creating Application Servers
resource "zpa_application_server" "sales" {
  name                          = "sales.securitygeek.io"
  description                   = "sales.securitygeek.io"
  address                       = "sales.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.browser_access_apps.id ]
}

resource "zpa_application_server" "intranet" {
  name                          = "intranet.securitygeek.io"
  description                   = "intranet.securitygeek.io"
  address                       = "intranet.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.sgio_intranet_web_apps.id ]
}

resource "zpa_application_server" "qa" {
  name                          = "qa.securitygeek.io"
  description                   = "qa.securitygeek.io"
  address                       = "qa.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.sgio_intranet_web_apps.id ]
}

resource "zpa_application_server" "jenkins" {
  name                          = "jenkins.securitygeek.io"
  description                   = "jenkins.securitygeek.io"
  address                       = "jenkins.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.sgio_devops_servers.id ]
}

// Home Lab Servers
resource "zpa_application_server" "vcenter" {
  name                          = "vcenter.securitygeek.io"
  description                   = "vcenter.securitygeek.io"
  address                       = "vcenter.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.sgio_vcenter_servers.id ]
}

resource "zpa_application_server" "cahlesx01" {
  name                          = "cahlesx01.securitygeek.io"
  description                   = "cahlesx01.securitygeek.io"
  address                       = "cahlesx01.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.sgio_vcenter_servers.id ]
}

resource "zpa_application_server" "cahlesx02" {
  name                          = "cahlesx02.securitygeek.io"
  description                   = "cahlesx02.securitygeek.io"
  address                       = "cahlesx02.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.sgio_vcenter_servers.id ]
}

resource "zpa_application_server" "pan220" {
  name                          = "pan220.securitygeek.io"
  description                   = "pan220.securitygeek.io"
  address                       = "pan220.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.all_other_services.id ]
}

resource "zpa_application_server" "trafficgen" {
  name                          = "trafficgen.securitygeek.io"
  description                   = "trafficgen.securitygeek.io"
  address                       = "trafficgen.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.all_other_services.id ]
}

resource "zpa_application_server" "zpa131" {
  name                          = "zpa131.securitygeek.io"
  description                   = "zpa131.securitygeek.io"
  address                       = "zpa131.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.all_other_services.id ]
}

resource "zpa_application_server" "splunk" {
  name                          = "splunk.securitygeek.io"
  description                   = "splunk.securitygeek.io"
  address                       = "splunk.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.all_other_services.id ]
}

resource "zpa_application_server" "nss128" {
  name                          = "nss128.securitygeek.io"
  description                   = "nss128.securitygeek.io"
  address                       = "nss128.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.all_other_services.id ]
}

resource "zpa_application_server" "rdp125" {
  name                          = "rdp125.securitygeek.io"
  description                   = "rdp125.securitygeek.io"
  address                       = "rdp125.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.all_other_services.id ]
}

resource "zpa_application_server" "rdp126" {
  name                          = "rdp126.securitygeek.io"
  description                   = "rdp126.securitygeek.io"
  address                       = "rdp126.securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.all_other_services.id ]
}

resource "zpa_application_server" "vcd125-ad01" {
  name                          = "vcd125-ad01.securitygeek.io"
  description                   = "vcd125-ad01.securitygeek.io"
  address                       = "vcd125-ad01securitygeek.io"
  enabled                       = true
  app_server_group_ids             = [ zpa_server_group.all_other_services.id ]
}