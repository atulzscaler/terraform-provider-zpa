package zscaler

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBaCertificate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBaCertificateRead,
		Schema: map[string]*schema.Schema{
			"cname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cert_chain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"issued_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issued_to": {
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
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"san": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"serial_no": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"valid_from_in_epochsec": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"valid_to_in_epochsec": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceBaCertificateRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id, err := strconv.ParseInt(d.Get("id").(string), 10, 64)
	if err != nil {
		return err
	}

	resp, _, err := zClient.bacertificate.Get(id)
	if err != nil {
		return err
	}

	// d.SetId(strconv.Itoa(resp.ID))
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("cname", resp.CName)
	_ = d.Set("cert_chain", resp.CertChain)
	_ = d.Set("certificate", resp.Certificate)
	_ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("issued_by", resp.IssuedBy)
	_ = d.Set("issued_to", resp.IssuedTo)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modified_time", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("public_key", resp.PublicKey)
	_ = d.Set("san", resp.San)
	_ = d.Set("serial_no", resp.SerialNo)
	_ = d.Set("status", resp.Status)
	_ = d.Set("valid_from_in_epochsec", resp.ValidFromInEpochSec)
	_ = d.Set("valid_to_in_epochsec", resp.ValidToInEpochSec)

	return nil
}

/*
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
*/
