package zscaler

import (
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/servergroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerGroupRead,
		Schema: map[string]*schema.Schema{
			"applications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"appconnectorgroups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"citycountry": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"countrycode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creationtime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dnsquerytype": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"geolocationid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"latitude": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"longitude": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"modifiedby": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"modifiedtime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connectors": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"creationtime": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"fingerprint": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"issuedcertid": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"modifiedby": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"modifiedtime": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"upgradeattempt": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"servergroups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configspace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"creationtime": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"dynamicdiscovery": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"modifiedby": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"modifiedtime": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"siemappconnectorgroup": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"upgradetimeinsecs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"upgradeday": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"versionprofileid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"configspace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creationtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ipanchored": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dynamicdiscovery": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"modifiedby": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modifiedtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"appservergroupids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"configspace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creationtime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"modifiedby": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"modifiedtime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceServerGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for server group %s\n", id)

	resp, _, err := zClient.servergroup.Get(id)
	if err != nil {
		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("configspace", resp.ConfigSpace)
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("dynamicdiscovery", resp.DynamicDiscovery)
	_ = d.Set("ipanchored", resp.IpAnchored)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("applications", flattenServerGroupApplications(resp.Applications))
	_ = d.Set("servers", flattenServers(resp.Servers))
	//_ = d.Set("appconnectorgroups", flattenAppConnectorGroups(resp))

	if err := d.Set("appconnectorgroups", flattenAppConnectorGroups(resp)); err != nil {
		return err
	}

	return nil
}

func flattenServerGroupApplications(applications []servergroup.Applications) []interface{} {
	serverGroupApplications := make([]interface{}, len(applications))
	for i, srvApplication := range applications {
		serverGroupApplications[i] = map[string]interface{}{
			"id":   srvApplication.ID,
			"name": srvApplication.Name,
		}
	}

	return serverGroupApplications
}

func flattenAppConnectorGroups(appConnectorGroup *servergroup.ServerGroup) []interface{} {
	appConnectorGroups := make([]interface{}, len(appConnectorGroup.AppConnectorGroups))
	for i, appConnectorGroup := range appConnectorGroup.AppConnectorGroups {
		appConnectorGroups[i] = map[string]interface{}{
			"citycountry":           appConnectorGroup.Citycountry,
			"countrycode":           appConnectorGroup.CountryCode,
			"creationtime":          appConnectorGroup.CreationTime,
			"description":           appConnectorGroup.Description,
			"dnsquerytype":          appConnectorGroup.DnsqueryType,
			"enabled":               appConnectorGroup.Enabled,
			"geolocationid":         appConnectorGroup.GeolocationId,
			"id":                    appConnectorGroup.ID,
			"latitude":              appConnectorGroup.Latitude,
			"location":              appConnectorGroup.Location,
			"longitude":             appConnectorGroup.Longitude,
			"modifiedby":            appConnectorGroup.ModifiedBy,
			"modifiedtime":          appConnectorGroup.ModifiedTime,
			"name":                  appConnectorGroup.Name,
			"siemappconnectorgroup": appConnectorGroup.SiemAppconnectorGroup,
			"upgradeday":            appConnectorGroup.UpgradeDay,
			"upgradetimeinsecs":     appConnectorGroup.UpgradeTimeinSecs,
			"versionprofileid":      appConnectorGroup.VersionProfileId,
			"servergroups":          flattenAppConnectorServerGroups(appConnectorGroup),
			"connectors":            flattenAppConnectors(appConnectorGroup),
		}
	}

	return appConnectorGroups
}

func flattenAppConnectorServerGroups(serverGroup servergroup.AppConnectorGroups) []interface{} {
	serverGroups := make([]interface{}, len(serverGroup.AppServerGroups))
	for i, serverGroup := range serverGroup.AppServerGroups {
		serverGroups[i] = map[string]interface{}{
			"configSpace":      serverGroup.ConfigSpace,
			"creationtime":     serverGroup.CreationTime,
			"description":      serverGroup.Description,
			"enabled":          serverGroup.Enabled,
			"id":               serverGroup.ID,
			"dynamicdiscovery": serverGroup.DynamicDiscovery,
			"modifiedby":       serverGroup.ModifiedBy,
			"modifiedtime":     serverGroup.ModifiedTime,
			"name":             serverGroup.Name,
		}
	}

	return serverGroups
}

func flattenAppConnectors(connector servergroup.AppConnectorGroups) []interface{} {
	appConnectors := make([]interface{}, len(connector.Connectors))
	for i, appConnector := range connector.Connectors {
		appConnectors[i] = map[string]interface{}{
			"creationtime": appConnector.CreationTime,
			"description":  appConnector.Description,
			"enabled":      appConnector.Enabled,
			"id":           appConnector.ID,
			"modifiedby":   appConnector.ModifiedBy,
			"modifiedtime": appConnector.ModifiedTime,
			"name":         appConnector.Name,
		}
	}

	return appConnectors
}

func flattenServers(applicationServer []servergroup.ApplicationServer) []interface{} {
	applicationServers := make([]interface{}, len(applicationServer))
	for i, appServerItem := range applicationServer {
		applicationServers[i] = map[string]interface{}{
			"address":           appServerItem.Address,
			"appservergroupids": appServerItem.AppServerGroupIds,
			"configspace":       appServerItem.ConfigSpace,
			"creationtime":      appServerItem.CreationTime,
			"description":       appServerItem.Description,
			"enabled":           appServerItem.Enabled,
			"id":                appServerItem.ID,
			"modifiedby":        appServerItem.ModifiedBy,
			"modifiedtime":      appServerItem.ModifiedTime,
			"name":              appServerItem.Name,
		}
	}
	return applicationServers
}
