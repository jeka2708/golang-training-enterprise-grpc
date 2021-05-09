package data

import (
	"fmt"
	"gorm.io/gorm"
)

type Client struct {
	Id           int `json:"id,string,omitempty"`
	FirstNameC   string
	LastNameC    string
	MiddleNameC  string
	PhoneNumberC string
}

type ClientData struct {
	db *gorm.DB
}

func NewClientData(db *gorm.DB) *ClientData {
	return &ClientData{db: db}
}

func (cc ClientData) ReadAllClients() ([]Client, error) {
	var clients []Client
	result := cc.db.Find(&clients)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read Clients from database %w", result.Error)
	}
	return clients, nil
}

func (cc ClientData) AddClient(c Client) (int, error) {
	result := cc.db.Create(&c)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create Clients to database: %w", result.Error)
	}
	return c.Id, nil
}

func (cc ClientData) UpdateClient(c Client) error {
	result := cc.db.Model(&c).Updates(&c)
	if result.Error != nil {
		return fmt.Errorf("can`t update Clients by id = %d, erorr: %w", c.Id, result.Error)
	}
	return nil
}

func (cc ClientData) DeleteByIdClient(id int) error {
	result := cc.db.Delete(&Client{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from Clients by id = %d, error: %w", id, result.Error)
	}
	return nil
}
