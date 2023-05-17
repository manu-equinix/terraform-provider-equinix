package equinix

import (
	"context"
	"fmt"
	"testing"

	v4 "github.com/equinix-labs/fabric-go/fabric/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFabricGatewayCreate(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkFabricGatewayDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricGatewayCreateConfig("PRO"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "name", fmt.Sprint("fg_tf_acc_test")),
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "package.code", fmt.Sprint("PRO")),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccFabricGatewayCreateConfig("LAB"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "name", fmt.Sprint("fg_tf_acc_test")),
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "package.code", fmt.Sprint("LAB")),
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
		if rs.Type != "equinix_fabric_gateway" {
			continue
		}
		err := waitUntilFGDeprovisioned(rs.Primary.ID, testAccProvider.Meta(), ctx)
		if err != nil {
			return fmt.Errorf("API call failed while waiting for resource deletion")
		}
	}
	return nil
}

func testAccFabricGatewayCreateConfig(name string) string {
	return fmt.Sprintf(`resource "equinix_fabric_gateway" "test"{
			type = "XF_GATEWAY"
			name = "fg_tf_acc_test"
			location{
			  metro_code  = "SV"
			}
			package{
				  code = "%s"
			}
			order{
				purchase_order_number = "1-234567"
			}
			   notifications{
				  type = "ALL"
				  emails = [
					"test@equinix.com",
					"test1@equinix.com"
				]
			}
				project{
				   project_id = "776847000642406"
			}
			  account {
				account_number = 203612
			}
		}`, name)
}

func TestAccFabricGatewayRead(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricGatewayReadConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "name", fmt.Sprint("fg_tf_acc_test")),
				),
			},
		},
	})
}


func testAccFabricGatewayReadConfig() string {
	return fmt.Sprint(`data "equinix_fabric_gateway" "test" {
		uuid = "3e91216d-526a-45d2-9029-0c8c8ba48b60"
	}`)
}

func TestAccFabricGatewayUpdate(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricGatewayUpdateConfig("fg_tf_test_Update"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"equinix_fabric_gateway.test", "name", fmt.Sprint("fg_tf_test_Update")),
				),
			},
		},
	})
}

func testAccFabricGatewayUpdateConfig(name string) string {
	return fmt.Sprintf(`data "equinix_fabric_gateway" "test"{
		uuid = "3e91216d-526a-45d2-9029-0c8c8ba48b60"
		op = "replace"
		path = "/name"
		value = %s
	}`, name)
}
