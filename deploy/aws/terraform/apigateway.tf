resource "aws_apigatewayv2_api" "api" {
    name = "${var.application_name}-api"
    protocol_type = "HTTP"
}