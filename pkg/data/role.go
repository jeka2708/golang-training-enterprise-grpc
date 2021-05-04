package data

import (
	"fmt"
)

type Role struct {
	Id         int
	Name       string
	DivisionId int
	Division   Division
}

type ResultRoles struct {
	Id           int `json:"id,string,omitempty"`
	Name         string
	DivisionName string
}

func (de DataEnterprise) ReadAllRoles() ([]ResultRoles, error) {

	var roles []ResultRoles
	result := de.db.Model(&Role{}).Select("roles.id, roles.name, divisions.division_name").
		Joins("left join divisions on divisions.id = roles.division_id").Scan(&roles)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read roles from database %w", result.Error)
	}
	return roles, nil
}

func (de DataEnterprise) addDivision(name string) (int, error) {
	var division = Division{
		DivisionName: name,
	}
	res := de.db.FirstOrCreate(&division, &division)

	if res.Error != nil {
		return -1, fmt.Errorf("can`t create division to database: %w", res.Error)
	}
	return division.Id, nil
}
func (de DataEnterprise) AddRole(nameRole, divisionName string) (int, error) {
	divisionId, err := de.addDivision(divisionName)
	if err != nil {
		return -1, err
	}
	var addRole = Role{
		Name:       nameRole,
		DivisionId: divisionId,
	}
	result := de.db.Create(&addRole)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create role to database: %w", result.Error)
	}
	return addRole.Id, nil
}

func (de DataEnterprise) UpdateRole(idRole int, nameRole, divisionName string) error {
	divisionId, err := de.addDivision(divisionName)
	if err != nil {
		return err
	}
	var addRole = Role{
		Name:       nameRole,
		DivisionId: divisionId,
	}
	result := de.db.Model(&Role{}).Where("id=?", idRole).Updates(&addRole)
	if result.Error != nil {
		return fmt.Errorf("can`t update role by id = %d, erorr: %w", idRole, result.Error)
	}
	return nil
}

func (de DataEnterprise) DeleteByIdRole(id int) error {
	result := de.db.Delete(&Role{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from roles by id = %d, error: %w", id, result.Error)
	}
	return nil
}
