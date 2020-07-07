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

2. DestinationRule

-- In Development --

```hcl
resource "istio_destination_rule" "test" {
  metadata {
    name = "terraform-example"
    namespace = "test"
  }
  spec {
    host = "ratings.prod.svc.cluster.local"
    trafficpolicy {
      loadbalancer {
        simple = "LEAST_CONN"
      }
    }
    subsets {
        name = "testversion"
        labels = {
          version = "v3"
        }
        trafficpolicy {
          loadbalancer {
            simple = "ROUND_ROBIN"
          }
        }
      }
   }
}
```

3. Sidecar

-- In Development --

```hcl
resource "istio_sidecar" "test" {
metadata {
  name = "terraform-example"
  namespace = "test"
 }
spec {
  workloadselector {
      labels = {
        app = "productpage"
      }
    }
  ingress {
      bind = "172.16.1.32"
      port {
          number = 80
          protocol = "HTTP"
          name = "somename"
      }
      defaultendpoint = "127.0.0.1:8080"
      capturemode = "NONE"
    }
  egress {
      capturemode = "IPTABLES"
      hosts = ["*/*"]
    }
  }
}
```

3. ServiceEntry

-- Yet to start --

```hcl
resource "istio_service_entry" "test" {
}
```

4. Gateway

-- Yet to start --

```hcl
resource "istio_gateway" "test" {
}
```

5. Envoyfilter

-- Yet to start --

```hcl
resource "istio_envoy_filter" "test" {
}
```

6. Workloadentry

-- Yet to start --

```hcl
resource "istio_workload_entry" "test" {
}
```


## Contribution to the provider
Appreciate your contribution to the provider
