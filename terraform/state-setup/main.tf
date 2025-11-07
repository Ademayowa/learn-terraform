# Run this ONCE to create S3 bucket and DynamoDB table for state management
# Command: cd terraform/state-setup & run terraform init && terraform apply

resource "aws_s3_bucket" "terraform_state" {
  bucket = "learn-terraform-state-files"
}

resource "aws_s3_bucket_versioning" "terraform_state" {
  bucket = aws_s3_bucket.terraform_state.id

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_dynamodb_table" "terraform_lock" {
  name         = "terraform-state-lock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
