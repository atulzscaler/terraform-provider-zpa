package zscaler

import (
	"log"
	"strconv"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/appconnectorgroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAppConnectorGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceConnectorGroupRead,
		Schema: map[string]*schema.Schema{
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
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
			"siemappconnectorgroup": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"upgradeday": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgradetimeinsecs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"versionprofileid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceConnectorGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for app connector group %s\n", id)

	resp, _, err := zClient.appconnectorgroup.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("citycountry", resp.CityCountry)
	_ = d.Set("countrycode", resp.CountryCode)
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("dnsquerytype", resp.DNSQueryType)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("latitude", resp.Latitude)
	_ = d.Set("location", resp.Location)
	_ = d.Set("longitude", resp.Longitude)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("siemappconnectorgroup", resp.SiemAppConnectorGroup)
	_ = d.Set("upgradeday", resp.UpgradeDay)
	_ = d.Set("upgradetimeinsecs", resp.UpgradeTimeInSecs)
	_ = d.Set("versionprofileid", resp.VersionProfileID)
	_ = d.Set("connectors", flattenConnectors(resp))
	_ = d.Set("servergroups", flattenServerGroups(resp))

	return nil

}

func flattenConnectors(appConnector *appconnectorgroup.AppConnectorGroupRequest) []interface{} {
	appConnectors := make([]interface{}, len(appConnector.Connectors))
	for i, appConnector := range appConnector.Connectors {
		appConnectors[i] = map[string]interface{}{
			"creationtime": appConnector.CreationTime,
			"description":  appConnector.Description,
			"enabled":      appConnector.Enabled,
			"fingerprint":  appConnector.Fingerprint,
			"id":           appConnector.ID,
			"issuedcertid": appConnector.IssuedCertID,
			"modifiedby":   appConnector.ModifiedBy,
			"modifiedtime": appConnector.ModifiedTime,
			"name":         appConnector.Name,
		}
	}

	return appConnectors
}

func flattenServerGroups(serverGroup *appconnectorgroup.AppConnectorGroupRequest) []interface{} {
	serverGroups := make([]interface{}, len(serverGroup.AppServerGroup))
	for i, serverGroup := range serverGroup.AppServerGroup {
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
