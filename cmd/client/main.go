package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

var (
	serverAddr = os.Getenv("backend")
	listen     = os.Getenv("LISTEN")
)

func init() {
	if serverAddr == "" {
		serverAddr = "localhost:8080"
	}
	if listen == "" {
		listen = "localhost:8081"
	}
}

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()
	api.RegisterAllServiceHandler(context.Background(), grpcMux, conn)
	log.Fatal(http.ListenAndServe(listen, grpcMux))
}
