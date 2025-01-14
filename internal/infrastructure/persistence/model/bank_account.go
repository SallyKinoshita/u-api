package persistencemodel

import (
	"time"

	"github.com/uptrace/bun"
)

type BankAccount struct {
	bun.BaseModel `bun:"table:bank_account"`

	ID                int       `bun:"id,pk,autoincrement"`
	BusinessPartnerID int       `bun:"business_partner_id,notnull"`
	BankName          string    `bun:"bank_name,notnull"`
	BranchName        string    `bun:"branch_name,notnull"`
	AccountNumber     string    `bun:"account_number,notnull"`
	AccountName       string    `bun:"account_name,notnull"`
	CreatedAt         time.Time `bun:"created_at,notnull"`
	UpdatedAt         time.Time `bun:"updated_at,notnull"`
}
