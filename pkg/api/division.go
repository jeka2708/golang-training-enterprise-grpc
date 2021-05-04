package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/jeka2708/golang-training-enterprise/pkg/data"
)

type pageDataDivision struct {
	Title     string
	Divisions []data.Division
}

func (a dataAPI) getAllDivision(writer http.ResponseWriter, request *http.Request) {
	dv, err := a.data.ReadAllDivision()

	pD := pageDataDivision{
		Title:     "Список отделов",
		Divisions: dv,
	}
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get divisions"))
		if err != nil {
			log.Println(err)
		}
	}
	tmpl, _ := template.ParseFiles("web/division.html")
	err = tmpl.Execute(writer, pD)
}
func (a dataAPI) CreateDivision(writer http.ResponseWriter, request *http.Request) {
	dv := new(data.Division)
	err := json.NewDecoder(request.Body).Decode(&dv)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if dv == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.AddDivision(dv.DivisionName)
	if err != nil {
		log.Println("division hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) DeleteDivision(writer http.ResponseWriter, request *http.Request) {
	dv := new(data.Division)
	err := json.NewDecoder(request.Body).Decode(&dv)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if dv == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.DeleteByIdDivision(dv.Id)
	if err != nil {
		log.Println("division hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) UpdateDivision(writer http.ResponseWriter, request *http.Request) {
	dv := new(data.Division)
	err := json.NewDecoder(request.Body).Decode(&dv)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if dv == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateDivision(*dv)
	if err != nil {
		log.Println("division hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
