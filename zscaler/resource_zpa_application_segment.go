package zscaler

import (
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/applicationsegment"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"segmentgroupid": {
				Type:     schema.TypeInt,
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
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether this application is enabled or not.",
			},
			"healthchecktype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"healthreporting": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.",
			},
			"icmpaccesstype": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
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
			"passivehealthenabled": {
				Type:     schema.TypeBool,
				Optional: true,
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
		},
	}
}

func resourceApplicationSegmentCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	if zClient == nil {
		return resourceNotSupportedError()
	}

	req := expandApplicationSegmentRequest(d)
	log.Printf("[INFO] Creating application segment request\n%+v\n", req)

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
	_ = d.Set("segmentgroupId", resp.SegmentGroupId)
	_ = d.Set("segmentgroupname", resp.SegmentGroupName)
	_ = d.Set("bypasstype", resp.BypassType)
	_ = d.Set("configspace", resp.ConfigSpace)
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("defaultidletimeout", resp.DefaultIdleTimeout)
	_ = d.Set("defaultmaxage", resp.DefaultMaxAge)
	_ = d.Set("description", resp.Description)
	_ = d.Set("domainnames", resp.DomainNames)
	_ = d.Set("doubleencrypt", resp.DoubleEncrypt)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("healthchecktype", resp.HealthCheckType)
	_ = d.Set("healthreporting", resp.HealthReporting)
	_ = d.Set("icmpaccesstype", resp.IcmpAccessType)
	_ = d.Set("id", resp.ID)
	_ = d.Set("ipanchored", resp.IpAnchored)
	_ = d.Set("iscnameenabled", resp.IsCnameEnabled)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("passivehealthenabled", resp.PassiveHealthEnabled)
	_ = d.Set("ipanchored", resp.IpAnchored)
	_ = d.Set("tcpportranges", resp.TcpPortRanges)
	_ = d.Set("udpportranges", resp.UdpPortRanges)

	if err := d.Set("servergroups", flattenAppServerGroups(resp)); err != nil {
		return err
	}

	return nil
}

func resourceApplicationSegmentUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	if zClient == nil {
		return resourceNotSupportedError()
	}

	id := d.Id()
	log.Printf("[INFO] Updating role ID: %v\n", id)
	req := expandApplicationSegmentRequest(d)

	if _, err := zClient.applicationsegment.Update(id, req); err != nil {
		return err
	}

	return resourceApplicationSegmentRead(d, m)
}

func resourceApplicationSegmentDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	log.Printf("[INFO] Deleting application segment with id %v\n", d.Id())

	if _, err := zClient.applicationsegment.Delete(d.Id()); err != nil {
		return err
	}

	return nil
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
		SegmentGroupId:   d.Get("segmentgroupid").(int),
		SegmentGroupName: d.Get("segmentgroupname").(string),
		BypassType:       d.Get("bypasstype").(string),
		Description:      d.Get("description").(string),
		DomainNames:      expandStringInSlice(d, "domainnames"),
		DoubleEncrypt:    d.Get("doubleencrypt").(bool),
		Enabled:          d.Get("enabled").(bool),
		HealthReporting:  d.Get("healthreporting").(string),
		IpAnchored:       d.Get("ipanchored").(bool),
		IsCnameEnabled:   d.Get("iscnameenabled").(bool),
		Name:             d.Get("name").(string),
		TcpPortRanges:    d.Get("tcpportranges").([]interface{}),
		UdpPortRanges:    d.Get("udpportranges").([]interface{}),
		ServerGroups:     expandAppServerGroups(d),
	}
}

func expandAppServerGroups(d *schema.ResourceData) []applicationsegment.AppServerGroups {
	var serverGroups []applicationsegment.AppServerGroups
	if serverGroupInterface, ok := d.GetOk("servergroups"); ok {
		servers := serverGroupInterface.([]interface{})
		serverGroups = make([]applicationsegment.AppServerGroups, len(servers))
		for i, srvGroup := range servers {
			serverGroup := srvGroup.(map[string]interface{})
			serverGroups[i] = applicationsegment.AppServerGroups{
				// Name: serverGroup["name"].(string),
				ID: serverGroup["id"].(int),
				// ConfigSpace:      serverGroup["configspace"].(string),
				// CreationTime:     serverGroup["creationtime"].(int32),
				// Description:      serverGroup["description"].(string),
				// Enabled:          serverGroup["enabled"].(bool),
				// DynamicDiscovery: serverGroup["dynamicdiscovery"].(bool),
				// ModifiedBy:       serverGroup["modifiedby"].(int64),
				// ModifiedTime:     serverGroup["modifiedtime"].(int32),
			}
		}
	}

	return serverGroups
}
