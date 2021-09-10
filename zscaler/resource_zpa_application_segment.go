package zscaler

import (
	"fmt"
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/applicationsegment"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/segmentgroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationSegment() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationSegmentCreate,
		Read:   resourceApplicationSegmentRead,
		Update: resourceApplicationSegmentUpdate,
		Delete: resourceApplicationSegmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"segment_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"segment_group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bypass_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Indicates whether users can bypass ZPA to access applications.",
				ValidateFunc: validation.StringInSlice([]string{
					"ALWAYS",
					"NEVER",
					"ON_NET",
				}, false),
			},
			"tcp_port_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TCP port ranges used to access the app.",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"udp_port_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "UDP port ranges used to access the app.",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"config_space": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"DEFAULT",
					"SIEM",
				}, false),
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the application.",
			},
			"domain_names": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of domains and IPs.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"double_encrypt": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether Double Encryption is enabled or disabled for the app.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether this application is enabled or not.",
			},
			"health_check_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_reporting": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.",
				ValidateFunc: validation.StringInSlice([]string{
					"NONE",
					"ON_ACCESS",
					"CONTINUOUS",
				}, false),
			},
			"icmp_access_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"PING_TRACEROUTING",
					"PING",
					"NONE",
				}, false),
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_anchored": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"log_features": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"skip_discovery",
					"full_wildcard",
				}, false),
			},
			"is_cname_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of the application.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"passive_health_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"server_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of the server group IDs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceApplicationSegmentCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandApplicationSegmentRequest(d)
	log.Printf("[INFO] Creating application segment request\n%+v\n", req)
	if req.SegmentGroupId == "" {
		log.Println("[ERROR] Please provde a valid segment group for the application segment")
		return fmt.Errorf("please provde a valid segment group for the application segment")
	}
	resp, _, err := zClient.applicationsegment.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created application segment request. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceApplicationSegmentRead(d, m)
}

func resourceApplicationSegmentRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	resp, _, err := zClient.applicationsegment.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing application segment %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Reading application segment and settings states: %+v\n", resp)
	_ = d.Set("segment_group_id", resp.SegmentGroupId)
	_ = d.Set("segment_group_name", resp.SegmentGroupName)
	_ = d.Set("bypass_type", resp.BypassType)
	_ = d.Set("config_space", resp.ConfigSpace)
	_ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("default_idle_timeout", resp.DefaultIdleTimeout)
	_ = d.Set("default_max_age", resp.DefaultMaxAge)
	_ = d.Set("description", resp.Description)
	_ = d.Set("domain_names", resp.DomainNames)
	_ = d.Set("double_encrypt", resp.DoubleEncrypt)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("health_check_type", resp.HealthCheckType)
	_ = d.Set("health_reporting", resp.HealthReporting)
	_ = d.Set("icmp_access_type", resp.IcmpAccessType)
	_ = d.Set("id", resp.ID)
	_ = d.Set("ip_anchored", resp.IpAnchored)
	_ = d.Set("is_cname_enabled", resp.IsCnameEnabled)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modified_time", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("passive_health_enabled", resp.PassiveHealthEnabled)
	_ = d.Set("ip_anchored", resp.IpAnchored)
	_ = d.Set("tcp_port_ranges", resp.TcpPortRanges)
	_ = d.Set("udp_port_ranges", resp.UdpPortRanges)
	_ = d.Set("server_groups", flattenAppServerGroups(resp))

	return nil
}

func resourceApplicationSegmentUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Id()
	log.Printf("[INFO] Updating role ID: %v\n", id)
	req := expandApplicationSegmentRequest(d)
	if d.HasChange("segment_group_id") && req.SegmentGroupId == "" {
		log.Println("[ERROR] Please provde a valid segment group for the application segment")
		return fmt.Errorf("please provde a valid segment group for the application segment")
	}
	if _, err := zClient.applicationsegment.Update(id, req); err != nil {
		return err
	}

	return resourceApplicationSegmentRead(d, m)
}

func resourceApplicationSegmentDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	id := d.Id()
	log.Printf("[INFO] Deleting application segment with id %v\n", id)
	segmentGroupId, ok := d.GetOk("segment_group_id")
	if ok && segmentGroupId != nil {
		gID, ok := segmentGroupId.(string)
		if ok && gID != "" {
			// detach it from segment group first
			if err := detachAppSegmentFromGroup(zClient, id, gID); err != nil {
				return err
			}
		}
	}

	if _, err := zClient.applicationsegment.Delete(id); err != nil {
		return err
	}

	return nil
}

func detachAppSegmentFromGroup(client *Client, segmentID, segmentGroupId string) error {
	log.Printf("[INFO] Detaching application segment  %s from segment group: %s\n", segmentID, segmentGroupId)
	segGroup, _, err := client.segmentgroup.Get(segmentGroupId)
	if err != nil {
		log.Printf("[error] Error while getting segment group id: %s", segmentGroupId)
		return err
	}
	adaptedApplications := []segmentgroup.Application{}
	for _, app := range segGroup.Applications {
		if app.ID != segmentID {
			adaptedApplications = append(adaptedApplications, app)
		}
	}
	segGroup.Applications = adaptedApplications
	_, err = client.segmentgroup.Update(segmentGroupId, segGroup)
	return err

}
func expandStringInSlice(d *schema.ResourceData, key string) []string {
	applicationSegments := d.Get(key).([]interface{})
	applicationSegmentList := make([]string, len(applicationSegments))
	for i, applicationSegment := range applicationSegments {
		applicationSegmentList[i] = applicationSegment.(string)
	}

	return applicationSegmentList
}

func expandApplicationSegmentRequest(d *schema.ResourceData) applicationsegment.ApplicationSegmentResource {
	return applicationsegment.ApplicationSegmentResource{
		SegmentGroupId:   d.Get("segment_group_id").(string),
		SegmentGroupName: d.Get("segment_group_name").(string),
		BypassType:       d.Get("bypass_type").(string),
		Description:      d.Get("description").(string),
		DomainNames:      expandStringInSlice(d, "domain_names"),
		DoubleEncrypt:    d.Get("double_encrypt").(bool),
		Enabled:          d.Get("enabled").(bool),
		HealthReporting:  d.Get("health_reporting").(string),
		IpAnchored:       d.Get("ip_anchored").(bool),
		IsCnameEnabled:   d.Get("is_cname_enabled").(bool),
		Name:             d.Get("name").(string),
		TcpPortRanges:    d.Get("tcp_port_ranges").([]interface{}),
		UdpPortRanges:    d.Get("udp_port_ranges").([]interface{}),
		ServerGroups:     expandAppServerGroups(d),
	}
}

func expandAppServerGroups(d *schema.ResourceData) []applicationsegment.AppServerGroups {
	appServerGroupsInterface, ok := d.GetOk("server_groups")
	if ok {
		appServer := appServerGroupsInterface.(*schema.Set)
		log.Printf("[INFO] app server groups data: %+v\n", appServer)
		var appServerGroups []applicationsegment.AppServerGroups
		for _, appServerGroup := range appServer.List() {
			appServerGroup, _ := appServerGroup.(map[string]interface{})
			if appServerGroup != nil {
				for _, id := range appServerGroup["id"].([]interface{}) {
					appServerGroups = append(appServerGroups, applicationsegment.AppServerGroups{
						ID: id.(string),
					})
				}
			}
		}
		return appServerGroups
	}

	return []applicationsegment.AppServerGroups{}
}
