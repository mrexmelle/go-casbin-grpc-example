package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/mrexmelle/go-casbin-grpc-example/proto/authz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <user> <resource> <action>\n", os.Args[0])
	}

	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot dial to 127.0.0.1:3000")
	}
	defer conn.Close()

	client := pb.NewAuthzClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Verify(ctx, &pb.VerificationRequest{
		Id:       os.Args[1],
		Resource: os.Args[2],
		Method:   os.Args[3],
	})

	if response.Authorized {
		log.Printf("Access allowed")
	} else {
		log.Printf("Access denied")
	}
}
