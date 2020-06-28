package istio

import (
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
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
			"istio_virtual_service": resourceVirtualService(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// Config ...
type Config struct {
	Client    dynamic.Interface
	Clientset *kubernetes.Clientset
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var data []byte
	var config *rest.Config
	var err error
	path := d.Get("config_path").(string)
	data, _ = readKubeconfigFile(path)
	config, err = clientcmd.RESTConfigFromKubeConfig(data)
	if err != nil {
		// if neither worked we fall back to an empty default config
		config = &rest.Config{}
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &Config{client, clientset}, nil
}

func readKubeconfigFile(s string) ([]byte, error) {
	p, err := homedir.Expand(s)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	return data, nil
}
