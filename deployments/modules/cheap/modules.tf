module "database_cost_saver_name" {
  source = "../name"

  region      = var.region
  environment = var.environment
  project     = var.project
  name        = "database-cost-saver"
}