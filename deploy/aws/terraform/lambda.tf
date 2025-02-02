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