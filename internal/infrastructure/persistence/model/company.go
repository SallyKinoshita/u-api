package persistencemodel

import (
	"time"

	"github.com/uptrace/bun"
)

type Company struct {
	bun.BaseModel `bun:"table:company"`

	ID             int       `bun:"id,pk,autoincrement"`
	CompanyName    string    `bun:"company_name,notnull"`
	Representative string    `bun:"representative,notnull"`
	Phone          string    `bun:"phone,notnull"`
	PostalCode     string    `bun:"postal_code,notnull"`
	Address        string    `bun:"address,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull"`
	UpdatedAt      time.Time `bun:"updated_at,notnull"`
}
