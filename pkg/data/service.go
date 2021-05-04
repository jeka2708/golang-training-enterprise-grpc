package data

import (
	"fmt"
)

type Service struct {
	Id   int `json:"id,string,omitempty"`
	Name string
	Cost int `json:"cost,string,omitempty"`
}

func (de DataEnterprise) ReadAllServices() ([]Service, error) {
	var services []Service
	result := de.db.Find(&services)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read Services from database %w", result.Error)
	}
	return services, nil
}

func (de DataEnterprise) AddService(s Service) (int, error) {
	result := de.db.Create(&s)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create Service to database: %w", result.Error)
	}
	return s.Id, nil
}

func (de DataEnterprise) UpdateService(s Service) error {
	result := de.db.Model(&s).Updates(s)
	if result.Error != nil {
		return fmt.Errorf("can`t update Service by id = %d, erorr: %w", s.Id, result.Error)
	}
	return nil
}

func (de DataEnterprise) DeleteByIdService(id int) error {
	result := de.db.Delete(&Service{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from Services by id = %d, error: %w", id, result.Error)
	}
	return nil
}
