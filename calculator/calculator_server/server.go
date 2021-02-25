package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KyriakosMilad/go-grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {}

func main() {
	fmt.Println("Hello...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
