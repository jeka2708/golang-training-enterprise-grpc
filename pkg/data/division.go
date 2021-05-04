package data

import (
	"fmt"
)

type Division struct {
	Id           int `json:"id,string,omitempty"`
	DivisionName string
}

func (de DataEnterprise) ReadAllDivision() ([]Division, error) {
	var Divisions []Division
	result := de.db.Find(&Divisions)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read Divisions from database %w", result.Error)
	}
	return Divisions, nil
}

func (de DataEnterprise) AddDivision(divisionName string) (int, error) {
	addDivision := Division{
		DivisionName: divisionName,
	}
	result := de.db.Create(&addDivision)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create Division to database: %w", result.Error)
	}
	return addDivision.Id, nil
}

func (de DataEnterprise) UpdateDivision(d Division) error {
	result := de.db.Model(&d).Updates(&d)
	if result.Error != nil {
		return fmt.Errorf("can`t update Division by id = %d, erorr: %w", d.Id, result.Error)
	}
	return nil
}

func (de DataEnterprise) DeleteByIdDivision(id int) error {
	result := de.db.Delete(&Division{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from Division by id = %d, error: %w", id, result.Error)
	}
	return nil
}
