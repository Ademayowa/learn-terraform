# Property API

A serverless API built with Go, AWS Lambda, API Gateway, and DynamoDB deployed using Terraform.

## Prerequisites

- Go 1.25+
- AWS CLI configured
- Terraform 1.0+

## Project Structure

```
learn-terraform/
├── build.sh              # Build script
├── lambda/
│   └── main.go          # Lambda handler
├── db/
│   └── db.go            # DynamoDB client
├── models/
│   └── property.go      # Property model
├── routes/
│   ├── routes.go        # Route registration
│   └── properties.go    # Property handlers
└── terraform/
    ├── provider.tf      # AWS provider configuration
    ├── main.tf          # Infrastructure resources
    ├── variables.tf     # Input variables
    └── outputs.tf       # Output values
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

### 3. Deploy Infrastructure

```bash
cd terraform
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

After making code changes:

```bash
./build.sh
cd terraform
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
