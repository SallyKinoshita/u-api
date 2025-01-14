package persistencemodel

import (
	"time"

	"github.com/uptrace/bun"
)

type Invoice struct {
	bun.BaseModel `bun:"table:invoice"`

	ID                int       `bun:"id,pk,autoincrement"`
	CompanyID         int       `bun:"company_id,notnull"`
	BusinessPartnerID int       `bun:"business_partner_id,notnull"`
	IssueDate         time.Time `bun:"issue_date,notnull"`
	Amount            float64   `bun:"amount,notnull"`
	Fee               float64   `bun:"fee,notnull"`
	FeeRate           float64   `bun:"fee_rate,notnull"`
	Tax               float64   `bun:"tax,notnull"`
	TaxRate           float64   `bun:"tax_rate,notnull"`
	TotalAmount       float64   `bun:"total_amount,notnull"`
	DueDate           time.Time `bun:"due_date,notnull"`
	Status            string    `bun:"status,notnull"`
	CreatedAt         time.Time `bun:"created_at,notnull"`
	UpdatedAt         time.Time `bun:"updated_at,notnull"`
}
