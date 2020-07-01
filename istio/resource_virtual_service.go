package istio

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkgApi "k8s.io/apimachinery/pkg/types"
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

func resourceVirtualServiceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating New VirtualService resourceVirtualServiceCreate")
	ic := meta.(*Config).istioClientset
	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	spec, err := expandVirtualServiceSpec(d.Get("spec").([]interface{}))
	if err != nil {
		return err
	}
	vs := v1alpha3.VirtualService{
		ObjectMeta: metadata,
		Spec:       *spec,
	}
	out, err := ic.NetworkingV1alpha3().VirtualServices(metadata.Namespace).Create(context.TODO(), &vs, metav1.CreateOptions{})
	if err != nil {
		log.Printf("[DEBUG] Failed to create VirtualService: %#v", err)
	}
	d.SetId(buildID(out.ObjectMeta))
	// Add wait time of creation
	time.Sleep(5 * time.Second)

	log.Printf("[INFO] Created New VirtualService %#v", out)

	return resourceVirtualServiceRead(d, meta)
}

func resourceVirtualServiceRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Read VirtualService resourceVirtualServiceRead")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	vs, err := ic.NetworkingV1alpha3().VirtualServices(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	err = d.Set("metadata", flattenMetadata(vs.ObjectMeta, d))
	if err != nil {
		return err
	}

	return nil
}

func resourceVirtualServiceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[INFO] Exists VirtualService resourceVirtualServiceExists")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	_, err = ic.NetworkingV1alpha3().VirtualServices(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, nil
}

func resourceVirtualServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] update VirtualService resourceVirtualServiceUpdate")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	ops := patchMetadata("metadata.0.", "/metadata/", d)
	if d.HasChange("spec") {
		spec, err := expandVirtualServiceSpec(d.Get("spec").([]interface{}))
		if err != nil {
			return err
		}

		ops = append(ops, &ReplaceOperation{
			Path:  "/spec",
			Value: spec,
		})
	}
	data, err := ops.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Failed to marshal update operations: %s", err)
	}
	out, err := ic.NetworkingV1alpha3().VirtualServices(namespace).Patch(context.TODO(), name, pkgApi.JSONPatchType, data, metav1.PatchOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	log.Printf("[INFO] Submitted updated virtualservice: %#v", out)
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return resourceVirtualServiceRead(d, meta)
}

func resourceVirtualServiceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] delete VirtualService resourceVirtualServiceRead")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	err = ic.NetworkingV1alpha3().VirtualServices(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("[DEBUG] VirtualService delete error: %#v", err)
	}
	log.Printf("[INFO] VirtualService deleted")
	d.SetId("")
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return nil
}
