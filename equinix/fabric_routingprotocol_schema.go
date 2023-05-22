package equinix

import (
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)


var createDirectConnectionIpv4Res = &schema.Resource{
	Schema: createDirectConnectionIpv4Sch(),
}

func createDirectConnectionIpv4Sch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"equinixIfaceIp": {
			Type: schema.TypeString,
			Required: true,
			Description: "Equinix side Interface IP address",
		},
	}
}

var createDirectConnectionIpv6Res = &schema.Resource{
	Schema: createDirectConnectionIpv6Sch(),
}

func createDirectConnectionIpv6Sch() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"equinixIfaceIp": {
			Type: schema.TypeString,
			Required: false,
			Description: "Equinix side Interface IP address\n\n",
		},
	}
}
var createBgpConnectionIpv4Res = &schema.Resource{
	Schema: createBgpConnectionIpv4Sch(),
}

func createBgpConnectionIpv4Sch() map[string]*schema.Schema {
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
			Required: false,
			Description: "Admin status for the BGP session",
		},
	}
}

var createBgpConnectionIpv6Res = &schema.Resource{
	Schema: createBgpConnectionIpv6Sch(),
}

func createBgpConnectionIpv6Sch() map[string]*schema.Schema {
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
			Required:    false,
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
			Required: false,
			Description: "Bidirectional Forwarding Detection enablement",
		},
		"interval": {
			Type: schema.TypeString,
			Optional: false,
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
		"description": {
			Type: schema.TypeString,
			Computed: true,
			Description: "Current state of latest Routing Protocol change",
		},
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

// todo:routingProtocol /changes schema
////var createChangesRes = &schema.Resource{
////	Schema: createChangesSch(),
////}
//func createChangesSch() map[string]*schema.Schema {
//	return map[string]*schema.Schema{
//		"href": {
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Routing Protocol Changes URI information",
//		},
//		"type": {
//			Type: ,
//			Computed: ,
//			Description: ,
//		},
//		"uuid": {
//			Type: ,
//			Computed: ,
//			Description: ,
//		},
//		"status": {
//			Type: ,
//			Computed: ,
//			Description: ,
//		},
//		"created_by": {
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Created by User Key",
//		},
//		"created_by_full_name": {
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Created by User Full Name",
//		},
//		"created_by_email": {
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Created by User Email Address",
//		},
//		"created_date_time": {
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Created by Date and Time",
//		},
//		"updated_by": {
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Updated by User Key",
//		},
//		"updated_by_full_name": {
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Updated by User Full Name",
//		},
//		"updated_by_email": {
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Updated by User Email Address",
//		},
//		"updated_date_time": {
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Updated by Date and Time",
//		},
//		"data": {	// fixme
//			Type:        schema.TypeString,
//			Computed:    true,
//			Description: "Change data",
//		},
//	}
//}

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
		//"uuid": {
		//	Type: ,
		//	Computed: ,
		//	Description: ,
		//},
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
			Description: "", // todo: add
			Elem: &schema.Resource{
				Schema: createRoutingProtocolChangeSch(),
			},
		},
		"direct_ipv4": {
			Type: schema.TypeSet,
			Required: true,
			Description: "", // todo: add
			Elem: &schema.Resource{
				Schema: createDirectConnectionIpv4Sch(),
			},
		},
		"direct_ipv6": {
			Type: schema.TypeSet,
			Optional: true,
			Description: "", // todo: add
			Elem: &schema.Resource{
				Schema: createDirectConnectionIpv6Sch(),
			},
		},
		"bgp_ipv4": {
			Type: schema.TypeSet,
			Required: true,
			Description: "", // todo: add
			Elem: &schema.Resource{
				Schema: createBgpConnectionIpv4Sch(),
			},
		},
		"bgp_ipv6": {
			Type: schema.TypeSet,
			Optional: true,
			Description: "", // todo: add
			Elem: &schema.Resource{
				Schema: createBgpConnectionIpv6Sch(),
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
			Required: false,
			Description: "BGP authorization key",
		},
		"bfd": {
			Type: schema.TypeSet,
			Required: false,
			Description: "",
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