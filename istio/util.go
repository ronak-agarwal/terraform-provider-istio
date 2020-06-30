package istio

import (
	"strings"

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

func expandStringMap(m map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		result[k] = v.(string)
	}
	return result
}
