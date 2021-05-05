package data

import (
	"fmt"
	"gorm.io/gorm"
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

type RoleData struct {
	db *gorm.DB
}

func NewRoleData(db *gorm.DB) *RoleData {
	return &RoleData{db: db}
}

func (r RoleData) ReadAllRoles() ([]ResultRoles, error) {

	var roles []ResultRoles
	result := r.db.Model(&Role{}).Select("roles.id, roles.name, divisions.division_name").
		Joins("left join divisions on divisions.id = roles.division_id").Scan(&roles)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read roles from database %w", result.Error)
	}
	return roles, nil
}

func (r RoleData) addDivision(name string) (int, error) {
	var division = Division{
		DivisionName: name,
	}
	res := r.db.FirstOrCreate(&division, &division)

	if res.Error != nil {
		return -1, fmt.Errorf("can`t create division to database: %w", res.Error)
	}
	return division.Id, nil
}
func (r RoleData) AddRole(nameRole, divisionName string) (int, error) {
	divisionId, err := r.addDivision(divisionName)
	if err != nil {
		return -1, err
	}
	var addRole = Role{
		Name:       nameRole,
		DivisionId: divisionId,
	}
	result := r.db.Create(&addRole)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create role to database: %w", result.Error)
	}
	return addRole.Id, nil
}

func (r RoleData) UpdateRole(idRole int, nameRole, divisionName string) error {
	divisionId, err := r.addDivision(divisionName)
	if err != nil {
		return err
	}
	var addRole = Role{
		Name:       nameRole,
		DivisionId: divisionId,
	}
	result := r.db.Model(&Role{}).Where("id=?", idRole).Updates(&addRole)
	if result.Error != nil {
		return fmt.Errorf("can`t update role by id = %d, erorr: %w", idRole, result.Error)
	}
	return nil
}

func (r RoleData) DeleteByIdRole(id int) error {
	result := r.db.Delete(&Role{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from roles by id = %d, error: %w", id, result.Error)
	}
	return nil
}
