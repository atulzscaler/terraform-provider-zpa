package zscaler

import (
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/browseraccess"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"segmentgroupid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"segmentgroupname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bypasstype": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Indicates whether users can bypass ZPA to access applications.",
			},
			"tcpportranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TCP port ranges used to access the app.",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"udpportranges": {
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
			"domainnames": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of domains and IPs.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"doubleencrypt": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether Double Encryption is enabled or disabled for the app.",
			},
			"healthchecktype": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"healthreporting": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.",
			},
			"ipanchored": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"iscnameenabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the application.",
			},
			"clientlessapps": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowoptions": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						// "appid": {
						// 	Type:     schema.TypeInt,
						// 	Computed: true,
						// },
						"applicationport": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"applicationprotocol": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"certificateid": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"certificatename": {
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
						"localdomain": {
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
						"trustuntrustedcert": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			// Server Group only takes one ID as int64
			"servergroups": {
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

	if zClient == nil {
		return resourceNotSupportedError()
	}

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
	_ = d.Set("segmentgroupid", resp.SegmentGroupId)
	_ = d.Set("segmentgroupname", resp.SegmentGroupName)
	_ = d.Set("bypasstype", resp.BypassType)
	_ = d.Set("configspace", resp.ConfigSpace)
	_ = d.Set("domainnames", resp.DomainNames)
	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	// _ = d.Set("creationtime", resp.CreationTime)
	// _ = d.Set("modifiedby", resp.ModifiedBy)
	// _ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("passivehealthenabled", resp.PassiveHealthEnabled)
	_ = d.Set("doubleencrypt", resp.DoubleEncrypt)
	_ = d.Set("healthchecktype", resp.HealthCheckType)
	_ = d.Set("iscnameenabled", resp.IsCnameEnabled)
	_ = d.Set("ipanchored", resp.IpAnchored)
	_ = d.Set("healthreporting", resp.HealthReporting)
	_ = d.Set("tcpportranges", resp.TcpPortRanges)
	_ = d.Set("udpportranges", resp.UdpPortRanges)
	// _ = d.Set("clientlessapps", resp.ClientlessApps)

	if err := d.Set("clientlessapps", flattenBaClientlessApps(resp)); err != nil {
		return err
	}
	if err := d.Set("servergroups", flattenClientlessAppServerGroups(resp)); err != nil {
		return err
	}

	return nil

}

func resourceBrowserAccessUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	if zClient == nil {
		return resourceNotSupportedError()
	}

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

	return nil
}

func expandBrowserAccess(d *schema.ResourceData) browseraccess.BrowserAccess {
	return browseraccess.BrowserAccess{
		SegmentGroupId:   d.Get("segmentgroupid").(string),
		SegmentGroupName: d.Get("segmentgroupname").(string),
		BypassType:       d.Get("bypasstype").(string),
		Description:      d.Get("description").(string),
		DomainNames:      expandStringInSlice(d, "domainnames"),
		DoubleEncrypt:    d.Get("doubleencrypt").(bool),
		Enabled:          d.Get("enabled").(bool),
		HealthReporting:  d.Get("healthreporting").(string),
		IpAnchored:       d.Get("ipanchored").(bool),
		IsCnameEnabled:   d.Get("iscnameenabled").(bool),
		Name:             d.Get("name").(string),
		TcpPortRanges:    d.Get("tcpportranges").([]interface{}),
		UdpPortRanges:    d.Get("udpportranges").([]interface{}),
		ClientlessApps:   expandClientlessApps(d),
		AppServerGroups:  expandClientlessAppServerGroups(d),
	}
}

func expandClientlessApps(d *schema.ResourceData) []browseraccess.ClientlessApps {
	var clientlessApps []browseraccess.ClientlessApps
	if clientlessInterface, ok := d.GetOk("clientlessapps"); ok {
		clientless := clientlessInterface.([]interface{})
		clientlessApps = make([]browseraccess.ClientlessApps, len(clientless))
		for i, app := range clientless {
			clientlessApp := app.(map[string]interface{})
			clientlessApps[i] = browseraccess.ClientlessApps{
				AllowOptions: clientlessApp["allowoptions"].(bool),
				//AppId:               clientlessApp["appid"].(int),
				ApplicationPort:     clientlessApp["applicationport"].(int),
				ApplicationProtocol: clientlessApp["applicationprotocol"].(string),
				CertificateId:       clientlessApp["certificateid"].(int),
				CertificateName:     clientlessApp["certificatename"].(string),
				Cname:               clientlessApp["cname"].(string),
				Description:         clientlessApp["description"].(string),
				Domain:              clientlessApp["domain"].(string),
				Enabled:             clientlessApp["enabled"].(bool),
				Hidden:              clientlessApp["hidden"].(bool),
				LocalDomain:         clientlessApp["localdomain"].(string),
				Name:                clientlessApp["name"].(string),
				Path:                clientlessApp["path"].(string),
				TrustUntrustedCert:  clientlessApp["trustuntrustedcert"].(bool),
			}
		}
	}

	return clientlessApps
}

func expandClientlessAppServerGroups(d *schema.ResourceData) []browseraccess.AppServerGroups {
	var serverGroups []browseraccess.AppServerGroups
	if serverGroupInterface, ok := d.GetOk("servergroups"); ok {
		servers := serverGroupInterface.([]interface{})
		serverGroups = make([]browseraccess.AppServerGroups, len(servers))
		for i, srvGroup := range servers {
			serverGroup := srvGroup.(map[string]interface{})
			serverGroups[i] = browseraccess.AppServerGroups{
				// Name: serverGroup["name"].(string),
				ID: serverGroup["id"].(int),
				// ConfigSpace:      serverGroup["configspace"].(string),
				// CreationTime:     serverGroup["creationtime"].(int32),
				// Description:      serverGroup["description"].(string),
				// Enabled:          serverGroup["enabled"].(bool),
				// DynamicDiscovery: serverGroup["dynamicdiscovery"].(bool),
				// ModifiedBy:       serverGroup["modifiedby"].(int64),
				// ModifiedTime:     serverGroup["modifiedtime"].(int32),
			}
		}
	}

	return serverGroups
}

func flattenBaClientlessApps(clientlessApp *browseraccess.BrowserAccess) []interface{} {
	clientlessApps := make([]interface{}, len(clientlessApp.ClientlessApps))
	for i, clientlessApp := range clientlessApp.ClientlessApps {
		clientlessApps[i] = map[string]interface{}{
			"allowoptions": clientlessApp.AllowOptions,
			//"appid":               clientlessApp.AppId,
			"applicationport":     clientlessApp.ApplicationPort,
			"applicationprotocol": clientlessApp.ApplicationProtocol,
			"certificateid":       clientlessApp.CertificateId,
			"certificatename":     clientlessApp.CertificateName,
			"cname":               clientlessApp.Cname,
			//"creationtime":        clientlessApp.CreationTime,
			"description": clientlessApp.Description,
			"domain":      clientlessApp.Domain,
			"enabled":     clientlessApp.Enabled,
			"hidden":      clientlessApp.Hidden,
			"id":          clientlessApp.ID,
			"localdomain": clientlessApp.LocalDomain,
			//"modifiedby":          clientlessApp.ModifiedBy,
			//"modifiedtime":        clientlessApp.ModifiedTime,
			"name":               clientlessApp.Name,
			"path":               clientlessApp.Path,
			"trustuntrustedcert": clientlessApp.TrustUntrustedCert,
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
			// "configspace":      val.ConfigSpace,
			// "creationtime":     val.CreationTime,
			// "description":      val.Description,
			// "enabled":          val.Enabled,
			// "dynamicdiscovery": val.DynamicDiscovery,
			// "modifiedby":       val.ModifiedBy,
			// "modifiedtime":     val.ModifiedTime,
		}
	}

	return serverGroups
}
