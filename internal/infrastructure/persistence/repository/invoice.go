package persistencerepository

import (
	"context"
	"time"

	"github.com/SallyKinoshita/u-api/internal/domain/model"
	"github.com/SallyKinoshita/u-api/internal/domain/repository"
	persistencemodel "github.com/SallyKinoshita/u-api/internal/infrastructure/persistence/model"
)

type Invoice struct {
	dbInvoice persistencemodel.Invoice
}

func NewInvoice() *Invoice {
	return &Invoice{}
}

func (i *Invoice) List(ctx context.Context, conn repository.DBConn, startDate, endDate time.Time, page, limit int) ([]*model.Invoice, int, error) {
	var dbInvoices []*persistencemodel.Invoice
	var totalCount int

	countQuery := conn.NewSelect().
		Model((*persistencemodel.Invoice)(nil)).
		Where("due_date BETWEEN ? AND ?", startDate, endDate)

	totalCount, err := countQuery.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	query := conn.NewSelect().
		Model(&dbInvoices).
		Where("due_date BETWEEN ? AND ?", startDate, endDate).
		Order("due_date DESC").
		Limit(limit).
		Offset((page - 1) * limit)

	if err := query.Scan(ctx); err != nil {
		return nil, 0, err
	}

	invoices := make([]*model.Invoice, len(dbInvoices))
	for i, dbInvoice := range dbInvoices {
		invoices[i] = &model.Invoice{
			ID:                dbInvoice.ID,
			CompanyID:         dbInvoice.CompanyID,
			BusinessPartnerID: dbInvoice.BusinessPartnerID,
			IssueDate:         dbInvoice.IssueDate,
			Amount:            dbInvoice.Amount,
			Fee:               dbInvoice.Fee,
			FeeRate:           dbInvoice.FeeRate,
			Tax:               dbInvoice.Tax,
			TaxRate:           dbInvoice.TaxRate,
			TotalAmount:       dbInvoice.TotalAmount,
			DueDate:           dbInvoice.DueDate,
			Status:            model.InvoiceStatus(dbInvoice.Status),
			CreatedAt:         dbInvoice.CreatedAt,
			UpdatedAt:         dbInvoice.UpdatedAt,
		}
	}
	return invoices, totalCount, nil
}

func (i *Invoice) Create(ctx context.Context, conn repository.DBConn, now time.Time, invoice *model.Invoice) error {
	dbInvoice := &persistencemodel.Invoice{
		ID:                invoice.ID,
		CompanyID:         invoice.CompanyID,
		BusinessPartnerID: invoice.BusinessPartnerID,
		IssueDate:         invoice.IssueDate,
		Amount:            invoice.Amount,
		Fee:               invoice.Fee,
		FeeRate:           invoice.FeeRate,
		Tax:               invoice.Tax,
		TaxRate:           invoice.TaxRate,
		TotalAmount:       invoice.TotalAmount,
		DueDate:           invoice.DueDate,
		Status:            string(invoice.Status),
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	_, err := conn.NewInsert().
		Model(dbInvoice).
		Exec(ctx)
	return err
}
