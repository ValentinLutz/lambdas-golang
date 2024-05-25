module "vpc_name" {
  source = "../../../../modules/name"

  region      = var.region
  environment = var.environment
  project     = var.project
  name        = "main"
}