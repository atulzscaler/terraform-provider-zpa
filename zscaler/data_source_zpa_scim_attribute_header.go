package zscaler

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceScimAttributeHeader() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceScimAttributeHeaderRead,
		Schema: map[string]*schema.Schema{
			"canonical_values": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"case_sensitive": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"data_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"idp_id": {
				Type:     schema.TypeInt,
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
			"multivalued": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"mutability": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"returned": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schema_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uniqueness": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceScimAttributeHeaderRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	// id := d.Get("id").(string)
	// log.Printf("[INFO] Getting data for scim group %s\n", id)

	id, err := strconv.ParseInt(d.Get("id").(string), 10, 64)
	if err != nil {
		return err
	}

	resp, _, err := zClient.scimattributeheader.Get(id)
	if err != nil {
		return err
	}

	// d.SetId(resp.ID)
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("canonical_values", resp.CanonicalValues)
	_ = d.Set("case_sensitive", resp.CaseSensitive)
	_ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("data_type", resp.DataType)
	_ = d.Set("description", resp.Description)
	_ = d.Set("idp_id", resp.IdpId)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modified_time", resp.ModifiedTime)
	_ = d.Set("multivalued", resp.MultiValued)
	_ = d.Set("mutability", resp.Mutability)
	_ = d.Set("name", resp.Name)
	_ = d.Set("required", resp.Required)
	_ = d.Set("returned", resp.Returned)
	_ = d.Set("schema_uri", resp.SchemaURI)
	_ = d.Set("uniqueness", resp.Uniqueness)

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
