data "aws_iam_policy_document" "lambda_execute_assume_policy" {
    statement {
        effect = "Allow"
        principals {
            type        = "Service"
            identifiers = ["lambda.amazonaws.com"]
        }
        actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "lambda_execute_role" {
  name = "${var.application_name}-${var.environment_name}-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_execute_assume_policy.json
}