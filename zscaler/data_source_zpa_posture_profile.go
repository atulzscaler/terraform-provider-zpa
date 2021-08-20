package zscaler

import (
	"log"
	"strconv"

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
			"postureudid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zscalercloud": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zscalercustomerid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"creationtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"modifiedtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"modifiedby": {
				Type:     schema.TypeInt,
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

	//d.SetId(strconv.Itoa(resp.ID))
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("domain", resp.Domain)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("postureudid", resp.PostureudId)
	_ = d.Set("zscalercloud", resp.ZscalerCloud)
	_ = d.Set("zscalercustomerid", resp.ZscalerCustomerId)

	return nil
}

/*
func flattenPostureProfile(postureProfile []postureprofile.PostureProfile) []interface{} {
	postureProfiles := make([]interface{}, len(postureProfile))
	for i, postureProfileItem := range postureProfile {
		postureProfiles[i] = map[string]interface{}{
			"id":                postureProfileItem.ID,
			"name":              postureProfileItem.Name,
			"creationtime":      postureProfileItem.CreationTime,
			"domain":            postureProfileItem.Domain,
			"modifiedby":        postureProfileItem.ModifiedBy,
			"modifiedtime":      postureProfileItem.ModifiedTime,
			"postureudid":       postureProfileItem.PostureudId,
			"zscalercloud":      postureProfileItem.ZscalerCloud,
			"zscalercustomerid": postureProfileItem.ZscalerCustomerId,
		}
	}
	return postureProfiles
}
*/
