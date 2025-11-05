package models

import (
	"context"
	"os"

	"github.com/Ademayowa/learn-terraform/db"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

type Property struct {
	ID       string `json:"id" dynamodbav:"id"`
	Title    string `json:"title" dynamodbav:"title"`
	Location string `json:"location" dynamodbav:"location"`
}

func (p *Property) Save() error {
	p.ID = uuid.New().String()

	av, err := attributevalue.MarshalMap(p)
	if err != nil {
		return err
	}

	tableName := os.Getenv("DYNAMODB_TABLE")

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &tableName,
	}

	_, err = db.DynamoDB.PutItem(context.TODO(), input)
	return err
}

func GetAllProperties() ([]Property, error) {
	tableName := os.Getenv("DYNAMODB_TABLE")

	input := &dynamodb.ScanInput{
		TableName: &tableName,
	}

	result, err := db.DynamoDB.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var properties []Property
	err = attributevalue.UnmarshalListOfMaps(result.Items, &properties)
	if err != nil {
		return nil, err
	}

	return properties, nil
}
