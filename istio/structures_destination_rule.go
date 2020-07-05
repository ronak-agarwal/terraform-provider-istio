package istio

import (
	"github.com/gogo/protobuf/types"
	v1alpha3spec "istio.io/api/networking/v1alpha3"
)

func expandDestinationRulesSpec(destinationrule []interface{}) (*v1alpha3spec.DestinationRule, error) {
	obj := &v1alpha3spec.DestinationRule{}
	if len(destinationrule) == 0 || destinationrule[0] == nil {
		return obj, nil
	}
	in := destinationrule[0].(map[string]interface{})
	if host, ok := in["host"].(string); ok && in["host"] != "" {
		obj.Host = host
	}
	if v, ok := in["trafficpolicy"].([]interface{}); ok && len(v) > 0 {
		tp, err := expandTrafficPolicy(v)
		if err != nil {
			return obj, err
		}
		obj.TrafficPolicy = tp
	}
	if v, ok := in["subsets"].([]interface{}); ok && len(v) > 0 {
		subs, err := expandSubsets(v)
		if err != nil {
			return obj, err
		}
		obj.Subsets = subs
	}
	if v, ok := in["exportto"].([]interface{}); ok && len(v) > 0 {
		obj.ExportTo = expandStringList(v)
	}

	return obj, nil
}

func expandTrafficPolicy(trafficpolicy []interface{}) (*v1alpha3spec.TrafficPolicy, error) {
	obj := &v1alpha3spec.TrafficPolicy{}
	if len(trafficpolicy) == 0 || trafficpolicy[0] == nil {
		return obj, nil
	}
	tp := trafficpolicy[0].(map[string]interface{})
	if lb, ok := tp["loadbalancer"].([]interface{}); ok && len(lb) > 0 {
		lber, err := expandLoadBalancer(lb)
		if err != nil {
			return obj, err
		}
		obj.LoadBalancer = lber
	}
	if cp, ok := tp["connectionpool"].([]interface{}); ok && len(cp) > 0 {
		cpool, err := expandConnectionPool(cp)
		if err != nil {
			return obj, err
		}
		obj.ConnectionPool = cpool
	}
	if od, ok := tp["outlierdetection"].([]interface{}); ok && len(od) > 0 {
		odtect, err := expandOutlierDetection(od)
		if err != nil {
			return obj, err
		}
		obj.OutlierDetection = odtect
	}
	if tls, ok := tp["tls"].([]interface{}); ok && len(tls) > 0 {
		tl, err := expandTLS(tls)
		if err != nil {
			return obj, err
		}
		obj.Tls = tl
	}
	if pset, ok := tp["portlevelsettings"].([]interface{}); ok && len(pset) > 0 {
		pls, err := expandPortLevelSettings(pset)
		if err != nil {
			return obj, err
		}
		obj.PortLevelSettings = pls
	}
	return obj, nil
}

func getSimpleLB(s string) (v int32) {
	switch s {
	case "ROUND_ROBIN":
		return 0
	case "LEAST_CONN":
		return 1
	case "RANDOM":
		return 2
	case "PASSTHROUGH":
		return 3
	}
	return 0
}

func expandLoadBalancer(loadbalancer []interface{}) (*v1alpha3spec.LoadBalancerSettings, error) {
	obj := &v1alpha3spec.LoadBalancerSettings{}
	if len(loadbalancer) == 0 || loadbalancer[0] == nil {
		return obj, nil
	}
	lb := loadbalancer[0].(map[string]interface{})
	if simple, ok := lb["simple"].(string); ok && lb["simple"] != "" {
		objsimple := &v1alpha3spec.LoadBalancerSettings_Simple{}
		objsimple.Simple = v1alpha3spec.LoadBalancerSettings_SimpleLB(getSimpleLB(simple))
		obj.LbPolicy = objsimple
	}
	if hash, ok := lb["consistenthash"].([]interface{}); ok && len(hash) > 0 {
		objhash := &v1alpha3spec.LoadBalancerSettings_ConsistentHash{}
		objhashLB := &v1alpha3spec.LoadBalancerSettings_ConsistentHashLB{}
		h := hash[0].(map[string]interface{})
		//HttpHeaderName
		if hdname, ok := h["httpheadername"]; ok {
			objhashheader := &v1alpha3spec.LoadBalancerSettings_ConsistentHashLB_HttpHeaderName{}
			objhashheader.HttpHeaderName = hdname.(string)
			objhashLB.HashKey = objhashheader
		}
		//HttpCookie
		if hcookie, ok := h["httpcookie"].([]interface{}); ok && len(hcookie) > 0 {
			objhashcookie := &v1alpha3spec.LoadBalancerSettings_ConsistentHashLB_HttpCookie{}
			c := hcookie[0].(map[string]interface{})
			ttl := types.Duration{Nanos: c["ttl"].(int32)}
			a := v1alpha3spec.LoadBalancerSettings_ConsistentHashLB_HTTPCookie{
				Name: c["name"].(string),
				Path: c["path"].(string),
				Ttl:  &ttl,
			}
			objhashcookie.HttpCookie = &a
			objhashLB.HashKey = objhashcookie
		}
		//UseSourceIp
		if sip, ok := h["usesourceip"]; ok && h["usesourceip"] != "" {
			objsiplb := &v1alpha3spec.LoadBalancerSettings_ConsistentHashLB_UseSourceIp{}
			objsiplb.UseSourceIp = sip.(bool)
			objhashLB.HashKey = objsiplb
		}
		//HttpQueryParameterName
		if hqpn, ok := h["httpqueryparametername"]; ok && h["httpqueryparametername"] != "" {
			objhqpn := &v1alpha3spec.LoadBalancerSettings_ConsistentHashLB_HttpQueryParameterName{}
			objhqpn.HttpQueryParameterName = hqpn.(string)
			objhashLB.HashKey = objhqpn
		}
		objhash.ConsistentHash = objhashLB
		obj.LbPolicy = objhash
	}
	if lbs, ok := lb["localitylbsetting"].([]interface{}); ok && len(lbs) > 0 {
		objlbsetting := &v1alpha3spec.LocalityLoadBalancerSetting{}

		obj.LocalityLbSetting = objlbsetting
	}
	return obj, nil
}

func expandConnectionPool(connectionpool []interface{}) (*v1alpha3spec.ConnectionPoolSettings, error) {
	obj := &v1alpha3spec.ConnectionPoolSettings{}
	return obj, nil
}

func expandOutlierDetection(outlierdetection []interface{}) (*v1alpha3spec.OutlierDetection, error) {
	obj := &v1alpha3spec.OutlierDetection{}
	return obj, nil
}

func expandTLS(tls []interface{}) (*v1alpha3spec.ClientTLSSettings, error) {
	obj := &v1alpha3spec.ClientTLSSettings{}
	return obj, nil
}

func expandPortLevelSettings(portlevelsettings []interface{}) ([]*v1alpha3spec.TrafficPolicy_PortTrafficPolicy, error) {
	objs := make([]*v1alpha3spec.TrafficPolicy_PortTrafficPolicy, len(portlevelsettings))
	if len(portlevelsettings) == 0 || portlevelsettings[0] == nil {
		return objs, nil
	}
	return objs, nil
}

func expandSubsets(subsets []interface{}) ([]*v1alpha3spec.Subset, error) {
	objs := make([]*v1alpha3spec.Subset, len(subsets))
	if len(subsets) == 0 || subsets[0] == nil {
		return objs, nil
	}
	for i, s := range subsets {
		subset := s.(map[string]interface{})
		obj := v1alpha3spec.Subset{}
		if name, ok := subset["name"]; ok {
			obj.Name = name.(string)
		}
		if label, ok := subset["labels"].(map[string]interface{}); ok && len(label) > 0 {
			obj.Labels = expandStringMap(label)
		}
		if v, ok := subset["trafficpolicy"].([]interface{}); ok && len(v) > 0 {
			tp, err := expandTrafficPolicy(v)
			if err != nil {
				return objs, err
			}
			obj.TrafficPolicy = tp
		}
		objs[i] = &obj
	}
	return objs, nil
}
