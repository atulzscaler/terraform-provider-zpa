
####################### Application Segments Outputs ####################### 
// Application Segment All Other Services
output "zpa_application_segment_all_other_services" {
  value = zpa_application_segment.all_other_services
}

// Application Segment SGIO DevOps
output "zpa_application_segment_as_sgio_devops" {
  value = zpa_application_segment.as_sgio_devops
}

// Application Segment SGIO vCenter Servers
output "zpa_application_segment_as_vcenter_servers" {
  value = zpa_application_segment.as_vcenter_servers
}

// Application Segment SGIO Intranet Web Apps
output "zpa_application_segment_as_intranet_web_apps"  {
  value = zpa_application_segment.as_intranet_web_apps
}

####################### Segment Groups Outputs #######################
// Segment Group All Other Services
output "zpa_segment_group_sg_all_other_services" {
  value = zpa_segment_group.sg_all_other_services
}

// Segment Group Browser Access
output "zpa_segment_group_sg_sgio_browser_access" {
  value = zpa_segment_group.sg_sgio_browser_access
}

// Segment Group SGIO DevOps
output "zpa_segment_group_sg_sgio_devops" {
  value = zpa_segment_group.sg_sgio_devops
}

// Segment Group SGIO Intranet Web Apps
output "zpa_segment_group_sg_sgio_intranet_web_apps" {
  value = zpa_segment_group.sg_sgio_intranet_web_apps
}

// Segment Group SGIO Intranet Web Apps
output "zpa_segment_group_sg_sgio_vcenter_servers" {
  value = zpa_segment_group.sg_sgio_vcenter_servers
}

####################### Server Groups Outputs ####################### 
// Segment Group SGIO All Other Servers
output "zpa_server_group_all_other_services"{
  value = zpa_server_group.all_other_services
}

// Segment Group SGIO Browser Access Apps
output "zpa_server_group_browser_access_apps"{
  value = zpa_server_group.browser_access_apps
}

// Segment Group SGIO DevOps Servers
output "zpa_server_group_sgio_devops_servers"{
  value = zpa_server_group.sgio_devops_servers
}

// Segment Group SGIO Intranet Web Apps
output "zpa_server_group_sgio_intranet_web_apps" {
  value = zpa_server_group.sgio_intranet_web_apps
}

// Segment Group SGIO vCenter Servers
output "zpa_server_group_sgio_vcenter_servers" {
  value = zpa_server_group.sgio_vcenter_servers
}


####################### Access Policy Outputs ####################### 
// Access to DevOps Servers
output "zpa_policyset_rule_as_sgio_devops" {
  value = zpa_policyset_rule.as_sgio_devops
}

// Access to vCenter Server Rule
output "zpa_policyset_rule_as_vcenter_servers"  {
  value = zpa_policyset_rule.as_vcenter_servers
}

// Access to Intranet Web Apps
output "zpa_policyset_rule_as_intranet_web_apps"   {
  value = zpa_policyset_rule.as_intranet_web_apps
}

// Browser Access Rule
output "zpa_policyset_rule_browser_access_apps"  {
  value = zpa_policyset_rule.browser_access_apps
}

// Access to all other Apps
output "zpa_policyset_rule_all_other_services" {
  value = zpa_policyset_rule.all_other_services
}