package istio

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	v1alpha3spec "istio.io/api/networking/v1alpha3"
)

func flattenVirtualServiceSpec(in v1alpha3spec.VirtualService, d *schema.ResourceData) ([]interface{}, error) {
	att := make(map[string]interface{})
	return []interface{}{att}, nil
}

func expandVirtualServiceSpec(virtualservice []interface{}) (*v1alpha3spec.VirtualService, error) {
	obj := &v1alpha3spec.VirtualService{}
	if len(virtualservice) == 0 || virtualservice[0] == nil {
		return obj, nil
	}
	//	log.Printf("[INFO] Creating New VirtualService expandVirtualServiceSpec obj1: %#v", obj)
	in := virtualservice[0].(map[string]interface{})
	if v, ok := in["hosts"].([]interface{}); ok && len(v) > 0 {
		obj.Hosts = expandStringList(v)
	}
	//	log.Printf("[INFO] Creating New VirtualService expandVirtualServiceSpec obj2: %#v", obj)
	if v, ok := in["http"].([]interface{}); ok && len(v) > 0 {
		httproute, err := expandHTTPRoute(v)
		if err != nil {
			return obj, err
		}
		obj.Http = httproute
	}
	return obj, nil
}

func expandHTTPRoute(httproutes []interface{}) ([]*v1alpha3spec.HTTPRoute, error) {
	objs := make([]*v1alpha3spec.HTTPRoute, len(httproutes))
	if len(httproutes) == 0 || httproutes[0] == nil {
		return objs, nil
	}
	//	log.Printf("[INFO] Creating New VirtualService expandHTTPRoute obj: %#v", objs)
	for i, r := range httproutes {
		route := r.(map[string]interface{})
		obj := v1alpha3spec.HTTPRoute{}
		if name, ok := route["name"]; ok {
			obj.Name = name.(string)
		}
		if match, ok := route["match"].([]interface{}); ok && len(match) > 0 {
			httpmatch, err := expandHTTPMatchRequest(match)
			if err != nil {
				return objs, err
			}
			obj.Match = httpmatch
		}
		if route, ok := route["route"].([]interface{}); ok && len(route) > 0 {
			httproute, err := expandHTTPRouteDestination(route)
			if err != nil {
				return objs, err
			}
			obj.Route = httproute
		}
		if redirect, ok := route["redirect"].([]interface{}); ok && len(redirect) > 0 {
			httpredirect, err := expandHTTPRedirect(redirect)
			if err != nil {
				return objs, err
			}
			obj.Redirect = httpredirect
		}
		if delegate, ok := route["delegate"].([]interface{}); ok && len(delegate) > 0 {
			httpdelegate, err := expandDelegate(delegate)
			if err != nil {
				return objs, err
			}
			obj.Delegate = httpdelegate
		}
		if rewrite, ok := route["rewrite"].([]interface{}); ok && len(rewrite) > 0 {
			httprewrite, err := expandHTTPRewrite(rewrite)
			if err != nil {
				return objs, err
			}
			obj.Rewrite = httprewrite
		}
		if retries, ok := route["retries"].([]interface{}); ok && len(retries) > 0 {
			httpretries, err := expandHTTPRetry(retries)
			if err != nil {
				return objs, err
			}
			obj.Retries = httpretries
		}
		if fault, ok := route["fault"].([]interface{}); ok && len(fault) > 0 {
			httpfault, err := expandHTTPFaultInjection(fault)
			if err != nil {
				return objs, err
			}
			obj.Fault = httpfault
		}
		if mirror, ok := route["mirror"].([]interface{}); ok && len(mirror) > 0 {
			httpmirror, err := expandDestination(mirror)
			if err != nil {
				return objs, err
			}
			obj.Mirror = httpmirror
		}
		if corspolicy, ok := route["corspolicy"].([]interface{}); ok && len(corspolicy) > 0 {
			httpcorspolicy, err := expandCorsPolicy(corspolicy)
			if err != nil {
				return objs, err
			}
			obj.CorsPolicy = httpcorspolicy
		}
		log.Printf("[INFO] Creating New VirtualService expandHTTPRoute obj: %#v", obj)
		objs[i] = &obj
	}
	//	log.Printf("[INFO] Creating New VirtualService expandHTTPRoute obj end: %#v", objs)
	return objs, nil
}

func expandHTTPMatchRequest(httpmatch []interface{}) ([]*v1alpha3spec.HTTPMatchRequest, error) {
	objs := make([]*v1alpha3spec.HTTPMatchRequest, len(httpmatch))
	if len(httpmatch) == 0 || httpmatch[0] == nil {
		return objs, nil
	}
	for i, r := range httpmatch {
		match := r.(map[string]interface{})
		//log.Printf("[INFO] Creating expandHTTPMatchRequest : %#v", match)
		obj := v1alpha3spec.HTTPMatchRequest{}
		if name, ok := match["name"]; ok {
			obj.Name = name.(string)
		}
		if uri, ok := match["uri"].([]interface{}); ok && len(uri) > 0 {
			stringmatchuri, err := expandStringMatch(uri)
			if err != nil {
				return objs, err
			}
			obj.Uri = stringmatchuri
		}
		objs[i] = &obj
	}
	return objs, nil
}

func expandStringMatch(stringmatchuri []interface{}) (*v1alpha3spec.StringMatch, error) {
	obj := &v1alpha3spec.StringMatch{}
	if len(stringmatchuri) == 0 || stringmatchuri[0] == nil {
		return obj, nil
	}
	uri := stringmatchuri[0].(map[string]interface{})
	if prefix, ok := uri["prefix"]; ok && uri["prefix"] != "" {
		objprefix := &v1alpha3spec.StringMatch_Prefix{}
		objprefix.Prefix = prefix.(string)
		obj.MatchType = objprefix
	}
	if exact, ok := uri["exact"]; ok && uri["exact"] != "" {
		objexact := &v1alpha3spec.StringMatch_Exact{}
		objexact.Exact = exact.(string)
		obj.MatchType = objexact
	}
	if regex, ok := uri["regex"]; ok && uri["regex"] != "" {
		objregex := &v1alpha3spec.StringMatch_Regex{}
		objregex.Regex = regex.(string)
		obj.MatchType = objregex
	}
	return obj, nil
}

func expandHTTPRouteDestination(httproutedest []interface{}) ([]*v1alpha3spec.HTTPRouteDestination, error) {
	objs := make([]*v1alpha3spec.HTTPRouteDestination, len(httproutedest))
	if len(httproutedest) == 0 || httproutedest[0] == nil {
		return objs, nil
	}
	for i, r := range httproutedest {
		routedest := r.(map[string]interface{})
		obj := v1alpha3spec.HTTPRouteDestination{}
		if dest, ok := routedest["destination"].([]interface{}); ok && len(dest) > 0 {
			destobj, err := expandDestination(dest)
			if err != nil {
				return objs, err
			}
			obj.Destination = destobj
		}
		objs[i] = &obj
	}
	return objs, nil
}

func expandHTTPRedirect(httpredirect []interface{}) (*v1alpha3spec.HTTPRedirect, error) {
	obj := &v1alpha3spec.HTTPRedirect{}
	if len(httpredirect) == 0 || httpredirect[0] == nil {
		return obj, nil
	}
	redirect := httpredirect[0].(map[string]interface{})
	if uri, ok := redirect["uri"]; ok && redirect["uri"] != "" {
		obj.Uri = uri.(string)
	}
	if authority, ok := redirect["authority"]; ok && redirect["authority"] != "" {
		obj.Authority = authority.(string)
	}
	return obj, nil
}

func expandDelegate(delegate []interface{}) (*v1alpha3spec.Delegate, error) {
	obj := &v1alpha3spec.Delegate{}
	return obj, nil
}

func expandHTTPRewrite(httprewrite []interface{}) (*v1alpha3spec.HTTPRewrite, error) {
	obj := &v1alpha3spec.HTTPRewrite{}
	if len(httprewrite) == 0 || httprewrite[0] == nil {
		return obj, nil
	}
	rewrite := httprewrite[0].(map[string]interface{})
	if uri, ok := rewrite["uri"]; ok && rewrite["uri"] != "" {
		obj.Uri = uri.(string)
	}
	if authority, ok := rewrite["authority"]; ok && rewrite["authority"] != "" {
		obj.Authority = authority.(string)
	}
	return obj, nil
}

func expandHTTPRetry(httpretry []interface{}) (*v1alpha3spec.HTTPRetry, error) {
	obj := &v1alpha3spec.HTTPRetry{}
	return obj, nil
}

func expandHTTPFaultInjection(httpfaultinject []interface{}) (*v1alpha3spec.HTTPFaultInjection, error) {
	obj := &v1alpha3spec.HTTPFaultInjection{}
	return obj, nil
}

func expandDestination(destination []interface{}) (*v1alpha3spec.Destination, error) {
	obj := &v1alpha3spec.Destination{}
	if len(destination) == 0 || destination[0] == nil {
		return obj, nil
	}
	dest := destination[0].(map[string]interface{})
	log.Printf("[INFO] Creating expandDestination dest : %#v", dest)
	if host, ok := dest["host"]; ok && dest["host"] != "" {
		obj.Host = host.(string)
	}
	if subset, ok := dest["subset"]; ok && dest["subset"] != "" {
		obj.Subset = subset.(string)
	}
	if port, ok := dest["port"].(int); ok && dest["port"].(int) > 0 {
		log.Printf("[INFO] Creating expandDestination port : %#v", port)
		objps := &v1alpha3spec.PortSelector{}
		objps.Number = uint32(port)
		obj.Port = objps
	}
	return obj, nil
}

func expandPercent(percent []interface{}) (*v1alpha3spec.Percent, error) {
	obj := &v1alpha3spec.Percent{}
	return obj, nil
}

func expandCorsPolicy(corspolicy []interface{}) (*v1alpha3spec.CorsPolicy, error) {
	obj := &v1alpha3spec.CorsPolicy{}
	return obj, nil
}
