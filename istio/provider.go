package istio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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
			"istio_virtual_service":  resourceVirtualService(),
			"istio_destination_rule": resourceDestinationRule(),
			"istio_sidecar":          resourceSidecar(),
			"istio_service_entry":    resourceServiceEntry(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// Config ...
type Config struct {
	istioClientset *istioclient.Clientset
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	//var data []byte
	var config *rest.Config
	var err error
	path := d.Get("config_path").(string)
	config, err = clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		// if neither worked we fall back to an empty default config
		config = &rest.Config{}
	}

	clientset, err := istioclient.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &Config{clientset}, nil
}
