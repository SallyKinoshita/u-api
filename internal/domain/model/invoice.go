package model

import (
	"errors"
	"time"
)

const (
	feeRateDefault = 4.0
	taxRateDefault = 10.0
)

type Invoice struct {
	ID                int           `json:"id"`
	CompanyID         int           `json:"company_id"`
	BusinessPartnerID int           `json:"business_partner_id"`
	IssueDate         time.Time     `json:"issue_date"`
	Amount            float64       `json:"amount"`
	Fee               float64       `json:"fee"`
	FeeRate           float64       `json:"fee_rate"`
	Tax               float64       `json:"tax"`
	TaxRate           float64       `json:"tax_rate"`
	TotalAmount       float64       `json:"total_amount"`
	DueDate           time.Time     `json:"due_date"`
	Status            InvoiceStatus `json:"status"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

func NewInvoice(companyID, businessPartnerID int, issueDate, dueDate time.Time, paymentAmount float64) (*Invoice, error) {
	// バリデーション
	if companyID <= 0 {
		return nil, errors.New("invalid company ID")
	}
	if businessPartnerID <= 0 {
		return nil, errors.New("invalid business partner ID")
	}
	if paymentAmount <= 0 {
		return nil, errors.New("payment amount must be greater than zero")
	}

	// Fee と Tax の計算
	fee := paymentAmount * (feeRateDefault / 100)
	tax := fee * (taxRateDefault / 100)

	// TotalAmount (請求金額) の計算
	totalAmount := paymentAmount + fee + tax

	// 現在時刻
	//TODO: 本当はcontextで時刻を管理したい
	now := time.Now()

	invoice := &Invoice{
		CompanyID:         companyID,
		BusinessPartnerID: businessPartnerID,
		IssueDate:         issueDate,
		Amount:            paymentAmount, // 支払金額
		Fee:               fee,
		FeeRate:           feeRateDefault,
		Tax:               tax,
		TaxRate:           taxRateDefault,
		TotalAmount:       totalAmount,
		DueDate:           dueDate,
		Status:            StatusUnpaid,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	return invoice, nil
}
