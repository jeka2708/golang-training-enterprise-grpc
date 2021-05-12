package data

import (
	"fmt"
	"gorm.io/gorm"
)

type Service struct {
	Id   int `json:"id,string,omitempty"`
	Name string
	Cost int `json:"cost,string,omitempty"`
}

type ServiceData struct {
	db *gorm.DB
}

func NewServiceData(db *gorm.DB) *ServiceData {
	return &ServiceData{db: db}
}

func (ss ServiceData) ReadAllServices() ([]Service, error) {
	var services []Service
	result := ss.db.Find(&services)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read Services from database %w", result.Error)
	}
	return services, nil
}

func (ss ServiceData) AddService(s Service) (int, error) {
	result := ss.db.Create(&s)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create Service to database: %w", result.Error)
	}
	return s.Id, nil
}

func (ss ServiceData) UpdateService(s Service) error {
	result := ss.db.Model(&s).Updates(s)
	if result.Error != nil {
		return fmt.Errorf("can`t update Service by id = %d, erorr: %w", s.Id, result.Error)
	}
	return nil
}

func (ss ServiceData) DeleteByIdService(id int) error {
	result := ss.db.Delete(&Service{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from Services by id = %d, error: %w", id, result.Error)
	}
	return nil
}
