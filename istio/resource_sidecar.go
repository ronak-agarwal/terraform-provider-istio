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

func resourceSidecar() *schema.Resource {
	return &schema.Resource{
		Create: resourceSidecarCreate,
		Read:   resourceSidecarRead,
		Exists: resourceSidecarExists,
		Update: resourceSidecarUpdate,
		Delete: resourceSidecarDelete,

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("sidecar"),
			"spec": {
				Type:        schema.TypeList,
				Description: "Spec defines the specification of the sidecar",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: sidecarSpecFields(),
				},
			},
		},
	}
}

func resourceSidecarCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating New Sidecar resourceSidecarCreate")
	ic := meta.(*Config).istioClientset
	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	spec, err := expandSidecarSpec(d.Get("spec").([]interface{}))
	if err != nil {
		return err
	}
	s := v1alpha3.Sidecar{
		ObjectMeta: metadata,
		Spec:       *spec,
	}
	out, err := ic.NetworkingV1alpha3().Sidecars(metadata.Namespace).Create(context.TODO(), &s, metav1.CreateOptions{})
	if err != nil {
		log.Printf("[DEBUG] Failed to create Sidecar: %#v", err)
	}
	d.SetId(buildID(out.ObjectMeta))
	// Add wait time of creation
	time.Sleep(5 * time.Second)

	log.Printf("[INFO] Created New Sidecar %#v", out)

	return resourceSidecarRead(d, meta)
}

func resourceSidecarRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Read Sidecar resourceSidecarRead")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	vs, err := ic.NetworkingV1alpha3().Sidecars(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	err = d.Set("metadata", flattenMetadata(vs.ObjectMeta, d))
	if err != nil {
		return err
	}

	return nil
}

func resourceSidecarExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[INFO] Exists Sidecars resourceSidecarExists")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	_, err = ic.NetworkingV1alpha3().Sidecars(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, nil
}

func resourceSidecarUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] update Sidecars resourceSidecarUpdate")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	ops := patchMetadata("metadata.0.", "/metadata/", d)
	if d.HasChange("spec") {
		spec, err := expandSidecarSpec(d.Get("spec").([]interface{}))
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
	out, err := ic.NetworkingV1alpha3().Sidecars(namespace).Patch(context.TODO(), name, pkgApi.JSONPatchType, data, metav1.PatchOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	log.Printf("[INFO] Submitted updated Sidecars: %#v", out)
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return resourceSidecarRead(d, meta)
}

func resourceSidecarDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] delete Sidecars resourceSidecarDelete")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	err = ic.NetworkingV1alpha3().Sidecars(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("[DEBUG] Sidecars delete error: %#v", err)
	}
	log.Printf("[INFO] Sidecars deleted")
	d.SetId("")
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return nil
}
