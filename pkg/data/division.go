package data

import (
	"fmt"
	"gorm.io/gorm"
)

type Division struct {
	Id           int `json:"id,string,omitempty"`
	DivisionName string
}

type DivisionData struct {
	db *gorm.DB
}

func NewDivisionData(db *gorm.DB) *DivisionData {
	return &DivisionData{db: db}
}

func (d DivisionData) ReadAllDivision() ([]Division, error) {
	var divisions []Division
	result := d.db.Find(&divisions)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read Divisions from database %w", result.Error)
	}
	return divisions, nil
}

func (d DivisionData) AddDivision(divisionName string) (int, error) {
	addDivision := Division{
		DivisionName: divisionName,
	}
	result := d.db.Create(&addDivision)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create Division to database: %w", result.Error)
	}
	return addDivision.Id, nil
}

func (d DivisionData) UpdateDivision(dd Division) error {
	result := d.db.Model(&d).Updates(&d)
	if result.Error != nil {
		return fmt.Errorf("can`t update Division by id = %d, erorr: %w", dd.Id, result.Error)
	}
	return nil
}

func (d DivisionData) DeleteByIdDivision(id int) error {
	result := d.db.Delete(&Division{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from Division by id = %d, error: %w", id, result.Error)
	}
	return nil
}
