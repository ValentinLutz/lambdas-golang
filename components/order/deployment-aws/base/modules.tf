module "api_name" {
  source = "../../../../modules/name"

  region      = var.region
  environment = var.environment
  project     = var.project
  name        = "orders"
}

module "lambda_v1_get_order" {
  source = "./lambda-v1-get-order"

  region      = var.region
  environment = var.environment
  project     = var.project

  api_gateway_arn = aws_api_gateway_rest_api.v1.execution_arn
}