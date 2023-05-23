package equinix

import (
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)


var createRoutingProtocolDirectIpv4Res = &schema.Resource{
	Schema: createRoutingProtocolDirectIpv4Sch(),
}

func createRoutingProtocolDirectIpv4Sch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"equinixIfaceIp": {
			Type: schema.TypeString,
			Required: true,
			Description: "Equinix side Interface IP address",
		},
	}
}

var createRoutingProtocolDirectIpv6Res = &schema.Resource{
	Schema: createRoutingProtocolDirectIpv6Sch(),
}

func createRoutingProtocolDirectIpv6Sch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"equinixIfaceIp": {
			Type: schema.TypeString,
			Required: false,
			Description: "Equinix side Interface IP address\n\n",
		},
	}
}
var createRoutingProtocolBgpIpv4Res = &schema.Resource{
	Schema: createRoutingProtocolBgpIpv4Sch(),
}

func createRoutingProtocolBgpIpv4Sch() map[string]*schema.Schema {
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

var createRoutingProtocolBgpIpv6Res = &schema.Resource{
	Schema: createRoutingProtocolBgpIpv6Sch(),
}

func createRoutingProtocolBgpIpv6Sch() map[string]*schema.Schema {
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

var createRoutingProtocolBfdRes = &schema.Resource{
	Schema: createRoutingProtocolBfdSch(),
}

func createRoutingProtocolBfdSch() map[string]*schema.Schema {
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
			// todo: validation
			Description: "Interval range between the received BFD control packets",
		},
	}
}

var createRoutingProtocolOperationRes = &schema.Resource{
	Schema: createRoutingProtocolOperationSch(),
}

func createRoutingProtocolOperationSch() map[string]*schema.Schema {
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

var createRoutingProtocolChangeRes = &schema.Resource{
	Schema: createRoutingProtocolChangeSch(),
}

func createRoutingProtocolChangeSch() map[string]*schema.Schema {
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
			ValidateFunc: validation.StringInSlice([]string{"ROUTING_PROTOCOL_UPDATE", "ROUTING_PROTOCOL_CREATION", "ROUTING_PROTOCOL_DELETION"}, true),
			Description: "Type of change",
		},
		"href": {
			Type: schema.TypeString,
			Computed: true,
			Description: "Routing Protocol Change URI",
		},
	}
}

func createFabricRoutingProtocolResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"href": {
			Type: 			schema.TypeString,
			Computed: 		true,
			Description: 	"Routing Protocol URI information",
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"BGP", "DIRECT"}, true),
			Description:  "Defines the routing protocol type like BGP or DIRECT",
		},
		"uuid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Equinix-assigned routing protocol identifier",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Customer-provided Fabric Routing Protocol description",
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			ValidateFunc: validation.StringInSlice([]string{"PROVISIONED", "DEPROVISIONED", "PROVISIONING", "DEPROVISIONING", "REPROVISIONING", "FAILED"}, true),
			Description: "Routing Protocol overall state",
		},
		"operation": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "Routing Protocol type-specific operational data",
			Elem: &schema.Resource{
				Schema: createRoutingProtocolOperationSch(),
			},
		},
		"change": {
			Type: schema.TypeSet,
			Computed: true,
			Description: "Routing Protocol configuration Changes",
			Elem: &schema.Resource{
				Schema: createRoutingProtocolChangeSch(),
			},
		},
		// fixme: questions about primative behaviors for ipv4 and ipv6
		"direct_ipv4": {
			Type: schema.TypeSet,
			Required: true,
			Description: "Routing Protocol Direct IPv4",
			Elem: &schema.Resource{
				Schema: createRoutingProtocolDirectIpv4Sch(),
			},
		},
		"direct_ipv6": {
			Type: schema.TypeSet,
			Optional: true,
			Description: "Routing Protocol Direct IPv6",
			Elem: &schema.Resource{
				Schema: createRoutingProtocolDirectIpv6Sch(),
			},
		},
		"bgp_ipv4": {
			Type: schema.TypeSet,
			Optional: true,
			Description: "Routing Protocol BGP IPv4",
			Elem: &schema.Resource{
				Schema: createRoutingProtocolBgpIpv4Sch(),
			},
		},
		"bgp_ipv6": {
			Type: schema.TypeSet,
			Optional: true,
			Description: "Routing Protocol BGP IPv6",
			Elem: &schema.Resource{
				Schema: createRoutingProtocolBgpIpv6Sch(),
			},
		},
		"customer_asn": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Customer-provided ASN",
		},
		"equinix_asn": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Equinix ASN",
		},
		"bgp_auth_key": {
			Type: schema.TypeString,
			Optional: true,
			Description: "BGP authorization key",
		},
		"bfd": {
			Type: schema.TypeSet,
			Optional: true,
			Description: "Bidirectional Forwarding Detection",
			Elem: &schema.Resource{
				Schema: createRoutingProtocolBfdSch(),
			},
		},
		"change_log": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "Captures Routing Protocol lifecycle change information",
			Elem: &schema.Resource{
				Schema: createChangeLogSch(),
			},
		},
	}
}