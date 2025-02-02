data "aws_iam_policy_document" "provider_search_assume_role" {
    statement {
        effect = "Allow"
        principals {
            type        = "Service"
            identifiers = ["lambda.amazonaws.com"]
        }
        actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "provider_search_role" {
  name = "${var.application_name}-${var.environment_name}-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.provider_search_assume_role.json
}

resource "aws_lambda_function" "providersearch_lambda" {
    filename = "${var.releases_path}/ProviderSearch_lambda_x86_64.zip"
    function_name = "provider_search"
    role = aws_iam_role.provider_search_role.arn
    runtime = "provided.al2023"
    handler = "bootstrap"
    architectures = [ "x86_64" ]
    source_code_hash = filemd5("${var.releases_path}/ProviderSearch_lambda_x86_64.zip")
}

resource "aws_apigatewayv2_integration" "providersearch_integration" {
  api_id = aws_apigatewayv2_api.api.id
  integration_type = "AWS_PROXY"
  integration_uri = aws_lambda_function.providersearch_lambda.arn
  payload_format_version = "1.0"
}

resource "aws_apigatewayv2_route" "providersearch" {
  api_id = aws_apigatewayv2_api.api.id
  route_key = "ANY /api/v1/providersearch/{proxy+}"
  target = "integrations/${aws_apigatewayv2_integration.providersearch_integration.id}"
}

resource "aws_lambda_permission" "providersearch_invoke" {
  statement_id = "AllowExecutionFromAPIGateway"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.providersearch_lambda.function_name
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_apigatewayv2_api.api.execution_arn}/*/*" 
}
