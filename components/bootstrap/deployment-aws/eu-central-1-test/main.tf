module "bootstrap" {
  source = "../base"

  region      = var.region
  environment = var.environment
  project     = var.project
}