package istio

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func loadbalancerFields() map[string]*schema.Schema {
	lb := map[string]*schema.Schema{
		"simple": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination route.",
		},
		"consistenthash": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"httpheadername": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Destination route.",
					},
					"httpcookie": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Destination route.",
								},
								"path": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Destination route.",
								},
								"ttl": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Destination route.",
								},
							},
						},
					},
					"usesourceip": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Destination route.",
					},
					"httpqueryparametername": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Destination route.",
					},
					"minimumringsize": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Destination route.",
					},
				},
			},
		},
		"localitylbsetting": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"distribute": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"from": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Destination route.",
								},
								"to": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "Destination route.",
								},
							},
						},
					},
					"failover": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"from": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Destination route.",
								},
								"to": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Destination route.",
								},
							},
						},
					},
					"enabled": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Destination route.",
					},
				},
			},
		},
	}
	return lb
}

func trafficPolicyFields() map[string]*schema.Schema {
	tp := map[string]*schema.Schema{
		"loadbalancer": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: loadbalancerFields(),
			},
		},
		"connectionpool": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: connectionPoolFields(),
			},
		},
		"outlierdetection": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: outlierDetectionFields(),
			},
		},
		"tls": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: tlsSettingsFields(),
			},
		},
		"portlevelsettings": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"port": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"number": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Destination route.",
								},
							},
						},
					},
					"loadbalancer": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: loadbalancerFields(),
						},
					},
					"connectionpool": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: connectionPoolFields(),
						},
					},
					"outlierdetection": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: outlierDetectionFields(),
						},
					},
					"tls": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: tlsSettingsFields(),
						},
					},
				},
			},
		},
	}
	return tp
}

func connectionPoolFields() map[string]*schema.Schema {
	cp := map[string]*schema.Schema{
		"tcp": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"maxconnections": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Destination route.",
					},
					"connecttimeout": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Destination route.",
					},
					"tcpkeepalive": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"probes": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Destination route.",
								},
								"time": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Destination route.",
								},
								"interval": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Destination route.",
								},
							},
						},
					},
				},
			},
		},
		"http": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"http1maxpendingrequests": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Destination route.",
					},
					"http2maxrequests": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Destination route.",
					},
					"maxrequestsperconnection": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Destination route.",
					},
					"maxretries": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Destination route.",
					},
					"idletimeout": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Destination route.",
					},
					"h2upgradepolicy": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Destination route.",
								},
							},
						},
					},
				},
			},
		},
	}
	return cp
}

func outlierDetectionFields() map[string]*schema.Schema {

	od := map[string]*schema.Schema{
		"consecutivegatewayerrors": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Destination route.",
		},
		"consecutive5xxerrors": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Destination route.",
		},
		"interval": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Destination route.",
		},
		"baseejectiontime": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Destination route.",
		},
		"maxejectionpercent": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Destination route.",
		},
		"minhealthpercent": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Destination route.",
		},
	}
	return od
}

func tlsSettingsFields() map[string]*schema.Schema {
	tls := map[string]*schema.Schema{
		"mode": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination route.",
		},
		"clientcertificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination route.",
		},
		"privatekey": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination route.",
		},
		"cacertificates": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination route.",
		},
		"subjectaltnames": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination route.",
		},
		"sni": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination route.",
		},
	}
	return tls
}

func destinationRuleSpecFields() map[string]*schema.Schema {
	log.Printf("[INFO] Creating New destinationRule destinationRuleSpecFields")
	dr := map[string]*schema.Schema{
		"host": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination route.",
		},
		"trafficpolicy": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: trafficPolicyFields(),
			},
		},
		"subsets": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify Destination Rule trafficpolicy",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Destination route.",
					},
					"labels": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "Destination route.",
					},
					"trafficpolicy": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Specify Destination Rule trafficpolicy",
						Elem: &schema.Resource{
							Schema: trafficPolicyFields(),
						},
					},
				},
			},
		},
		"exportto": {
			Type:        schema.TypeList,
			Description: "List of exports.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
	return dr
}
