// Creating Browser Access Application
resource "zpa_browser_access" "browser_access_apps" {
    name = "Browser Access Apps"
    description = "Browser Access Apps"
    enabled = true
    healthreporting = "ON_ACCESS"
    bypasstype = "NEVER"
    tcpportranges = ["80", "80"]
    domainnames = ["sales.securitygeek.io"]
    segmentgroupid = zpa_segment_group.sg_sgio_browser_access.id

    clientlessapps {
        name = "sales.securitygeek.io"
        applicationprotocol = "HTTP"
        applicationport = "80"
        certificateid = data.zpa_ba_certificate.sales_ba.id
        trustuntrustedcert = true
        enabled = true
        domain = "sales.securitygeek.io"
    }
    servergroups {
        id = zpa_server_group.browser_access_apps.id
    }
}