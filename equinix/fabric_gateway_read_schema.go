package equinix

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func readPackageSch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Type",
		},
		"href": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Unique Resource Identifier",
		},
		"code": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Fabric Gateway package code",
		},
	}
}

func readFabricGatewayResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Equinix-assigned Fabric Gateway identifier",
		},
		"href": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Fabric Gateway URI information",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Fabric Gateway name. An alpha-numeric 24 characters string which can include only hyphens and underscores",
		},
		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Customer-provided Fabric Gateway description",
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Connection overall state",
		},
		"equinixAsn": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Equinix ASN",
		},
		"bgpIpv4RoutesCount": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "",
		},
		"bgpIpv6RoutesCount": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "",
		},
		"connectionsCount": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "",
		},
		"package": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "Fabric Gateway package information",
			Elem: &schema.Resource{
				Schema: readPackageSch(),
			},
		},
		"change_log": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "Captures Fabric Gateway lifecycle change information",
			Elem: &schema.Resource{
				Schema: readChangeLogSch(),
			},
		},
		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Defines the Fabric Gateway type like XF_GATEWAY",
		},
		"location": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "Fabric Gateway location",
			Elem: &schema.Resource{
				Schema: readLocationSch(),
			},
		},
		"project": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Project information",
			Elem: &schema.Resource{
				Schema: createGatewayProjectSch(),
			},
		},
		"account": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "Customer account information that is associated with this connection",
			Elem: &schema.Resource{
				Schema: readAccountSch(),
			},
		},
		"notifications": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Preferences for notifications on Fabric Gateway configuration or status changes",
			Elem: &schema.Resource{
				Schema: readNotificationSch(),
			},
		},
	}
}
