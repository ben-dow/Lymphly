resource "aws_dynamodb_table" "lymphly-table" {
    name = "${var.application_name}-${var.environment_name}-table"
    billing_mode = "PAY_PER_REQUEST"
    
    attribute {
      name = "pk"
      type = "S"
    }
    
    attribute {
      name = "sk"
      type = "S"
    }

    
    hash_key = "pk"
    range_key = "sk"
}