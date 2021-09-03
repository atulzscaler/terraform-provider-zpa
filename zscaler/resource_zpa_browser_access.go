package zscaler

import (
	"fmt"
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/browseraccess"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
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
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"udp_port_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "UDP port ranges used to access the app.",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"configspace": {
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
						// "appid": {
						// 	Type:     schema.TypeInt,
						// 	Computed: true,
						// },
						"application_port": {
							Type:     schema.TypeInt,
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
							Type:     schema.TypeInt,
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
							Type:     schema.TypeInt,
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
			// Server Group only takes one ID as int64
			"server_groups": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "ID of the server group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
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

	resp, _, err := zClient.browseraccess.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created application segment request. ID: %v\n", resp.ID)
	d.SetId(resp.ID)
	//d.SetId(strconv.FormatInt(int64(applicationsegment.ID), 10))
	//d.SetId(strconv.Itoa(resp.ID))

	return resourceBrowserAccessRead(d, m)
}

func resourceBrowserAccessRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	resp, _, err := zClient.browseraccess.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing segment group %s from state because it no longer exists in ZPA", d.Id())
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
	// _ = d.Set("creationtime", resp.CreationTime)
	// _ = d.Set("modifiedby", resp.ModifiedBy)
	// _ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("passive_health_enabled", resp.PassiveHealthEnabled)
	_ = d.Set("double_encrypt", resp.DoubleEncrypt)
	_ = d.Set("health_check_type", resp.HealthCheckType)
	_ = d.Set("is_cname_enabled", resp.IsCnameEnabled)
	_ = d.Set("ip_anchored", resp.IpAnchored)
	_ = d.Set("health_reporting", resp.HealthReporting)
	_ = d.Set("tcp_port_ranges", resp.TcpPortRanges)
	_ = d.Set("udp_port_ranges", resp.UdpPortRanges)
	// _ = d.Set("clientlessapps", resp.ClientlessApps)

	if err := d.Set("clientless_apps", flattenBaClientlessApps(resp)); err != nil {
		return fmt.Errorf("failed to read clientless apps %s", err)
	}
	if err := d.Set("server_groups", flattenClientlessAppServerGroups(resp)); err != nil {
		return fmt.Errorf("failed to read app server groups %s", err)
	}

	return nil

}

func resourceBrowserAccessUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Id()
	log.Printf("[INFO] Updating browser access ID: %v\n", id)
	req := expandBrowserAccess(d)

	if _, err := zClient.browseraccess.Update(id, req); err != nil {
		return err
	}

	return resourceBrowserAccessRead(d, m)
}

func resourceBrowserAccessDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	log.Printf("[INFO] Deleting browser access application with id %v\n", d.Id())
	if _, err := zClient.browseraccess.Delete(d.Id()); err != nil {
		return err
	}

	d.SetId("")
	log.Printf("[INFO] browser access application deleted")

	return nil
}

func expandBrowserAccess(d *schema.ResourceData) browseraccess.BrowserAccess {
	return browseraccess.BrowserAccess{
		SegmentGroupId:   d.Get("segment_group_id").(string),
		SegmentGroupName: d.Get("segment_group_name").(string),
		BypassType:       d.Get("bypass_type").(string),
		Description:      d.Get("description").(string),
		DomainNames:      expandStringInSlice(d, "domain_names"),
		DoubleEncrypt:    d.Get("double_encrypt").(bool),
		Enabled:          d.Get("enabled").(bool),
		HealthReporting:  d.Get("health_reporting").(string),
		IpAnchored:       d.Get("ip_anchored").(bool),
		IsCnameEnabled:   d.Get("is_cname_enabled").(bool),
		Name:             d.Get("name").(string),
		TcpPortRanges:    d.Get("tcp_port_ranges").([]interface{}),
		UdpPortRanges:    d.Get("udp_port_ranges").([]interface{}),
		ClientlessApps:   expandClientlessApps(d),
		AppServerGroups:  expandClientlessAppServerGroups(d),
	}
}

func expandClientlessApps(d *schema.ResourceData) []browseraccess.ClientlessApps {
	var clientlessApps []browseraccess.ClientlessApps
	if clientlessInterface, ok := d.GetOk("clientless_apps"); ok {
		clientless := clientlessInterface.([]interface{})
		clientlessApps = make([]browseraccess.ClientlessApps, len(clientless))
		for i, app := range clientless {
			clientlessApp := app.(map[string]interface{})
			clientlessApps[i] = browseraccess.ClientlessApps{
				AllowOptions: clientlessApp["allow_options"].(bool),
				// AppId:               clientlessApp["app_id"].(int64),
				ApplicationPort:     clientlessApp["application_port"].(int),
				ApplicationProtocol: clientlessApp["application_protocol"].(string),
				CertificateId:       clientlessApp["certificate_id"].(int),
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
			}
		}
	}

	return clientlessApps
}

func expandClientlessAppServerGroups(d *schema.ResourceData) []browseraccess.AppServerGroups {
	var serverGroups []browseraccess.AppServerGroups
	if serverGroupInterface, ok := d.GetOk("server_groups"); ok {
		servers := serverGroupInterface.([]interface{})
		serverGroups = make([]browseraccess.AppServerGroups, len(servers))
		for i, srvGroup := range servers {
			serverGroup := srvGroup.(map[string]interface{})
			serverGroups[i] = browseraccess.AppServerGroups{
				// Name: serverGroup["name"].(string),
				ID: serverGroup["id"].(int),
				// ConfigSpace:      serverGroup["config_space"].(string),
				// CreationTime:     serverGroup["creation_time"].(int32),
				// Description:      serverGroup["description"].(string),
				// Enabled:          serverGroup["enabled"].(bool),
				// DynamicDiscovery: serverGroup["dynamic_discovery"].(bool),
				// ModifiedBy:       serverGroup["modifiedby"].(int64),
				// ModifiedTime:     serverGroup["modified_time"].(int32),
			}
		}
	}

	return serverGroups
}

func flattenBaClientlessApps(clientlessApp *browseraccess.BrowserAccess) []interface{} {
	clientlessApps := make([]interface{}, len(clientlessApp.ClientlessApps))
	for i, clientlessApp := range clientlessApp.ClientlessApps {
		clientlessApps[i] = map[string]interface{}{
			"allow_options": clientlessApp.AllowOptions,
			// "app_id":               clientlessApp.AppId,
			"application_port":     clientlessApp.ApplicationPort,
			"application_protocol": clientlessApp.ApplicationProtocol,
			"certificate_id":       clientlessApp.CertificateId,
			"certificate_name":     clientlessApp.CertificateName,
			"cname":                clientlessApp.Cname,
			//"creation_time":        clientlessApp.CreationTime,
			"description":  clientlessApp.Description,
			"domain":       clientlessApp.Domain,
			"enabled":      clientlessApp.Enabled,
			"hidden":       clientlessApp.Hidden,
			"id":           clientlessApp.ID,
			"local_domain": clientlessApp.LocalDomain,
			//"modifiedby":          clientlessApp.ModifiedBy,
			//"modified_time":        clientlessApp.ModifiedTime,
			"name":                 clientlessApp.Name,
			"path":                 clientlessApp.Path,
			"trust_untrusted_cert": clientlessApp.TrustUntrustedCert,
		}
	}

	return clientlessApps
}

func flattenClientlessAppServerGroups(serverGroup *browseraccess.BrowserAccess) []interface{} {
	serverGroups := make([]interface{}, len(serverGroup.AppServerGroups))
	for i, val := range serverGroup.AppServerGroups {
		serverGroups[i] = map[string]interface{}{
			// "name":             val.Name,
			"id": val.ID,
			// "config_space":      val.ConfigSpace,
			// "creation_time":     val.CreationTime,
			// "description":      val.Description,
			// "enabled":          val.Enabled,
			// "dynamic_discovery": val.DynamicDiscovery,
			// "modifiedby":       val.ModifiedBy,
			// "modified_time":     val.ModifiedTime,
		}
	}

	return serverGroups
}
