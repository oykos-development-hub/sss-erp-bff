package dto

import (
	"bff/structs"
	"time"
)

type InvoiceResponseItem struct {
	ID                    int                          `json:"id"`
	PassedToInventory     bool                         `json:"passed_to_inventory"`
	PassedToAccounting    bool                         `json:"passed_to_accounting"`
	IsInvoice             bool                         `json:"is_invoice"`
	InvoiceNumber         string                       `json:"invoice_number"`
	Type                  string                       `json:"type"`
	TypeOfSubject         DropdownSimple               `json:"type_of_subject"`
	TypeOfContract        DropdownSimple               `json:"type_of_contract"`
	SourceOfFunding       DropdownSimple               `json:"source_of_funding"`
	SupplierTitle         string                       `json:"supplier_title"`
	Status                string                       `json:"status"`
	GrossPrice            float64                      `json:"gross_price"`
	NetPrice              float64                      `json:"net_price"`
	VATPrice              float64                      `json:"vat_price"`
	Supplier              DropdownSimple               `json:"supplier"`
	TaxAuthorityCodebook  DropdownSimple               `json:"tax_authority_codebook"`
	Activity              DropdownSimple               `json:"activity"`
	OrderID               int                          `json:"order_id"`
	OrganizationUnit      DropdownSimple               `json:"organization_unit"`
	ProFormaInvoiceDate   *time.Time                   `json:"pro_forma_invoice_date"`
	ProFormaInvoiceNumber string                       `json:"pro_forma_invoice_number"`
	DateOfInvoice         *time.Time                   `json:"date_of_invoice"`
	ReceiptDate           *time.Time                   `json:"receipt_date"`
	DateOfPayment         *time.Time                   `json:"date_of_payment"`
	DateOfStart           *time.Time                   `json:"date_of_start"`
	SSSInvoiceReceiptDate *time.Time                   `json:"sss_invoice_receipt_date"`
	File                  FileDropdownSimple           `json:"file"`
	ProFormaInvoiceFile   FileDropdownSimple           `json:"pro_forma_invoice_file_id"`
	BankAccount           string                       `json:"bank_account"`
	Description           string                       `json:"description"`
	Articles              []InvoiceArticleResponse     `json:"articles"`
	AdditionalExpenses    []AdditionalExpensesResponse `json:"additional_expenses"`
	CreatedAt             time.Time                    `json:"created_at"`
	UpdatedAt             time.Time                    `json:"updated_at"`
}

type InvoiceArticleResponse struct {
	ID            int            `json:"id"`
	Title         string         `json:"title"`
	NetPrice      float64        `json:"net_price"`
	VatPrice      float64        `json:"vat_price"`
	VatPercentage int            `json:"vat_percentage"`
	Description   string         `json:"description"`
	Amount        int            `json:"amount"`
	Account       DropdownSimple `json:"account"`
	CostAccount   DropdownSimple `json:"cost_account"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type AdditionalExpensesResponse struct {
	ID               int                             `json:"id"`
	Title            string                          `json:"title"`
	Account          DropdownSimple                  `json:"account"`
	Price            float32                         `json:"price"`
	Subject          DropdownSimple                  `json:"subject"`
	BankAccount      string                          `json:"bank_account"`
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

type GetAdditionalExpensesListResponseMS struct {
	Data  []structs.AdditionalExpenses `json:"data"`
	Total int                          `json:"total"`
}

type GetInvoiceListInputMS struct {
	Search             *string `json:"search"`
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	Year               *int    `json:"year"`
	Status             *string `json:"status"`
	SupplierID         *int    `json:"supplier_id"`
	OrderID            *int    `json:"order_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Type               *string `json:"type"`
	PassedToInventory  *bool   `json:"passed_to_inventory"`
}

type InvoiceArticleFilterDTO struct {
	InvoiceID *int `json:"invoice_id"`
}

type AdditionalExpensesListInputMS struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	InvoiceID          *int    `json:"invoice_id"`
	SubjectID          *int    `json:"subject_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Year               *int    `json:"year"`
	Status             *int    `json:"status"`
	Search             *string `json:"search"`
}

type TaxAuthorityCodebookFilter struct {
	Search *string `json:"search"`
	Active *bool   `json:"active"`
}

type GetTaxAuthorityCodebooksResponseMS struct {
	Data  []structs.TaxAuthorityCodebook `json:"data"`
	Total int                            `json:"total"`
}

type GetTaxAuthorityCodebookResponseMS struct {
	Data structs.TaxAuthorityCodebook `json:"data"`
}
