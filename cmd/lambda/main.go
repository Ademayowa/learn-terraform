package main

import (
	"context"
	"log"

	"github.com/Ademayowa/learn-terraform/db"
	"github.com/Ademayowa/learn-terraform/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	db.InitDB()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Add debug middleware
	router.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	routes.RegisterRoutes(router)

	ginLambda = ginadapter.NewV2(router)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log.Printf("Received request: %s %s", req.RequestContext.HTTP.Method, req.RawPath)
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
