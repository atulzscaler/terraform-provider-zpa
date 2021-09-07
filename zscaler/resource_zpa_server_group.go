package zscaler

import (
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/servergroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"applications": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "This field is a json array of app-connector-id only.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"app_connector_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of app-connector IDs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"config_space": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"DEFAULT",
					"SIEM",
				}, false),
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
			"ip_anchored": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dynamic_discovery": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "This field controls dynamic discovery of the servers.",
			},
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
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceServerGroupCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandServerGroup(d)
	log.Printf("[INFO] Creating zpa server group with request\n%+v\n", req)

	resp, _, err := zClient.servergroup.Create(&req)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Created server group request. ID: %v\n", resp)
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
	_ = d.Set("configspace", resp.ConfigSpace)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("ip_anchored", resp.IpAnchored)
	_ = d.Set("dynamic_discovery", resp.DynamicDiscovery)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("name", resp.Name)
	_ = d.Set("appconnector_groups", flattenAppConnectorGroups(resp.AppConnectorGroups))
	_ = d.Set("applications", flattenServerGroupApplications(resp.Applications))
	_ = d.Set("servers", flattenServers(resp.Servers))

	return nil

}

func resourceServerGroupUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Id()
	log.Printf("[INFO] Updating server group ID: %v\n", id)
	req := expandServerGroup(d)

	if _, err := zClient.servergroup.Update(id, &req); err != nil {
		return err
	}
	return resourceServerGroupRead(d, m)
}

func resourceServerGroupDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	log.Printf("[INFO] Deleting server group ID: %v\n", d.Id())

	if _, err := zClient.servergroup.Delete(d.Id()); err != nil {
		return err
	}
	d.SetId("")
	log.Printf("[INFO] server group deleted")
	return nil
}

func expandServerGroup(d *schema.ResourceData) servergroup.ServerGroup {
	req := servergroup.ServerGroup{
		Enabled:            d.Get("enabled").(bool),
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		IpAnchored:         d.Get("ip_anchored").(bool),
		ConfigSpace:        d.Get("config_space").(string),
		DynamicDiscovery:   d.Get("dynamic_discovery").(bool),
		AppConnectorGroups: expandAppConnectorGroups(d.Get("app_connector_groups").([]interface{})),
		Applications:       expandServerGroupApplications(d.Get("applications").([]interface{})),
		Servers:            expandServers(d.Get("servers").([]interface{})),
	}

	return req
}

func expandServerGroupApplications(serverGroupAppRequest []interface{}) []servergroup.Applications {
	serverGroupApplications := make([]servergroup.Applications, len(serverGroupAppRequest))

	for i, serverGroupApplication := range serverGroupAppRequest {
		serverApplicationItem := serverGroupApplication.(map[string]interface{})
		serverGroupApplications[i] = servergroup.Applications{
			ID: serverApplicationItem["id"].(string),
		}
	}

	return serverGroupApplications
}

func expandAppConnectorGroups(appConnectorGroupRequest []interface{}) []servergroup.AppConnectorGroups {
	appConnectorGroups := make([]servergroup.AppConnectorGroups, len(appConnectorGroupRequest))

	for i, appConnectorGroup := range appConnectorGroupRequest {
		appConnectorGroupItem := appConnectorGroup.(map[string]interface{})
		appConnectorGroups[i] = servergroup.AppConnectorGroups{
			ID: appConnectorGroupItem["id"].(string),
		}

	}

	return appConnectorGroups
}

func expandServers(applicationServerRequest []interface{}) []servergroup.ApplicationServer {
	applicationServers := make([]servergroup.ApplicationServer, len(applicationServerRequest))

	for i, applicationServer := range applicationServerRequest {
		applicationServerItem := applicationServer.(map[string]interface{})
		applicationServers[i] = servergroup.ApplicationServer{
			ID: applicationServerItem["id"].(string),
		}
	}

	return applicationServers
}
