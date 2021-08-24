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
		Schema: map[string]*schema.Schema{
			"applications": {
				Type:        schema.TypeList,
				Description: "The App ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"name": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
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
			"modifiedby": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"modifiedtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the app group.",
				Required:    true,
			},
			"policymigrated": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
		Create: resourceSegmentGroupCreate,
		Read:   resourceSegmentGroupRead,
		Update: resourceSegmentGroupUpdate,
		Delete: resourceSegmentGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceSegmentGroupCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandCreateSegmentGroupRequest(d)
	log.Printf("[INFO] Creating segment group with request\n%+v\n", req)

	segmentgroup, _, err := zClient.segmentgroup.Create(req)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(int64(segmentgroup.ID), 10))
	//d.SetId(strconv.Itoa(segmentGroup.ID))
	return resourceSegmentGroupRead(d, m)

}

func resourceSegmentGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	resp, _, err := zClient.segmentgroup.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing segment group %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting segment group:\n%+v\n", resp)
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

	if err := d.Set("applications", flattenSegmentGroupApplications(resp)); err != nil {
		return err
	}
	return nil
}

func resourceSegmentGroupUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	log.Println("An updated occurred")

	if d.HasChange("name") {
		log.Println("The name or ID has been changed")

		if _, err := zClient.segmentgroup.Update(d.Id(), segmentgroup.SegmentGroupRequest{
			Name: d.Get("name").(string),
		}); err != nil {
			return err
		}
		return resourceSegmentGroupRead(d, m)
	}

	return nil
}

func resourceSegmentGroupDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	log.Printf("[INFO] Deleting segment group with id %v\n", d.Id())

	if _, err := zClient.segmentgroup.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func expandCreateSegmentGroupRequest(d *schema.ResourceData) segmentgroup.SegmentGroupRequest {
	segmentGroup := segmentgroup.SegmentGroupRequest{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Enabled:        d.Get("enabled").(bool),
		PolicyMigrated: d.Get("policymigrated").(bool),
		Applications:   expandSegmentGroupApplications(d),
		// CreationTime:   d.Get("creationtime").(int32),
		// ModifiedBy:     d.Get("modifiedby").(int64),
		// ModifiedTime:   d.Get("modifiedtime").(int32),

	}
	return segmentGroup
}

func expandSegmentGroupApplications(d *schema.ResourceData) []segmentgroup.Applications {
	var segmentGroupApplications []segmentgroup.Applications
	if applicationsInterface, ok := d.GetOk("applications"); ok {
		applications := applicationsInterface.([]interface{})
		segmentGroupApplications = make([]segmentgroup.Applications, len(applications))
		for i, application := range applications {
			segmentGroupApplication := application.(map[string]interface{})
			segmentGroupApplications[i] = segmentgroup.Applications{
				// BypassType:           segmentGroupApplication["bypasstype"].(string),
				// ConfigSpace:          segmentGroupApplication["configspace"].(string),
				// CreationTime:         segmentGroupApplication["creationtime"].(int32),
				// DefaultIdleTimeout:   segmentGroupApplication["defaultidletimeout"].(int32),
				// DefaultMaxAge:        segmentGroupApplication["defaultmaxage"].(int32),
				// Description:          segmentGroupApplication["description"].(string),
				// DomainName:           segmentGroupApplication["domainname"].(string),
				// DomainNames:          segmentGroupApplication["domainnames"].([]string),
				// DoubleEncrypt:        segmentGroupApplication["doubleencrypt"].(bool),
				// Enabled:              segmentGroupApplication["enabled"].(bool),
				// HealthCheckType:      segmentGroupApplication["healthchecktype"].(string),
				// IPAnchored:           segmentGroupApplication["ipanchored"].(bool),
				// LogFeatures:          segmentGroupApplication["logfeatures"].([]string),
				// ModifiedBy:           segmentGroupApplication["modifiedby"].(int64),
				// ModifiedTime:         segmentGroupApplication["modifiedtime"].(int32),
				Name: segmentGroupApplication["name"].([]interface{}),
				ID:   segmentGroupApplication["id"].(int64),
				// PassiveHealthEnabled: segmentGroupApplication["passivehealthenabled"].(bool),
				// TCPPortRanges:        segmentGroupApplication["tcpportranges"].([]int32),
				// TCPPortsIn:           segmentGroupApplication["tcpportsin"].([]int32),
				// TCPPortsOut:          segmentGroupApplication["tcpportsout"].([]int32),
				// UDPPortRanges:        segmentGroupApplication["udpportranges"].([]int32),
				AppServerGroup: expandServerGroups(d),
			}
		}
	}

	return segmentGroupApplications
}

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

func flattenSegmentGroupApplications(segmentGroup *segmentgroup.SegmentGroupResponse) []interface{} {
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

func flattenAppServerGroup(segmentGroup segmentgroup.Applications) []interface{} {
	segmentServerGroups := make([]interface{}, len(segmentGroup.AppServerGroup))
	for i, segmentServerGroup := range segmentGroup.AppServerGroup {
		segmentServerGroups[i] = map[string]interface{}{
			"configspace":  segmentServerGroup.ConfigSpace,
			"creationtime": segmentServerGroup.CreationTime,
			"description":  segmentServerGroup.Description,
			"enabled":      segmentServerGroup.Enabled,
			"id":           segmentServerGroup.ID,
			"modifiedby":   segmentServerGroup.ModifiedBy,
			"modifiedtime": segmentServerGroup.ModifiedTime,
			"name":         segmentServerGroup.Name,
		}
	}

	return segmentServerGroups
}
