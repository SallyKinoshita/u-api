package repository

import (
	"context"
	"time"

	"github.com/SallyKinoshita/u-api/internal/domain/model"
)

type Invoice interface {
	List(ctx context.Context, conn DBConn, startDate, endDate time.Time, page, limit int) ([]*model.Invoice, int, error)
	Create(ctx context.Context, conn DBConn, now time.Time, invoice *model.Invoice) error
}
