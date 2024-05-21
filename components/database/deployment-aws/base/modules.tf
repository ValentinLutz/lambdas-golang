module "database_name" {
  source = "../../../../modules/name"

  region      = var.region
  environment = var.environment
  project     = var.project
  name        = var.name
}

module "cost_reduction" {
  source = "./cost-reduction"

  region      = var.region
  environment = var.environment
  project     = var.project
}