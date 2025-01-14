package model

type InvoiceStatus string

const (
	StatusUnpaid InvoiceStatus = "unpaid"
	StatusPaid   InvoiceStatus = "paid"
	StatusPaying InvoiceStatus = "paying"
	StatusError  InvoiceStatus = "error"
)
