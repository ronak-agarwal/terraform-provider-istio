package istio

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func metadataFields(objectName string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: fmt.Sprintf("Name of the %s, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names", objectName),
			Optional:    true,
			ForceNew:    true,
			Computed:    true,
		},
		"resource_version": {
			Type:        schema.TypeString,
			Description: fmt.Sprintf("An opaque value that represents the internal version of this %s that can be used by clients to determine when %s has changed. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency", objectName, objectName),
			Computed:    true,
		},
		"self_link": {
			Type:        schema.TypeString,
			Description: fmt.Sprintf("A URL representing this %s.", objectName),
			Computed:    true,
		},
		"uid": {
			Type:        schema.TypeString,
			Description: fmt.Sprintf("The unique in time and space value for this %s. More info: http://kubernetes.io/docs/user-guide/identifiers#uids", objectName),
			Computed:    true,
		},
	}
}

func namespacedMetadataSchema(objectName string) *schema.Schema {
	log.Printf("[INFO] Creating New VirtualService namespacedMetadataSchema")
	fields := metadataFields(objectName)
	fields["namespace"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: fmt.Sprintf("Namespace defines the space within which name of the %s must be unique.", objectName),
		Optional:    true,
		ForceNew:    true,
		Default:     "default",
	}

	return &schema.Schema{
		Type:        schema.TypeList,
		Description: fmt.Sprintf("Standard %s's metadata", objectName),
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: fields,
		},
	}
}
