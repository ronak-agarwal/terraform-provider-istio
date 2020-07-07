package istio

import (
	v1alpha3spec "istio.io/api/networking/v1alpha3"
)

func expandPorts(ports []interface{}) ([]*v1alpha3spec.Port, error) {
	objs := make([]*v1alpha3spec.Port, len(ports))
	if len(ports) == 0 || ports[0] == nil {
		return objs, nil
	}
	for i, s := range ports {
		in := s.(map[string]interface{})
		obj := v1alpha3spec.Port{}
		if v, ok := in["number"].(int); ok && v > 0 {
			obj.Number = uint32(v)
		}
		if v, ok := in["protocol"]; ok && v != "" {
			obj.Protocol = v.(string)
		}
		if v, ok := in["name"]; ok && v != "" {
			obj.Name = v.(string)
		}
		objs[i] = &obj
	}
	return objs, nil
}

func expandServiceEntrySpec(serviceentry []interface{}) (*v1alpha3spec.ServiceEntry, error) {
	obj := &v1alpha3spec.ServiceEntry{}
	if len(serviceentry) == 0 || serviceentry[0] == nil {
		return obj, nil
	}
	in := serviceentry[0].(map[string]interface{})
	if v, ok := in["hosts"].([]interface{}); ok && len(v) > 0 {
		obj.Hosts = expandStringList(v)
	}
	if v, ok := in["addresses"].([]interface{}); ok && len(v) > 0 {
		obj.Addresses = expandStringList(v)
	}
	if port, ok := in["ports"].([]interface{}); ok && len(port) > 0 {
		p, err := expandPorts(port)
		if err != nil {
			return obj, err
		}
		obj.Ports = p
	}
	if loc, ok := in["location"].(string); ok && loc != "" {
		obj.Location = v1alpha3spec.ServiceEntry_Location(getServiceEntryLoc(loc))
	}
	if v, ok := in["resolution"].(string); ok && v != "" {
		obj.Resolution = v1alpha3spec.ServiceEntry_Resolution(getServiceEntryResolution(v))
	}
	if v, ok := in["endpoints"].([]interface{}); ok && len(v) > 0 {
		ep, err := expandWorkloadEntry(v)
		if err != nil {
			return obj, err
		}
		obj.Endpoints = ep
	}
	if v, ok := in["exportto"].([]interface{}); ok && len(v) > 0 {
		obj.ExportTo = expandStringList(v)
	}
	if v, ok := in["subjectaltnames"].([]interface{}); ok && len(v) > 0 {
		obj.SubjectAltNames = expandStringList(v)
	}
	return obj, nil
}

func expandWorkloadEntry(workloadentry []interface{}) ([]*v1alpha3spec.WorkloadEntry, error) {
	objs := make([]*v1alpha3spec.WorkloadEntry, len(workloadentry))
	if len(workloadentry) == 0 || workloadentry[0] == nil {
		return objs, nil
	}
	for i, s := range workloadentry {
		w := s.(map[string]interface{})
		obj := v1alpha3spec.WorkloadEntry{}
		if v, ok := w["address"].(string); ok && v != "" {
			obj.Address = v
		}
		if v, ok := w["ports"].(map[string]interface{}); ok && len(v) > 0 {
			obj.Ports = expandStringUintMap(v)
		}
		if v, ok := w["labels"].(map[string]interface{}); ok && len(v) > 0 {
			obj.Labels = expandStringMap(v)
		}
		if v, ok := w["network"].(string); ok && v != "" {
			obj.Network = v
		}
		if v, ok := w["locality"].(string); ok && v != "" {
			obj.Locality = v
		}
		if v, ok := w["weight"].(int); ok && v > 0 {
			obj.Weight = uint32(v)
		}
		if v, ok := w["serviceaccount"].(string); ok && v != "" {
			obj.ServiceAccount = v
		}
		objs[i] = &obj
	}
	return objs, nil
}

func getServiceEntryLoc(s string) (v int32) {
	switch s {
	case "MESH_EXTERNAL":
		return 0
	case "MESH_INTERNAL":
		return 1
	}
	return 0
}

func getServiceEntryResolution(s string) (v int32) {
	switch s {
	case "NONE":
		return 0
	case "STATIC":
		return 1
	case "DNS":
		return 2
	}
	return 0
}
