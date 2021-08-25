package zscaler

/*
import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
)

func resourceBrowserAccess() *schema.Resource {
	return &schema.Resource{
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
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "TCP port ranges used to access the app.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"udpportranges": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "UDP port ranges used to access the app.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
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
				Description: "Name of the application.",
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
			"clientlessapps": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowoptions": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"applicationid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
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
						"creationtime": {
							Type:     schema.TypeInt,
							Computed: true,
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
						"localdomain": {
							Type:     schema.TypeString,
							Optional: true,
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
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trustuntrustedcert": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
					Create: resourceBrowserAccessCreate,
					Read:   resourceBrowserAccessRead,
					Update: resourceBrowserAccessUpdate,
					Delete: resourceBrowserAccessDelete,
					Importer: &schema.ResourceImporter{
						State: schema.ImportStatePassthrough,
					},
				},
			},
		},
	}
}

func resourceBrowserAccessCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	//req := expandApplicationSegmentRequest(d)
	log.Printf("[INFO] Creating application segment request\n%+v\n", req)

	resp, _, err := zClient.applicationsegment.Create(req)
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
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("passivehealthenabled", resp.PassiveHealthEnabled)
	_ = d.Set("doubleencrypt", resp.DoubleEncrypt)
	_ = d.Set("healthchecktype", resp.HealthCheckType)
	_ = d.Set("iscnameenabled", resp.IsCnameEnabled)
	_ = d.Set("ipanchored", resp.IpAnchored)
	_ = d.Set("healthreporting", resp.HealthReporting)
	_ = d.Set("tcpportranges", resp.TcpPortRanges)
	_ = d.Set("udpportranges", resp.UdpPortRanges)
	_ = d.Set("clientlessapps", resp.ClientlessApps)

	return nil

}

func resourceBrowserAccessUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	return resourceBrowserAccessRead(d, m)
}

func resourceBrowserAccessDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

}
*/
/*
func expandClientlessApps(d *schema.ResourceData) []applicationsegment.ClientlessApps {
    var clientlessApps []applicationsegment.ClientlessApps
    if clientlessInterface, ok := d.GetOk("clientlessapps"); ok {
        clientless := clientlessInterface.([]interface{})
        clientlessApps = make([]applicationsegment.ClientlessApps, len(clientless))
        for i, app := range clientless {
            clientlessApp := app.(map[string]interface{})
            clientlessApps[i] = applicationsegment.ClientlessApps{
                AllowOptions:        clientlessApp["allowoptions"].(bool),
                AppId:               clientlessApp["appid"].(int),
                ApplicationPort:     clientlessApp["applicationport"].(int),
                ApplicationProtocol: clientlessApp["applicationprotocol"].(string), // │ Error: clientlessapps.0.applicationprotocol: '': source data must be an array or slice, got string
                CertificateId:       clientlessApp["certificateid"].(int),
                CertificateName:     clientlessApp["certificatename"].(string),
                Cname:               clientlessApp["cname"].(string),
                // CreationTime:        clientlessApp["creationtime"].(int32),
                Description: clientlessApp["description"].(string),
                Domain:      clientlessApp["domain"].(string),
                Enabled:     clientlessApp["enabled"].(bool),
                Hidden:      clientlessApp["hidden"].(bool),
                // ID:                  clientlessApp["id"].(int64),
                LocalDomain: clientlessApp["localdomain"].(string),
                // ModifiedBy:          clientlessApp["modifiedby"].(int64),
                // ModifiedTime:        clientlessApp["modifiedtime"].(int32),
                Name: clientlessApp["name"].(string),
                // Path:               clientlessApp["path"].(string),
                TrustUntrustedCert: clientlessApp["trustuntrustedcert"].(bool),
            }
        }
    }

    return clientlessApps
}

func flattenClientlessApps(clientlessApp *applicationsegment.ApplicationSegmentResponse) []interface{} {
    clientlessApps := make([]interface{}, len(clientlessApp.ClientlessApps))
    for i, clientlessApp := range clientlessApp.ClientlessApps {
        clientlessApps[i] = map[string]interface{}{
            "allowoptions":        clientlessApp.AllowOptions,
            "appid":               clientlessApp.AppId,
            "applicationport":     clientlessApp.ApplicationPort,
            "applicationprotocol": clientlessApp.ApplicationProtocol,
            "certificateid":       clientlessApp.CertificateId,
            "certificatename":     clientlessApp.CertificateName,
            "cname":               clientlessApp.Cname,
            "creationtime":        clientlessApp.CreationTime,
            "description":         clientlessApp.Description,
            "domain":              clientlessApp.Domain,
            "enabled":             clientlessApp.Enabled,
            "hidden":              clientlessApp.Hidden,
            "id":                  clientlessApp.ID,
            "localdomain":         clientlessApp.LocalDomain,
            "modifiedby":          clientlessApp.ModifiedBy,
            "modifiedtime":        clientlessApp.ModifiedTime,
            "name":                clientlessApp.Name,
            "path":                clientlessApp.Path,
            "trustuntrustedcert":  clientlessApp.TrustUntrustedCert,
        }
    }

    return clientlessApps
}
*/
