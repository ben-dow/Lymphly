resource "aws_apigatewayv2_api" "api" {
    name = "${var.application_name}-api"
    protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "default_stage" {
  api_id = aws_apigatewayv2_api.api.id
  name   = "$default"
  auto_deploy = true
}