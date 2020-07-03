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
