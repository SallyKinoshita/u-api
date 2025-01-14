package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/SallyKinoshita/u-api/internal/domain/model"
	"github.com/SallyKinoshita/u-api/internal/domain/repository"
	"github.com/uptrace/bun"
)

type Invoice interface {
	List(ctx context.Context, startDate, endDate time.Time, page, limit int) ([]*model.Invoice, int, error)
	Create(ctx context.Context, partnerID int, paymentAmount float64, issueDate, dueDate time.Time) (*model.Invoice, error)
}

type InvoiceImpl struct {
	db          *bun.DB
	invoiceRepo repository.Invoice
}

func NewInvoice(db *bun.DB, invoiceRepo repository.Invoice) Invoice {
	return &InvoiceImpl{
		db:          db,
		invoiceRepo: invoiceRepo,
	}
}

func (i *InvoiceImpl) Create(ctx context.Context, partnerID int, paymentAmount float64, issueDate, dueDate time.Time) (*model.Invoice, error) {
	//TODO: ログインユーザーからCOmpanyIDを取得する。仮の値
	companyID := 1
	//TODO: 本当はcontextで時刻を管理したい
	now := time.Now()

	invoice, err := model.NewInvoice(companyID, partnerID, issueDate, dueDate, paymentAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to new invoice: %w", err)
	}

	err = i.invoiceRepo.Create(ctx, i.db, now, invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (i *InvoiceImpl) List(ctx context.Context, startDate time.Time, endDate time.Time, page int, limit int) ([]*model.Invoice, int, error) {
	invoices, total, err := i.invoiceRepo.List(ctx, i.db, startDate, endDate, page, limit)
	if err != nil {
		return nil, 0, err
	}
	return invoices, total, nil
}
