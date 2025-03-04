locals {
  private_basepath = "/api/private/v1"
}

resource "aws_lambda_function" "private_labmda" {
    filename = "${var.releases_path}/private_lambda_x86_64.zip"
    function_name = "${var.application_name}_${var.environment_name}_private"
    role = aws_iam_role.lambda_execute_role.arn
    runtime = "provided.al2023"
    handler = "bootstrap"
    architectures = [ "x86_64" ]
    source_code_hash = filemd5("${var.releases_path}/private_lambda_x86_64.zip")
    environment {
      variables = {
        BASE_PATH = local.private_basepath
        APP_NAME = var.application_name
        ENV_NAME = var.environment_name
        REGION = var.deployment_region        
        TABLE_NAME = aws_dynamodb_table.lymphly-table.name
        LOG_LEVEL = "INFO"
        RADAR_PRIVATE_KEY = var.radar_secret_key
      }
    }
}

resource "aws_apigatewayv2_integration" "private_integration" {
  api_id = aws_apigatewayv2_api.api.id
  integration_type = "AWS_PROXY"
  integration_uri = aws_lambda_function.private_labmda.arn
  payload_format_version = "1.0"
}

resource "aws_apigatewayv2_route" "private" {
  api_id = aws_apigatewayv2_api.api.id
  route_key = "ANY ${local.private_basepath}/{proxy+}"
  target = "integrations/${aws_apigatewayv2_integration.private_integration.id}"
  authorization_type = "JWT"
  authorizer_id = aws_apigatewayv2_authorizer.authorizer.id
  authorization_scopes = ["aws.cognito.signin.user.admin"]
}

resource "aws_lambda_permission" "private_invoke" {
  statement_id = "AllowExecutionFromAPIGateway"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.private_labmda.function_name
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_apigatewayv2_api.api.execution_arn}/*/*" 
}
