locals {
  public_basepath = "/api/public/v1"
}

resource "aws_lambda_function" "public_lambda" {
    filename = "${var.releases_path}/public_lambda_x86_64.zip"
    function_name = "${var.application_name}_${var.environment_name}_public"
    role = aws_iam_role.lambda_execute_role.arn
    runtime = "provided.al2023"
    handler = "bootstrap"
    timeout = 10
    architectures = [ "x86_64" ]
    source_code_hash = filemd5("${var.releases_path}/public_lambda_x86_64.zip")
    environment {
      variables = {
        BASE_PATH = local.public_basepath
        APP_NAME = var.application_name
        ENV_NAME = var.environment_name
        REGION = var.deployment_region
        TABLE_NAME = aws_dynamodb_table.lymphly-table.name
        LOG_LEVEL = "INFO"
        RADAR_PRIVATE_KEY = var.radar_secret_key
      }
    }
}

resource "aws_apigatewayv2_integration" "public_integration" {
  api_id = aws_apigatewayv2_api.api.id
  integration_type = "AWS_PROXY"
  integration_uri = aws_lambda_function.public_lambda.arn
  payload_format_version = "1.0"
}

resource "aws_apigatewayv2_route" "public" {
  api_id = aws_apigatewayv2_api.api.id
  route_key = "ANY ${local.public_basepath}/{proxy+}"
  target = "integrations/${aws_apigatewayv2_integration.public_integration.id}"
}

resource "aws_lambda_permission" "public_invoke" {
  statement_id = "AllowExecutionFromAPIGateway"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.public_lambda.function_name
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_apigatewayv2_api.api.execution_arn}/*/*" 
}
