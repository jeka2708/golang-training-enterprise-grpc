package data

import (
	"fmt"
	"gorm.io/gorm"
)

type WorkClient struct {
	Id       int `json:"Id,string,omitempty"`
	ClientId int `json:"ClientId,string,omitempty"`
	WorkId   int `json:"WorkId,string,omitempty"`
	Client   Client
	Work     Work
}

type ResultClientWork struct {
	Id           int
	Name         string
	Cost         string
	FirstName    string
	LastName     string
	MiddleName   string
	PhoneNumber  string
	FirstNameC   string
	LastNameC    string
	MiddleNameC  string
	PhoneNumberC string
}

type WorkClientData struct {
	db *gorm.DB
}

func NewWorkClientData(db *gorm.DB) *WorkClientData {
	return &WorkClientData{db: db}
}

func (wrc WorkClientData) ReadAllWorkClients() ([]ResultClientWork, error) {

	var clientWorks []ResultClientWork
	ResultClientWork := wrc.db.Model(&WorkClient{}).Select("work_clients.id, services.name, services.cost, " +
		"workers.first_name, workers.last_name, workers.middle_name, workers.phone_number," +
		"clients.first_name_c, clients.last_name_c, clients.middle_name_c, clients.phone_number_c").
		Joins("left join works on work_clients.work_id = works.id").
		Joins("left join services on services.id = works.service_id").
		Joins("left join workers on workers.id = works.worker_id").
		Joins("left join clients on clients.id = work_clients.client_id").Scan(&clientWorks)
	if ResultClientWork.Error != nil {
		return nil, fmt.Errorf("can`t read client_works from database %w", ResultClientWork.Error)
	}
	return clientWorks, nil
}

func (wrc WorkClientData) AddWorkClient(clientId, workId int) (int, error) {
	addWorkClients := WorkClient{
		ClientId: clientId,
		WorkId:   workId,
	}
	ResultClientWork := wrc.db.Create(&addWorkClients)
	if ResultClientWork.Error != nil {
		return -1, fmt.Errorf("can`t create workClient to database: %w", ResultClientWork.Error)
	}
	return addWorkClients.Id, nil
}

func (wrc WorkClientData) UpdateWorkClient(wc WorkClient) error {
	ResultClientWork := wrc.db.Model(&wc).Updates(&wc)
	if ResultClientWork.Error != nil {
		return fmt.Errorf("can`t update workClient by id = %d, erorr: %w", wc.Id, ResultClientWork.Error)
	}
	return nil
}

func (wrc WorkClientData) DeleteByIdWorkClient(id int) error {
	ResultClientWork := wrc.db.Delete(&WorkClient{}, id)
	if ResultClientWork.Error != nil {
		return fmt.Errorf("can`t delete from client_works by id = %d, error: %w", id, ResultClientWork.Error)
	}
	return nil
}
