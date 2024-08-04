package main

import (
	"context"
	"github.com/keribend/grpc-sandbox/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

var idFactory uint64

type server struct {
	pb.UnimplementedNoteServer
}

func (s *server) Create(ctx context.Context, in *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	idFactory++
	return &pb.CreateNoteResponse{
		Id:          idFactory,
		Schema:      in.Schema,
		Text:        in.Text,
		CreatedTime: time.Now().Format("2006-01-02T15:04:05"),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	noteServer := &server{}

	pb.RegisterNoteServer(s, noteServer)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
