package istio

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func gatewaySpecFields() map[string]*schema.Schema {
	log.Printf("[INFO] Creating New serviceEntry serviceEntrySpecFields")
	gw := map[string]*schema.Schema{
		"servers": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify gateway",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"port": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Specify port",
						Elem: &schema.Resource{
							Schema: portFields(),
						},
					},
					"hosts": {
						Type:        schema.TypeList,
						Description: "List of hosts.",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"tls": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Specify port",
						Elem: &schema.Resource{
							Schema: serverTLSFields(),
						},
					},
				},
			},
		},
		"selector": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "gateway.",
		},
	}
	return gw
}

func serverTLSFields() map[string]*schema.Schema {
	log.Printf("[INFO] Creating New tls serverTlsFields")
	tls := map[string]*schema.Schema{
		"httpsredirect": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "gateway.",
		},
		"mode": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "gateway.",
		},
		"servercertificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "gateway.",
		},
		"privatekey": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "gateway.",
		},
		"cacertificates": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "gateway.",
		},
		"credentialname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "gateway.",
		},
		"subjectaltnames": {
			Type:        schema.TypeList,
			Description: "List of hosts.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"verifycertificatespki": {
			Type:        schema.TypeList,
			Description: "List of hosts.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"verifycertificatehash": {
			Type:        schema.TypeList,
			Description: "List of hosts.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"minprotocolversion": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "gateway.",
		},
		"maxprotocolversion": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "gateway.",
		},
		"ciphersuites": {
			Type:        schema.TypeList,
			Description: "List of hosts.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
	return tls
}
