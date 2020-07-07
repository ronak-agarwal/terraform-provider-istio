package istio

import (
	v1alpha3spec "istio.io/api/networking/v1alpha3"
)

func expandSidecarSpec(sidecar []interface{}) (*v1alpha3spec.Sidecar, error) {
	obj := &v1alpha3spec.Sidecar{}
	if len(sidecar) == 0 || sidecar[0] == nil {
		return obj, nil
	}
	in := sidecar[0].(map[string]interface{})
	if v, ok := in["workloadselector"].(map[string]interface{}); ok && len(v) > 0 {
		if label, ok := v["labels"].(map[string]interface{}); ok && len(label) > 0 {
			objsel := &v1alpha3spec.WorkloadSelector{}
			objsel.Labels = expandStringMap(label)
			obj.WorkloadSelector = objsel
		}
	}
	if v, ok := in["ingress"].([]interface{}); ok && len(v) > 0 {
		ing, err := expandIngress(v)
		if err != nil {
			return obj, err
		}
		obj.Ingress = ing
	}
	if v, ok := in["egress"].([]interface{}); ok && len(v) > 0 {
		eg, err := expandEgress(v)
		if err != nil {
			return obj, err
		}
		obj.Egress = eg
	}
	if tp, ok := in["outboundtrafficpolicy"].([]interface{}); ok && len(tp) > 0 {
		intp := tp[0].(map[string]interface{})
		if m, ok := intp["mode"].(string); ok && m != "" {
			objtp := &v1alpha3spec.OutboundTrafficPolicy{}
			objtp.Mode = v1alpha3spec.OutboundTrafficPolicy_Mode(getOutTrafficPolicyMode(m))
			obj.OutboundTrafficPolicy = objtp
		}
	}
	return obj, nil
}

func getOutTrafficPolicyMode(s string) (v int32) {
	switch s {
	case "REGISTRY_ONLY":
		return 0
	case "ALLOW_ANY":
		return 1
	}
	return 0
}

func expandEgress(egresslistener []interface{}) ([]*v1alpha3spec.IstioEgressListener, error) {
	objs := make([]*v1alpha3spec.IstioEgressListener, len(egresslistener))
	if len(egresslistener) == 0 || egresslistener[0] == nil {
		return objs, nil
	}
	for i, s := range egresslistener {
		eg := s.(map[string]interface{})
		obj := v1alpha3spec.IstioEgressListener{}
		if port, ok := eg["port"].([]interface{}); ok && len(port) > 0 {
			p, err := expandPort(port)
			if err != nil {
				return objs, err
			}
			obj.Port = p
		}
		if bind, ok := eg["bind"].(string); ok && bind != "" {
			obj.Bind = bind
		}
		if cm, ok := eg["capturemode"].(string); ok && cm != "" {
			obj.CaptureMode = v1alpha3spec.CaptureMode(getCaptureMode(cm))
		}
		if v, ok := eg["hosts"].([]interface{}); ok && len(v) > 0 {
			obj.Hosts = expandStringList(v)
		}
		objs[i] = &obj
	}
	return objs, nil
}

func expandIngress(ingresslistener []interface{}) ([]*v1alpha3spec.IstioIngressListener, error) {
	objs := make([]*v1alpha3spec.IstioIngressListener, len(ingresslistener))
	if len(ingresslistener) == 0 || ingresslistener[0] == nil {
		return objs, nil
	}
	for i, s := range ingresslistener {
		ing := s.(map[string]interface{})
		obj := v1alpha3spec.IstioIngressListener{}
		if port, ok := ing["port"].([]interface{}); ok && len(port) > 0 {
			p, err := expandPort(port)
			if err != nil {
				return objs, err
			}
			obj.Port = p
		}
		if bind, ok := ing["bind"].(string); ok && bind != "" {
			obj.Bind = bind
		}
		if cm, ok := ing["capturemode"].(string); ok && cm != "" {
			obj.CaptureMode = v1alpha3spec.CaptureMode(getCaptureMode(cm))
		}
		if dep, ok := ing["defaultendpoint"].(string); ok && dep != "" {
			obj.DefaultEndpoint = dep
		}
		objs[i] = &obj
	}
	return objs, nil
}

func getCaptureMode(s string) (v int32) {
	switch s {
	case "DEFAULT":
		return 0
	case "IPTABLES":
		return 1
	case "NONE":
		return 2
	}
	return 0
}

func expandPort(port []interface{}) (*v1alpha3spec.Port, error) {
	obj := &v1alpha3spec.Port{}
	if len(port) == 0 || port[0] == nil {
		return obj, nil
	}
	in := port[0].(map[string]interface{})
	if v, ok := in["number"].(int); ok && v > 0 {
		obj.Number = uint32(v)
	}
	if v, ok := in["protocol"]; ok && v != "" {
		obj.Protocol = v.(string)
	}
	if v, ok := in["name"]; ok && v != "" {
		obj.Name = v.(string)
	}
	return obj, nil
}
