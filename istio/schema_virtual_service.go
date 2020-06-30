package istio

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

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
	log.Printf("[INFO] Creating New VirtualService virtualServiceSpecFields")
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
								"sourcelabels": {
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
								"queryparams": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "HTTP match name.",
								},
								"ignoreuricase": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "HTTP match name.",
								},
								"withoutheaders": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "HTTP match name.",
								},
								"sourcenamespace": {
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
								"redirectcode": {
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
								"pertrytimeout": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP route HTTPRetry.",
								},
								"retryon": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP route HTTPRetry.",
								},
								"retryremotelocalities": {
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
											"fixeddelay": {
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
											"httpstatus": {
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
					"mirrorpercentage": {
						Type:        schema.TypeFloat,
						Optional:    true,
						Description: "HTTP route HTTPRoute.",
					},
					"corspolicy": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "HTTP route CorsPolicy.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"alloworigins": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "HTTP route allowOrigins CorsPolicy.",
									Elem: &schema.Resource{
										Schema: stringMatchFields(),
									},
								},
								"allowmethods": {
									Type:        schema.TypeList,
									Description: "List of allowMethods CorsPolicy.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"allowheaders": {
									Type:        schema.TypeList,
									Description: "List of allowHeaders CorsPolicy.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"exposeheaders": {
									Type:        schema.TypeList,
									Description: "List of exposeHeaders CorsPolicy.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"maxage": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP route CorsPolicy.",
								},
								"allowcredentials": {
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
					"mirrorpercent": {
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
								"snihosts": {
									Type:        schema.TypeList,
									Description: "List of TLSMatchAttributes route sniHosts.",
									Optional:    true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"destinationsubnets": {
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
								"sourcelabels": {
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
								"sourcenamespace": {
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
								"destinationsubnets": {
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
								"sourcelabels": {
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
								"sourcenamespace": {
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
		"exportto": {
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
