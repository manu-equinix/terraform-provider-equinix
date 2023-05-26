package equinix

import (
	"context"
	"log"
	"strings"
	"time"

	v4 "github.com/equinix-labs/fabric-go/fabric/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFabricRoutingProtocol() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(6 * time.Minute),
			Read:   schema.DefaultTimeout(6 * time.Minute),
		},
		ReadContext:   resourceFabricRoutingProtocolRead,
		CreateContext: resourceFabricRoutingProtocolCreate,
		//UpdateContext: resourceFabricRoutingProtocolUpdate,
		//DeleteContext: resourceFabricRoutingProtocolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: createFabricRoutingProtocolResourceSchema(),

		Description: "Fabric V4 API compatible resource allows creation and management of Equinix Fabric connection\n\n~> **Note** Equinix Fabric v4 resources and datasources are currently in Beta. The interfaces related to `equinix_fabric_` resources and datasources may change ahead of general availability. Please, do not hesitate to report any problems that you experience by opening a new [issue](https://github.com/equinix/terraform-provider-equinix/issues/new?template=bug.md)",
	}
}

func resourceFabricRoutingProtocolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	connId := "temp connection uuid"

	client := meta.(*Config).fabricClient
	ctx = context.WithValue(ctx, v4.ContextAccessToken, meta.(*Config).FabricAuthToken)
	fabricRoutingProtocol, _, err := client.RoutingProtocolsApi.GetConnectionRoutingProtocolByUuid(ctx, d.Id(), connId)
	if err != nil {
		log.Printf("[WARN] Routing Protocol %s not found , error %s", d.Id(), err)
		if !strings.Contains(err.Error(), "500") {
			d.SetId("")
		}
		return diag.FromErr(err)
	}
	d.SetId(fabricRoutingProtocol.RoutingProtocolDirectData.Uuid)
	return setFabricRoutingProtocolMap(d, fabricRoutingProtocol)
}

func resourceFabricRoutingProtocolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Config).fabricClient
	ctx = context.WithValue(ctx, v4.ContextAccessToken, meta.(*Config).FabricAuthToken)
	schemaBgpIpv4 := d.Get("bgp_ipv4").(*schema.Set).List()
	bgpIpv4 := routingProtocolBgpIpv4ToFabric(schemaBgpIpv4)
	schemaBgpIpv6 := d.Get("bgp_ipv6").(*schema.Set).List()
	bgpIpv6 := routingProtocolBgpIpv6ToFabric(schemaBgpIpv6)
	schemaDirectIpv4 := d.Get("direct_ipv4").(*schema.Set).List()
	directIpv4 := routingProtocolDirectIpv4ToFabric(schemaDirectIpv4)
	schemaDirectIpv6 := d.Get("direct_ipv6").(*schema.Set).List()
	DirectIpv6 := routingProtocolDirectIpv6ToFabric(schemaDirectIpv6)
	schemaBfd := d.Get("bfd").(*schema.Set).List()
	bfd := routingProtocolBfdToFabric(schemaBfd)

	var createRequest = v4.RoutingProtocolBase{
		Type_: d.Get("type").(string),
		OneOfRoutingProtocolBase: v4.OneOfRoutingProtocolBase{
			RoutingProtocolBgpType: v4.RoutingProtocolBgpType{
				Type_:       d.Get("type").(string),
				Name:        d.Get("name").(string),
				BgpIpv4:     &bgpIpv4,
				BgpIpv6:     &bgpIpv6,
				CustomerAsn: d.Get("customer_asn").(int64),
				EquinixAsn:  d.Get("equinix_asn").(int64),
				BgpAuthKey:  d.Get("bgp_auth_key").(string),
				Bfd:         &bfd,
			},
			RoutingProtocolDirectType: v4.RoutingProtocolDirectType{
				Type_:      d.Get("type").(string),
				Name:       d.Get("name").(string),
				DirectIpv4: &directIpv4,
				DirectIpv6: &DirectIpv6,
			},
		},
	}
	rp, _, err := client.RoutingProtocolsApi.CreateConnectionRoutingProtocol(ctx, createRequest, "testConnectionId") // fixme: get connectionId
	if err != nil {
		return diag.FromErr(err)
	}

	switch rp.Type_ {
	case "BGP":
		d.SetId(rp.RoutingProtocolBgpData.Uuid)
	case "DIRECT":
		d.SetId(rp.RoutingProtocolDirectData.Uuid)
	}

	if _, err = waitUntilFGIsProvisioned(d.Id(), meta, ctx); err != nil {
		return diag.Errorf("error waiting for RP (%s) to be created: %s", d.Id(), err)
	}

	return resourceFabricRoutingProtocolRead(ctx, d, meta)
}

func setFabricRoutingProtocolMap(d *schema.ResourceData, rp v4.RoutingProtocolData) diag.Diagnostics {
	diags := diag.Diagnostics{}

	if rp.Type_ == "BGP" {
		err := setMap(d, map[string]interface{}{
			"name":    rp.RoutingProtocolBgpData.Name,
			"href":    rp.RoutingProtocolBgpData.Href,
			"type":    rp.RoutingProtocolBgpData.Type_,
			"state":   rp.RoutingProtocolBgpData.State,
			"bgpIpv4": rp.BgpIpv4,
		})
		if err != nil {
			return diag.FromErr(err)
		}
		return diags
	}

	return diags
}
