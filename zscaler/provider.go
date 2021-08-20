package zscaler

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("ZPA_CLIENT_ID"),
				Description: "zpa client id",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: envDefaultFunc("ZPA_CLIENT_SECRET"),
				Description: "zpa client secret",
			},
			"customerid": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: envDefaultFunc("ZPA_CUSTOMER_ID"),
				Description: "zpa customer id",
			},
			"baseurl": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: envDefaultFunc("ZPA_BASE_URL"),
				Description: "zpa base url",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			/*
				terraform resource name: resource schema
				resource formation: provider-resourcename-subresource
			*/
			"zscaler_application_server":  resourceApplicationServer(),
			"zscaler_application_segment": resourceApplicationSegment(),
			"zscaler_server_group":        resourceServerGroup(),
			"zscaler_segment_group":       resourceSegmentGroup(),
			//"zscaler_browser_access": resourceBrowserAccess(),
			//"zscaler_policyset_rule": resourcePolicySetRule(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			// terraform date source name: data source schema
			"zscaler_posture_profile":       dataSourcePostureProfile(),
			"zscaler_trusted_network":       dataSourceTrustedNetwork(),
			"zscaler_saml_attribute":        dataSourceSamlAttribute(),
			"zscaler_scim_groups":           dataSourceScimGroup(),
			"zscaler_ba_certificate":        dataSourceBaCertificate(),
			"zscaler_machine_group":         dataSourceMachineGroup(),
			"zscaler_application_segment":   dataSourceApplicationSegment(),
			"zscaler_application_server":    dataSourceApplicationServer(),
			"zscaler_server_group":          dataSourceServerGroup(),
			"zscaler_cloud_connector_group": dataSourceCloudConnectorGroup(),
			"zscaler_app_connector_group":   dataSourceAppConnectorGroup(),
			"zscaler_segment_group":         dataSourceSegmentGroup(),
			"zscaler_idp_controller":        dataSourceIdpController(),
			"zscaler_policy_set_global":     dataSourcePolicySetGlobal(),
		},
		ConfigureFunc: zscalerConfigure,
	}
}

func envDefaultFunc(k string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(k); v != "" {
			return v, nil
		}

		return nil, nil
	}
}

func zscalerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ClientID:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
		CustomerID:   d.Get("customerid").(string),
		BaseURL:      d.Get("baseurl").(string),
	}

	return config.Client()
}
