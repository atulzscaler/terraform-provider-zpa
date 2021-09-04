package zscaler

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/scimgroup"
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
				Computed: true,
			},
			"idp_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idp_id": {
				Type:     schema.TypeInt,
				Required: true,
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

	var scimGroup int64
	scimGroup = 0
	if scimGroupInterface, ok := d.GetOk("id"); ok {
		scimGroup = int64(scimGroupInterface.(int))
	}

	if scimGroup != 0 {
		log.Printf("[INFO] Getting data for scim groups %d\n", scimGroup)
		// Getting specific saml attribute ID
		resp, _, err := zClient.scimgroup.Get(fmt.Sprintf("%d", scimGroup))
		if err != nil {
			return err
		}

		// Add the one saml attribute ID we received
		d.SetId(fmt.Sprintf("%d", resp.ID))
		// Now we make a slice of saml attributes, so we can add the one saml attribute to the resource list after flattening
		scimGroups := make([]scimgroup.ScimGroup, 0)
		scimGroups = append(scimGroups, *resp)
		d.Set("list", flattenScimGroups(scimGroups))
	} else {
		log.Printf("[INFO] Getting data for all scim groups \n")
		// Getting all saml attribute IDs
		resp, _, err := zClient.scimgroup.GetAll()
		if err != nil {
			return err
		}

		// In case of all saml attributes returned, I don't now which ID to set as ID here, so I add time from documentation
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		d.Set("list", flattenScimGroups(resp))
	}

	return nil
}

func flattenScimGroups(scimGroupResponse []scimgroup.ScimGroup) []interface{} {
	scimGroups := make([]interface{}, len(scimGroupResponse))
	for i, scimGroupItem := range scimGroupResponse {
		scimGroups[i] = map[string]interface{}{
			"creation_time": scimGroupItem.CreationTime,
			"id":            scimGroupItem.ID,
			"idp_id":        scimGroupItem.IdpId,
			"idp_group_id":  scimGroupItem.IdpGroupId,
			"modified_time": scimGroupItem.ModifiedTime,
			"name":          scimGroupItem.Name,
		}
	}
	return scimGroups
}
