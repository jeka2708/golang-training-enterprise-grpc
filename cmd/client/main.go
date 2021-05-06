package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()
	api.RegisterAllServiceHandler(context.Background(), grpcMux, conn)
	log.Fatal(http.ListenAndServe("localhost:8081", grpcMux))
}
