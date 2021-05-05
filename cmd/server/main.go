package main

import (
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/api"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	host     = "localhost"
	port     = ":8282"
	user     = "postgres"
	dbname   = "enterprise"
	password = "root"
	sslmode  = "disable"
)

func main() {
	conn, error := db.GetConnection(host, "5432", user, dbname, password, sslmode)
	if error != nil {
		log.Println(error)
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	api.RegisterAllServices(server, conn)

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
