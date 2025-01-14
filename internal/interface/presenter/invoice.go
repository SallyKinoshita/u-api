package presenter

import (
	"strconv"
	"time"

	"github.com/SallyKinoshita/u-api/internal/domain/model"
	"github.com/SallyKinoshita/u-api/internal/gen/openapi"
	"github.com/oapi-codegen/runtime/types"
)

func ToInvoiceListResponse(invoices []*model.Invoice, total int, page, perPage int) openapi.InvoiceListResponse {
	var responseInvoices []openapi.InvoiceResponse

	for _, invoice := range invoices {
		responseInvoices = append(responseInvoices, openapi.InvoiceResponse{
			Id:             intToStrPtr(invoice.ID),
			PartnerId:      intToStrPtr(invoice.BusinessPartnerID),
			PaymentAmount:  float64ToIntPtr(invoice.TotalAmount),
			Status:         (*openapi.InvoiceResponseStatus)(&invoice.Status),
			IssueDate:      timeToDatePtr(invoice.IssueDate),
			PaymentDueDate: timeToDatePtr(invoice.DueDate),
			CreatedAt:      &invoice.CreatedAt,
			UpdatedAt:      &invoice.UpdatedAt,
		})
	}

	return openapi.InvoiceListResponse{
		Invoices: &responseInvoices,
		Total:    &total,
		Page:     &page,
		PerPage:  &perPage,
	}
}

func intToStrPtr(int int) *string {
	str := strconv.Itoa(int)
	return &str
}

func float64ToIntPtr(float float64) *int {
	int := int(float)
	return &int
}

func timeToDatePtr(time time.Time) *types.Date {
	date := types.Date{Time: time}
	return &date
}
