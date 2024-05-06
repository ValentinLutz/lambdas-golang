module "bootstrap" {
  source = "../../../modules/database"

  region      = var.region
  environment = var.environment
  project     = var.project

  name = "master"
}