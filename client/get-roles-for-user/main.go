package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/mrexmelle/go-casbin-grpc-example/proto/authz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <user>\n", os.Args[0])
	}

	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot dial to 127.0.0.1:3000")
	}
	defer conn.Close()

	client := pb.NewAuthzClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetRolesForUser(ctx, &pb.RolesForUserRequest{
		Id: os.Args[1],
	})

	if err == nil && len(response.Roles) > 0 {
		log.Printf(strings.Join(response.Roles, "\n"))
	}
}
