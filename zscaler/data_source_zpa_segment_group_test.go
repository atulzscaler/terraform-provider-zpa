package zscaler

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSegmentGroupDataSource(t *testing.T) {
	rName := RandStringFromCharSet(10, CharSetAlphaNum)
	dataSourceName := "data.zpa_segment_group.test_segment_group_ds"
	resourceName := "zpa_segment_group.test_segment_group"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: segmentGroupDataSourceConfig(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func segmentGroupDataSourceConfig(rName string) string {
	return fmt.Sprintf(`
resource "zpa_segment_group" "test_segment_group" {
    name  = "%s"
	description = "test.acme.com"
	enabled = true
	policy_migrated = true
}
data "zpa_segment_group" "test_segment_group_ds" {
    name = zpa_segment_group.test_segment_group.id
}`, rName)
}
