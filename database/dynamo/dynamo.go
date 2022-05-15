package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDB struct {
	Client *dynamodb.DynamoDB
}

func NewDynamoDB() (DynamoDB, error) {
	db := DynamoDB{}

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("eu-central-1"),
	}))

	db.Client = dynamodb.New(sess)

	return db, nil
}
