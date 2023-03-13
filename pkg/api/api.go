package api

import (
	"github.com/gorilla/mux"
	"github.com/jeka2708/golang-training-enterprise/pkg/data"
)

type dataAPI struct {
	data *data.DataEnterprise
}

func ServeUserResource(r *mux.Router, data data.DataEnterprise) {
	api := &dataAPI{data: &data}
	r.HandleFunc("/clients", api.getAllClients).Methods("GET")
	r.HandleFunc("/divisions", api.getAllDivision).Methods("GET")
	r.HandleFunc("/roles", api.getAllRole).Methods("GET")
	r.HandleFunc("/services", api.getAllServices).Methods("GET")
	r.HandleFunc("/works", api.getAllWork).Methods("GET")
	r.HandleFunc("/works-clients", api.getAllWorkClient).Methods("GET")
	r.HandleFunc("/workers", api.getAllWorkers).Methods("GET")

	r.HandleFunc("/create-client", api.CreateUser).Methods("POST")
	r.HandleFunc("/create-division", api.CreateDivision).Methods("POST")
	r.HandleFunc("/create-role", api.CreateRole).Methods("POST")
	r.HandleFunc("/create-service", api.CreateService).Methods("POST")
	r.HandleFunc("/create-work", api.CreateWork).Methods("POST")
	r.HandleFunc("/create-work-clients", api.CreateWorkClient).Methods("POST")
	r.HandleFunc("/create-worker", api.CreateWorker).Methods("POST")

	r.HandleFunc("/delete-client", api.DeleteClient).Methods("DELETE")
	r.HandleFunc("/delete-division", api.DeleteDivision).Methods("DELETE")
	r.HandleFunc("/delete-role", api.DeleteRole).Methods("DELETE")
	r.HandleFunc("/delete-service", api.DeleteService).Methods("DELETE")
	r.HandleFunc("/delete-work", api.DeleteWork).Methods("DELETE")
	r.HandleFunc("/delete-work-clients", api.DeleteWorkClient).Methods("DELETE")
	r.HandleFunc("/delete-worker", api.DeleteWorker).Methods("DELETE")

	r.HandleFunc("/update-client", api.UpdateClient).Methods("PUT")
	r.HandleFunc("/update-division", api.UpdateDivision).Methods("PUT")
	r.HandleFunc("/update-service", api.UpdateService).Methods("PUT")
}
