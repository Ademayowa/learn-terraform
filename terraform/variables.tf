variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "learn-terraform"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "dev"
}

variable "dynamodb_table_name" {
  description = "DynamoDB table name for properties"
  type        = string
  default     = "properties"
}

variable "lambda_zip_path" {
  description = "Path to the Lambda function zip file"
  type        = string
  default     = "../lambda.zip"
}

variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}
