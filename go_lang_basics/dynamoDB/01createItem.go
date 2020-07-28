package main

import (
	"fmt"
	"github.com/BhaviD/Swiggy_Bootcamp/go_lang_basics/dynamoDB/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("us-east-1"),
	}))

	db := dynamodb.New(sess)

	music := types.Music{
		Artist: "Sonu Nigam",
		SongTitle: "Hello Again",
		AlbumTitle: "My Album",
		Awards: 2,
	}

	musicMap, err := dynamodbattribute.MarshalMap(music)
	if err != nil {
		panic("Cannot map the values given in music struct...")
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String("Music1"),
		Item: musicMap,
	}

	resp, err := db.PutItem(params)

	if err != nil {
		log.Fatalf("Some problem while inserting : %v", err)
	}
	// print the response
	fmt.Println("Record Inserted...")
	fmt.Println(resp)
}

