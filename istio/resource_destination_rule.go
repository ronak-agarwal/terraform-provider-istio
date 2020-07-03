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

func resourceDestinationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceDestinationRuleCreate,
		Read:   resourceDestinationRuleRead,
		Exists: resourceDestinationRuleExists,
		Update: resourceDestinationRuleUpdate,
		Delete: resourceDestinationRuleDelete,

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("destinationrule"),
			"spec": {
				Type:        schema.TypeList,
				Description: "Spec defines the specification of the DestinationRule",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: destinationRuleSpecFields(),
				},
			},
		},
	}
}

func resourceDestinationRuleCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating New DestinationRule resourceDestinationRuleCreate")
	ic := meta.(*Config).istioClientset
	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	spec, err := expandDestinationRulesSpec(d.Get("spec").([]interface{}))
	if err != nil {
		return err
	}
	dr := v1alpha3.DestinationRule{
		ObjectMeta: metadata,
		Spec:       *spec,
	}
	out, err := ic.NetworkingV1alpha3().DestinationRules(metadata.Namespace).Create(context.TODO(), &dr, metav1.CreateOptions{})
	if err != nil {
		log.Printf("[DEBUG] Failed to create DestinationRule: %#v", err)
	}
	d.SetId(buildID(out.ObjectMeta))
	// Add wait time of creation
	time.Sleep(5 * time.Second)

	log.Printf("[INFO] Created New DestinationRule %#v", out)

	return resourceDestinationRuleRead(d, meta)
}

func resourceDestinationRuleRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Read DestinationRule resourceDestinationRuleRead")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	dr, err := ic.NetworkingV1alpha3().DestinationRules(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	err = d.Set("metadata", flattenMetadata(dr.ObjectMeta, d))
	if err != nil {
		return err
	}

	return nil
}

func resourceDestinationRuleExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[INFO] Exists DestinationRule resourceDestinationRuleExists")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	_, err = ic.NetworkingV1alpha3().DestinationRules(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, nil
}

func resourceDestinationRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] update DestinationRule resourceDestinationRuleUpdate")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	ops := patchMetadata("metadata.0.", "/metadata/", d)
	if d.HasChange("spec") {
		spec, err := expandDestinationRulesSpec(d.Get("spec").([]interface{}))
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
	out, err := ic.NetworkingV1alpha3().DestinationRules(namespace).Patch(context.TODO(), name, pkgApi.JSONPatchType, data, metav1.PatchOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	log.Printf("[INFO] Submitted updated DestinationRule: %#v", out)
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return resourceDestinationRuleRead(d, meta)
}

func resourceDestinationRuleDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] delete DestinationRule resourceDestinationRuleDelete")
	ic := meta.(*Config).istioClientset
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}
	err = ic.NetworkingV1alpha3().DestinationRules(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("[DEBUG] DestinationRule delete error: %#v", err)
	}
	log.Printf("[INFO] DestinationRule deleted")
	d.SetId("")
	// Add wait time of creation
	time.Sleep(5 * time.Second)
	return nil
}
