package istio

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	v1alpha3spec "istio.io/api/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
)

func resourceVirtualService() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualServiceCreate,
		Read:   resourceVirtualServiceRead,
		Exists: resourceVirtualServiceExists,
		Update: resourceVirtualServiceUpdate,
		Delete: resourceVirtualServiceDelete,

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("virtualservice"),
			"spec": {
				Type:        schema.TypeList,
				Description: "Spec defines the specification of the VS",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: virtualServiceSpecFields(),
				},
			},
		},
	}
}

func expandVirtualServiceSpec(virtualservice []interface{}) (*v1alpha3spec.VirtualService, error) {
	obj := &v1alpha3spec.VirtualService{}
	if len(virtualservice) == 0 || virtualservice[0] == nil {
		return obj, nil
	}
	return obj, nil
}

func resourceVirtualServiceCreate(d *schema.ResourceData, meta interface{}) error {
	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	spec, err := expandVirtualServiceSpec(d.Get("spec").([]interface{}))
	if err != nil {
		return err
	}
	vs := v1alpha3.VirtualService{
		ObjectMeta: metadata,
		Spec:       *spec,
	}
	log.Printf("[INFO] Creating New VirtualService %#v", vs)
	return nil
}

func resourceVirtualServiceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceVirtualServiceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return true, nil
}

func resourceVirtualServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceVirtualServiceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
