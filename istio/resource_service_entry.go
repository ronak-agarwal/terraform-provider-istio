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

func resourceServiceEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceEntryCreate,
		Read:   resourceServiceEntryRead,
		Exists: resourceServiceEntryExists,
		Update: resourceServiceEntryUpdate,
		Delete: resourceServiceEntryDelete,

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("serviceentry"),
			"spec": {
				Type:        schema.TypeList,
				Description: "Spec defines the specification of the serviceentry",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: serviceEntrySpecFields(),
				},
			},
		},
	}
}

func resourceServiceEntryCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating New ServiceEntries resourceServiceEntryCreate")
	ic := meta.(*Config).istioClientset
	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	spec, err := expandServiceEntrySpec(d.Get("spec").([]interface{}))
	if err != nil {
		return err
	}
	s := v1alpha3.ServiceEntry{
		ObjectMeta: metadata,
		Spec:       *spec,
	}
	out, err := ic.NetworkingV1alpha3().ServiceEntries(metadata.Namespace).Create(context.TODO(), &s, metav1.CreateOptions{})
	if err != nil {
		log.Printf("[DEBUG] Failed to create ServiceEntries: %#v", err)
	}
	d.SetId(buildID(out.ObjectMeta))
	// Add wait time of creation
	time.Sleep(5 * time.Second)

	log.Printf("[INFO] Created New ServiceEntries %#v", out)

	return resourceServiceEntryRead(d, meta)
}

func resourceServiceEntryRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Read ServiceEntries resourceServiceEntryRead")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	vs, err := ic.NetworkingV1alpha3().ServiceEntries(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	err = d.Set("metadata", flattenMetadata(vs.ObjectMeta, d))
	if err != nil {
		return err
	}
	return nil
}

func resourceServiceEntryExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[INFO] Exists Sidecars resourceServiceEntryExists")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	_, err = ic.NetworkingV1alpha3().ServiceEntries(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, nil
}

func resourceServiceEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] update ServiceEntries resourceServiceEntryUpdate")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	ops := patchMetadata("metadata.0.", "/metadata/", d)
	if d.HasChange("spec") {
		spec, err := expandServiceEntrySpec(d.Get("spec").([]interface{}))
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
	out, err := ic.NetworkingV1alpha3().ServiceEntries(namespace).Patch(context.TODO(), name, pkgApi.JSONPatchType, data, metav1.PatchOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	log.Printf("[INFO] Submitted updated ServiceEntries: %#v", out)
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return resourceServiceEntryRead(d, meta)
}

func resourceServiceEntryDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] delete ServiceEntries resourceServiceEntryDelete")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	err = ic.NetworkingV1alpha3().ServiceEntries(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("[DEBUG] ServiceEntries delete error: %#v", err)
	}
	log.Printf("[INFO] ServiceEntries deleted")
	d.SetId("")
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return nil
}
