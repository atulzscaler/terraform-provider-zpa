package zscaler

import (
	"fmt"
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/servergroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerGroupCreate,
		Read:   resourceServerGroupRead,
		Update: resourceServerGroupUpdate,
		Delete: resourceServerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"configspace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"creationtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
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
			"modifiedby": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modifiedtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This field defines the name of the server group.",
			},
		},
	}
}

func resourceServerGroupCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	if zClient == nil {
		return resourceNotSupportedError()
	}

	req := expandCreateAppServerGroupRequest(d)
	// log.Printf("[INFO] Creating zpa server group with request\n%+v\n", req)

	resp, _, err := zClient.servergroup.Create(req)
	if err != nil {
		return err
	}

	// log.Printf("[INFO] Created zpa server group request. ID: %v\n", resp)
	d.SetId(resp.ID)

	return resourceServerGroupRead(d, m)
}

func resourceServerGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	if zClient == nil {
		return resourceNotSupportedError()
	}

	id := d.Id()
	if id == "" {
		return fmt.Errorf("error obtaining server group id")
	}

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

	_ = d.Set("configspace", resp.ConfigSpace)
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("ipanchored", resp.IpAnchored)
	_ = d.Set("dynamicdiscovery", resp.DynamicDiscovery)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	// _ = d.Set("appconnectorgroups", flattenAppConnectorGroups(resp.AppConnectorGroups))
	// _ = d.Set("applications", flattenServerGroupApplications(resp.Applications))
	// _ = d.Set("servers", flattenServers(resp.Servers))

	// if err := d.Set("applications", flattenServerGroupApplications(resp.Applications)); err != nil {
	// 	return err
	// }

	// if err := d.Set("appconnectorgroups", flattenAppConnectorGroups(resp.AppConnectorGroups)); err != nil {
	// 	return err
	// }

	// if err := d.Set("servers", flattenServers(resp.Servers)); err != nil {
	// 	return err
	// }
	return nil

}

func resourceServerGroupUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	if zClient == nil {
		return resourceNotSupportedError()
	}

	id := d.Id()
	log.Printf("[INFO] Updating server group ID: %v\n", id)
	req := expandCreateAppServerGroupRequest(d)

	if _, err := zClient.servergroup.Update(id, req); err != nil {
		return err
	}
	return resourceServerGroupRead(d, m)
}

func resourceServerGroupDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	if zClient == nil {
		return resourceNotSupportedError()
	}

	id := d.Id()
	if id == "" {
		return fmt.Errorf("error obtaining server group id")
	}

	log.Printf("[INFO] Deleting server group ID: %v\n", d.Id())
	if _, err := zClient.servergroup.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func expandCreateAppServerGroupRequest(d *schema.ResourceData) servergroup.ServerGroup {
	return servergroup.ServerGroup{
		// serverGroup := servergroup.ServerGroup{
		ID:               d.Get("id").(string),
		Enabled:          d.Get("enabled").(bool),
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		IpAnchored:       d.Get("ipanchored").(bool),
		ConfigSpace:      d.Get("configspace").(string),
		DynamicDiscovery: d.Get("dynamicdiscovery").(bool),
		CreationTime:     d.Get("creationtime").(int32),
		ModifiedBy:       d.Get("modifiedby").(string),
		ModifiedTime:     d.Get("Modifiedtime").(int32),
		// Applications:       expandServerGroupApplications(d),
		// AppConnectorGroups: expandAppConnectorGroups(d),
		// Servers:            expandServers(d),
	}
	// return serverGroup
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
				// GeolocationId:         connectorGroup["geolocationid"].(int64),
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

func expandServers(d *schema.ResourceData) []servergroup.ApplicationServer {
	var applicationServers []servergroup.ApplicationServer
	if appServersInterface, ok := d.GetOk("servers"); ok {
		appservers := appServersInterface.([]interface{})
		applicationServers = make([]servergroup.ApplicationServer, len(appservers))
		for i, appserver := range appservers {
			appSrv := appserver.(map[string]interface{})
			applicationServers[i] = servergroup.ApplicationServer{
				// Address:           appSrv["address"].(string),
				// AppServerGroupIds: appSrv["appservergroupids"].([]string),
				// ConfigSpace:       appSrv["configspace"].(string),
				// CreationTime:      appSrv["creationtime"].(int32),
				// Description: appSrv["description"].(string),
				ID: appSrv["id"].(int64),
				// Enabled:     appSrv["enabled"].(bool),
				// ModifiedBy:        appSrv["modifiedby"].(int64),
				// ModifiedTime:      appSrv["modifiedtime"].(int32),
				// Name: appSrv["name"].(string),
			}
		}
	}

	return applicationServers
}
*/
