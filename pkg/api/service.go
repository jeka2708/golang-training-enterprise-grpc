package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/jeka2708/golang-training-enterprise/pkg/data"
)

type pageDataService struct {
	Title   string
	Service []data.Service
}

func (a dataAPI) getAllServices(writer http.ResponseWriter, request *http.Request) {
	r, err := a.data.ReadAllServices()
	pD := pageDataService{
		Title:   "Список выполняемых работ",
		Service: r,
	}
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get service"))
		if err != nil {
			log.Println(err)
		}
	}
	tmpl, _ := template.ParseFiles("web/service.html")
	err = tmpl.Execute(writer, pD)
}
func (a dataAPI) CreateService(writer http.ResponseWriter, request *http.Request) {
	s := new(data.Service)
	err := json.NewDecoder(request.Body).Decode(&s)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if s == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.AddService(*s)
	if err != nil {
		log.Println("service hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) DeleteService(writer http.ResponseWriter, request *http.Request) {
	s := new(data.Service)
	err := json.NewDecoder(request.Body).Decode(&s)
	log.Println(s)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if s == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.DeleteByIdService(s.Id)
	if err != nil {
		log.Println("service hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) UpdateService(writer http.ResponseWriter, request *http.Request) {
	s := new(data.Service)
	err := json.NewDecoder(request.Body).Decode(&s)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if s == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateService(*s)
	if err != nil {
		log.Println("service hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
