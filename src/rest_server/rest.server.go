package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	api "../grpc"
)

var (
	pingRpcEndpoint = flag.String("ping", "localhost:9090", "/v1/ping")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterApiHandlerFromEndpoint(ctx, mux, *pingRpcEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	fmt.Println("Compiler Tool HTTP/2 REST GRPC Server starting up on PORT 8080")

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
