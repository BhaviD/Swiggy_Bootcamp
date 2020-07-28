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

	// query parameters
	artist := "Sonu Nigam"
	songTitle := "Hello Again"

	params := &dynamodb.GetItemInput{
		TableName: aws.String("Music1"),
		Key: map[string]*dynamodb.AttributeValue{
			"Artist" : {
				S: aws.String(artist),
			},
			"SongTitle" : {
				S: aws.String(songTitle),
			},
		},
	}

	resp, err := db.GetItem(params)
	if err != nil {
		log.Fatalf("Sorry item not found : %v", err)
	}
	// print the response
	fmt.Println(resp)

	var music = types.Music {}
	err = dynamodbattribute.UnmarshalMap(resp.Item, &music)
	fmt.Printf("Unmarshalled Item %v: ", music)
}

