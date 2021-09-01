package zscaler

import (
	"fmt"
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/applicationsegment"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationSegment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplicationSegmentRead,
		Schema: map[string]*schema.Schema{
			"segment_group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"segment_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bypass_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"clientless_apps": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_options": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"appid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"application_port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"application_protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"certificate_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creation_time": {
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
						"local_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"modifiedby": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"modified_time": {
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
						"trust_untrusted_cert": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"config_space": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_idle_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_max_age": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"double_encrypt": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"health_checktype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_reporting": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_anchored": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_cname_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"modifiedby": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"passive_health_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
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
						"dynamic_discovery": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"modifiedby": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"modified_time": {
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
			"tcp_port_ranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"udp_port_ranges": {
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
	_ = d.Set("segment_group_id", resp.SegmentGroupId)
	_ = d.Set("segment_group_name", resp.SegmentGroupName)
	_ = d.Set("bypass_type", resp.BypassType)
	_ = d.Set("config_space", resp.ConfigSpace)
	_ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("domain_names", resp.DomainNames)
	_ = d.Set("double_encrypt", resp.DoubleEncrypt)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("health_checktype", resp.HealthCheckType)
	_ = d.Set("health_reporting", resp.HealthReporting)
	_ = d.Set("ip_anchored", resp.IpAnchored)
	_ = d.Set("is_cname_enabled", resp.IsCnameEnabled)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modified_time", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("passive_health_enabled", resp.PassiveHealthEnabled)
	_ = d.Set("tcp_port_ranges", resp.TcpPortRanges)
	_ = d.Set("udp_port_ranges", resp.UdpPortRanges)

	if err := d.Set("clientless_apps", flattenClientlessApps(resp)); err != nil {
		return fmt.Errorf("failed to read clientless apps %s", err)
	}
	if err := d.Set("server_groups", flattenAppServerGroups(resp)); err != nil {
		return fmt.Errorf("failed to read app server groups %s", err)
	}

	return nil

}

func flattenClientlessApps(clientlessApp *applicationsegment.ApplicationSegmentResource) []interface{} {
	clientlessApps := make([]interface{}, len(clientlessApp.ClientlessApps))
	for i, clientlessApp := range clientlessApp.ClientlessApps {
		clientlessApps[i] = map[string]interface{}{
			"allow_options":        clientlessApp.AllowOptions,
			"appid":                clientlessApp.AppId,
			"application_port":     clientlessApp.ApplicationPort,
			"application_protocol": clientlessApp.ApplicationProtocol,
			"certificate_id":       clientlessApp.CertificateId,
			"certificate_name":     clientlessApp.CertificateName,
			"cname":                clientlessApp.Cname,
			"creationtime":         clientlessApp.CreationTime,
			"description":          clientlessApp.Description,
			"domain":               clientlessApp.Domain,
			"enabled":              clientlessApp.Enabled,
			"hidden":               clientlessApp.Hidden,
			"id":                   clientlessApp.ID,
			"local_domain":         clientlessApp.LocalDomain,
			"modifiedby":           clientlessApp.ModifiedBy,
			"modified_time":        clientlessApp.ModifiedTime,
			"name":                 clientlessApp.Name,
			"path":                 clientlessApp.Path,
			"trust_untrusted_cert": clientlessApp.TrustUntrustedCert,
		}
	}

	return clientlessApps
}

func flattenAppServerGroups(serverGroup *applicationsegment.ApplicationSegmentResource) []interface{} {
	serverGroups := make([]interface{}, len(serverGroup.ServerGroups))
	for i, val := range serverGroup.ServerGroups {
		serverGroups[i] = map[string]interface{}{
			"name":              val.Name,
			"id":                val.ID,
			"config_space":      val.ConfigSpace,
			"creation_time":     val.CreationTime,
			"description":       val.Description,
			"enabled":           val.Enabled,
			"dynamic_discovery": val.DynamicDiscovery,
			"modifiedby":        val.ModifiedBy,
			"modified_time":     val.ModifiedTime,
		}
	}

	return serverGroups
}
