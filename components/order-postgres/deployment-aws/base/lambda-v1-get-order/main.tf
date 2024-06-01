data "archive_file" "lambda" {
  type        = "zip"
  source_file = "${path.module}/../../../lambda-v1-get-order/bootstrap"
  output_path = "${path.root}/.terraform/files/lambda-v1-get-order.zip"
}

data "aws_iam_policy_document" "lambda" {
  version = "2012-10-17"

  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda" {
  name               = module.name.name
  assume_role_policy = data.aws_iam_policy_document.lambda.json
}

resource "aws_lambda_permission" "lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${var.api_gateway_arn}/*"
}

resource "aws_lambda_function" "lambda" {
  function_name    = module.name.name
  role             = aws_iam_role.lambda.arn
  handler          = "bootstrap"
  runtime          = "provided.al2023"
  architectures    = ["arm64"]
  memory_size      = 128
  timeout          = 10
  filename         = data.archive_file.lambda.output_path
  source_code_hash = data.archive_file.lambda.output_base64sha256

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

resource "aws_cloudwatch_log_group" "logs" {
  name              = "/aws/lambda/${aws_lambda_function.lambda.function_name}"
  retention_in_days = 30
}

data "aws_iam_policy_document" "logs" {
  version = "2012-10-17"

  statement {
    effect  = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "${aws_cloudwatch_log_group.logs.arn}:*"
    ]
  }
}

resource "aws_iam_role_policy" "logs" {
  name   = "AllowCloudWatchLogs"
  role   = aws_iam_role.lambda.id
  policy = data.aws_iam_policy_document.logs.json
}

data "aws_iam_policy_document" "xray" {
  version = "2012-10-17"

  statement {
    effect  = "Allow"
    actions = [
      "xray:PutTraceSegments",
      "xray:PutTelemetryRecords",
      "xray:GetSamplingRules",
      "xray:GetSamplingTargets",
      "xray:GetSamplingStatisticSummaries"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "xray" {
  name   = "AllowXRayTracing"
  role   = aws_iam_role.lambda.id
  policy = data.aws_iam_policy_document.xray.json
}