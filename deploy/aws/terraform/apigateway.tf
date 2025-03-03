resource "aws_apigatewayv2_api" "api" {
    name = "${var.application_name}-${var.environment_name}-apigateway"
    protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "default_stage" {
  api_id = aws_apigatewayv2_api.api.id
  name   = "$default"
  auto_deploy = true
}

resource "aws_apigatewayv2_authorizer" "authorizer" {
  api_id = aws_apigatewayv2_api.api.id
  authorizer_type = "JWT"
  identity_sources = ["$request.header.Authorization"]
  name = "${var.application_name}-${var.environment_name}-authorizer"

  jwt_configuration {
    audience = [aws_cognito_user_pool_client.client.id, aws_cognito_user_pool_client.programatic.id]
    issuer = "https://${aws_cognito_user_pool.userpool.endpoint}"
    
  }
}