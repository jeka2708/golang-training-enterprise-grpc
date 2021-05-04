package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeka2708/golang-training-enterprise/pkg/api"
	"github.com/jeka2708/golang-training-enterprise/pkg/data"
	"github.com/jeka2708/golang-training-enterprise/pkg/db"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "enterprise"
	password = "root"
	sslmode  = "disable"
)

func main() {
	conn, error := db.GetConnection(host, port, user, dbname, password, sslmode)
	if error != nil {
		log.Println(error)
	}
	r := mux.NewRouter()
	enterprise := data.NewDataEnterprise(conn)
	// 4. send enterprise layer to api layer
	api.ServeUserResource(r, *enterprise)
	// 5. cors for making requests from any domain
	r.Use(mux.CORSMethodMiddleware(r))
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web/"))))
	// 6. start server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server Listen port...")
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}
