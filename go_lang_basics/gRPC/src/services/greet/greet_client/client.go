package main

import (
	"fmt"
	"github.com/BhaviD/Swiggy_Bootcamp/go_lang_basics/gRPC/src/services/greet/greetpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hi I'm in client...")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	callGreet(c);
}

func callGreet(c greetpb.GreetServiceClient)  {
	fmt.Println("In Call Greet Function...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting {
			FirstName: "Bhavi",
			LastName: "Dhingra",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet : %v", err)
	}
	fmt.Println("Response from the Greet ", res.Result)
}
