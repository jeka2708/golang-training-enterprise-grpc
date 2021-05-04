package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/jeka2708/golang-training-enterprise/pkg/data"
)

type pageDataWork struct {
	Title string
	Works []data.ResultWork
}

func (a dataAPI) getAllWork(writer http.ResponseWriter, request *http.Request) {
	w, err := a.data.ReadAllWorks()
	pD := pageDataWork{
		Title: "Список работ",
		Works: w,
	}
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get works"))
		if err != nil {
			log.Println(err)
		}
	}
	tmpl, _ := template.ParseFiles("web/work.html")
	err = tmpl.Execute(writer, pD)
}
func (a dataAPI) CreateWork(writer http.ResponseWriter, request *http.Request) {
	w := new(data.Work)
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
	_, err = a.data.AddWork(w.WorkerId, w.ServiceId)
	if err != nil {
		log.Println("work hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) DeleteWork(writer http.ResponseWriter, request *http.Request) {
	w := new(data.Work)
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
	err = a.data.DeleteByIdWork(w.Id)
	if err != nil {
		log.Println("work hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
