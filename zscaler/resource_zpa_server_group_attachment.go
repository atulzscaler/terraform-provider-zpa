package zscaler

import (
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/servergroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerGroupAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerGroupAttachmentCreate,
		Read:   resourceServerGroupAttachmentRead,
		Delete: resourceServerGroupAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "This field defines the name of the server group.",
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "This field defines if the server group is enabled or disabled.",
			},
			"dynamic_discovery": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},
			"app_connector_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "List of app-connector IDs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
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

func resourceServerGroupAttachmentCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandServerGroupAttachment(d)
	log.Printf("[INFO] Creating zpa server group attachment with request\n%+v\n", req)

	resp, _, err := zClient.servergroup.Create(&req)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Created server group attachment request. ID: %v\n", resp)
	d.SetId(resp.ID)

	return resourceServerGroupAttachmentRead(d, m)
}

func resourceServerGroupAttachmentRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	resp, _, err := zClient.servergroup.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing server group attachment %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting server group attachment:\n%+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("dynamic_discovery", resp.DynamicDiscovery)
	_ = d.Set("appconnector_groups", flattenAppConnectorGroups(resp.AppConnectorGroups))

	return nil

}

func resourceServerGroupAttachmentDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	log.Printf("[INFO] Deleting server group attachment ID: %v\n", d.Id())

	if _, err := zClient.servergroup.Delete(d.Id()); err != nil {
		return err
	}
	d.SetId("")
	log.Printf("[INFO] server group attachment deleted")
	return nil
}

func expandServerGroupAttachment(d *schema.ResourceData) servergroup.ServerGroup {
	return servergroup.ServerGroup{
		Name:               d.Get("name").(string),
		ID:                 d.Get("id").(string),
		Enabled:            d.Get("enabled").(bool),
		DynamicDiscovery:   d.Get("dynamic_discovery").(bool),
		AppConnectorGroups: expandAppConnectorGroups(d),
	}

}
