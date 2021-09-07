package zscaler

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePostureProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePostureProfileRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"posture_udid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zscaler_cloud": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zscaler_customer_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modifiedby": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePostureProfileRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data user with id %s\n", id)

	resp, _, err := zClient.postureprofile.Get(id)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Getting Policy Set Global Rules:\n%+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("domain", resp.Domain)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modified_time", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("posture_udid", resp.PostureudId)
	_ = d.Set("zscaler_cloud", resp.ZscalerCloud)
	_ = d.Set("zscaler_customer_id", resp.ZscalerCustomerId)

	return nil
}
