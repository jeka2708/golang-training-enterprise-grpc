package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/jeka2708/golang-training-enterprise/pkg/data"
)

type pageData struct {
	Title   string
	Clients []data.Client
}

func (a dataAPI) getAllClients(writer http.ResponseWriter, request *http.Request) {
	clients, err := a.data.ReadAllClients()
	pD := pageData{
		Title:   "Список клиентов",
		Clients: clients,
	}
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get users"))
		if err != nil {
			log.Println(err)
		}
	}
	tmpl, _ := template.ParseFiles("web/clients.html")
	err = tmpl.Execute(writer, pD)
}
func (a dataAPI) CreateUser(writer http.ResponseWriter, request *http.Request) {
	client := new(data.Client)
	err := json.NewDecoder(request.Body).Decode(&client)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if client == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.AddClient(*client)
	if err != nil {
		log.Println("user hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) DeleteClient(writer http.ResponseWriter, request *http.Request) {
	client := new(data.Client)

	if id, err := strconv.Atoi(request.FormValue("id")); err == nil {
		client.Id = id
	}
	err := a.data.DeleteByIdClient(client.Id)
	if err != nil {
		log.Println("user hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) UpdateClient(writer http.ResponseWriter, request *http.Request) {
	client := new(data.Client)
	err := json.NewDecoder(request.Body).Decode(&client)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if client == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateClient(*client)
	if err != nil {
		log.Println("user hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
