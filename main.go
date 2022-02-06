package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jeremyseow/backend-assignment-grpc/config"
	"github.com/jeremyseow/backend-assignment-grpc/internal/storage"
	"github.com/jeremyseow/backend-assignment-grpc/internal/storage/file"
	"github.com/jeremyseow/backend-assignment-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
)

type server struct {
	pb.UnimplementedEventServiceServer
	storage storage.Storage
}

func main() {
	listener, err := net.Listen("tcp", config.Conf.GrpcPort)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	fw := file.NewWriter()

	pb.RegisterEventServiceServer(grpcServer, &server{storage: fw})
	reflection.Register(grpcServer)

	if e := grpcServer.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) SendEvent(_ context.Context, request *pb.EventRequest) (*pb.EventResponse, error) {
	log.Printf("size of request: %d", len(request.Events))

	currTime := time.Now().UTC()

	bytes, err := proto.Marshal(request)
	if err != nil {
		log.Printf("error when unmarshalling: %s", err)
		return &pb.EventResponse{Result: "Failed"}, err
	}

	err = s.storage.Write("output", fmt.Sprintf("%s-%d.txt", "message", currTime.UnixMilli()), bytes)
	if err != nil {
		log.Printf("error when writing: %s", err)
		return &pb.EventResponse{Result: "Failed"}, err
	}

	return &pb.EventResponse{Result: "Success"}, nil
}
