package istio

import (
	"log"

	v1alpha3spec "istio.io/api/networking/v1alpha3"
)

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
		log.Printf("[INFO] Creating expandHTTPMatchRequest : %#v", match)
		obj := v1alpha3spec.HTTPMatchRequest{}
		objs[i] = &obj
	}
	return objs, nil
}

func expandHTTPRouteDestination(httproutedest []interface{}) ([]*v1alpha3spec.HTTPRouteDestination, error) {
	objs := make([]*v1alpha3spec.HTTPRouteDestination, len(httproutedest))
	if len(httproutedest) == 0 || httproutedest[0] == nil {
		return objs, nil
	}
	for i, r := range httproutedest {
		routedest := r.(map[string]interface{})
		log.Printf("[INFO] Creating expandHTTPRouteDestination : %#v", routedest)
		obj := v1alpha3spec.HTTPRouteDestination{}
		objs[i] = &obj
	}
	return objs, nil
}

func expandHTTPRedirect(httpredirect []interface{}) (*v1alpha3spec.HTTPRedirect, error) {
	obj := &v1alpha3spec.HTTPRedirect{}
	return obj, nil
}

func expandDelegate(delegate []interface{}) (*v1alpha3spec.Delegate, error) {
	obj := &v1alpha3spec.Delegate{}
	return obj, nil
}

func expandHTTPRewrite(httprewrite []interface{}) (*v1alpha3spec.HTTPRewrite, error) {
	obj := &v1alpha3spec.HTTPRewrite{}
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
