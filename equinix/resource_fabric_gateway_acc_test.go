package equinix

import (
	"context"
	"fmt"
	"testing"

	v4 "github.com/equinix-labs/fabric-go/fabric/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFabricGatewayCreateConnection(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkFabricGatewayDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricGatewayCreateConfig("LAB"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "name", fmt.Sprint("fabric_gateway_tf_acc_test")),
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "bandwidth", fmt.Sprint("50")),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccFabricGatewayCreateConfig("PRO"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "name", fmt.Sprint("fabric_gateway_tf_acc_test")),
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "bandwidth", fmt.Sprint("100")),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func checkFabricGatewayDelete(s *terraform.State) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, v4.ContextAccessToken, testAccProvider.Meta().(*Config).FabricAuthToken)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "equinix_fabric_connection" {
			continue
		}
		err := waitUntilFGDeprovisioned(rs.Primary.ID, testAccProvider.Meta(), ctx)
		if err != nil {
			return fmt.Errorf("API call failed while waiting for resource deletion")
		}
	}
	return nil
}

// To-do: Add config
func testAccFabricGatewayCreateConfig(packageCode string) string {
	return fmt.Sprintf(`%s---`, packageCode)
}

func TestAccFabricGatewayReadConnection(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricGatewayReadConnectionConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.equinix_fabric_connection.test", "name", fmt.Sprint("fabric_tf_acc_test")),
				),
			},
		},
	})
}

func testAccFabricGatewayReadConnectionConfig() string {
	return fmt.Sprint(`data "equinix_fabric_gateway" "test" {
	uuid = "3e91216d-526a-45d2-9029-0c8c8ba48b60"
	}`)
}
