package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Ademayowa/learn-terraform/db"
	"github.com/Ademayowa/learn-terraform/routes"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
)

type mockDynamoDB struct{}

func (m *mockDynamoDB) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, nil
}

func (m *mockDynamoDB) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	return &dynamodb.ScanOutput{Items: []map[string]types.AttributeValue{}}, nil
}

func TestMain(m *testing.M) {
	os.Setenv("DYNAMODB_TABLE", "test-properties")
	db.DynamoDB = &mockDynamoDB{}
	os.Exit(m.Run())
}

func TestCreateProperty(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"title":"Test House","location":"Lagos"}`
	c.Request = httptest.NewRequest("POST", "/properties", bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")

	routes.CreateProperty(c)

	if w.Code != http.StatusCreated {
		t.Errorf("expected %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetProperties(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/properties", nil)

	routes.GetProperties(c)

	if w.Code != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, w.Code)
	}

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if _, ok := resp["properties"]; !ok {
		t.Error("response should contain 'properties' key")
	}
}
