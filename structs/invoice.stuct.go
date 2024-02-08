package structs

import "time"

type Invoice struct {
	ID                    int               `json:"id"`
	InvoiceNumber         string            `json:"invoice_number"`
	Status                string            `json:"status"`
	GrossPrice            float64           `json:"gross_price"`
	VATPrice              float64           `json:"vat_price"`
	SupplierID            int               `json:"supplier_id"`
	OrderID               int               `json:"order_id"`
	OrganizationUnitID    int               `json:"organization_unit_id"`
	DateOfInvoice         time.Time         `json:"date_of_invoice"`
	ReceiptDate           time.Time         `json:"receipt_date"`
	DateOfPayment         time.Time         `json:"date_of_payment"`
	SSSInvoiceReceiptDate *time.Time        `json:"sss_invoice_receipt_date"`
	FileID                int               `json:"file_id"`
	BankAccount           string            `json:"bank_account"`
	Description           string            `json:"description"`
	Articles              []InvoiceArticles `json:"articles"`
	CreatedAt             time.Time         `json:"created_at"`
	UpdatedAt             time.Time         `json:"updated_at"`
}

type InvoiceArticles struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	NetPrice      float64   `json:"net_price"`
	VatPrice      float64   `json:"vat_price"`
	Description   string    `json:"description"`
	InvoiceID     int       `json:"invoice_id"`
	AccountID     int       `json:"account_id"`
	CostAccountID int       `json:"cost_account_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
