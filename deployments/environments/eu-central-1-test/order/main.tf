module "order" {
  source = "../../../../services/order/deployment-aws"

  region      = var.region
  environment = var.environment
  project     = var.project
}