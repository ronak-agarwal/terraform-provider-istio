package istio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider ...
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"config_path": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc(
					[]string{
						"KUBE_CONFIG",
						"KUBECONFIG",
					},
					"~/.kube/config"),
				Description: "Path to the kube config file, defaults to ~/.kube/config",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"istio_virtual_service": resourceVirtualService(),
		},
		ConfigureFunc: providerConfigure,
	}
}
