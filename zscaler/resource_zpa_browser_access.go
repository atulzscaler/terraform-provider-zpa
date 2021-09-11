package zscaler

import (
	"fmt"
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/browseraccess"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/segmentgroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceBrowserAccess() *schema.Resource {
	return &schema.Resource{
		Create: resourceBrowserAccessCreate,
		Read:   resourceBrowserAccessRead,
		Update: resourceBrowserAccessUpdate,
		Delete: resourceBrowserAccessDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"segment_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"segment_group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bypass_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Indicates whether users can bypass ZPA to access applications.",
			},
			"tcp_port_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TCP port ranges used to access the app.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"udp_port_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "UDP port ranges used to access the app.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"config_space": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the application.",
			},
			"domain_names": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of domains and IPs.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"double_encrypt": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether Double Encryption is enabled or disabled for the app.",
			},
			"health_check_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"passive_health_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"health_reporting": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.",
			},
			"ip_anchored": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_cname_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the application.",
			},
			"clientless_apps": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_options": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"app_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"application_port": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"application_protocol": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								"HTTP",
								"HTTPS",
								"FTP",
								"RDP",
							}, false),
						},
						"certificate_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"certificate_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cname": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"hidden": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"local_domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"trust_untrusted_cert": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"server_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of the server group IDs.",
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

func resourceBrowserAccessCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandBrowserAccess(d)
	log.Printf("[INFO] Creating browser access request\n%+v\n", req)

	if req.SegmentGroupId == "" {
		log.Println("[ERROR] Please provde a valid segment group for the application segment")
		return fmt.Errorf("please provde a valid segment group for the application segment")
	}

	browseraccess, _, err := zClient.browseraccess.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created browser access request. ID: %v\n", browseraccess.ID)
	d.SetId(browseraccess.ID)

	return resourceBrowserAccessRead(d, m)
}

func resourceBrowserAccessRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	resp, _, err := zClient.browseraccess.Get(d.Id())
	if err != nil {
		if errResp, ok := err.(*client.ErrorResponse); ok && errResp.IsObjectNotFound() {
			log.Printf("[WARN] Removing browser access %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	log.Printf("[INFO] Getting browser access:\n%+v\n", resp)
	_ = d.Set("segment_group_id", resp.SegmentGroupId)
	_ = d.Set("segment_group_name", resp.SegmentGroupName)
	_ = d.Set("bypass_type", resp.BypassType)
	_ = d.Set("config_space", resp.ConfigSpace)
	_ = d.Set("domain_names", resp.DomainNames)
	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("id", resp.ID)
	_ = d.Set("passive_health_enabled", resp.PassiveHealthEnabled)
	_ = d.Set("double_encrypt", resp.DoubleEncrypt)
	_ = d.Set("health_check_type", resp.HealthCheckType)
	_ = d.Set("is_cname_enabled", resp.IsCnameEnabled)
	_ = d.Set("ip_anchored", resp.IpAnchored)
	_ = d.Set("health_reporting", resp.HealthReporting)
	_ = d.Set("tcp_port_ranges", resp.TcpPortRanges)
	_ = d.Set("udp_port_ranges", resp.UdpPortRanges)

	if err := d.Set("clientless_apps", flattenBaClientlessApps(resp)); err != nil {
		return fmt.Errorf("failed to read clientless apps %s", err)
	}
	if err := d.Set("server_groups", flattenClientlessAppServerGroups(resp.AppServerGroups)); err != nil {
		return fmt.Errorf("failed to read app server groups %s", err)
	}

	return nil

}

func resourceBrowserAccessUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Id()
	log.Printf("[INFO] Updating browser access ID: %v\n", id)
	req := expandBrowserAccess(d)

	if d.HasChange("segment_group_id") && req.SegmentGroupId == "" {
		log.Println("[ERROR] Please provde a valid segment group for the browser access application segment")
		return fmt.Errorf("please provde a valid segment group for the browser access application segment")
	}

	if _, err := zClient.browseraccess.Update(id, &req); err != nil {
		return err
	}

	return resourceBrowserAccessRead(d, m)
}

func resourceBrowserAccessDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	id := d.Id()
	segmentGroupId, ok := d.GetOk("segment_group_id")
	if ok && segmentGroupId != nil {
		gID, ok := segmentGroupId.(string)
		if ok && gID != "" {
			// detach it from segment group first
			if err := detachBrowserAccessFromGroup(zClient, id, gID); err != nil {
				return err
			}
		}
	}
	log.Printf("[INFO] Deleting browser access application with id %v\n", id)
	if _, err := zClient.browseraccess.Delete(id); err != nil {
		return err
	}

	return nil
}

func detachBrowserAccessFromGroup(client *Client, segmentID, segmentGroupId string) error {
	log.Printf("[INFO] Detaching browser access  %s from segment group: %s\n", segmentID, segmentGroupId)
	segGroup, _, err := client.segmentgroup.Get(segmentGroupId)
	if err != nil {
		log.Printf("[error] Error while getting segment group id: %s", segmentGroupId)
		return err
	}
	adaptedApplications := []segmentgroup.Application{}
	for _, app := range segGroup.Applications {
		if app.ID != segmentID {
			adaptedApplications = append(adaptedApplications, app)
		}
	}
	segGroup.Applications = adaptedApplications
	_, err = client.segmentgroup.Update(segmentGroupId, segGroup)
	return err

}

func expandBrowserAccess(d *schema.ResourceData) browseraccess.BrowserAccess {
	details := browseraccess.BrowserAccess{
		BypassType:      d.Get("bypass_type").(string),
		Description:     d.Get("description").(string),
		DoubleEncrypt:   d.Get("double_encrypt").(bool),
		Enabled:         d.Get("enabled").(bool),
		HealthReporting: d.Get("health_reporting").(string),
		IpAnchored:      d.Get("ip_anchored").(bool),
		IsCnameEnabled:  d.Get("is_cname_enabled").(bool),
		DomainNames:     expandStringInSlice(d, "domain_names"),
		SegmentGroupId:  d.Get("segment_group_id").(string),
	}
	if d.HasChange("name") {
		details.Name = d.Get("name").(string)
	}
	if d.HasChange("segment_group_name") {
		details.SegmentGroupName = d.Get("segment_group_name").(string)
	}
	if d.HasChange("server_groups") {
		details.AppServerGroups = expandClientlessAppServerGroups(d)
	}
	if d.HasChange("clientless_apps") {
		details.ClientlessApps = expandClientlessApps(d)
	}
	if d.HasChange("udp_port_ranges") {
		details.UdpPortRanges = convertToListString(d.Get("udp_port_ranges"))
	}
	if d.HasChange("tcp_port_ranges") {
		details.TcpPortRanges = convertToListString(d.Get("tcp_port_ranges"))
	}
	return details
}
func convertToListString(obj interface{}) []string {
	listI, ok := obj.([]interface{})
	if ok && len(listI) > 0 {
		list := make([]string, len(listI))
		for i, e := range listI {
			s, ok := e.(string)
			if ok {
				list[i] = e.(string)
			} else {
				log.Printf("[WARN] invalid type: %v\n", s)
			}
		}
		return list
	}
	return []string{}
}
func expandClientlessApps(d *schema.ResourceData) []browseraccess.ClientlessApps {
	clientlessInterface, ok := d.GetOk("clientless_apps")
	if ok {
		clientless := clientlessInterface.([]interface{})
		log.Printf("[INFO] clientless apps data: %+v\n", clientless)
		var clientlessApps []browseraccess.ClientlessApps
		for _, clientlessApp := range clientless {
			clientlessApp, ok := clientlessApp.(map[string]interface{})
			if ok {
				clientlessApps = append(clientlessApps, browseraccess.ClientlessApps{
					AllowOptions:        clientlessApp["allow_options"].(bool),
					AppId:               clientlessApp["app_id"].(string),
					ApplicationPort:     clientlessApp["application_port"].(string),
					ApplicationProtocol: clientlessApp["application_protocol"].(string),
					CertificateId:       clientlessApp["certificate_id"].(string),
					CertificateName:     clientlessApp["certificate_name"].(string),
					Cname:               clientlessApp["cname"].(string),
					Description:         clientlessApp["description"].(string),
					Domain:              clientlessApp["domain"].(string),
					Enabled:             clientlessApp["enabled"].(bool),
					Hidden:              clientlessApp["hidden"].(bool),
					LocalDomain:         clientlessApp["local_domain"].(string),
					Name:                clientlessApp["name"].(string),
					Path:                clientlessApp["path"].(string),
					TrustUntrustedCert:  clientlessApp["trust_untrusted_cert"].(bool),
				})
			}
		}
		return clientlessApps
	}

	return []browseraccess.ClientlessApps{}
}

func expandClientlessAppServerGroups(d *schema.ResourceData) []browseraccess.AppServerGroups {
	serverGroupsInterface, ok := d.GetOk("server_groups")
	if ok {
		serverGroup := serverGroupsInterface.(*schema.Set)
		log.Printf("[INFO] app server groups data: %+v\n", serverGroup)
		var serverGroups []browseraccess.AppServerGroups
		for _, appServerGroup := range serverGroup.List() {
			appServerGroup, _ := appServerGroup.(map[string]interface{})
			if appServerGroup != nil {
				for _, id := range appServerGroup["id"].([]interface{}) {
					serverGroups = append(serverGroups, browseraccess.AppServerGroups{
						ID: id.(string),
					})
				}
			}
		}
		return serverGroups
	}

	return []browseraccess.AppServerGroups{}
}

func flattenBaClientlessApps(clientlessApp *browseraccess.BrowserAccess) []interface{} {
	clientlessApps := make([]interface{}, len(clientlessApp.ClientlessApps))
	for i, clientlessApp := range clientlessApp.ClientlessApps {
		clientlessApps[i] = map[string]interface{}{
			"allow_options":        clientlessApp.AllowOptions,
			"app_id":               clientlessApp.AppId,
			"application_port":     clientlessApp.ApplicationPort,
			"application_protocol": clientlessApp.ApplicationProtocol,
			"certificate_id":       clientlessApp.CertificateId,
			"certificate_name":     clientlessApp.CertificateName,
			"cname":                clientlessApp.Cname,
			"description":          clientlessApp.Description,
			"domain":               clientlessApp.Domain,
			"enabled":              clientlessApp.Enabled,
			"hidden":               clientlessApp.Hidden,
			"id":                   clientlessApp.ID,
			"local_domain":         clientlessApp.LocalDomain,
			"name":                 clientlessApp.Name,
			"path":                 clientlessApp.Path,
			"trust_untrusted_cert": clientlessApp.TrustUntrustedCert,
		}
	}

	return clientlessApps
}

func flattenClientlessAppServerGroups(appServerGroup []browseraccess.AppServerGroups) []interface{} {
	result := make([]interface{}, 1)
	mapIds := make(map[string]interface{})
	ids := make([]string, len(appServerGroup))
	for i, serverGroup := range appServerGroup {
		ids[i] = serverGroup.ID
	}
	mapIds["id"] = ids
	result[0] = mapIds
	return result
}
