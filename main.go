package main

import (
	"context"
	"log"
	"net"

	"github.com/jeremyseow/backend-assignment-grpc/config"
	"github.com/jeremyseow/backend-assignment-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedEventServiceServer
}

func main() {
	listener, err := net.Listen("tcp", config.Conf.GrpcPort)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterEventServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	if e := grpcServer.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) SendEvent(_ context.Context, request *pb.EventRequest) (*pb.EventResponse, error) {
	result := request.String()
	log.Printf("size of request: %d", len(request.Events))

	return &pb.EventResponse{Result: result}, nil
}
