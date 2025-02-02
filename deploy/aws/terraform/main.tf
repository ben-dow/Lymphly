terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~>5.84"
    }
  }

  backend "s3" {}

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = var.deployment_region
  default_tags {
    tags = {
      Application = var.application_name
      Environment = var.environment_name
    }
  }
}