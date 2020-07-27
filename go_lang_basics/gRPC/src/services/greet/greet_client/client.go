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
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	c := greetpb.NewGreetServiceClient(conn)

	callGreet(c);
	callGreetFullName(c);
}

func  callGreet(c greetpb.GreetServiceClient)  {
	fmt.Println("In callGreet()... ")

	req:= &greetpb.GreetRequest {
		Greeting: &greetpb.Greeting {
			FirstName: "Bhavi",
			LastName: "Dhingra",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While calling Greet : %v", err)
	}

	log.Println("Response From the Greet ", res.Result)
}

func  callGreetFullName(c greetpb.GreetServiceClient)  {
	fmt.Println("In callGreetFullName()... ")

	req:= &greetpb.GreetRequest {
		Greeting: &greetpb.Greeting {
			FirstName: "Bhavi",
			LastName: "Dhingra",
		},
	}
	res, err := c.GreetFullName(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	log.Println("Response From the GreetFullName:", res.Greeting.FirstName + " " + res.Greeting.LastName)
}
