resource "aws_cognito_user_pool" "userpool" {
    name = "${var.application_name}-${var.environment_name}"
}

resource "aws_cognito_user_pool_client" "client" {
  name = "${var.application_name}-${var.environment_name}-default"
  user_pool_id = aws_cognito_user_pool.userpool.id
  explicit_auth_flows = [
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH"
  ]
}

resource "aws_cognito_user_pool_client" "programatic" {
  name = "${var.application_name}-${var.environment_name}-programatic"
  user_pool_id = aws_cognito_user_pool.userpool.id
  explicit_auth_flows = [
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH"
  ]
}

locals {
  auth_basepath = "/api/v1/auth"
}

resource "aws_lambda_function" "auth_lambda" {
    filename = "${var.releases_path}/auth_lambda_x86_64.zip"
    function_name = "${var.application_name}_${var.environment_name}_auth"
    role = aws_iam_role.lambda_execute_role.arn
    runtime = "provided.al2023"
    handler = "bootstrap"
    timeout = 10
    architectures = [ "x86_64" ]
    source_code_hash = filemd5("${var.releases_path}/auth_lambda_x86_64.zip")
    environment {
      variables = {
        BASE_PATH = local.auth_basepath
        APP_NAME = var.application_name
        ENV_NAME = var.environment_name
        REGION = var.deployment_region
        LOG_LEVEL = "INFO"
        POOL_ID = aws_cognito_user_pool.userpool.id
        CLIENT_ID = aws_cognito_user_pool_client.programatic.id
      }
    }
}

resource "aws_apigatewayv2_integration" "auth_lambda" {
  api_id = aws_apigatewayv2_api.api.id
  integration_type = "AWS_PROXY"
  integration_uri = aws_lambda_function.auth_lambda.arn
  payload_format_version = "1.0"
}

resource "aws_apigatewayv2_route" "auth_lambda" {
  api_id = aws_apigatewayv2_api.api.id
  route_key = "ANY ${local.auth_basepath}/{proxy+}"
  target = "integrations/${aws_apigatewayv2_integration.auth_lambda.id}"
}

resource "aws_lambda_permission" "auth_lambda_invoke" {
  statement_id = "AllowExecutionFromAPIGateway"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.auth_lambda.function_name
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_apigatewayv2_api.api.execution_arn}/*/*" 
}
