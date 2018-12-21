package main

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"github.com/aws/aws-sdk-go/service/lambda"
	awsLambda "github.com/aws/aws-lambda-go/lambda"
	"log"
	"fmt"
	"os"
)

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return Run(req)
	default:
		response := events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
		}
		return response, nil
	}
}


func getLambdaNamesFromEnv() []string {
	lambdaNames := []string{}
	for index := 1; index < 10; index++ {
		nextEnvVar := fmt.Sprintf("LAMBDA_FUNCTION%d", index)
		nextName := os.Getenv(nextEnvVar)

		if nextName == "" {
			return lambdaNames
		}

		lambdaNames = append(lambdaNames, nextName)
	}

	return lambdaNames
}


func Run(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sleeper := lambda.Lambda{

	}

	invocationType := "Event" // Asynchronous

	for _, lambdaName := range getLambdaNamesFromEnv() {
		input := lambda.InvokeInput{
			FunctionName:   &lambdaName,
			InvocationType: &invocationType,
		}

		_, err := sleeper.Invoke(&input)
		if err != nil {
			log.Fatalf("Error invoking lambda function %s\n. %s\n", lambdaName, err.Error())
		}
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}

	return response, nil
}


func main() {
	awsLambda.Start(router)
}