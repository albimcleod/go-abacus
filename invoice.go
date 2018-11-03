package goabacus

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
	ReferenceID   string `json:"referenceId"`
	InvoiceNumber string `json:"invoiceNumber"`
	Status        string `json:"status"`
}

//Invoices defines a list of invoices from Abacus
type Invoices struct {
	Pagination Pagination `json:"pagination"`
	Invoices   []Invoice  `json:"invoices"`
}
