package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"services/greet/greetpb"
)

type server struct {

}

func (*server) Greet (ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Function called...")

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Hello : " + firstName + ", "  + lastName
	res := &greetpb.GreetResponse {
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World from Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v: ", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if s.Serve(lis); err != nil {
		log.Fatalf("failed to Serve %v", err)
	}
}

