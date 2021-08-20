package zscaler

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/bacertificate"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBaCertificate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBaCertificateRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"certificates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certchain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creationtime": {
							Type:     schema.TypeInt,
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
						"issuedby": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issuedto": {
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
						"publickey": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"san": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"serialno": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validfrominepochsec": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"validtoinepochsec": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceBaCertificateRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	var baCertificateID int64
	baCertificateID = 0
	if baCertificateIDInterface, ok := d.GetOk("id"); ok {
		baCertificateID = int64(baCertificateIDInterface.(int))
	}

	if baCertificateID != 0 {
		log.Printf("[INFO] Getting data for Ba Certificate %d\n", baCertificateID)
		// Getting specific Ba Certificate ID
		resp, _, err := zClient.bacertificate.Get(fmt.Sprintf("%d", baCertificateID))
		if err != nil {
			return err
		}

		// Add the one Ba Certificate ID we received
		d.SetId(fmt.Sprintf("%d", resp.ID))
		// Now we make a slice of Ba Certificates, so we can add the one Ba Certificate to the resource list after flattening
		baCertificates := make([]bacertificate.BaCertificate, 0)
		baCertificates = append(baCertificates, *resp)
		d.Set("certificates", flattenBaCertificates(baCertificates))
	} else {
		log.Printf("[INFO] Getting data for all Ba Certificates \n")
		// Getting all Ba Certificate IDs
		resp, _, err := zClient.bacertificate.GetAll()
		if err != nil {
			return err
		}

		// In case of all Ba Certificates returned, I don't now which ID to set as ID here, so I add time from documentation
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		d.Set("certificates", flattenBaCertificates(resp))
	}

	return nil
}

func flattenBaCertificates(baCertificateResponse []bacertificate.BaCertificate) []interface{} {
	baCertificates := make([]interface{}, len(baCertificateResponse))
	for i, baCertificateItem := range baCertificateResponse {
		baCertificates[i] = map[string]interface{}{
			"cname":               baCertificateItem.CName,
			"certchain":           baCertificateItem.CertChain,
			"certificate":         baCertificateItem.Certificate,
			"creationtime":        baCertificateItem.CreationTime,
			"description":         baCertificateItem.Description,
			"id":                  baCertificateItem.ID,
			"issuedby":            baCertificateItem.IssuedBy,
			"issuedto":            baCertificateItem.IssuedTo,
			"modifiedby":          baCertificateItem.ModifiedBy,
			"modifiedtime":        baCertificateItem.ModifiedTime,
			"name":                baCertificateItem.Name,
			"publickey":           baCertificateItem.PublicKey,
			"san":                 baCertificateItem.San,
			"serialno":            baCertificateItem.SerialNo,
			"status":              baCertificateItem.Status,
			"validfrominepochsec": baCertificateItem.ValidFromInEpochSec,
			"validtoinepochsec":   baCertificateItem.ValidToInEpochSec,
		}
	}
	return baCertificates
}
