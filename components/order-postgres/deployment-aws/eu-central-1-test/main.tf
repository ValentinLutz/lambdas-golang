module "order" {
  source = "../base"

  region      = var.region
  environment = var.environment
  project     = var.project
}