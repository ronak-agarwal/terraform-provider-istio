package istio

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func destinationFields() map[string]*schema.Schema {
	d := map[string]*schema.Schema{
		"host": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "HTTP route Destination.",
		},
		"subset": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "HTTP route Destination.",
		},
		"port": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "HTTP route Destination.",
		},
	}
	return d
}

func stringMatchFields() map[string]*schema.Schema {
	sm := map[string]*schema.Schema{
		"exact": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "exact.",
		},
		"prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "prefix.",
		},
		"regex": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "regex.",
		},
	}
	return sm
}

func headersFields() map[string]*schema.Schema {
	h := map[string]*schema.Schema{
		"request": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "HTTP route Headers.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"set": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "HTTP route request Headers.",
					},
					"add": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "HTTP route request Headers.",
					},
					"remove": {
						Type:        schema.TypeList,
						Description: "List of request Headers.",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
		"response": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "HTTP route response Headers.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"set": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "HTTP route response Headers.",
					},
					"add": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "HTTP route response Headers.",
					},
					"remove": {
						Type:        schema.TypeList,
						Description: "List of response Headers.",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
	return h
}

func virtualServiceSpecFields() map[string]*schema.Schema {
	vs := map[string]*schema.Schema{
		"hosts": {
			Type:        schema.TypeList,
			Description: "List of hosts.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"gateways": {
			Type:        schema.TypeList,
			Description: "List of gateways.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"http": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify HTTP Routes",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "HTTP route name.",
					},
					"match": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route HTTPMatchRequest.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP HTTPMatchRequest name.",
								},
								"uri": {
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Description: "HTTP route HTTPMatchRequest.",
									Elem: &schema.Resource{
										Schema: stringMatchFields(),
									},
								},
								"scheme": {
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Description: "HTTP route HTTPMatchRequest.",
									Elem: &schema.Resource{
										Schema: stringMatchFields(),
									},
								},
								"method": {
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Description: "HTTP route HTTPMatchRequest.",
									Elem: &schema.Resource{
										Schema: stringMatchFields(),
									},
								},
								"authority": {
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Description: "HTTP route HTTPMatchRequest.",
									Elem: &schema.Resource{
										Schema: stringMatchFields(),
									},
								},
								"headers": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "HTTP match name.",
								},
								"port": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP match name.",
								},
								"sourceLabels": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "HTTP match name.",
								},
								"gateways": {
									Type:        schema.TypeList,
									Description: "List of gateways.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"queryParams": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "HTTP match name.",
								},
								"ignoreUriCase": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "HTTP match name.",
								},
								"withoutHeaders": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "HTTP match name.",
								},
								"sourceNamespace": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP match name.",
								},
							},
						},
					},
					"route": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route HTTPRouteDestination.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"destination": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "HTTP route HTTPRouteDestination.",
									Elem: &schema.Resource{
										Schema: destinationFields(),
									},
								},
								"weight": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP route HTTPRouteDestination.",
								},
								"headers": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "HTTP route HTTPRouteDestination Headers.",
									Elem: &schema.Resource{
										Schema: headersFields(),
									},
								},
							},
						},
					},
					"redirect": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route HTTPRedirect.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"uri": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP route HTTPRedirect.",
								},
								"authority": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP route HTTPRedirect.",
								},
								"redirectCode": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP route HTTPRedirect.",
								},
							},
						},
					},
					"delegate": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route Delegate.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP route Delegate.",
								},
								"namespace": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP route Delegate.",
								},
							},
						},
					},
					"rewrite": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route HTTPRewrite.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"uri": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP route HTTPRewrite.",
								},
								"authority": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP route HTTPRewrite.",
								},
							},
						},
					},
					"timeout": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "HTTP route.",
					},
					"retries": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route HTTPRetry.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"attempts": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP route HTTPRetry.",
								},
								"perTryTimeout": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP route HTTPRetry.",
								},
								"retryOn": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP route HTTPRetry.",
								},
								"retryRemoteLocalities": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "HTTP route HTTPRetry.",
								},
							},
						},
					},
					"fault": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route HTTPFaultInjection.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"delay": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "HTTP route delay HTTPFaultInjection.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"fixedDelay": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "HTTP route delay HTTPFaultInjection.",
											},
											"percentage": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "HTTP route delay HTTPFaultInjection.",
											},
											"percent": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "HTTP route delay HTTPFaultInjection.",
											},
										},
									},
								},
								"abort": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "HTTP route abort HTTPFaultInjection.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"httpStatus": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "HTTP route abort HTTPFaultInjection.",
											},
											"percent": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "HTTP route abort HTTPFaultInjection.",
											},
										},
									},
								},
							},
						},
					},
					"mirror": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route Destination.",
						Elem: &schema.Resource{
							Schema: destinationFields(),
						},
					},
					"mirrorPercentage": {
						Type:        schema.TypeFloat,
						Optional:    true,
						Description: "HTTP route HTTPRoute.",
					},
					"corsPolicy": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route CorsPolicy.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"allowOrigins": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "HTTP route allowOrigins CorsPolicy.",
									Elem: &schema.Resource{
										Schema: stringMatchFields(),
									},
								},
								"allowMethods": {
									Type:        schema.TypeList,
									Description: "List of allowMethods CorsPolicy.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"allowHeaders": {
									Type:        schema.TypeList,
									Description: "List of allowHeaders CorsPolicy.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"exposeHeaders": {
									Type:        schema.TypeList,
									Description: "List of exposeHeaders CorsPolicy.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"maxAge": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP route CorsPolicy.",
								},
								"allowCredentials": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "HTTP route CorsPolicy.",
								},
							},
						},
					},
					"headers": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route Headers.",
						Elem: &schema.Resource{
							Schema: headersFields(),
						},
					},
					"mirrorPercent": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "HTTP route match.",
					},
				},
			},
		},
		"tls": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify TLS",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"match": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify TLS TLSMatchAttributes",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"sniHosts": {
									Type:        schema.TypeList,
									Description: "List of TLSMatchAttributes route sniHosts.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"destinationSubnets": {
									Type:        schema.TypeList,
									Description: "List of TLSMatchAttributes route destinationSubnets.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"port": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Specify TLSMatchAttributes route.",
								},
								"sourceLabels": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "Specify TLSMatchAttributes route.",
								},
								"gateways": {
									Type:        schema.TypeList,
									Description: "List of TLSMatchAttributes route gateways.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"sourceNamespace": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Specify TLSMatchAttributes route.",
								},
							},
						},
					},
					"route": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify TLS RouteDestination",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"destination": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "HTTP TLS RouteDestination.",
									Elem: &schema.Resource{
										Schema: destinationFields(),
									},
								},
								"weight": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Specify TLS RouteDestination route.",
								},
							},
						},
					},
				},
			},
		},
		"tcp": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify TCP",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"match": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify TCP L4MatchAttributes match.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"destinationSubnets": {
									Type:        schema.TypeList,
									Description: "List of TCP L4MatchAttributes match destinationSubnets.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"port": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Specify TCP L4MatchAttributes match.",
								},
								"sourceLabels": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "Specify TCP L4MatchAttributes match",
								},
								"gateways": {
									Type:        schema.TypeList,
									Description: "List of TCP L4MatchAttributes match gateways.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"sourceNamespace": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Specify TLS RouteDestination route.",
								},
							},
						},
					},
					"route": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Specify TCP RouteDestination",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"destination": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "HTTP TCP RouteDestination.",
									Elem: &schema.Resource{
										Schema: destinationFields(),
									},
								},
								"weight": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Specify TCP RouteDestination route.",
								},
							},
						},
					},
				},
			},
		},
		"exportTo": {
			Type:        schema.TypeList,
			Description: "List of exports.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	return vs
}
