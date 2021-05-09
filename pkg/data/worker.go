package data

import (
	"fmt"
	"gorm.io/gorm"
)

type Worker struct {
	Id          int `json:"id,string,omitempty"`
	FirstName   string
	LastName    string
	MiddleName  string
	PhoneNumber string
	RoleId      int `json:"RoleId,string,omitempty"`
	Role        Role
}

type ResultWorker struct {
	Id          int
	FirstName   string
	LastName    string
	MiddleName  string
	PhoneNumber string
	Name        string
}

type WorkerData struct {
	db *gorm.DB
}

func NewWorkerData(db *gorm.DB) *WorkerData {
	return &WorkerData{db: db}
}

func (ww WorkerData) ReadAllWorkers() ([]ResultWorker, error) {

	var workers []ResultWorker
	result := ww.db.Model(&Worker{}).Select(
		"workers.id, " +
			"workers.first_name, " +
			"workers.last_name," +
			"workers.middle_name, " +
			"workers.phone_number, " +
			"roles.name").
		Joins("left join roles on roles.id = workers.role_id").Scan(&workers)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read worker from database %w", result.Error)
	}
	return workers, nil
}

func (ww WorkerData) AddWorker(w Worker) (int, error) {
	result := ww.db.Create(&w)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create worker to database: %w", result.Error)
	}
	return w.Id, nil
}

func (ww WorkerData) UpdateWorker(w Worker) error {
	result := ww.db.Model(&w).Updates(&w)
	if result.Error != nil {
		return fmt.Errorf("can`t update worker by id = %d, erorr: %w", w.Id, result.Error)
	}
	return nil
}

func (ww WorkerData) DeleteByIdWorker(id int) error {
	result := ww.db.Delete(&Worker{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from worker by id = %d, error: %w", id, result.Error)
	}
	return nil
}
