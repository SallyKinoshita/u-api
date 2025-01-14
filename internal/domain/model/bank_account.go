package model

import "time"

type BankAccount struct {
	ID                int       `json:"id"`
	BusinessPartnerID int       `json:"business_partner_id"`
	BankName          string    `json:"bank_name"`
	BranchName        string    `json:"branch_name"`
	AccountNumber     string    `json:"account_number"`
	AccountName       string    `json:"account_name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
