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
						"applicationstarttime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"appconnectorgroupid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"appconnectorgroupname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"controlchannelstatus": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creationtime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ctrlbrokername": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"currentversion": {
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
						"expectedupgradetime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"expectedversion": {
							Type:     schema.TypeString,
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
						"ipacl": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issuedcertid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"lastbrokerconnecttime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"lastbrokerdisconnecttime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"lastupgradetime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"latitude": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"longitude": {
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
						"platform": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"previousversion": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"privateip": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"publicip": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"upgradeattempt": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"signingcert": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

	//d.SetId(strconv.Itoa(resp.ID))
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
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
	//_ = d.Set("connectors", flattenConnectors(resp))

	if err := d.Set("connectors", flattenConnectors(resp)); err != nil {
		return err
	}
	if err := d.Set("servergroups", flattenServerGroups(resp)); err != nil {
		return err
	}

	return nil

}

func flattenConnectors(appConnector *appconnectorgroup.AppConnectorGroup) []interface{} {
	appConnectors := make([]interface{}, len(appConnector.Connectors))
	for i, appConnector := range appConnector.Connectors {
		appConnectors[i] = map[string]interface{}{
			"applicationstarttime":     appConnector.ApplicationStartTime,
			"appconnectorgroupid":      appConnector.AppConnectorGroupID,
			"controlchannelstatus":     appConnector.ControlChannelStatus,
			"creationtime":             appConnector.CreationTime,
			"ctrlbrokername":           appConnector.CtrlBrokerName,
			"currentversion":           appConnector.CurrentVersion,
			"description":              appConnector.Description,
			"enabled":                  appConnector.Enabled,
			"expectedupgradetime":      appConnector.ExpectedUpgradeTime,
			"expectedversion":          appConnector.ExpectedVersion,
			"fingerprint":              appConnector.Fingerprint,
			"id":                       appConnector.ID,
			"ipacl":                    appConnector.IpAcl,
			"issuedcertid":             appConnector.IssuedCertID,
			"lastbrokerconnecttime":    appConnector.LastBrokerConnectTime,
			"lastbrokerdisconnecttime": appConnector.LastBrokerDisconnectTime,
			"lastupgradetime":          appConnector.LastUpgradeTime,
			"latitude":                 appConnector.Latitude,
			"location":                 appConnector.Location,
			"longitude":                appConnector.Longitude,
			"modifiedby":               appConnector.ModifiedBy,
			"modifiedtime":             appConnector.ModifiedTime,
			"name":                     appConnector.Name,
			"platform":                 appConnector.Platform,
			"previousversion":          appConnector.PreviousVersion,
			"privateip":                appConnector.PrivateIp,
			"publicip":                 appConnector.PublicIp,
			"signingcert":              appConnector.SigningCert,
			"upgradeattempt":           appConnector.UpgradeAttempt,
			"upgradestatus":            appConnector.UpgradeStatus,
		}
	}

	return appConnectors
}

func flattenServerGroups(serverGroup *appconnectorgroup.AppConnectorGroup) []interface{} {
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
