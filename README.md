# terraform-provider-istio

This plugin supports istio 1.6.x

STATUS - In Development

## Requirements
1. You need Istio to be installed on your Kubernetes cluster (in istio-system namespace) using Operator https://istio.io/latest/docs/setup/install/standalone-operator/ (Although this plugin can be extended to Istio on VMs but need to be tested)
2. Terraform 0.12.x

## Description
Using this plugin you can perform regular terraform operation to create, delete, update, read different resources of Istio:
1. NetworkingV1alpha3 (VirtualService, DestinationRule, Sidecar, ServiceEntry, Gateway, Envoyfilter, Workloadentry)
2. SecurityV1beta1 (AuthorizationPolicy, PeerAuthentication, RequestAuthentication)

## Usage

1. VirtualService

-- In Development --

```hcl
provider istio {
  config_path = "/Users/ronagarw/.kube/config"
}

resource "istio_virtual_service" "example"{
  metadata {
    name = "terraform-example"
    namespace = "test"
  }

  spec {
    hosts = [
      "reviews.prod.svc.cluster.local",
    ]
    http  {
        name = "reviews-v2-routes"
        match  {
            uri {
              prefix = "/wpcatalog"
            }
          }
        match  {
            uri {
              prefix = "/consumercatalog"
            }
          }
        rewrite {
          uri = "/newcatalog"
        }
        route  {
            destination {
              host = "reviews.prod.svc.cluster.local"
              subset = "v2"
            }
          }
      }
    http  {
        name = "reviews-v1-route"
        route  {
            destination {
              host = "reviews.prod.svc.cluster.local"
              subset = "v1"
            }
          }
       }
    }
}
```

##Contributing to the provider
Appreciate your contribution to the provider
