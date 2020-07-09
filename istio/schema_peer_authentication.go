package istio

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func peerAuthenticationSpecFields() map[string]*schema.Schema {
	log.Printf("[INFO] Creating New peerAuthentication peerAuthenticationSpecFields")
	pa := map[string]*schema.Schema{
		"selector": {
			Type:        schema.TypeList,
			Description: "A label query over pods that should match the Replicas count.",
			Optional:    true,
			ForceNew:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"matchlabels": {
						Type:        schema.TypeMap,
						Description: "A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
						Optional:    true,
						ForceNew:    true,
					},
				},
			},
		},
		"mtls": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify peerAuthentication",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"mode": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "gateway.",
					},
				},
			},
		},
		"portlevelmtls": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "peerAuthentication.",
		},
	}
	return pa
}
