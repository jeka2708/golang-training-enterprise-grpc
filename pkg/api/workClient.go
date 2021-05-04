package api

import (
	"encoding/json"
	"github.com/jeka2708/golang-training-enterprise/pkg/data"
	"html/template"
	"log"
	"net/http"
)

type pageDataWorkClient struct {
	Title        string
	WorksClients []data.ResultClientWork
}

func (a dataAPI) getAllWorkClient(writer http.ResponseWriter, request *http.Request) {
	wc, err := a.data.ReadAllWorkClients()
	pD := pageDataWorkClient{
		Title:        "Список заказов клиентов",
		WorksClients: wc,
	}
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get worksClients"))
		if err != nil {
			log.Println(err)
		}
	}
	tmpl, _ := template.ParseFiles("web/works-clients.html")
	err = tmpl.Execute(writer, pD)
}
func (a dataAPI) CreateWorkClient(writer http.ResponseWriter, request *http.Request) {
	wc := new(data.WorkClient)
	err := json.NewDecoder(request.Body).Decode(&wc)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if wc == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.AddWorkClient(wc.ClientId, wc.WorkId)
	if err != nil {
		log.Println("workClients hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) DeleteWorkClient(writer http.ResponseWriter, request *http.Request) {
	wc := new(data.WorkClient)
	err := json.NewDecoder(request.Body).Decode(&wc)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if wc == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.DeleteByIdWorkClient(wc.Id)
	if err != nil {
		log.Println("work hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
