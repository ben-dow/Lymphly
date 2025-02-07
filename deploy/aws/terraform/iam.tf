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

resource "aws_iam_role_policy_attachment" "lambda_execute_attach" {
  policy_arn = aws_iam_policy.lambda_execution_policy.arn
  role = aws_iam_role.lambda_execute_role.name
}

resource "aws_iam_policy" "lambda_execution_policy" {
  name = "${var.application_name}-${var.environment_name}-lambda-policy"
  policy = data.aws_iam_policy_document.lambda_execute_policy.json
}

data "aws_iam_policy_document" "lambda_execute_policy" {
    statement {
        effect = "Allow"
        actions = ["ssm:GetParameter", "ssm:GetParametersByPath"]
        resources = [ "arn:aws:ssm:${var.deployment_region}:${data.aws_caller_identity.current.account_id}:parameter/${data.aws_caller_identity.current.account_id}:${var.application_name}/${var.environment_name}/*"  ]
    }
}

data "aws_caller_identity" "current" {}
