package data

import (
	"fmt"
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

func (de DataEnterprise) ReadAllWorkers() ([]ResultWorker, error) {

	var workers []ResultWorker
	result := de.db.Model(&Worker{}).Select(
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

func (de DataEnterprise) AddWorker(w Worker) (int, error) {
	result := de.db.Create(&w)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create worker to database: %w", result.Error)
	}
	return w.Id, nil
}

func (de DataEnterprise) UpdateWorker(w Worker) error {
	result := de.db.Model(&w).Updates(&w)
	if result.Error != nil {
		return fmt.Errorf("can`t update worker by id = %d, erorr: %w", w.Id, result.Error)
	}
	return nil
}

func (de DataEnterprise) DeleteByIdWorker(id int) error {
	result := de.db.Delete(&Worker{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from worker by id = %d, error: %w", id, result.Error)
	}
	return nil
}
