module "bootstrap" {
  source = "../../../modules/bootstrap"

  region      = var.region
  environment = var.environment
  project     = var.project
}