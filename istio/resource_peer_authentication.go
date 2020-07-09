package istio

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"istio.io/client-go/pkg/apis/security/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkgApi "k8s.io/apimachinery/pkg/types"
)

func resourcePeerAuthentication() *schema.Resource {
	return &schema.Resource{
		Create: resourcePeerAuthenticationCreate,
		Read:   resourcePeerAuthenticationRead,
		Exists: resourcePeerAuthenticationExists,
		Update: resourcePeerAuthenticationUpdate,
		Delete: resourcePeerAuthenticationDelete,

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("peerauthentication"),
			"spec": {
				Type:        schema.TypeList,
				Description: "Spec defines the specification of the PeerAuthentication",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: peerAuthenticationSpecFields(),
				},
			},
		},
	}
}

func resourcePeerAuthenticationCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating New PeerAuthentications resourcePeerAuthenticationCreate")
	ic := meta.(*Config).istioClientset
	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	spec, err := expandPeerAuthenticationSpec(d.Get("spec").([]interface{}))
	if err != nil {
		return err
	}
	s := v1beta1.PeerAuthentication{
		ObjectMeta: metadata,
		Spec:       *spec,
	}
	out, err := ic.SecurityV1beta1().PeerAuthentications(metadata.Namespace).Create(context.TODO(), &s, metav1.CreateOptions{})
	if err != nil {
		log.Printf("[DEBUG] Failed to create PeerAuthentications: %#v", err)
	}
	d.SetId(buildID(out.ObjectMeta))
	// Add wait time of creation
	time.Sleep(5 * time.Second)

	log.Printf("[INFO] Created New PeerAuthentications %#v", out)

	return resourcePeerAuthenticationRead(d, meta)
}

func resourcePeerAuthenticationRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Read PeerAuthentications resourcePeerAuthenticationRead")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	vs, err := ic.SecurityV1beta1().PeerAuthentications(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	err = d.Set("metadata", flattenMetadata(vs.ObjectMeta, d))
	if err != nil {
		return err
	}
	return nil
}

func resourcePeerAuthenticationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[INFO] Exists PeerAuthentications resourcePeerAuthenticationExists")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	_, err = ic.SecurityV1beta1().PeerAuthentications(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, nil
}

func resourcePeerAuthenticationUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] update PeerAuthentications resourcePeerAuthenticationUpdate")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	ops := patchMetadata("metadata.0.", "/metadata/", d)
	if d.HasChange("spec") {
		spec, err := expandPeerAuthenticationSpec(d.Get("spec").([]interface{}))
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
	out, err := ic.SecurityV1beta1().PeerAuthentications(namespace).Patch(context.TODO(), name, pkgApi.JSONPatchType, data, metav1.PatchOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	log.Printf("[INFO] Submitted updated PeerAuthentications: %#v", out)
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return resourcePeerAuthenticationRead(d, meta)
}

func resourcePeerAuthenticationDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] delete PeerAuthentications resourcePeerAuthenticationDelete")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	err = ic.SecurityV1beta1().PeerAuthentications(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("[DEBUG] ServiceEntries delete error: %#v", err)
	}
	log.Printf("[INFO] ServiceEntries deleted")
	d.SetId("")
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return nil
}
