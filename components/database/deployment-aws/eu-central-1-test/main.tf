module "database" {
  source = "../base"

  region      = var.region
  environment = var.environment
  project     = var.project

  name = "master"
}