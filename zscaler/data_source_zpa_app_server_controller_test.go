package zscaler

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/SecurityGeekIO/terraform-provider-zpa/zscaler/common/resourcetype"
	"github.com/SecurityGeekIO/terraform-provider-zpa/zscaler/common/testing/method"
	"github.com/SecurityGeekIO/terraform-provider-zpa/zscaler/common/testing/variable"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceRoleBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ApplicationServer)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApplicationServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckApplicationServerConfigure(resourceTypeAndName, generatedName, variable.ApplicationServerName, variable.ApplicationServerStatus),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "description", resourceTypeAndName, "description"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "permit_alert_actions", strconv.FormatBool(variable.ApplicationServerStatus)),
				),
			},
		},
	})
}

func testAccCheckApplicationServerDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.ApplicationServer {
			continue
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			log.Println("Failed in conversion with error:", err)
			return err
		}
		ipList, _, err := apiClient.appservercontroller.Get(id)

		if err == nil {
			return fmt.Errorf("id %d already exists", id)
		}

		if ipList != nil {
			return fmt.Errorf("iplist with id %d exists and wasn't destroyed", id)
		}
	}

	return nil
}

func testAccCheckApplicationServerConfigure(resourceTypeAndName, generatedName, description string, toPermittedAlertActions bool) string {
	return fmt.Sprintf(`
// role resource
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// resource variables
		RoleResourceHCL(generatedName, description, toPermittedAlertActions),

		// data source variables
		resourcetype.ApplicationServer,
		generatedName,
		resourceTypeAndName,
	)
}
