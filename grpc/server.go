package main

import (
	"context"
	"fmt"
	pb "grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Employee struct {
	id   int32
	name string
	dept string
}

var Emps = []Employee{
	{
		id:   1,
		name: "Mahir",
		dept: "IT",
	},
	{
		id:   2,
		name: "Riham",
		dept: "Chemistry",
	},
	{
		id:   3,
		name: "Roshan",
		dept: "Physics",
	},
}

type server struct {
	pb.UnimplementedEmployeeServer
}

func (s *server) GetEmpDetails(ctx context.Context, req *pb.EmpRequest) (*pb.EmpResponse, error) {
	// log.Printf("Recieved ID: %d", req.Id)
	for _, emp := range Emps {
		if emp.id == req.Id {
			return &pb.EmpResponse{
				Name: emp.name,
				Dept: emp.dept,
			}, nil
		}
	}
	return &pb.EmpResponse{
		// Name: "",
		// Dept: "",
	}, fmt.Errorf("no employee found")
}

func (s *server) PostEmpDetails(ctx context.Context, req *pb.EmpResponse) (*pb.EmpRequest, error) {
	// if req.Dept == "" && req.Name == "" {
	// 	return &pb.EmpRequest{}, fmt.Errorf("invalid arguments")
	// }

	Emps = append(Emps, Employee{
		id:   int32(len(Emps) + 1),
		name: req.Name,
		dept: req.Dept,
	})

	return &pb.EmpRequest{
		Id: int32(len(Emps)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed To Listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterEmployeeServer(s, &server{})

	log.Println("starting server at port 50051!!")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
