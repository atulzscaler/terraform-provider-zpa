// Creating Segment Groups
 resource "zpa_segment_group" "sg_all_other_services" {
   name = "All Other Services"
   description = "All Other Services"
   enabled = true
   policy_migrated = true
 }

  resource "zpa_segment_group" "sg_sgio_browser_access" {
   name = "Browser Access Apps"
   description = "Browser Access Apps"
   enabled = true
   policy_migrated = true
 }

  resource "zpa_segment_group" "sg_sgio_devops" {
   name = "SGIO DevOps Servers"
   description = "SGIO DevOps Servers"
   enabled = true
   policy_migrated = true
 }

   resource "zpa_segment_group" "sg_sgio_intranet_web_apps" {
   name = "SGIO Intranet Web Apps"
   description = "SGIO Intranet Web Apps"
   enabled = true
   policy_migrated = true
 }

    resource "zpa_segment_group" "sg_sgio_vcenter_servers" {
   name = "SGIO vCenter Servers"
   description = "SGIO vCenter Servers"
   enabled = true
   policy_migrated = true
 }