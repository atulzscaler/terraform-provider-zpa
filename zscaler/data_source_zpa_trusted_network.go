package zscaler

import (
	"log"
	"strconv"
	"time"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/trustednetwork"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTrustedNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTrustedNetworkRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creation_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Required: true,
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
						"network_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"zscaler_cloud": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTrustedNetworkRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	trustedNetworkID := ""
	if trustedNetworkIDInterface, ok := d.GetOk("id"); ok {
		trustedNetworkID = trustedNetworkIDInterface.(string)
	}

	if len(trustedNetworkID) != 0 {
		log.Printf("[INFO] Getting data for posture profile %s\n", trustedNetworkID)
		// Getting specific posture ID
		resp, _, err := zClient.trustednetwork.Get(trustedNetworkID)
		if err != nil {
			return err
		}

		// Add the one posture profile ID we received
		d.SetId(resp.ID)
		// Now we make a slice of posture profiles, so we can add the one posture profile to the resource list after flattening
		trustedNetworks := make([]trustednetwork.TrustedNetwork, 0)
		trustedNetworks = append(trustedNetworks, *resp)
		d.Set("list", flattenTrustedNetwork(trustedNetworks))
	} else {
		log.Printf("[INFO] Getting data for all posture profiles \n")
		// Getting all posture IDs
		resp, _, err := zClient.trustednetwork.GetAll()
		if err != nil {
			return err
		}

		// In case of all posture profiles returned, I don't now which ID to set as ID here, so I add time from documentation
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		d.Set("list", flattenTrustedNetwork(resp))
	}

	return nil
}

func flattenTrustedNetwork(trustedNetwork []trustednetwork.TrustedNetwork) []interface{} {
	trustedNetworks := make([]interface{}, len(trustedNetwork))
	for i, trustedNetworksItem := range trustedNetwork {
		trustedNetworks[i] = map[string]interface{}{
			"creation_time": trustedNetworksItem.CreationTime,
			"domain":        trustedNetworksItem.Domain,
			"id":            trustedNetworksItem.ID,
			"modifiedby":    trustedNetworksItem.ModifiedBy,
			"modified_time": trustedNetworksItem.ModifiedTime,
			"name":          trustedNetworksItem.Name,
			"network_id":    trustedNetworksItem.NetworkId,
			"zscaler_cloud": trustedNetworksItem.ZscalerCloud,
		}
	}
	return trustedNetworks
}
