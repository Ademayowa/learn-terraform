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
├── cmd/lambda/
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
├── tests/
│   └── properties_test.go
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

### 3. Deploy Infrastructure on dev

```bash
cd terraform/envs/dev
terraform init
terraform apply
```

## terraform/state-setup

⚠️ **WARNING: DO NOT DESTROY THIS INFRASTRUCTURE** ⚠️

This folder creates the S3 bucket and DynamoDB table for storing
Terraform state across all environments.

**Already created:** ✅ S3 bucket: `learn-terraform-state-files`
**Already created:** ✅ DynamoDB table: `terraform-state-lock`

**DO NOT RUN:** `terraform destroy` (will break all environments)
**IF DELETED:** All dev/prod state files will be lost

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

Copy the API endpoint from the output.

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
