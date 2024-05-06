module "bootstrap" {
  source = "../../../modules/cheap"

  region      = var.region
  environment = var.environment
  project     = var.project
}