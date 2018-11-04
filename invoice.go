package goabacus

import "time"

// InvoiceStatus defines the status of an invoice
type InvoiceStatus int

const (
	//NonConfirmed is a
	NonConfirmed InvoiceStatus = 1 + iota

	//ConfirmedAccept is a
	ConfirmedAccept

	//ConfirmedReject is a
	ConfirmedReject
)

//Invoice defines an Invoice from Abacus POS
type Invoice struct {
	InvoiceNumber      string        `json:"invoiceNumber"`
	Status             InvoiceStatus `json:"status"`
	TotalExcludeTax    float64       `json:"totalExcludeTax"`
	Total              float64       `json:"total"`
	DiscountExcludeTax float64       `json:"discountExcludeTax"`
	DiscountTax        float64       `json:"discountTax"`
	CreatedAt          time.Time     `json:"createdAt"`
}

//Invoices defines a list of invoices from Abacus
type Invoices struct {
	Pagination Pagination `json:"pagination"`
	Invoices   []Invoice  `json:"invoices"`
}
