package data

import (
	"fmt"
	"gorm.io/gorm"
)

type Work struct {
	Id        int `json:"Id,string,omitempty"`
	WorkerId  int `json:"WorkerId,string,omitempty"`
	ServiceId int `json:"ServiceId,string,omitempty"`
	Worker    Worker
	Service   Service
}

type ResultWork struct {
	Id          int
	Name        string
	Cost        string
	FirstName   string
	LastName    string
	MiddleName  string
	PhoneNumber string
}

type WorkData struct {
	db *gorm.DB
}

func NewWorkData(db *gorm.DB) *WorkData {
	return &WorkData{db: db}
}

func (ww WorkData) ReadAllWorks() ([]ResultWork, error) {

	var works []ResultWork
	result := ww.db.Model(&Work{}).Select("works.id, services.name, services.cost, " +
		"workers.first_name, workers.last_name, workers.middle_name, workers.phone_number").
		Joins("left join services on services.id = works.service_id").
		Joins("left join workers on works.worker_id = workers.id").Scan(&works)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read works from database %w", result.Error)
	}
	return works, nil
}

func (ww WorkData) AddWork(workerId, serviceId int) (int, error) {
	addWork := Work{
		WorkerId:  workerId,
		ServiceId: serviceId,
	}
	result := ww.db.Create(&addWork)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create work to database: %w", result.Error)
	}
	return addWork.Id, nil
}

func (ww WorkData) UpdateWork(w Work) error {
	result := ww.db.Model(&w).Updates(w)
	if result.Error != nil {
		return fmt.Errorf("can`t update works by id = %d, erorr: %w", w.Id, result.Error)
	}
	return nil
}

func (ww WorkData) DeleteByIdWork(id int) error {
	result := ww.db.Delete(&Work{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from works by id = %d, error: %w", id, result.Error)
	}
	return nil
}
