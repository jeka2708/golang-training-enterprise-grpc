package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/jeka2708/golang-training-enterprise/pkg/data"
)

type pageDataRole struct {
	Title string
	Roles []data.ResultRoles
}

func (a dataAPI) getAllRole(writer http.ResponseWriter, request *http.Request) {
	r, err := a.data.ReadAllRoles()
	pD := pageDataRole{
		Title: "Список должностей",
		Roles: r,
	}
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get roles"))
		if err != nil {
			log.Println(err)
		}
	}
	tmpl, _ := template.ParseFiles("web/role.html")
	err = tmpl.Execute(writer, pD)
}
func (a dataAPI) CreateRole(writer http.ResponseWriter, request *http.Request) {
	r := new(data.ResultRoles)
	err := json.NewDecoder(request.Body).Decode(&r)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if r == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.AddRole(r.Name, r.DivisionName)
	if err != nil {
		log.Println("role hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
func (a dataAPI) DeleteRole(writer http.ResponseWriter, request *http.Request) {
	r := new(data.ResultRoles)
	err := json.NewDecoder(request.Body).Decode(&r)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if r == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.DeleteByIdRole(r.Id)
	if err != nil {
		log.Println("role hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
