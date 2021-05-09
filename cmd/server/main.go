package main

import (
	"context"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/api"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/db"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"os"

	"time"
)

var (
	listen   = os.Getenv("LISTEN")
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if listen == "" {
		listen = ":8080"
	}
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
}

func main() {
	ctx := context.Background()
	ctx, error := context.WithCancel(ctx)
	ctx, error = context.WithTimeout(ctx, time.Second*30)
	if error != nil {
		log.Println(error)
	}

	listener, err := net.Listen("tcp", listen)
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
