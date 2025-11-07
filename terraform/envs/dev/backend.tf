terraform {
  backend "s3" {
    bucket         = "learn-terraform-state-files"
    key            = "terraform/envs/dev/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform-state-lock"
    encrypt        = true
  }
}
