package equinix

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	v4 "github.com/equinix-labs/fabric-go/fabric/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

connId := "temp connection uuid"

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
		UpdateContext: resourceFabricRoutingProtocolUpdate,
		DeleteContext: resourceFabricRoutingProtocolDelete,
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
	d.SetId(fabricRoutingProtocol.Uuid,)
	return setFabricRoutingProtocolMap(d, fabricRoutingProtocol)
}

func resourceFabricRoutingProtocolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Config).fabricClient
	ctx = context.WithValue(ctx, v4.ContextAccessToken, meta.(*Config).FabricAuthToken)
	schemaNotifications := d.Get("notifications").([]interface{})
	notifications := notificationToFabric(schemaNotifications)
	schemaOrder := d.Get("order").(*schema.Set).List()
	order := orderToFabric(schemaOrder)
	schemaAccount := d.Get("account").(*schema.Set).List()
	account := accountToFabricGateway(schemaAccount)
	schemaLocation := d.Get("location").(*schema.Set).List()
	location := locationToFabricGateway(schemaLocation)
	project := v4.Project{}
	schemaProject := d.Get("project").(*schema.Set).List()
	if len(schemaProject) != 0 {
		project = projectToFabricGateway(schemaProject)
	}
	schemaPackage := d.Get("package").(*schema.Set).List()
	packages := packageToFabricGateway(schemaPackage)

	createRequest := v4.FabricGatewayPostRequest{
		Name:          d.Get("name").(string),
		Type_:         d.Get("type").(string),
		Order:         &order,
		Location:      &location,
		Notifications: notifications,
		Package_:      &packages,
		Account:       &account,
		Project: 	   &project,
	}

	rp, _, err := client.GatewaysApi.CreateGateway(ctx, createRequest)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fg.Uuid)

	if _, err = waitUntilFGIsProvisioned(d.Id(), meta, ctx); err != nil {
		return diag.Errorf("error waiting for FG (%s) to be created: %s", d.Id(), err)
	}

	return resourceFabricGatewayRead(ctx, d, meta)
}

func setFabricRoutingProtocolMap(d *schema.ResourceData, rp v4.RoutingProtocolData) diag.Diagnostics {
	diags := diag.Diagnostics{}

	if rp.Type_ == "BGP" {
		err := setMap(d, map[string]interface{}{
			"name":          rp.Name,
			"href":          rp.Href,
			"type":          rp.Type_,
			"state":         rp.State,
			"directIpv4":    rp.DirectIpv4,
			"package":       fabricGatewayPackageToTerra(rp.Package_),
			"location":      locationFGToTerra(fg.Location),
			"change_log":    changeLogToTerra(fg.ChangeLog),
			"notifications": notificationToTerra(fg.Notifications),
			"project":       projectToTerra(fg.Project),
		})
		if err != nil {
			return diag.FromErr(err)
		}
		return diags
	}


	return diags
}