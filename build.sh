#!/bin/bash

# Build the Go binary for Lambda
cd lambda
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap main.go

# Zip the binary
zip ../build/lambda.zip bootstrap

# Clean up
rm bootstrap

echo "✅ Lambda built successfully at build/lambda.zip"