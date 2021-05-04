package data

import (
	"fmt"
)

type Client struct {
	Id           int `json:"id,string,omitempty"`
	FirstNameC   string
	LastNameC    string
	MiddleNameC  string
	PhoneNumberC string
}

func (de DataEnterprise) ReadAllClients() ([]Client, error) {
	var clients []Client
	result := de.db.Find(&clients)
	if result.Error != nil {
		return nil, fmt.Errorf("can`t read Clients from database %w", result.Error)
	}
	return clients, nil
}

func (de DataEnterprise) AddClient(c Client) (int, error) {
	result := de.db.Create(&c)
	if result.Error != nil {
		return -1, fmt.Errorf("can`t create Clients to database: %w", result.Error)
	}
	return c.Id, nil
}

func (de DataEnterprise) UpdateClient(c Client) error {
	result := de.db.Model(&c).Updates(&c)
	if result.Error != nil {
		return fmt.Errorf("can`t update Clients by id = %d, erorr: %w", c.Id, result.Error)
	}
	return nil
}

func (de DataEnterprise) DeleteByIdClient(id int) error {
	result := de.db.Delete(&Client{}, id)
	if result.Error != nil {
		return fmt.Errorf("can`t delete from Clients by id = %d, error: %w", id, result.Error)
	}
	return nil
}
