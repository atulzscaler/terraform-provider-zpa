
// Creating Browser Access Application
resource "zpa_browser_access" "browser_access_apps" {
    name = "Browser Access Apps"
    description = "Browser Access Apps"
    enabled = true
    health_reporting = "ON_ACCESS"
    bypass_type = "NEVER"
    tcp_port_ranges = ["80", "80"]
    domain_names = ["sales.securitygeek.io"]
    segment_group_id = zpa_segment_group.sg_sgio_browser_access.id

    clientless_apps {
        name = "sales.securitygeek.io"
        application_protocol = "HTTP"
        application_port = "80"
        certificate_id = data.zpa_ba_certificate.sales_ba.id
        trust_untrusted_cert = true
        enabled = true
        domain = "sales.securitygeek.io"
    }
    server_groups {
        id = [zpa_server_group.browser_access_apps.id]
    }
}