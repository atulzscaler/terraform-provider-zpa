package zscaler

import (
	"log"
	"strconv"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/segmentgroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSegmentGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSegmentGroupCreate,
		Read:   resourceSegmentGroupRead,
		Update: resourceSegmentGroupUpdate,
		Delete: resourceSegmentGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"applications": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bypasstype": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"configspace": {
							Type:     schema.TypeString,
							Optional: true,
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
						"ipanchored": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"logfeatures": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"passivehealthenabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"healthchecktype": {
							Type:     schema.TypeString,
							Computed: true,
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
			"configspace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// "creationtime": {
			// 	Type:     schema.TypeInt,
			// 	Computed: true,
			// },
			"description": {
				Type:        schema.TypeString,
				Description: "Description of the app group.",
				Optional:    true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Whether this app group is enabled or not.",
				Optional:    true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// "modifiedby": {
			// 	Type:     schema.TypeInt,
			// 	Computed: true,
			// },
			// "modifiedtime": {
			// 	Type:     schema.TypeInt,
			// 	Computed: true,
			// },
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the app group.",
				Required:    true,
			},
			"policymigrated": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tcpkeepaliveenabled": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceSegmentGroupCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandSegmentGroup(d)
	log.Printf("[INFO] Creating segment group with request\n%+v\n", req)

	segmentgroup, _, err := zClient.segmentgroup.Create(req)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(int64(segmentgroup.ID), 10))
	return resourceSegmentGroupRead(d, m)

}

func resourceSegmentGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	// resp, _, err := zClient.segmentgroup.Get(d.Id())

	id, err := strconv.ParseInt(d.Get("id").(string), 10, 64)
	if err != nil {
		return err
	}

	resp, _, err := zClient.segmentgroup.Get(id)
	if err != nil {
		return err
	}

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing segment group %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting segment group:\n%+v\n", resp)
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("configspace", resp.ConfigSpace)
	// _ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	// _ = d.Set("modifiedby", resp.ModifiedBy)
	// _ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("policymigrated", resp.PolicyMigrated)
	_ = d.Set("tcpkeepaliveenabled", resp.TcpKeepAliveEnabled)
	_ = d.Set("appconnectorgroups", flattenSegmentGroupApplications(resp))
	// if err := d.Set("applications", flattenSegmentGroupApplications(resp)); err != nil {
	// 	return fmt.Errorf("failed to read segment group applications %s", err)
	// }
	return nil
}

func resourceSegmentGroupUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	segmentGroupRequest := expandSegmentGroup(d)
	segmentGroupRequest.ID = id
	log.Printf("[INFO] Updating IpList with name %s\n", segmentGroupRequest.Name)

	if _, err := zClient.segmentgroup.Update(id, segmentGroupRequest); err != nil {
		return err
	}
	return resourceSegmentGroupRead(d, m)
}

func resourceSegmentGroupDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting segment group with id %v\n", id)

	if _, err := zClient.segmentgroup.Delete(id); err != nil {
		return err
	}

	return nil
}

func expandSegmentGroup(d *schema.ResourceData) segmentgroup.SegmentGroup {
	segmentGroup := segmentgroup.SegmentGroup{
		Name:                d.Get("name").(string),
		Description:         d.Get("description").(string),
		Enabled:             d.Get("enabled").(bool),
		PolicyMigrated:      d.Get("policymigrated").(bool),
		TcpKeepAliveEnabled: d.Get("tcpkeepaliveenabled").(int),
		Applications:        expandSegmentGroupApplications(d),
		// Applications: expandSegmentGroupApplications(d.Get("applications").([]interface{})),
		// CreationTime:   d.Get("creationtime").(int32),
		// ModifiedBy:     d.Get("modifiedby").(int64),
		// ModifiedTime:   d.Get("modifiedtime").(int32),

	}
	return segmentGroup
}

/*
func expandSegmentGroupApplications(segmentGroupAppRequest []interface{}) []segmentgroup.Application {
	segmentGroupApps := make([]segmentgroup.Application, len(segmentGroupAppRequest))

	for i, segmentGroupApp := range segmentGroupAppRequest {
		segmentGroupAppItem := segmentGroupApp.(map[string]interface{})
		segmentGroupApps[i] = segmentgroup.Application{
			ID: segmentGroupAppItem["id"].(int64),
			// TCPPortRanges: segmentGroupAppItem["tcpportranges"].([]int32),
			// TCPPortsIn:    segmentGroupAppItem["tcpportsin"].([]int32),
			// ModifiedBy:    segmentGroupAppItem["modifiedby"].(int64),
			// ModifiedTime:  segmentGroupAppItem["modifiedtime"].(int32),
			// Name:          segmentGroupAppItem["name"].(string),
		}
	}

	return segmentGroupApps
}
*/

func expandSegmentGroupApplications(d *schema.ResourceData) []segmentgroup.Application {
	var segmentGroupApplications []segmentgroup.Application
	if applicationsInterface, ok := d.GetOk("applications"); ok {
		applications := applicationsInterface.([]interface{})
		segmentGroupApplications = make([]segmentgroup.Application, len(applications))
		for i, application := range applications {
			segmentGroupApplication := application.(map[string]interface{})
			segmentGroupApplications[i] = segmentgroup.Application{
				BypassType:           segmentGroupApplication["bypasstype"].(string),
				ConfigSpace:          segmentGroupApplication["configspace"].(string),
				CreationTime:         segmentGroupApplication["creationtime"].(int32),
				DefaultIdleTimeout:   segmentGroupApplication["defaultidletimeout"].(int32),
				DefaultMaxAge:        segmentGroupApplication["defaultmaxage"].(int32),
				Description:          segmentGroupApplication["description"].(string),
				DomainName:           segmentGroupApplication["domainname"].(string),
				DomainNames:          segmentGroupApplication["domainnames"].([]string),
				DoubleEncrypt:        segmentGroupApplication["doubleencrypt"].(bool),
				Enabled:              segmentGroupApplication["enabled"].(bool),
				HealthCheckType:      segmentGroupApplication["healthchecktype"].(string),
				IPAnchored:           segmentGroupApplication["ipanchored"].(bool),
				LogFeatures:          segmentGroupApplication["logfeatures"].([]string),
				ModifiedBy:           segmentGroupApplication["modifiedby"].(int64),
				ModifiedTime:         segmentGroupApplication["modifiedtime"].(int32),
				Name:                 segmentGroupApplication["name"].(string),
				ID:                   segmentGroupApplication["id"].(int64),
				PassiveHealthEnabled: segmentGroupApplication["passivehealthenabled"].(bool),
				TCPPortRanges:        segmentGroupApplication["tcpportranges"].([]int32),
				TCPPortsIn:           segmentGroupApplication["tcpportsin"].([]int32),
				TCPPortsOut:          segmentGroupApplication["tcpportsout"].([]int32),
				UDPPortRanges:        segmentGroupApplication["udpportranges"].([]int32),
			}
		}
	}

	return segmentGroupApplications
}

/*
func expandServerGroups(d *schema.ResourceData) []segmentgroup.AppServerGroup {
	var segmentServerGroups []segmentgroup.AppServerGroup
	if serverGroupsInterface, ok := d.GetOk("servergroups"); ok {
		serverGroups := serverGroupsInterface.([]interface{})
		segmentServerGroups = make([]segmentgroup.AppServerGroup, len(serverGroups))
		for i, serverGroup := range serverGroups {
			segmentServerGroup := serverGroup.(map[string]interface{})
			segmentServerGroups[i] = segmentgroup.AppServerGroup{
				ConfigSpace:  segmentServerGroup["configspace"].(string),
				CreationTime: segmentServerGroup["creationtime"].(int32),
				Description:  segmentServerGroup["description"].(string),
				Enabled:      segmentServerGroup["enabled"].(bool),
				ID:           segmentServerGroup["id"].(int64),
				ModifiedBy:   segmentServerGroup["modifiedby"].(int64),
				ModifiedTime: segmentServerGroup["modifiedtime"].(int32),
				Name:         segmentServerGroup["name"].(string),
			}
		}
	}

	return segmentServerGroups
}
*/
