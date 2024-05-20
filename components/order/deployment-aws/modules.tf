module "lambda_v1_get_order" {
  source = "./lambda-v1-get-order"

  region      = var.region
  environment = var.environment
  project     = var.project
}