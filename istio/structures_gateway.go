package istio

import (
	v1alpha3spec "istio.io/api/networking/v1alpha3"
)

func expandGatewaySpec(gateway []interface{}) (*v1alpha3spec.Gateway, error) {
	obj := &v1alpha3spec.Gateway{}
	if len(gateway) == 0 || gateway[0] == nil {
		return obj, nil
	}
	in := gateway[0].(map[string]interface{})
	if v, ok := in["servers"].([]interface{}); ok && len(v) > 0 {
		s, err := expandServers(v)
		if err != nil {
			return obj, err
		}
		obj.Servers = s
	}
	if v, ok := in["selector"].(map[string]interface{}); ok && len(v) > 0 {
		obj.Selector = expandStringMap(v)
	}
	return obj, nil
}

func expandServers(servers []interface{}) ([]*v1alpha3spec.Server, error) {
	objs := make([]*v1alpha3spec.Server, len(servers))
	if len(servers) == 0 || servers[0] == nil {
		return objs, nil
	}
	for i, s := range servers {
		in := s.(map[string]interface{})
		obj := v1alpha3spec.Server{}
		if port, ok := in["port"].([]interface{}); ok && len(port) > 0 {
			p, err := expandPort(port)
			if err != nil {
				return objs, err
			}
			obj.Port = p
		}
		if v, ok := in["hosts"].([]interface{}); ok && len(v) > 0 {
			obj.Hosts = expandStringList(v)
		}
		if v, ok := in["tls"].([]interface{}); ok && len(v) > 0 {
			tlsSet, err := expandServerTLSSettings(v)
			if err != nil {
				return objs, err
			}
			obj.Tls = tlsSet
		}
		objs[i] = &obj
	}
	return objs, nil
}

func expandServerTLSSettings(servertlsettings []interface{}) (*v1alpha3spec.ServerTLSSettings, error) {
	obj := &v1alpha3spec.ServerTLSSettings{}
	if len(servertlsettings) == 0 || servertlsettings[0] == nil {
		return obj, nil
	}
	in := servertlsettings[0].(map[string]interface{})
	if v, ok := in["httpsredirect"].(bool); ok {
		obj.HttpsRedirect = v
	}
	if v, ok := in["mode"].(string); ok && v != "" {
		obj.Mode = v1alpha3spec.ServerTLSSettings_TLSmode(getTLSMode(v))
	}
	if v, ok := in["servercertificate"].(string); ok && v != "" {
		obj.ServerCertificate = v
	}
	if v, ok := in["privatekey"].(string); ok && v != "" {
		obj.PrivateKey = v
	}
	if v, ok := in["cacertificates"].(string); ok && v != "" {
		obj.CaCertificates = v
	}
	if v, ok := in["credentialname"].(string); ok && v != "" {
		obj.CredentialName = v
	}
	if v, ok := in["subjectaltnames"].([]interface{}); ok && len(v) > 0 {
		obj.SubjectAltNames = expandStringList(v)
	}
	if v, ok := in["verifycertificatespki"].([]interface{}); ok && len(v) > 0 {
		obj.VerifyCertificateSpki = expandStringList(v)
	}
	if v, ok := in["verifycertificatehash"].([]interface{}); ok && len(v) > 0 {
		obj.VerifyCertificateHash = expandStringList(v)
	}
	if v, ok := in["minprotocolversion"].(string); ok && v != "" {
		obj.MinProtocolVersion = v1alpha3spec.ServerTLSSettings_TLSProtocol(getTLSProtocol(v))
	}
	if v, ok := in["maxprotocolversion"].(string); ok && v != "" {
		obj.MaxProtocolVersion = v1alpha3spec.ServerTLSSettings_TLSProtocol(getTLSProtocol(v))
	}
	if v, ok := in["ciphersuites"].([]interface{}); ok && len(v) > 0 {
		obj.CipherSuites = expandStringList(v)
	}
	return obj, nil
}

func getTLSMode(s string) (v int32) {
	switch s {
	case "PASSTHROUGH":
		return 0
	case "SIMPLE":
		return 1
	case "MUTUAL":
		return 2
	case "AUTO_PASSTHROUGH":
		return 3
	case "ISTIO_MUTUAL":
		return 4
	}
	return 0
}

func getTLSProtocol(s string) (v int32) {
	switch s {
	case "TLS_AUTO":
		return 0
	case "TLSV1_0":
		return 1
	case "TLSV1_1":
		return 2
	case "TLSV1_2":
		return 3
	case "TLSV1_3":
		return 4
	}
	return 0
}
