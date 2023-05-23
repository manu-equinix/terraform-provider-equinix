package equinix

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func readDirectConnectionIpv4Sch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"equinixIfaceIp": {
			Type: schema.TypeString,
			Required: true,
			Description: "Equinix side Interface IP address",
		},
	}
}
func readDirectConnectionIpv6Sch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"equinixIfaceIp": {
			Type: schema.TypeString,
			Required: false,
			Description: "Equinix side Interface IP address\n\n",
		},
	}
}
func readBgpConnectionIpv4Sch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"customerPeerIp": {
			Type: schema.TypeString,
			Required: false,
			Description: "Customer side peering ip",
		},
		"equinixPeerIp": {
			Type: schema.TypeString,
			Computed: true,
			Description: "Equinix side peering ip",
		},
		"enabled": {
			Type: schema.TypeBool,
			Optional: true,
			Default: true,
			Description: "Admin status for the BGP session",
		},
	}
}
func readBgpConnectionIpv6Sch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"customerPeerIp": {
			Type:        schema.TypeString,
			Required:    false,
			Description: "Customer side peering ip",
		},
		"equinixPeerIp": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Equinix side peering ip",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default: 	 true,
			Description: "Admin status for the BGP session",
		},
	}
}
func readRoutingProtocolBfdSch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type: schema.TypeBool,
			Required: true,
			Description: "Bidirectional Forwarding Detection enablement",
		},
		"interval": {
			Type: schema.TypeString,
			Optional: false,
			Default: 100,
			Description: "Interval range between the received BFD control packets",
		},
	}
}
func readRoutingProtocolOperationSch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"errors": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Errors occurred",
			Elem: &schema.Resource{
				Schema: createOperationalErrorSch(),
			},
		},
	}
}
func readRoutingProtocolChangeSch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		//"description": {
		//	Type: schema.TypeString,
		//	Computed: true,
		//	Description: "Details of latest Routing Protocol change",
		//},
		"uuid": {
			Type: schema.TypeString,
			Computed: true,
			Description: "Uniquely identifies a change",
		},
		"type": {
			Type: schema.TypeString,
			Computed: true,
			Description: "Type of change",
		},
		"href": {
			Type: schema.TypeString,
			Computed: true,
			Description: "Routing Protocol Change URI",
		},
	}
}

// todo: copy and rename createFabricRoutingProtocolResourceSchema() from schema.go