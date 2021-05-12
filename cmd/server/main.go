package main

import (
	"context"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	"log"
	"net"
	"os"

	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/api"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/db"
	"time"
)

func main() {
	var listen = os.Getenv("LISTEN")
	if listen == "" {
		listen = ":8080"
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx, cancel = context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	conn, err := connectToDbWithTimeout(ctx)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}

	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	api.RegisterAllServices(server, conn)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
func connectToDbWithTimeout(ctx context.Context) (*gorm.DB, error) {

	var host = os.Getenv("DB_USERS_HOST")
	var port = os.Getenv("DB_USERS_PORT")
	var user = os.Getenv("DB_USERS_USER")
	var dbname = os.Getenv("DB_USERS_DBNAME")
	var password = os.Getenv("DB_USERS_PASSWORD")
	var sslmode = os.Getenv("DB_USERS_SSL")

	if host == "" {
		host = "db-enterprise"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "enterprise"
	}
	if password == "" {
		password = "root"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
	for {
		time.Sleep(2 * time.Second)
		conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
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
