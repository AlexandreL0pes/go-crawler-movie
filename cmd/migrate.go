package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamo *dynamodb.DynamoDB

var TableName = "movies"

// func connectDynamo() (db *dynamodb.DynamoDB) {
// 	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
// 		Region: aws.String("eu-central-1"),
// 	})))
// }

func createTable() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("eu-central-1"),
	}))

	svc := dynamodb.New(sess)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   aws.String(TableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		log.Fatalf("Database setup failed\nError: %s", err)
	}

	fmt.Println("Database setup finalized", TableName)
}

func main() {
	createTable()
}
