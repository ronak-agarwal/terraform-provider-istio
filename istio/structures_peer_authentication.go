package istio

import (
	"strconv"

	securityv1beta1 "istio.io/api/security/v1beta1"
	"istio.io/api/type/v1beta1"
)

func expandPeerAuthenticationSpec(peerauth []interface{}) (*securityv1beta1.PeerAuthentication, error) {
	obj := &securityv1beta1.PeerAuthentication{}
	if len(peerauth) == 0 || peerauth[0] == nil {
		return obj, nil
	}
	in := peerauth[0].(map[string]interface{})
	if v, ok := in["selector"].([]interface{}); ok && len(v) > 0 {
		obj.Selector = expandLabelSelector(v)
	}
	if v, ok := in["mtls"].([]interface{}); ok && len(v) > 0 {
		obj.Mtls = expandMutualTLS(v)
	}
	if v, ok := in["portlevelmtls"].(map[string]interface{}); ok && len(v) > 0 {
		obj.PortLevelMtls = expandPortMutualTLS(v)
	}
	return obj, nil
}

func expandPortMutualTLS(in map[string]interface{}) map[uint32]*securityv1beta1.PeerAuthentication_MutualTLS {
	if len(in) == 0 {
		return nil
	}
	result := make(map[uint32]*securityv1beta1.PeerAuthentication_MutualTLS)
	for k, v := range in {
		key, _ := strconv.Atoi(k)
		objm := &securityv1beta1.PeerAuthentication_MutualTLS{}
		objm.Mode = securityv1beta1.PeerAuthentication_MutualTLS_Mode(getMTLSModes(v.(string)))
		result[uint32(key)] = objm
	}
	return result
}

func expandMutualTLS(l []interface{}) *securityv1beta1.PeerAuthentication_MutualTLS {
	if len(l) == 0 || l[0] == nil {
		return &securityv1beta1.PeerAuthentication_MutualTLS{}
	}
	in := l[0].(map[string]interface{})
	obj := &securityv1beta1.PeerAuthentication_MutualTLS{}
	if v, ok := in["mode"].(string); ok && v != "" {
		obj.Mode = securityv1beta1.PeerAuthentication_MutualTLS_Mode(getMTLSModes(v))
	}
	return obj
}

func getMTLSModes(s string) (v int32) {
	switch s {
	case "UNSET":
		return 0
	case "DISABLE":
		return 1
	case "PERMISSIVE":
		return 2
	case "STRICT":
		return 3
	}
	return 0
}

func expandLabelSelector(l []interface{}) *v1beta1.WorkloadSelector {
	if len(l) == 0 || l[0] == nil {
		return &v1beta1.WorkloadSelector{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1beta1.WorkloadSelector{}
	if v, ok := in["matchlabels"].(map[string]interface{}); ok && len(v) > 0 {
		obj.MatchLabels = expandStringMap(v)
	}
	return obj
}
