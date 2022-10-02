package main

import (
	"context"
	"log"
	"net"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/mrexmelle/go-casbin-grpc-example/proto/authz"
	"google.golang.org/grpc"
)

type authzServer struct {
	pb.UnimplementedAuthzServer
	enforcer *casbin.Enforcer
}

func newServer() *authzServer {
	a, _ := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	e, _ := casbin.NewEnforcer("model.conf", a)
	e.LoadPolicy()

	e.AddPolicy("alice", "data1", "read")
	e.AddPolicy("data2_admin", "data2", "read")
	e.AddPolicy("data2_admin", "data2", "write")
	e.AddGroupingPolicy("alice", "data2_admin")

	e.SavePolicy()

	return &authzServer{enforcer: e}
}

func (s *authzServer) Verify(ctx context.Context, in *pb.VerificationRequest) (*pb.VerificationResponse, error) {
	res, err := s.enforcer.Enforce(in.Id, in.Resource, in.Method)
	if res {
		return &pb.VerificationResponse{Authorized: true}, err
	} else {
		return &pb.VerificationResponse{Authorized: false}, err
	}
}

func (s *authzServer) GetRolesForUser(ctx context.Context, in *pb.RolesForUserRequest) (*pb.RolesForUserResponse, error) {
	res, err := s.enforcer.GetRolesForUser(in.Id)
	return &pb.RolesForUserResponse{Roles: res}, err
}

func main() {

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthzServer(grpcServer, newServer())
	log.Printf("Listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
