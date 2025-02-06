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
  role = aws_iam_role.lambda_execute_role.arn
}

resource "aws_iam_policy" "lambda_execution_policy" {
  name = "${var.application_name}-${var.environment_name}-lambda-policy"
  policy = data.aws_iam_policy_document.lambda_execute_policy.json
}

data "aws_iam_policy_document" "lambda_execute_policy" {
    statement {
        effect = "Allow"
        actions = ["ssm:GetParameter"]
        resources = [ aws_ssm_parameter.radarSecret.arn ]
  }
}