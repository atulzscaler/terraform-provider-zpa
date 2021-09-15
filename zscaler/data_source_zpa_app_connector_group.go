package zscaler

import (
	"fmt"
	"log"

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
						"application_start_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"appconnector_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"appconnector_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"control_channel_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creation_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ctrl_broker_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"current_version": {
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
						"expected_upgrade_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"expected_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipacl": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issued_cert_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_broker_connect_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_broker_disconnect_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_upgrade_time": {
							Type:     schema.TypeString,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"modified_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"previous_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"upgrade_attempt": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"signing_cert": {
							Type:     schema.TypeMap,
							Elem:     schema.TypeString,
							Computed: true,
						},
						"upgrade_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"city_country": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"country_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_query_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"geolocation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_space": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creation_time": {
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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dynamic_discovery": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"modifiedby": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"modified_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"siem_appconnector_group": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"upgrade_day": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_time_in_secs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version_profile_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceConnectorGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	var resp *appconnectorgroup.AppConnectorGroup
	id, ok := d.Get("id").(string)
	if ok && id != "" {
		log.Printf("[INFO] Getting data for app connector group  %s\n", id)
		res, _, err := zClient.appconnectorgroup.Get(id)
		if err != nil {
			return err
		}
		resp = res
	}
	name, ok := d.Get("name").(string)
	if ok && name != "" {
		log.Printf("[INFO] Getting data for app connector group name %s\n", name)
		res, _, err := zClient.appconnectorgroup.GetByName(name)
		if err != nil {
			return err
		}
		resp = res
	}
	if resp != nil {
		d.SetId(resp.ID)
		_ = d.Set("city_country", resp.CityCountry)
		_ = d.Set("country_code", resp.CountryCode)
		_ = d.Set("creation_time", resp.CreationTime)
		_ = d.Set("description", resp.Description)
		_ = d.Set("dns_query_type", resp.DNSQueryType)
		_ = d.Set("enabled", resp.Enabled)
		_ = d.Set("latitude", resp.Latitude)
		_ = d.Set("location", resp.Location)
		_ = d.Set("longitude", resp.Longitude)
		_ = d.Set("modifiedby", resp.ModifiedBy)
		_ = d.Set("modified_time", resp.ModifiedTime)
		_ = d.Set("name", resp.Name)
		_ = d.Set("siem_appconnector_group", resp.SiemAppConnectorGroup)
		_ = d.Set("upgrade_day", resp.UpgradeDay)
		_ = d.Set("upgrade_time_in_secs", resp.UpgradeTimeInSecs)
		_ = d.Set("version_profile_id", resp.VersionProfileID)
		_ = d.Set("connectors", flattenConnectors(resp))

		if err := d.Set("server_groups", flattenServerGroups(resp)); err != nil {
			return fmt.Errorf("failed to read server groups %s", err)
		}
	} else {
		return fmt.Errorf("couldn't find any app connector group with name '%s' or id '%s'", name, id)
	}

	return nil
}

func flattenConnectors(appConnector *appconnectorgroup.AppConnectorGroup) []interface{} {
	appConnectors := make([]interface{}, len(appConnector.Connectors))
	for i, appConnector := range appConnector.Connectors {
		appConnectors[i] = map[string]interface{}{
			"application_start_time":      appConnector.ApplicationStartTime,
			"appconnector_group_id":       appConnector.AppConnectorGroupID,
			"control_channel_status":      appConnector.ControlChannelStatus,
			"creation_time":               appConnector.CreationTime,
			"ctrl_broker_name":            appConnector.CtrlBrokerName,
			"current_version":             appConnector.CurrentVersion,
			"description":                 appConnector.Description,
			"enabled":                     appConnector.Enabled,
			"expected_upgrade_time":       appConnector.ExpectedUpgradeTime,
			"expected_version":            appConnector.ExpectedVersion,
			"fingerprint":                 appConnector.Fingerprint,
			"id":                          appConnector.ID,
			"ipacl":                       appConnector.IpAcl,
			"issued_cert_id":              appConnector.IssuedCertID,
			"last_broker_connect_time":    appConnector.LastBrokerConnectTime,
			"last_broker_disconnect_time": appConnector.LastBrokerDisconnectTime,
			"last_upgrade_time":           appConnector.LastUpgradeTime,
			"latitude":                    appConnector.Latitude,
			"location":                    appConnector.Location,
			"longitude":                   appConnector.Longitude,
			"modifiedby":                  appConnector.ModifiedBy,
			"modified_time":               appConnector.ModifiedTime,
			"name":                        appConnector.Name,
			"platform":                    appConnector.Platform,
			"previous_version":            appConnector.PreviousVersion,
			"private_ip":                  appConnector.PrivateIP,
			"public_ip":                   appConnector.PublicIP,
			"signing_cert":                appConnector.SigningCert,
			"upgrade_attempt":             appConnector.UpgradeAttempt,
			"upgrade_status":              appConnector.UpgradeStatus,
		}
	}

	return appConnectors
}

func flattenServerGroups(serverGroup *appconnectorgroup.AppConnectorGroup) []interface{} {
	serverGroups := make([]interface{}, len(serverGroup.AppServerGroup))
	for i, serverGroup := range serverGroup.AppServerGroup {
		serverGroups[i] = map[string]interface{}{
			"config_space":      serverGroup.ConfigSpace,
			"creation_time":     serverGroup.CreationTime,
			"description":       serverGroup.Description,
			"enabled":           serverGroup.Enabled,
			"id":                serverGroup.ID,
			"dynamic_discovery": serverGroup.DynamicDiscovery,
			"modifiedby":        serverGroup.ModifiedBy,
			"modified_time":     serverGroup.ModifiedTime,
			"name":              serverGroup.Name,
		}
	}

	return serverGroups
}
