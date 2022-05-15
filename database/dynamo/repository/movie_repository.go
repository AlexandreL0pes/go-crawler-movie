package repository

import (
	"fmt"
	"go-crawler-movie/domain/entity"

	"go-crawler-movie/database/dynamo"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const TABLE_NAME = "movies"

type MoviesRepository struct {
	dynamoDB dynamo.DynamoDB
}

func Initialize(db dynamo.DynamoDB) *MoviesRepository {
	return &MoviesRepository{
		dynamoDB: db,
	}
}

func (s *MoviesRepository) Insert(movie *entity.Movie) error {
	mm, err := dynamodbattribute.MarshalMap(movie)

	if err != nil {
		fmt.Printf("Got error marshalling map:\n\t %s", err.Error())
	}

	fmt.Printf("Marshed Message = %v\n", mm)

	input := &dynamodb.PutItemInput{
		Item:      mm,
		TableName: aws.String(TABLE_NAME),
	}

	_, err = s.dynamoDB.Client.PutItem(input)

	if err != nil {
		fmt.Printf("Got an error calling PutItem:\n\t%s", err.Error())
		return nil
	}

	fmt.Println("Successfully movie created")

	return nil
}
