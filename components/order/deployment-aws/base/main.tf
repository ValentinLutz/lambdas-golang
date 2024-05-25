data "aws_iam_policy_document" "v1" {
  version = "2012-10-17"

  statement {
    effect    = "Allow"
    actions   = ["execute-api:Invoke"]
    resources = ["*"]

    principals {
      type        = "AWS"
      identifiers = ["489721517942"]
    }
  }
}

resource "aws_api_gateway_rest_api" "v1" {
  name = module.api_name.name
  body = templatefile("${path.module}/../../api-definition/order-api-v1.yaml", {
    region                   = var.region
    get_orders_function_arn  = module.lambda_v1_get_order.arn
    post_orders_function_arn = module.lambda_v1_get_order.arn
    get_order_function_arn   = module.lambda_v1_get_order.arn
  })

  policy = data.aws_iam_policy_document.v1.json

  endpoint_configuration {
    types = ["PRIVATE"]
  }
}

resource "aws_api_gateway_deployment" "v1" {
  rest_api_id = aws_api_gateway_rest_api.v1.id
  stage_name  = var.environment
}