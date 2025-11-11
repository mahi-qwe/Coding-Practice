package main

import (
	"context"
	pb "grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to build a connection %v", err)
	}
	defer conn.Close()

	client := pb.NewEmployeeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.PostEmpDetails(ctx, &pb.EmpResponse{
		Name: "Rohit",
		Dept: "Maths",
	})
	if err != nil {
		log.Fatalf("request failed %v", err)
	}

	log.Printf("id: %d", res.Id)
}
