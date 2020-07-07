package istio

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func expandStringList(configured []interface{}) []string {
	if configured[0] != nil && len(configured) == 1 && strings.Contains(configured[0].(string), ",") {
		return strings.Split(configured[0].(string), ",")
	}

	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		if v != nil {
			vs = append(vs, v.(string))
		}
	}
	return vs
}

func idParts(id string) (string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		err := fmt.Errorf("unexpected id format (%q), expected %q", id, "namespace/name")
		return "", "", err
	}

	return parts[0], parts[1], nil
}

func buildID(meta metav1.ObjectMeta) string {
	return meta.Namespace + "/" + meta.Name
}

func expandMetadata(in []interface{}) metav1.ObjectMeta {
	meta := metav1.ObjectMeta{}
	if len(in) < 1 {
		return meta
	}
	m := in[0].(map[string]interface{})

	if v, ok := m["name"]; ok {
		meta.Name = v.(string)
	}
	if v, ok := m["namespace"]; ok {
		meta.Namespace = v.(string)
	}

	return meta
}

func expandStringUintMap(m map[string]interface{}) map[string]uint32 {
	result := make(map[string]uint32)
	for k, v := range m {
		val, _ := strconv.Atoi(v.(string))
		result[k] = uint32(val)
	}
	return result
}

func expandStringMap(m map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		result[k] = v.(string)
	}
	return result
}

func patchMetadata(keyPrefix, pathPrefix string, d *schema.ResourceData) PatchOperations {
	ops := make([]PatchOperation, 0, 0)
	if d.HasChange(keyPrefix + "annotations") {
		oldV, newV := d.GetChange(keyPrefix + "annotations")
		diffOps := diffStringMap(pathPrefix+"annotations", oldV.(map[string]interface{}), newV.(map[string]interface{}))
		ops = append(ops, diffOps...)
	}
	if d.HasChange(keyPrefix + "labels") {
		oldV, newV := d.GetChange(keyPrefix + "labels")
		diffOps := diffStringMap(pathPrefix+"labels", oldV.(map[string]interface{}), newV.(map[string]interface{}))
		ops = append(ops, diffOps...)
	}
	return ops
}

func flattenMetadata(meta metav1.ObjectMeta, d *schema.ResourceData, metaPrefix ...string) []interface{} {
	m := make(map[string]interface{})
	m["name"] = meta.Name
	m["resource_version"] = meta.ResourceVersion
	m["self_link"] = meta.SelfLink
	m["uid"] = fmt.Sprintf("%v", meta.UID)
	//	m["generation"] = meta.Generation

	if meta.Namespace != "" {
		m["namespace"] = meta.Namespace
	}

	return []interface{}{m}
}
