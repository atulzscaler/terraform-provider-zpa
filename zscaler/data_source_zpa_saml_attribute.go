package zscaler

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/samlattribute"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSamlAttribute() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSamlAttributeRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creationtime": {
							Type:     schema.TypeInt,
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
						"idp_name": {
							Type:     schema.TypeString,
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
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"saml_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_attribute": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSamlAttributeRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	var samlAttribute int64
	samlAttribute = 0
	if samlAttributeInterface, ok := d.GetOk("id"); ok {
		samlAttribute = int64(samlAttributeInterface.(int))
	}

	if samlAttribute != 0 {
		log.Printf("[INFO] Getting data for saml attribute %d\n", samlAttribute)
		// Getting specific saml attribute ID
		resp, _, err := zClient.samlattribute.Get(fmt.Sprintf("%d", samlAttribute))
		if err != nil {
			return err
		}

		// Add the one saml attribute ID we received
		d.SetId(fmt.Sprintf("%d", resp.ID))
		// Now we make a slice of saml attributes, so we can add the one saml attribute to the resource list after flattening
		samlAttributes := make([]samlattribute.SamlAttribute, 0)
		samlAttributes = append(samlAttributes, *resp)
		d.Set("list", flattenSamlAttributes(samlAttributes))
	} else {
		log.Printf("[INFO] Getting data for all saml attributes \n")
		// Getting all saml attribute IDs
		resp, _, err := zClient.samlattribute.GetAll()
		if err != nil {
			return err
		}

		// In case of all saml attributes returned, I don't now which ID to set as ID here, so I add time from documentation
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		d.Set("list", flattenSamlAttributes(resp))
	}

	return nil
}

func flattenSamlAttributes(samlAttributeResponse []samlattribute.SamlAttribute) []interface{} {
	samlAttributes := make([]interface{}, len(samlAttributeResponse))
	for i, samlAttributeItem := range samlAttributeResponse {
		samlAttributes[i] = map[string]interface{}{
			"creation_time":  samlAttributeItem.CreationTime,
			"id":             samlAttributeItem.ID,
			"idp_id":         samlAttributeItem.IdpId,
			"idp_name":       samlAttributeItem.IdpName,
			"modifiedby":     samlAttributeItem.ModifiedBy,
			"modified_time":  samlAttributeItem.ModifiedTime,
			"name":           samlAttributeItem.Name,
			"saml_name":      samlAttributeItem.SamlName,
			"user_attribute": samlAttributeItem.UserAttribute,
		}
	}
	return samlAttributes
}
