package zscaler

import (
	"log"
	"strconv"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/cloudconnectorgroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudConnectorGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudConnectorGroupRead,
		Schema: map[string]*schema.Schema{
			"creationtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloudconnectors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creationtime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ipacl": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issuedcertid": {
							Type:     schema.TypeInt,
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
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"geolocationid": {
				Type:     schema.TypeInt,
				Computed: true,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"ziacloud": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ziaorgid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceCloudConnectorGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for cloud connector group %s\n", id)

	resp, _, err := zClient.cloudconnectorgroup.Get()
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("geolocationid", resp.GeolocationId)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("ziacloud", resp.ZiaCloud)
	_ = d.Set("ziaorgid", resp.ZiaOrgid)
	_ = d.Set("cloudconnectors", flattenCloudConnectors(resp))

	return nil
}

func flattenCloudConnectors(cloudConnectors *cloudconnectorgroup.CloudConnectorGroup) []interface{} {
	connectorItems := make([]interface{}, len(cloudConnectors.CloudConnectors))
	for i, connectorItem := range cloudConnectors.CloudConnectors {
		connectorItems[i] = map[string]interface{}{
			"creationtime": connectorItem.CreationTime,
			"description":  connectorItem.Description,
			"enabled":      connectorItem.Enabled,
			"fingerprint":  connectorItem.Fingerprint,
			"id":           connectorItem.ID,
			"Ipacl":        connectorItem.IpAcl,
			"issuedcertid": connectorItem.IssuedCertId,
			"modifiedby":   connectorItem.ModifiedBy,
			"modifiedtime": connectorItem.ModifiedTime,
			"name":         connectorItem.Name,
		}
	}

	return connectorItems
}
