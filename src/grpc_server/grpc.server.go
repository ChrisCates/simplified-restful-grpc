package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	api "../grpc"
)

type APIServer struct {
}

func (s *APIServer) Ping(ctx context.Context, message *api.PingMessage) (*api.PingMessage, error) {

	ping := "pong"
	time := int64(time.Now().Unix())

	return &api.PingMessage{ping, time, struct{}{}, nil, 64}, nil
}

func newAPIServer() *APIServer {
	server := &APIServer{}
	return server
}

func main() {
	flag.Parse()
	tcpServer, err := net.Listen("tcp", fmt.Sprintf(":9090"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterApiServer(grpcServer, newAPIServer())

	fmt.Println("Compiler Tool GRPC Server starting up on PORT 9090")
	grpcServer.Serve(tcpServer)
}
