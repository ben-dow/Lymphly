resource "aws_ssm_parameter" "radarSecret" {
    name = "/${var.application_name}/${var.environment_name}/key/radar/private"
    value = var.radar_secret_key
    type="SecureScring"
}