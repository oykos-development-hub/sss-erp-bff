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
	Order                 DropdownSimple               `json:"order"`
	OrganizationUnit      DropdownSimple               `json:"organization_unit"`
	ProFormaInvoiceDate   *time.Time                   `json:"pro_forma_invoice_date"`
	ProFormaInvoiceNumber string                       `json:"pro_forma_invoice_number"`
	DateOfInvoice         *time.Time                   `json:"date_of_invoice"`
	ReceiptDate           *time.Time                   `json:"receipt_date"`
	DateOfPayment         *time.Time                   `json:"date_of_payment"`
	DateOfStart           *time.Time                   `json:"date_of_start"`
	SSSInvoiceReceiptDate *time.Time                   `json:"sss_invoice_receipt_date"`
	File                  FileDropdownSimple           `json:"file"`
	ProFormaInvoiceFile   FileDropdownSimple           `json:"pro_forma_invoice_file"`
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

type TaxAuthorityCodebookResponse struct {
	ID                                   int            `json:"id"`
	Title                                string         `json:"title"`
	Code                                 string         `json:"code"`
	Active                               bool           `json:"active"`
	TaxPercentage                        float64        `json:"tax_percentage"`
	TaxSupplier                          DropdownSimple `json:"tax_supplier"`
	ReleasePercentage                    float64        `json:"release_percentage"`
	PioPercentage                        float64        `json:"pio_percentage"`
	PioSupplier                          DropdownSimple `json:"pio_supplier"`
	PioPercentageEmployerPercentage      float64        `json:"pio_percentage_employer_percentage"`
	PioEmployerSupplier                  DropdownSimple `json:"pio_employer_supplier"`
	PioPercentageEmployeePercentage      float64        `json:"pio_percentage_employee_percentage"`
	PioEmployeeSupplier                  DropdownSimple `json:"pio_employee_supplier"`
	UnemploymentPercentage               float64        `json:"unemployment_percentage"`
	UnemploymentSupplier                 DropdownSimple `json:"unemployment_supplier"`
	UnemploymentEmployerPercentage       float64        `json:"unemployment_employer_percentage"`
	UnemploymentEmployerSupplier         DropdownSimple `json:"unemployment_employer_supplier"`
	UnemploymentEmployeePercentage       float64        `json:"unemployment_employee_percentage"`
	UnemploymentEmployeeSupplier         DropdownSimple `json:"unemployment_employee_supplier"`
	LaborFund                            float64        `json:"labor_fund"`
	LaborFundSupplier                    DropdownSimple `json:"labor_fund_supplier"`
	PreviousIncomePercentageLessThan700  float64        `json:"previous_income_percentage_less_than_700"`
	PreviousIncomePercentageLessThan1000 float64        `json:"previous_income_percentage_less_than_1000"`
	PreviousIncomePercentageMoreThan1000 float64        `json:"previous_income_percentage_more_than_1000"`
	Coefficient                          float64        `json:"coefficient"`
	CoefficientLess700                   float64        `json:"coefficient_less_700"`
	CoefficientLess1000                  float64        `json:"coefficient_less_1000"`
	CoefficientMore1000                  float64        `json:"coefficient_more_1000"`
	CreatedAt                            time.Time      `json:"created_at"`
	UpdatedAt                            time.Time      `json:"updated_at"`
}
