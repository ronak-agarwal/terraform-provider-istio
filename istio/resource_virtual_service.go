package istio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	v1alpha3spec "istio.io/api/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func virtualService() *v1alpha3.VirtualService {
	return &v1alpha3.VirtualService{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "hostname",
			Namespace: "route.Namespace",
		},
		Spec: v1alpha3spec.VirtualService{
			Hosts:    []string{"hostname"},
			Gateways: []string{"route.Name"},
			Http: []*v1alpha3spec.HTTPRoute{{
				Match: []*v1alpha3spec.HTTPMatchRequest{{
					Uri: &v1alpha3spec.StringMatch{
						MatchType: &v1alpha3spec.StringMatch_Prefix{
							Prefix: "/",
						},
					},
				}},
				Route: []*v1alpha3spec.HTTPRouteDestination{{
					Destination: &v1alpha3spec.Destination{
						Port: &v1alpha3spec.PortSelector{
							Number: uint32(22),
						},
						Host: "deployment.Name",
					},
				}},
			}},
		},
	}
}

func resourceVirtualServiceCreate(d *schema.ResourceData, meta interface{}) error {

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
