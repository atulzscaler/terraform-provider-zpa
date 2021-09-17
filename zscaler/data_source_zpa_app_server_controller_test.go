package zscaler

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccApplicationServerDataSource(t *testing.T) {
	rName := RandStringFromCharSet(10, CharSetAlphaNum)
	dataSourceName := "data.zpa_application_server.test_app_server_ds"
	resourceName := "zpa_application_server.test_app_server"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: applicationServerDataSourceConfig(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func applicationServerDataSourceConfig(rName string) string {
	return fmt.Sprintf(`
resource "zpa_application_server" "test_app_server" {
    name  = "%s"
	description = "test.acme.com"
    address = "test.acme.com"
}
data "zpa_application_server" "test_app_server_ds" {
    name = zpa_application_server.test_app_server.id
}`, rName)
}
