package istio

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func portFields() map[string]*schema.Schema {
	log.Printf("[INFO] Creating New sidecar portFields")
	p := map[string]*schema.Schema{
		"number": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "sidecar.",
		},
		"protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "sidecar.",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "sidecar.",
		},
	}
	return p
}

func sidecarSpecFields() map[string]*schema.Schema {
	log.Printf("[INFO] Creating New sidecar sidecarSpecFields")
	sc := map[string]*schema.Schema{
		"workloadselector": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify sidecar",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"labels": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "sidecar.",
					},
				},
			},
		},
		"ingress": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify sidecar",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"port": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Specify sidecar",
						Elem: &schema.Resource{
							Schema: portFields(),
						},
					},
					"bind": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "sidecar.",
					},
					"capturemode": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "sidecar.",
					},
					"defaultendpoint": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "sidecar.",
					},
				},
			},
		},
		"egress": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify sidecar",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"port": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "sidecar.",
						Elem: &schema.Resource{
							Schema: portFields(),
						},
					},
					"bind": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "sidecar.",
					},
					"capturemode": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "sidecar.",
					},
					"hosts": {
						Type:        schema.TypeList,
						Description: "sidecar.",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
		"outboundtrafficpolicy": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify sidecar",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"mode": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "sidecar.",
					},
				},
			},
		},
	}
	return sc
}
