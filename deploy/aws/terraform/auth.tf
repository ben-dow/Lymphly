resource "aws_cognito_user_pool" "userpool" {
    name = "${var.application_name}-${var.environment_name}"
}

resource "aws_cognito_user_pool_client" "client" {
  name = "${var.application_name}-${var.environment_name}-default"
  user_pool_id = aws_cognito_user_pool.pool.id
  explicit_auth_flows = [
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_USER_SRP_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH"
  ]
}