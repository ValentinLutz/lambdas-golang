module "database" {
  source = "../database"
}

module "v1_get_order" {
  source = "./v1-get-order"
}