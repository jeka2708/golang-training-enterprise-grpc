package main

import (
	"context"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/api"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/db"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"

	"time"
)

var (
	host     = "db-enterprise"
	port     = ":8080"
	user     = "postgres"
	dbname   = "enterprise"
	password = "root"
	sslmode  = "disable"
)

func main() {
	ctx := context.Background()
	ctx, error := context.WithCancel(ctx)
	ctx, error = context.WithTimeout(ctx, time.Second*30)
	if error != nil {
		log.Println(error)
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	conn, err := connectToDbWithTimeout(ctx)
	if err != nil {
		log.Fatal(err)
	}
	api.RegisterAllServices(server, conn)

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
func connectToDbWithTimeout(ctx context.Context) (*gorm.DB, error) {
	for {
		time.Sleep(2 * time.Second)
		conn, err := db.GetConnection(host, "5432", user, dbname, password, sslmode)
		if err == nil {
			return conn, nil
		}
		select {
		case <-ctx.Done():
			return nil, err
		default:
			continue
		}
	}
}
