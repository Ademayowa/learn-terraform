terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.19" // keeps the aws provider stable in other to avoid breaking changes
    }
  }
}

provider "aws" {
  region = var.aws_region
}
