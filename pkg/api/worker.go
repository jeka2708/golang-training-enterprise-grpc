package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/jeka2708/golang-training-enterprise/pkg/data"
)

type pageDataWorker struct {
	Title        string
	WorksClients []data.ResultWorker
}

func (a dataAPI) getAllWorkers(writer http.ResponseWriter, request *http.Request) {
	w, err := a.data.ReadAllWorkers()
	pD := pageDataWorker{
		Title:        "Список работников",
		WorksClients: w,
	}
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get workers"))
		if err != nil {
			log.Println(err)
		}
	}
	tmpl, _ := template.ParseFiles("web/workers.html")
	err = tmpl.Execute(writer, pD)
}
func (a dataAPI) CreateWorker(writer http.ResponseWriter, request *http.Request) {
	w := new(data.Worker)
	err := json.NewDecoder(request.Body).Decode(&w)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if w == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.AddWorker(*w)
	if err != nil {
		log.Println("workClients hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) DeleteWorker(writer http.ResponseWriter, request *http.Request) {
	w := new(data.Worker)
	err := json.NewDecoder(request.Body).Decode(&w)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if w == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.DeleteByIdWorker(w.Id)
	if err != nil {
		log.Println("work hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
