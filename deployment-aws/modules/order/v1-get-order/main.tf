resource "random_id" "suffix" {
  byte_length = 4
}

data "archive_file" "v1_get_order" {
  type        = "zip"
  source_file = "${path.root}/../../../../services/order/lambda-v1-get-order/bootstrap"
  output_path = "${path.root}/.terraform/files/v1-get-order.zip"
}

data "aws_iam_policy_document" "v1_get_order_role_policy" {
  version = "2012-10-17"

  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "v1_get_order" {
  name               = "v1-get-order-${random_id.suffix.hex}"
  assume_role_policy = data.aws_iam_policy_document.v1_get_order_role_policy.json
}

resource "aws_lambda_function" "v1_get_order" {
  function_name    = "v1-get-order-${random_id.suffix.hex}"
  role             = aws_iam_role.v1_get_order.arn
  handler          = "bootstrap"
  runtime          = "provided.al2023"
  architectures    = ["arm64"]
  memory_size      = 128
  timeout          = 10
  filename         = data.archive_file.v1_get_order.output_path
  source_code_hash = data.archive_file.v1_get_order.output_base64sha256

  tracing_config {
    mode = "Active"
  }

  environment {
    variables = {
      DB_HOST      = "todo"
      DB_PORT      = "5432"
      DB_NAME      = "todo"
      DB_SECRET_ID = "todo"
    }
  }
}

resource "aws_cloudwatch_log_group" "v1_get_order" {
  name              = "/aws/lambda/${aws_lambda_function.v1_get_order.function_name}"
  retention_in_days = 30
}

data "aws_iam_policy_document" "v1_get_order_policy" {
  version = "2012-10-17"

  statement {
    effect  = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      aws_cloudwatch_log_group.v1_get_order.arn
    ]
  }
}

resource "aws_iam_role_policy" "v1_get_order" {
  role   = aws_iam_role.v1_get_order.id
  policy = data.aws_iam_policy_document.v1_get_order_policy.json
}