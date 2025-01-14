package model

import "time"

type BusinessPartner struct {
	ID             int       `json:"id"`
	CompanyID      int       `json:"company_id"`
	CompanyName    string    `json:"company_name"`
	Representative string    `json:"representative"`
	Phone          string    `json:"phone"`
	PostalCode     string    `json:"postal_code"`
	Address        string    `json:"address"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
