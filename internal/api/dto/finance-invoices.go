package dto

import (
	"bff/structs"
	"time"
)

type InvoiceResponseItem struct {
	ID                    int                          `json:"id"`
	InvoiceNumber         string                       `json:"invoice_number"`
	Type                  string                       `json:"type"`
	TypeOfSubject         DropdownSimple               `json:"type_of_subject"`
	TypeOfContract        DropdownSimple               `json:"type_of_contract"`
	SourceOfFunding       DropdownSimple               `json:"source_of_funding"`
	SupplierTitle         string                       `json:"supplier_title"`
	Status                string                       `json:"status"`
	GrossPrice            float64                      `json:"gross_price"`
	VATPrice              float64                      `json:"vat_price"`
	Supplier              DropdownSimple               `json:"supplier"`
	Activity              DropdownSimple               `json:"activity"`
	OrderID               int                          `json:"order_id"`
	OrganizationUnit      DropdownSimple               `json:"organization_unit"`
	DateOfInvoice         time.Time                    `json:"date_of_invoice"`
	ReceiptDate           time.Time                    `json:"receipt_date"`
	DateOfPayment         time.Time                    `json:"date_of_payment"`
	DateOfStart           time.Time                    `json:"date_of_start"`
	SSSInvoiceReceiptDate *time.Time                   `json:"sss_invoice_receipt_date"`
	File                  FileDropdownSimple           `json:"file"`
	BankAccount           string                       `json:"bank_account"`
	Description           string                       `json:"description"`
	Articles              []InvoiceArticleResponse     `json:"articles"`
	AdditionalExpenses    []AdditionalExpensesResponse `json:"additional_expenses"`
	CreatedAt             time.Time                    `json:"created_at"`
	UpdatedAt             time.Time                    `json:"updated_at"`
}

type InvoiceArticleResponse struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	NetPrice    float64        `json:"net_price"`
	VatPrice    float64        `json:"vat_price"`
	Description string         `json:"description"`
	Account     DropdownSimple `json:"account"`
	CostAccount DropdownSimple `json:"cost_account"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type AdditionalExpensesResponse struct {
	ID               int                             `json:"id"`
	Title            string                          `json:"title"`
	Account          DropdownSimple                  `json:"account"`
	Price            float32                         `json:"price"`
	Subject          DropdownSimple                  `json:"subject"`
	BankAccount      int                             `json:"bank_account"`
	Invoice          DropdownSimple                  `json:"invoice"`
	OrganizationUnit DropdownSimple                  `json:"organization_unit"`
	Status           structs.AdditionalExpenseStatus `json:"status"`
	CreatedAt        time.Time                       `json:"created_at"`
	UpdatedAt        time.Time                       `json:"updated_at"`
}

type GetInvoiceResponseMS struct {
	Data structs.Invoice `json:"data"`
}

type GetInvoiceListResponseMS struct {
	Data  []structs.Invoice `json:"data"`
	Total int               `json:"total"`
}

type GetInvoiceArticleResponseMS struct {
	Data structs.InvoiceArticles `json:"data"`
}

type GetInvoiceArticleListResponseMS struct {
	Data  []structs.InvoiceArticles `json:"data"`
	Total int                       `json:"total"`
}

type GetInvoiceListInputMS struct {
	Search             *string `json:"search"`
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	Year               *int    `json:"year"`
	Status             *string `json:"status"`
	SupplierID         *int    `json:"supplier_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Type               *string `json:"type"`
}

type InvoiceArticleFilterDTO struct {
	InvoiceID *int `json:"invoice_id"`
}
