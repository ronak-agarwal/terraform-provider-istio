package istio

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func serviceEntrySpecFields() map[string]*schema.Schema {
	log.Printf("[INFO] Creating New serviceEntry serviceEntrySpecFields")
	se := map[string]*schema.Schema{
		"hosts": {
			Type:        schema.TypeList,
			Description: "List of hosts.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"addresses": {
			Type:        schema.TypeList,
			Description: "List of hosts.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ports": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify port",
			Elem: &schema.Resource{
				Schema: portFields(),
			},
		},
		"location": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "serviceEntry.",
		},
		"resolution": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "serviceEntry.",
		},
		"endpoints": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify serviceEntry",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"address": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "serviceEntry.",
					},
					"ports": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "serviceEntry.",
					},
					"labels": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "serviceEntry.",
					},
					"network": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "serviceEntry.",
					},
					"locality": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "serviceEntry.",
					},
					"weight": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "serviceEntry.",
					},
					"serviceaccount": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "serviceEntry.",
					},
				},
			},
		},
		"workloadselector": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify serviceEntry",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"labels": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "serviceEntry.",
					},
				},
			},
		},
		"exportto": {
			Type:        schema.TypeList,
			Description: "List of exportto.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"subjectaltnames": {
			Type:        schema.TypeList,
			Description: "List of SAN.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
	return se
}
