package zscaler

import (
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/applicationsegment"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationSegment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplicationSegmentRead,
		Schema: map[string]*schema.Schema{
			"segmentgroupid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"segmentgroupname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bypasstype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"clientlessapps": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowoptions": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"appid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"applicationport": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"applicationprotocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificateid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"certificatename": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cname": {
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
						"domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"hidden": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"localdomain": {
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
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trustuntrustedcert": {
							Type:     schema.TypeBool,
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
			"defaultidletimeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"defaultmaxage": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domainnames": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"doubleencrypt": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"healthchecktype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"healthreporting": {
				Type:     schema.TypeString,
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
			"iscnameenabled": {
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
				Computed: true,
			},
			"passivehealthenabled": {
				Type:     schema.TypeBool,
				Computed: true,
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
						"dynamicdiscovery": {
							Type:     schema.TypeBool,
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
			"tcpportranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"udpportranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceApplicationSegmentRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for server group %s\n", id)

	resp, _, err := zClient.applicationsegment.Get(id)
	if err != nil {
		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("segmentgroupid", resp.SegmentGroupId)
	_ = d.Set("segmentgroupname", resp.SegmentGroupName)
	_ = d.Set("bypasstype", resp.BypassType)
	_ = d.Set("configSpace", resp.ConfigSpace)
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("domainnames", resp.DomainNames)
	_ = d.Set("doubleencrypt", resp.DoubleEncrypt)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("healthchecktype", resp.HealthCheckType)
	_ = d.Set("healthreporting", resp.HealthReporting)
	_ = d.Set("ipanchored", resp.IpAnchored)
	_ = d.Set("iscnameenabled", resp.IsCnameEnabled)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("passivehealthenabled", resp.PassiveHealthEnabled)
	_ = d.Set("tcpportranges", resp.TcpPortRanges)
	_ = d.Set("udpportranges", resp.UdpPortRanges)

	if err := d.Set("clientlessapps", flattenClientlessApps(resp)); err != nil {
		return err
	}
	if err := d.Set("servergroups", flattenAppServerGroups(resp)); err != nil {
		return err
	}

	return nil

}

func flattenClientlessApps(clientlessApp *applicationsegment.ApplicationSegmentResource) []interface{} {
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

func flattenAppServerGroups(serverGroup *applicationsegment.ApplicationSegmentResource) []interface{} {
	serverGroups := make([]interface{}, len(serverGroup.ServerGroups))
	for i, val := range serverGroup.ServerGroups {
		serverGroups[i] = map[string]interface{}{
			"name":             val.Name,
			"id":               val.ID,
			"configspace":      val.ConfigSpace,
			"creationtime":     val.CreationTime,
			"description":      val.Description,
			"enabled":          val.Enabled,
			"dynamicdiscovery": val.DynamicDiscovery,
			"modifiedby":       val.ModifiedBy,
			"modifiedtime":     val.ModifiedTime,
		}
	}

	return serverGroups
}
