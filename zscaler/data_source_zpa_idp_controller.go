package zscaler

import (
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/idpcontroller"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIdpController() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIdpControllerRead,
		Schema: map[string]*schema.Schema{
			"admin_metadata": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sp_entity_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sp_metadata_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sp_post_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"auto_provision": {
				Type:     schema.TypeString,
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
						"serial_no": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid_from_in_sec": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid_to_in_sec": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_saml_based_policy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enable_scim_based_policy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Default value if null is True",
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_entity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login_name_attribute": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modifiedby": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reauth_on_user_update": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"redirect_binding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"scim_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"scim_service_provider_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scim_shared_secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scim_shared_secret_exists": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"sign_saml_request": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sso_type": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"use_custom_sp_metadata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"user_metadata": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sp_entity_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sp_metadata_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sp_post_url": {
							Type:     schema.TypeString,
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
	// d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	d.SetId(resp.ID)
	_ = d.Set("auto_provision", resp.AutoProvision)
	_ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("disable_saml_based_policy", resp.DisableSamlBasedPolicy)
	_ = d.Set("domain_list", resp.Domainlist)
	_ = d.Set("enable_scim_based_policy", resp.EnableScimBasedPolicy)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("idp_entity_id", resp.IdpEntityId)
	_ = d.Set("login_name_attribute", resp.LoginNameAttribute)
	_ = d.Set("login_url", resp.LoginUrl)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modified_time", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("reauth_on_user_update", resp.ReauthOnUserUpdate)
	_ = d.Set("redirect_binding", resp.RedirectBinding)
	_ = d.Set("scim_enabled", resp.ScimEnabled)
	_ = d.Set("scim_service_provider_endpoint", resp.ScimServiceProviderEndpoint)
	_ = d.Set("scim_shared_secret", resp.ScimSharedSecret)
	_ = d.Set("scim_shared_secret_exists", resp.ScimSharedSecretExists)
	_ = d.Set("sign_saml_request", resp.SignSamlRequest)
	_ = d.Set("sso_type", resp.SsoType)
	_ = d.Set("use_custom_sp_metadata", resp.UseCustomSpMetadata)
	_ = d.Set("user_metadata.certificate_url", resp.UserMetadata.CertificateUrl)
	_ = d.Set("user_metadata.sp_entity_id", resp.UserMetadata.SpEntityId)
	_ = d.Set("user_metadata.sp_metadata_url", resp.UserMetadata.SpMetadataUrl)
	_ = d.Set("user_metadata.sp_post_url", resp.UserMetadata.SpPostUrl)
	_ = d.Set("admin_metadata.certificate_url", resp.AdminMetadata.CertificateUrl)
	_ = d.Set("admin_metadata.sp_entity_id", resp.AdminMetadata.SpEntityId)
	_ = d.Set("admin_metadata.sp_metadata_url", resp.AdminMetadata.SpMetadataUrl)
	_ = d.Set("admin_metadata.sp_post_url", resp.AdminMetadata.SpPostUrl)

	if err := d.Set("certificates", flattenIdpCertificates(resp.Certificates)); err != nil {
		return err
	}

	return nil
}

func flattenIdpCertificates(idpCertificate []idpcontroller.Certificates) []interface{} {
	idpCertificates := make([]interface{}, len(idpCertificate))
	for i, idpCertificateItems := range idpCertificate {
		idpCertificates[i] = map[string]interface{}{
			"certificate":       idpCertificateItems.Certificate,
			"cname":             idpCertificateItems.Cname,
			"serial_no":         idpCertificateItems.SerialNo,
			"valid_from_in_sec": idpCertificateItems.ValidFrominSec,
			"valid_to_in_sec":   idpCertificateItems.ValidToinSec,
		}
	}

	return idpCertificates
}
