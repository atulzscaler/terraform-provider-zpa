package zscaler

import (
	"strconv"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/segmentgroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSegmentGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSegmentGroupRead,
		Schema: map[string]*schema.Schema{
			"applications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypasstype": {
							Type:     schema.TypeString,
							Computed: true,
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
						"domainname": {
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
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ipanchored": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"logfeatures": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
						"tcpportsin": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"tcpportsout": {
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
				Computed: true,
			},
			"policymigrated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tcpkeepaliveenabled": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSegmentGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	// id := d.Get("id").(string)
	// log.Printf("[INFO] Getting data for segment group %s\n", id)

	// resp, _, err := zClient.segmentgroup.Get(id)
	// if err != nil {
	// 	return err
	// }

	id, err := strconv.ParseInt(d.Get("id").(string), 10, 64)
	if err != nil {
		return err
	}

	resp, _, err := zClient.segmentgroup.Get(id)
	if err != nil {
		return err
	}

	//d.SetId(strconv.Itoa(resp.ID))
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("configspace", resp.ConfigSpace)
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("policymigrated", resp.PolicyMigrated)
	_ = d.Set("tcpkeepaliveenabled", resp.TcpKeepAliveEnabled)

	if err := d.Set("applications", flattenSegmentGroupApplications(resp)); err != nil {
		return err
	}

	return nil
}

func flattenSegmentGroupApplications(segmentGroup *segmentgroup.SegmentGroup) []interface{} {
	segmentGroupApplications := make([]interface{}, len(segmentGroup.Applications))
	for i, segmentGroupApplication := range segmentGroup.Applications {
		segmentGroupApplications[i] = map[string]interface{}{
			"bypasstype":           segmentGroupApplication.BypassType,
			"configspace":          segmentGroupApplication.ConfigSpace,
			"creationtime":         segmentGroupApplication.CreationTime,
			"defaultidletimeout":   segmentGroupApplication.DefaultIdleTimeout,
			"defaultmaxage":        segmentGroupApplication.DefaultMaxAge,
			"description":          segmentGroupApplication.Description,
			"domainname":           segmentGroupApplication.DomainName,
			"domainnames":          segmentGroupApplication.DomainNames,
			"doubleencrypt":        segmentGroupApplication.DoubleEncrypt,
			"enabled":              segmentGroupApplication.Enabled,
			"healthchecktype":      segmentGroupApplication.HealthCheckType,
			"ipanchored":           segmentGroupApplication.IPAnchored,
			"logfeatures":          segmentGroupApplication.LogFeatures,
			"modifiedby":           segmentGroupApplication.ModifiedBy,
			"modifiedtime":         segmentGroupApplication.ModifiedTime,
			"name":                 segmentGroupApplication.Name,
			"id":                   segmentGroupApplication.ID,
			"passivehealthenabled": segmentGroupApplication.PassiveHealthEnabled,
			"tcpportranges":        segmentGroupApplication.TCPPortRanges,
			"tcpportsin":           segmentGroupApplication.TCPPortsIn,
			"tcpportsout":          segmentGroupApplication.TCPPortsOut,
			"servergroups":         flattenAppServerGroup(segmentGroupApplication),
		}
	}

	return segmentGroupApplications
}

func flattenAppServerGroup(segmentGroup segmentgroup.Application) []interface{} {
	segmentServerGroups := make([]interface{}, len(segmentGroup.ServerGroup))
	for i, segmentServerGroup := range segmentGroup.ServerGroup {
		segmentServerGroups[i] = map[string]interface{}{
			"configspace":  segmentServerGroup.ConfigSpace,
			"creationtime": segmentServerGroup.CreationTime,
			"description":  segmentServerGroup.Description,
			"enabled":      segmentServerGroup.Enabled,
			//"id":           segmentServerGroup.ID,
			"modifiedby":   segmentServerGroup.ModifiedBy,
			"modifiedtime": segmentServerGroup.ModifiedTime,
			"name":         segmentServerGroup.Name,
		}
	}

	return segmentServerGroups
}
