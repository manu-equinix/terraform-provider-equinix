package equinix

import (
	"context"
	"fmt"
	"testing"

	v4 "github.com/equinix-labs/fabric-go/fabric/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFabricCreateRoutingProtocol(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkRoutingProtocolDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricCreateRoutingProtocolDirectConfig("10.10.100.1/26"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"equinix_fabric_routingprotocol.test", "name", fmt.Sprint("fabric_tf_acc_test_rpDirect")),
					resource.TestCheckResourceAttr(
						"equinix_fabric_routingprotocol.test", "direct_ipv4.equinix_iface_ip", fmt.Sprint("10.10.100.1/26")),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccFabricCreateRoutingProtocolDirectConfig(ipv4 string) string {
	return fmt.Sprintf(`resource "equinix_fabric_routingprotocol" "test" {
		type = "DIRECT"
		name = "fabric_tf_acc_test_rpDirect"
		direct_ipv4{
			equinix_iface_ip = "%s"
		}
	}`, ipv4)
}

func checkRoutingProtocolDelete(s *terraform.State) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, v4.ContextAccessToken, testAccProvider.Meta().(*Config).FabricAuthToken)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "equinix_fabric_routingprotocol" {
			continue
		}
		err := waitUntilRoutingProtocolDeprovisioned(rs.Primary.ID, rs.Primary.Attributes["connUuid"], testAccProvider.Meta(), ctx) // fixme: connUuid
		if err != nil {
			return fmt.Errorf("API call failed while waiting for resource deletion")
		}
	}
	return nil
}

func TestAccFabricReadRoutingProtocol(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricReadRoutingProtocolConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.equinix_fabric_routingprotocol.test", "name", fmt.Sprint("fabric_tf_acc_test")),
				),
			},
		},
	})
}

func testAccFabricReadRoutingProtocolConfig() string {
	return fmt.Sprint(`data "equinix_fabric_routingprotocol" "test" {
	uuid = "3e91216d-526a-45d2-9029-0c8c8ba48b60"
	connection_uuid = "f67896ea-fba1-11ed-be56-0242ac120002"
	}`) // todo: can you add connUuid as an extra and use it in resource??? - update i think this is it
}
