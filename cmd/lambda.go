package main

import (
	"context"

	"smith.lambda/service"
  "smith.lambda/swagger-generate/restapi"
	"smith.lambda/swagger-generate/restapi/operations"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-openapi/loads"
)

var httpAdapter *httpadapter.HandlerAdapter

func getHttpAdapter() (*httpadapter.HandlerAdapter, error) {
  if httpAdapter == nil {
    swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
    if err != nil {
      return nil, err
    }

		api := operations.NewLambdaAppAPI(swaggerSpec)
    server := restapi.NewServer(api)
		service.Route(api)
    server.ConfigureAPI()

    // see https://github.com/go-swagger/go-swagger/issues/962#issuecomment-478382896
    httpAdapter = httpadapter.New(server.GetHandler())
  }

  return httpAdapter, nil
}

// Handler handles API requests
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  adapter, err := getHttpAdapter()
  if err != nil {
    panic(err)
  }
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
