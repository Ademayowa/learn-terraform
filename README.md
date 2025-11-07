# Property API

A serverless API built with Go, AWS Lambda, API Gateway, and DynamoDB deployed using Terraform.

## Prerequisites

- Go 1.25+
- AWS CLI configured
- Terraform 1.0+

## Project Structure

```
learn-terraform/
├── api-test/
├── build/
│   └── lambda.zip
├── db/
├── lambda/
├── models/
├── routes/
├── terraform/
│   ├── envs/
│   │   ├── dev/
│   │   │   ├── backend.tf
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   ├── provider.tf
│   │   │   ├── terraform.tfvars
│   │   │   └── variables.tf
│   │   └── prod/
│   │       ├── backend.tf
│   │       ├── main.tf
│   │       ├── outputs.tf
│   │       ├── provider.tf
│   │       ├── terraform.tfvars
│   │       └── variables.tf
│   └── state-setup/
│       └── main.tf
├── .gitignore
├── build.sh
├── go.mod
├── go.sum
├── makefile
└── README.md
```

## Setup

### 1. Install Dependencies

```bash
go mod download
```

### 2. Build Lambda

```bash
chmod +x build.sh
./build.sh
```

### 4. Setup Terraform

Run this ONCE to create S3 bucket and DynamoDB table for state management.

```bash
cd terraform/state-setup
terraform init && terraform apply
```

### 5. Deploy Infrastructure on dev

```bash
cd terraform/envs/dev
terraform init
terraform apply
```

Copy the API endpoint from the output.

## API Endpoints

**Base URL**: `https://YOUR-API-ID.execute-api.us-east-1.amazonaws.com`

### Create Property

```bash
curl -X POST https://YOUR-API-ID.execute-api.us-east-1.amazonaws.com/properties \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Beach House",
    "location": "Miami, FL"
  }'
```

### Get All Properties

```bash
curl https://YOUR-API-ID.execute-api.us-east-1.amazonaws.com/properties
```

## Development

After making code changes on dev:

```bash
./build.sh
cd terraform/envs/dev
terraform apply
```

## Cleanup

```bash
cd terraform
terraform destroy
```

## Tech Stack

- **Runtime**: Go with AWS Lambda
- **Database**: DynamoDB
- **API Gateway**: AWS HTTP API
- **Infrastructure**: Terraform
- **Framework**: Gin
