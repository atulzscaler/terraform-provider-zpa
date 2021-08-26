package zscaler

/*
import (
	"strconv"
	"testing"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/appservercontroller"
	"github.com/SecurityGeekIO/terraform-provider-zpa/zscaler/common/testing/method"
	"github.com/SecurityGeekIO/terraform-provider-zpa/zscaler/common/resourcetype"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceApplicationServerBasic(t *testing.T) {
	var role appservercontroller.ApplicationServer
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ApplicationServer)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApplicationServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckApplicationServerConfigure(resourceTypeAndName, generatedName, variable.ApplicationServerName, variable.RoleToPermittedAlertActions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationServerExists(resourceTypeAndName, &role),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.RoleName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "description", variable.RoleDescription),
					resource.TestCheckResourceAttr(resourceTypeAndName, "permit_alert_actions", strconv.FormatBool(variable.RoleToPermittedAlertActions)),
				),
			},

			// Update test
			{
				Config: testAccCheckRoleConfigure(resourceTypeAndName, generatedName, variable.RoleUpdateDescription, variable.RoleUpdateToPermittedAlertActions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleExists(resourceTypeAndName, &role),
					resource.TestCheckResourceAttr(resourceTypeAndName, "description", variable.RoleUpdateDescription),
					resource.TestCheckResourceAttr(resourceTypeAndName, "permit_alert_actions", strconv.FormatBool(variable.RoleUpdateToPermittedAlertActions)),
				),
			},
		},
	})
}

*/
