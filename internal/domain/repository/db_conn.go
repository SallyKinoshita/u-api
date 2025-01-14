package repository

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type DBConn interface {
	NewInsert() *bun.InsertQuery
	NewUpdate() *bun.UpdateQuery
	NewSelect() *bun.SelectQuery
	NewDelete() *bun.DeleteQuery
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}
