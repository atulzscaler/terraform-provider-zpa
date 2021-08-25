package zscaler

import (
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/servergroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"applications": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "This field is a json array of app-connector-id only.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"appconnectorgroups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "This field is a json array of app-connector-id only.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			// "configspace": {
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			// "creationtime": {
			// 	Type:     schema.TypeInt,
			// 	Computed: true,
			// },
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This field is the description of the server group.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "This field defines if the server group is enabled or disabled.",
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipanchored": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dynamicdiscovery": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "This field controls dynamic discovery of the servers.",
			},
			// "modifiedby": {
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			// "modifiedtime": {
			// 	Type:     schema.TypeInt,
			// 	Computed: true,
			// },
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This field defines the name of the server group.",
			},
			"servers": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API. For existing servers, pass only the serverId.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
					Create: resourceServerGroupCreate,
					Read:   resourceServerGroupRead,
					Update: resourceServerGroupUpdate,
					Delete: resourceServerGroupDelete,
					Importer: &schema.ResourceImporter{
						State: schema.ImportStatePassthrough,
					},
				},
			},
		},
	}
}

func resourceServerGroupCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandCreateAppServerGroupRequest(d)
	log.Printf("[INFO] Creating zpa server group with request\n%+v\n", req)

	resp, _, err := zClient.servergroup.Create(req)
	if err != nil {
		return err
	}
	d.SetId(resp.ID)

	return resourceServerGroupRead(d, m)
}

func resourceServerGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	resp, _, err := zClient.servergroup.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing server group %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting server group:\n%+v\n", resp)
	d.SetId(resp.ID)
	// _ = d.Set("configspace", resp.ConfigSpace)
	// _ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("ipanchored", resp.IpAnchored)
	_ = d.Set("dynamicdiscovery", resp.DynamicDiscovery)
	_ = d.Set("enabled", resp.Enabled)
	// _ = d.Set("modifiedby", resp.ModifiedBy)
	// _ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("applications", flattenServerGroupApplications(resp.Applications))
	_ = d.Set("servers", flattenServers(resp.ApplicationServers))

	if err := d.Set("appconnectorgroups", flattenAppConnectorGroups(resp)); err != nil {
		return err
	}
	//_ = d.Set("appconnectorgroups", flattenAppConnectorGroups(resp))
	return nil

}

func resourceServerGroupUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	log.Println("An updated occurred")

	if d.HasChange("name") {
		log.Println("The name or ID has been changed")

		if _, err := zClient.servergroup.Update(d.Id(), servergroup.ServerGroupRequest{
			Name: d.Get("name").(string),
		}); err != nil {
			return err
		}
		return resourceServerGroupRead(d, m)
	}

	return nil
}

func resourceServerGroupDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	log.Printf("[INFO] Updating server group ID: %v\n", d.Id())

	if _, err := zClient.servergroup.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func expandCreateAppServerGroupRequest(d *schema.ResourceData) servergroup.ServerGroupRequest {
	serverGroup := servergroup.ServerGroupRequest{
		//ID:                 d.Get("id").(string),
		Enabled:          d.Get("enabled").(bool),
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		IpAnchored:       d.Get("ipanchored").(bool),
		ConfigSpace:      d.Get("configspace").(string),
		DynamicDiscovery: d.Get("dynamicdiscovery").(bool),
		//Applications:       expandServerGroupApplications(d),
		AppConnectorGroups: expandAppConnectorGroups(d),
		//ApplicationServers: expandServers(d),
	}
	return serverGroup
}

/*
func expandServerGroupApplications(d *schema.ResourceData) []servergroup.Applications {
	var serverGroupApplications []servergroup.Applications
	if applicationsInterface, ok := d.GetOk("applications"); ok {
		applications := applicationsInterface.([]interface{})
		serverGroupApplications = make([]servergroup.Applications, len(applications))
		for i, application := range applications {
			srvApplication := application.(map[string]interface{})
			serverGroupApplications[i] = servergroup.Applications{
				ID:   srvApplication["id"].(int64),
				Name: srvApplication["name"].(string),
			}
		}
	}

	return serverGroupApplications
}
*/
func expandAppConnectorGroups(d *schema.ResourceData) []servergroup.AppConnectorGroups {
	var appConnectorGroups []servergroup.AppConnectorGroups
	if appConnectorGroupsInterface, ok := d.GetOk("appconnectorgroups"); ok {
		groups := appConnectorGroupsInterface.([]interface{})
		appConnectorGroups = make([]servergroup.AppConnectorGroups, len(groups))
		for i, group := range groups {
			connectorGroup := group.(map[string]interface{})
			appConnectorGroups[i] = servergroup.AppConnectorGroups{
				// Citycountry:  connectorGroup["citycountry"].(string),
				// CountryCode:  connectorGroup["countrycode"].(string),
				// CreationTime: connectorGroup["creationtime"].(int),
				// Description:  connectorGroup["description"].(string),
				// DnsqueryType: connectorGroup["dnsquerytype"].(string),
				// Enabled:      connectorGroup["enabled"].(bool),
				//GeolocationId:         connectorGroup["geolocationid"].(int64),
				ID: connectorGroup["id"].(int64),
				// Latitude:              connectorGroup["latitude"].(string),
				// Location:              connectorGroup["location"].(string),
				// Longitude:             connectorGroup["longitude"].(string),
				// ModifiedBy:            connectorGroup["modifiedby"].(int64),
				// ModifiedTime:          connectorGroup["modifiedtime"].(int32),
				// Name:                  connectorGroup["name"].(string),
				// SiemAppconnectorGroup: connectorGroup["siemappconnectorgroup"].(bool),
				// UpgradeDay:            connectorGroup["upgradeday"].(string),
				// UpgradeTimeinSecs:     connectorGroup["upgradetimeinsecs"].(string),
				// VersionProfileId:      connectorGroup["versionprofileid"].(int64),
			}
		}
	}

	return appConnectorGroups
}

/*
func expandServers(d *schema.ResourceData) []servergroup.ApplicationServers {
	var applicationServers []servergroup.ApplicationServers
	if appServersInterface, ok := d.GetOk("servers"); ok {
		appservers := appServersInterface.([]interface{})
		applicationServers = make([]servergroup.ApplicationServers, len(appservers))
		for i, appserver := range appservers {
			appSrv := appserver.(map[string]interface{})
			applicationServers[i] = servergroup.ApplicationServers{
				Address:           appSrv["address"].(string),
				AppServerGroupIds: appSrv["appservergroupids"].([]string),
				ConfigSpace:       appSrv["configspace"].(string),
				// CreationTime:      appSrv["creationtime"].(int32),
				Description: appSrv["description"].(string),
				ID:          appSrv["id"].(int64),
				Enabled:     appSrv["enabled"].(bool),
				// ModifiedBy:        appSrv["modifiedby"].(int64),
				// ModifiedTime:      appSrv["modifiedtime"].(int32),
				Name: appSrv["name"].(string),
			}
		}
	}

	return applicationServers
}
*/
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

func flattenAppConnectorGroups(appConnectorGroup *servergroup.ServerGroupResponse) []interface{} {
	appConnectorGroups := make([]interface{}, len(appConnectorGroup.AppConnectorGroups))
	for i, appConnectorGroup := range appConnectorGroup.AppConnectorGroups {
		appConnectorGroups[i] = map[string]interface{}{
			"citycountry":  appConnectorGroup.Citycountry,
			"countrycode":  appConnectorGroup.CountryCode,
			"creationtime": appConnectorGroup.CreationTime,
			"description":  appConnectorGroup.Description,
			"dnsquerytype": appConnectorGroup.DnsqueryType,
			"enabled":      appConnectorGroup.Enabled,
			//"geolocationid":         appConnectorGroup.GeolocationId,
			//"id":                    appConnectorGroup.ID,
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
			"configSpace":  serverGroup.ConfigSpace,
			"creationtime": serverGroup.CreationTime,
			"description":  serverGroup.Description,
			"enabled":      serverGroup.Enabled,
			//"id":               serverGroup.ID,
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
			//"id":           appConnector.ID,
			"modifiedby":   appConnector.ModifiedBy,
			"modifiedtime": appConnector.ModifiedTime,
			"name":         appConnector.Name,
		}
	}

	return appConnectors
}

func flattenServers(applicationServer []servergroup.ApplicationServers) []interface{} {
	applicationServers := make([]interface{}, len(applicationServer))
	for i, appServerItem := range applicationServer {
		applicationServers[i] = map[string]interface{}{
			"address":           appServerItem.Address,
			"appservergroupids": appServerItem.AppServerGroupIds,
			"configspace":       appServerItem.ConfigSpace,
			"creationtime":      appServerItem.CreationTime,
			"description":       appServerItem.Description,
			"enabled":           appServerItem.Enabled,
			//"id":                appServerItem.ID,
			"modifiedby":   appServerItem.ModifiedBy,
			"modifiedtime": appServerItem.ModifiedTime,
			"name":         appServerItem.Name,
		}
	}
	return applicationServers
}
