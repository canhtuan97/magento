package main

import (
	"context"
	"flag"
	"github.com/canhtuan97/magento/proto/cart"
	"github.com/canhtuan97/magento/proto/customer"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:50069", "gRPC server endpoint")
)

func run() error {
	log.Println("Proxy is running")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := customerPb.RegisterCustomerHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err1 := cartPb.RegisterAddItemProductHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err1 != nil {
		return err
	}
	// Start HTTP server (and proto_demo calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}