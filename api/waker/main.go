package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/lambda"
	awsLambda "github.com/aws/aws-lambda-go/lambda"
	"log"
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return Run(req)
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
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	targetLambdas := getLambdaNamesFromEnv()
	log.Printf("Starting process to invoke %d lambdas\n", len(targetLambdas))

	invocationType := "Event" // Asynchronous
	lambdaCount := 0

	for _, lambdaName := range targetLambdas {
		input := lambda.InvokeInput{
			FunctionName:   &lambdaName,
			InvocationType: &invocationType,
		}

		_, err := client.Invoke(&input)
		if err != nil {
			log.Fatalf("Error invoking lambda function %s\n. %s\n", lambdaName, err.Error())
		}
		log.Printf("Just invoked lambda named %s\n", lambdaName)
		lambdaCount += 1
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}
	log.Printf("Invoked %d lambdas\n", lambdaCount)

	return response, nil
}


func main() {
	awsLambda.Start(router)
}