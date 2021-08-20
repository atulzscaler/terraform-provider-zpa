package zscaler

import (
	"log"
	"strconv"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/idpcontroller"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIdpController() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIdpControllerRead,
		Schema: map[string]*schema.Schema{
			"adminmetadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificateurl": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spentityid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spmetadataurl": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spposturl": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"autoprovision": {
				Type:     schema.TypeInt,
				Computed: true,
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
						"certificate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serialno": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validfrominsec": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"validtoinsec": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"creationtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disablesamlbasedpolicy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domainlist": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enablescimbasedpolicy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"idpentityid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"loginnameattribute": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"loginurl": {
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
				Optional: true,
			},
			"reauthonuserupdate": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"redirectbinding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"scimenabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"scimserviceproviderendpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scimsharedsecret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scimsharedsecretexists": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"signsamlrequest": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ssotype": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"usecustomspmetadata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"usermetadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificateurl": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spentityiid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spmetadataurl": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spposturl": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIdpControllerRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for machine group %s\n", id)

	resp, _, err := zClient.idpcontroller.Get(id)
	if err != nil {
		return err
	}

	//d.SetId(strconv.Itoa(resp.ID))
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("autoprovision", resp.AutoProvision)
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("disablesamlbasedpolicy", resp.DisableSamlBasedPolicy)
	_ = d.Set("domainlist", resp.Domainlist)
	_ = d.Set("enablescimbasedpolicy", resp.EnableScimBasedPolicy)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("idpentityid", resp.IdpEntityId)
	_ = d.Set("loginnameattribute", resp.LoginNameAttribute)
	_ = d.Set("loginurl", resp.LoginUrl)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("reauthonuserupdate", resp.ReauthOnUserUpdate)
	_ = d.Set("redirectbinding", resp.RedirectBinding)
	_ = d.Set("scimenabled", resp.ScimEnabled)
	_ = d.Set("scimserviceproviderendpoint", resp.ScimServiceProviderEndpoint)
	_ = d.Set("scimsharedsecret", resp.ScimSharedSecret)
	_ = d.Set("scimsharedsecretexists", resp.ScimSharedSecretExists)
	_ = d.Set("signsamlrequest", resp.SignSamlRequest)
	_ = d.Set("ssotype", resp.SsoType)
	_ = d.Set("usecustomspmetadata", resp.UseCustomSpMetadata)
	_ = d.Set("usermetadata", resp.UserMetadata)
	_ = d.Set("spentityid", resp.UserMetadata.SpEntityId)
	_ = d.Set("spmetadataurl", resp.UserMetadata.SpMetadataUrl)
	_ = d.Set("spposturl", resp.UserMetadata.SpPostUrl)
	_ = d.Set("certificates", flattenIdpCertificates(resp.Certificates))
	// _ = d.Set("usermetadata", flattenIdpUserMetadata(resp.UserMetadata))

	return nil
}

func flattenIdpCertificates(idpCertificate []idpcontroller.Certificates) []interface{} {
	idpCertificates := make([]interface{}, len(idpCertificate))
	for i, idpCertificateItems := range idpCertificate {
		idpCertificates[i] = map[string]interface{}{
			"certificate":    idpCertificateItems.Certificate,
			"cname":          idpCertificateItems.Cname,
			"serialno":       idpCertificateItems.Serialno,
			"validfrominsec": idpCertificateItems.ValidFrominSec,
			"validtoinsec":   idpCertificateItems.ValidToinSec,
		}
	}
	return idpCertificates
}

// func flattenIdpUserMetadata(idpUserMetadata idpcontroller.UserMetadata) []interface{} {
// 	userMetadata := make([]interface{}, len())
// 	for i, userMetadataItems := range idpUserMetadata {
// 		userMetadata[i] = map[string]interface{}{
// 			"certificateurl": userMetadataItems.CertificateUrl,
// 			"spentityid":     userMetadataItems.SpEntityId,
// 			"spmetadataurl":  userMetadataItems.SpMetadataUrl,
// 			"spposturl":      userMetadataItems.SpPostUrl,
// 		}
// 	}
// 	return userMetadata
// }
