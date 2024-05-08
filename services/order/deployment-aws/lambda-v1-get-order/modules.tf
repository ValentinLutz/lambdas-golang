module "name" {
  source = "../../../../deployments/modules/name"

  region      = var.region
  environment = var.environment
  project     = var.project
  name        = "v1-get-order"
}