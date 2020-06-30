provider istio {
  config_path = "~/.kube/config"
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
