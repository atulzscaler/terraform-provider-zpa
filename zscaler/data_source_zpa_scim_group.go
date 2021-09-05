package zscaler

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceScimGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceScimGroupRead,
		Schema: map[string]*schema.Schema{
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idp_id": {
				Type:     schema.TypeInt,
				Optional: true,
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
	}
}

func dataSourceScimGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	// id := d.Get("id").(string)
	// log.Printf("[INFO] Getting data for scim group %s\n", id)

	id, err := strconv.ParseInt(d.Get("id").(string), 10, 64)
	if err != nil {
		return err
	}

	resp, _, err := zClient.scimgroup.Get(id)
	if err != nil {
		return err
	}

	// d.SetId(resp.ID)
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("idp_group_id", resp.IdpGroupId)
	_ = d.Set("idp_id", resp.IdpId)
	_ = d.Set("modified_time", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)

	return nil

}

// func flattenScimGroups(scimGroupResponse []scimgroup.ScimGroup) []interface{} {
// 	scimGroups := make([]interface{}, len(scimGroupResponse))
// 	for i, scimGroupItem := range scimGroupResponse {
// 		scimGroups[i] = map[string]interface{}{
// 			"creation_time": scimGroupItem.CreationTime,
// 			"id":            scimGroupItem.ID,
// 			"idp_id":        scimGroupItem.IdpId,
// 			"idp_group_id":  scimGroupItem.IdpGroupId,
// 			"modified_time": scimGroupItem.ModifiedTime,
// 			"name":          scimGroupItem.Name,
// 		}
// 	}
// 	return scimGroups
// }
