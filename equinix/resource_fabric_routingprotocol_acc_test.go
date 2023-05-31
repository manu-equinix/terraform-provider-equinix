package equinix

import (
	"context"
	"fmt"
	"testing"

	v4 "github.com/equinix-labs/fabric-go/fabric/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFabricCreateDirectRoutingProtocol(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkRoutingProtocolDelete,
		Steps: []resource.TestStep{
			{
				// og conn_uuid: 3d205bfa-1064-4d5b-a199-2908de84999e
				Config: testAccFabricCreateRoutingProtocolDirectConfig("1f2ea948-8dd7-4aec-be2e-666874f046c3", "10.10.100.1/26", "172::1:1/126"),
				//Check: resource.ComposeTestCheckFunc(
				//	resource.TestCheckResourceAttr(
				//		"equinix_fabric_routingprotocol.test", "direct_ipv4.equinix_iface_ip", fmt.Sprint("10.10.100.1/26")),
				//),
				ExpectNonEmptyPlan: true,
			}, {
				Config: testAccFabricCreateRoutingProtocolDirectConfig("1f2ea948-8dd7-4aec-be2e-666874f046c3", "192.168.200.201/26", "172::1:1/126"),
				//Check: resource.ComposeTestCheckFunc(
				//	resource.TestCheckResourceAttr(
				//		"equinix_fabric_routingprotocol.test", "direct_ipv4.equinix_iface_ip", fmt.Sprint("192.168.200.201/26")),
				//),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccFabricCreateRoutingProtocolDirectConfig(connectionUuid string, ipv4 string, ipv6 string) string {
	return fmt.Sprintf(`	resource "equinix_fabric_routingprotocol" "test" {
		connection_uuid = "%s"

		type = "DIRECT"
		name = "fabric_tf_acc_test_rpDirect"
		direct_ipv4{
			equinix_iface_ip = "%s"
		}
		direct_ipv6{
			equinix_iface_ip = "%s"
		}
	}`, connectionUuid, ipv4, ipv6)
}

func TestAccFabricCreateBgpRoutingProtocol(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkRoutingProtocolDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricCreateRoutingProtocolBgpConfig("1f2ea948-8dd7-4aec-be2e-666874f046c3", "180.1.1.2", "172::1:2"),
				//Check: resource.ComposeTestCheckFunc(
				//	resource.TestCheckResourceAttr(
				//		"equinix_fabric_routingprotocol.test", "bgpIpv4.customerPeerIp", fmt.Sprint("10.10.100.2")),
				//),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccFabricCreateRoutingProtocolBgpConfig("1f2ea948-8dd7-4aec-be2e-666874f046c3", "180.1.1.3", "172::1:2"),
				//Check: resource.ComposeTestCheckFunc(
				//	resource.TestCheckResourceAttr(
				//		"equinix_fabric_routingprotocol.test", "bgpIpv4.customerPeerIp", fmt.Sprint("10.10.100.2")),
				//),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccFabricCreateRoutingProtocolBgpConfig(connectionUuid string, ipv4 string, ipv6 string) string {
	return fmt.Sprintf(`	resource "equinix_fabric_routingprotocol" "test" {
		connection_uuid = "%s"

		type = "BGP"
		bgp_ipv4{
			customer_peer_ip = "%s"
		}
		bgp_ipv6{
			customer_peer_ip = "%s"
		}
		customer_asn = "100"
	}`, connectionUuid, ipv4, ipv6)
}

func checkRoutingProtocolDelete(s *terraform.State) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, v4.ContextAccessToken, testAccProvider.Meta().(*Config).FabricAuthToken)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "equinix_fabric_routingprotocol" {
			continue
		}
		err := waitUntilRoutingProtocolIsDeprovisioned(rs.Primary.ID, rs.Primary.Attributes["connection_uuid"], testAccProvider.Meta(), ctx)
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
				Config: testAccFabricReadRoutingProtocolConfig("1f2ea948-8dd7-4aec-be2e-666874f046c3", "00f48313-ab13-4524-aaad-93c31b5b8848"),
				//Check: resource.ComposeTestCheckFunc(
				//	resource.TestCheckResourceAttr(
				//		"equinix_fabric_routingprotocol.test", "type", fmt.Sprint("DIRECT")),
				//	resource.TestCheckResourceAttr(
				//		"equinix_fabric_routingprotocol.test", "state", fmt.Sprint("PROVISIONED")),
				//),
			},
		},
	})
}

func testAccFabricReadRoutingProtocolConfig(connectionUuid string, routingProtocolUuid string) string {
	return fmt.Sprintf(`data "equinix_fabric_routingprotocol" "test" {
	connection_uuid = "%s"
	uuid = "%s"
	}`, connectionUuid, routingProtocolUuid)
}
