locals {
  providersearch_basepath = "/api/v1/providersearch"
}

resource "aws_lambda_function" "providersearch_lambda" {
    filename = "${var.releases_path}/providersearch_lambda_x86_64.zip"
    function_name = "provider_search"
    role = aws_iam_role.lambda_execute_role.arn
    runtime = "provided.al2023"
    handler = "bootstrap"
    architectures = [ "x86_64" ]
    source_code_hash = filemd5("${var.releases_path}/providersearch_lambda_x86_64.zip")
    environment {
      variables = {
        BASE_PATH = local.providersearch_basepath
      }
    }
}

resource "aws_apigatewayv2_integration" "providersearch_integration" {
  api_id = aws_apigatewayv2_api.api.id
  integration_type = "AWS_PROXY"
  integration_uri = aws_lambda_function.providersearch_lambda.arn
  payload_format_version = "1.0"
}

resource "aws_apigatewayv2_route" "providersearch" {
  api_id = aws_apigatewayv2_api.api.id
  route_key = "ANY ${local.providersearch_basepath}/{proxy+}"
  target = "integrations/${aws_apigatewayv2_integration.providersearch_integration.id}"
}

resource "aws_lambda_permission" "providersearch_invoke" {
  statement_id = "AllowExecutionFromAPIGateway"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.providersearch_lambda.function_name
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_apigatewayv2_api.api.execution_arn}/*/*" 
}
