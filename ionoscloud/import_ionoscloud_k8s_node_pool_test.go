package ionoscloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIonosCloudk8sNodepool_ImportBasic(t *testing.T) {
	resourceName := "example"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDIonosCloudk8sNodepoolDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckIonosCloudk8sNodepoolConfigBasic, resourceName),
			},
			{
				ResourceName:            fmt.Sprintf("ionoscloud_k8s_node_pool.%s", resourceName),
				ImportStateIdFunc:       testAccIonosCloudk8sNodepoolImportStateID,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"maintenance_window.0.time"},
			},
		},
	})
}

func testAccIonosCloudk8sNodepoolImportStateID(s *terraform.State) (string, error) {
	var importID string = ""

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ionoscloud_k8s_node_pool" {
			continue
		}

		importID = fmt.Sprintf("%s/%s", rs.Primary.Attributes["k8s_cluster_id"], rs.Primary.ID)
	}

	return importID, nil
}