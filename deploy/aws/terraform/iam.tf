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