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
				Type:     schema.TypeString,
				Optional: true,
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
	_ = d.Set("certchain", resp.CertChain)
	_ = d.Set("certificate", resp.Certificate)
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("issuedby", resp.IssuedBy)
	_ = d.Set("Issuedto", resp.IssuedTo)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("publickey", resp.PublicKey)
	_ = d.Set("san", resp.San)
	_ = d.Set("serialno", resp.SerialNo)
	_ = d.Set("status", resp.Status)
	_ = d.Set("validfrominepochsec", resp.ValidFromInEpochSec)
	_ = d.Set("validtoinepochsec", resp.ValidToInEpochSec)

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
