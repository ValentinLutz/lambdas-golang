module "database_name" {
  source = "../name"

  region      = var.region
  environment = var.environment
  project     = var.project
  name        = var.name
}