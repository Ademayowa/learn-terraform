terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# Default VPC
data "aws_vpc" "default" {
  default = true
}

# Create one new subnet in a different AZ
resource "aws_subnet" "extra" {
  vpc_id                  = data.aws_vpc.default.id
  cidr_block              = "172.31.100.0/24"
  availability_zone       = "us-east-1b"
  map_public_ip_on_launch = true
}

# --- Security Groups ---

# Lambda SG
resource "aws_security_group" "lambda" {
  vpc_id = data.aws_vpc.default.id
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# RDS SG (allow only Lambda)
resource "aws_security_group" "rds" {
  vpc_id = data.aws_vpc.default.id
  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.lambda.id]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# --- RDS ---

resource "aws_db_subnet_group" "default" {
  name       = "${var.app_name}-db-subnet-group"
  subnet_ids = ["subnet-03ec7a79d1062e313", aws_subnet.extra.id]
}

resource "aws_db_instance" "postgres" {
  identifier             = "${var.app_name}-db"
  engine                 = "postgres"
  engine_version         = "13.22"
  instance_class         = "db.t3.micro"
  allocated_storage      = 20
  db_name                = var.db_name
  username               = var.db_username
  password               = var.db_password
  skip_final_snapshot    = true
  publicly_accessible    = false
  db_subnet_group_name   = aws_db_subnet_group.default.name
  vpc_security_group_ids = [aws_security_group.rds.id]
}

# --- Lambda ---

resource "aws_lambda_function" "api" {
  filename         = "deployment.zip"
  function_name    = var.app_name
  role             = aws_iam_role.lambda.arn
  handler          = "bootstrap"
  source_code_hash = filebase64sha256("deployment.zip")
  runtime          = "provided.al2023"
  timeout          = 30
  memory_size      = 512

  # Connect Lambda to VPC
  vpc_config {
    subnet_ids         = [aws_subnet.extra.id]
    security_group_ids = [aws_security_group.lambda.id]
  }

  environment {
    variables = {
      DB_HOST     = aws_db_instance.postgres.address
      DB_PORT     = "5432"
      DB_USER     = var.db_username
      DB_PASSWORD = var.db_password
      DB_NAME     = var.db_name
    }
  }

  depends_on = [aws_db_instance.postgres]
}

# --- IAM ---

resource "aws_iam_role_policy_attachment" "lambda_vpc_access" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
  role       = aws_iam_role.lambda.name
}


resource "aws_iam_role" "lambda" {
  name = "${var.app_name}-lambda-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect    = "Allow"
      Action    = "sts:AssumeRole"
      Principal = { Service = "lambda.amazonaws.com" }
    }]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_basic" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda.name
}

# --- API Gateway ---

resource "aws_apigatewayv2_api" "main" {
  name          = "${var.app_name}-api"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_integration" "lambda" {
  api_id                 = aws_apigatewayv2_api.main.id
  integration_type       = "AWS_PROXY"
  integration_uri        = aws_lambda_function.api.invoke_arn
  integration_method     = "POST"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "default" {
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.lambda.id}"
}

resource "aws_apigatewayv2_stage" "default" {
  api_id      = aws_apigatewayv2_api.main.id
  name        = "$default"
  auto_deploy = true
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.api.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.main.execution_arn}/*/*"
}

# --- Outputs ---

output "api_endpoint" {
  value = aws_apigatewayv2_stage.default.invoke_url
}

output "db_endpoint" {
  value     = aws_db_instance.postgres.endpoint
  sensitive = true
}
