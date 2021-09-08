terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}

data "zpa_saml_attribute" "email_user_sso" {
    name = "Email_User SSO"
}

output "zpa_saml_attribute" {
    value = data.zpa_saml_attribute.email_user_sso
}