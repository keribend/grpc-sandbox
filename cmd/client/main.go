package main

import (
	"context"
	"flag"
	"github.com/keribend/grpc-sandbox/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	noteServerAddr := flag.String("server", "localhost:8080", "The server address in the format of host:port")
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *noteServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewNoteClient(conn)
	createResp, err := client.Create(ctx, &pb.CreateNoteRequest{
		Schema: "schema-slug",
		Text:   "Some note text",
	})

	if err != nil {
		log.Fatalf("failed to create note: %v", err)
	}

	log.Printf("created note: %+v", createResp)
}
