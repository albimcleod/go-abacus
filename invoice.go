package goabacus

// InvoiceStatus defines the status of an invoice
type InvoiceStatus int

const (
	//NonConfirmed
	NonConfirmed InvoiceStatus = 1 + iota

	//ConfirmedAccept
	ConfirmedAccept

	//ConfirmedReject
	ConfirmedReject
)

//Invoice defines an Invoice from Abacus POS
type Invoice struct {
	ReferenceID   string `json:"referenceId"`
	InvoiceNumber string `json:"invoiceNumber"`
	Status        string `json:"status"`
}
