package structs

import "time"

type Invoice struct {
	ID                            int                  `json:"id"`
	PassedToInventory             bool                 `json:"passed_to_inventory"`
	PassedToAccounting            bool                 `json:"passed_to_accounting"`
	IsInvoice                     bool                 `json:"is_invoice"`
	Registred                     bool                 `json:"registred"`
	Issuer                        string               `json:"issuer"`
	InvoiceNumber                 string               `json:"invoice_number"`
	Status                        string               `json:"status"`
	Type                          string               `json:"type"`
	TypeOfDecision                int                  `json:"type_of_decision"`
	TypeOfSubject                 int                  `json:"type_of_subject"`
	TypeOfContract                int                  `json:"type_of_contract"`
	SourceOfFunding               string               `json:"source_of_funding"`
	MunicipalityID                int                  `json:"municipality_id"`
	Supplier                      string               `json:"supplier"`
	GrossPrice                    float64              `json:"gross_price"`
	VATPrice                      float64              `json:"vat_price"`
	NetPrice                      float64              `json:"net_price"`
	SupplierID                    int                  `json:"supplier_id"`
	OrderID                       int                  `json:"order_id"`
	OrganizationUnitID            int                  `json:"organization_unit_id"`
	ActivityID                    int                  `json:"activity_id"`
	TaxAuthorityCodebookID        int                  `json:"tax_authority_codebook_id"`
	DateOfInvoice                 *time.Time           `json:"date_of_invoice"`
	ReceiptDate                   *time.Time           `json:"receipt_date"`
	DateOfPayment                 *time.Time           `json:"date_of_payment"`
	DateOfStart                   *time.Time           `json:"date_of_start"`
	SSSInvoiceReceiptDate         *time.Time           `json:"sss_invoice_receipt_date"`
	SSSProFormaInvoiceReceiptDate *time.Time           `json:"sss_pro_forma_invoice_receipt_date"`
	FileID                        int                  `json:"file_id"`
	ProFormaInvoiceFileID         int                  `json:"pro_forma_invoice_file_id"`
	BankAccount                   string               `json:"bank_account"`
	Description                   string               `json:"description"`
	ProFormaInvoiceDate           *time.Time           `json:"pro_forma_invoice_date"`
	ProFormaInvoiceNumber         string               `json:"pro_forma_invoice_number"`
	Articles                      []InvoiceArticles    `json:"articles"`
	AdditionalExpenses            []AdditionalExpenses `json:"additional_expenses"`
	CreatedAt                     time.Time            `json:"created_at"`
	UpdatedAt                     time.Time            `json:"updated_at"`
}

type InvoiceArticles struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	NetPrice      float64   `json:"net_price"`
	VatPrice      float64   `json:"vat_price"`
	VatPercentage int       `json:"vat_percentage"`
	Description   string    `json:"description"`
	InvoiceID     int       `json:"invoice_id"`
	Amount        int       `json:"amount"`
	AccountID     int       `json:"account_id"`
	CostAccountID int       `json:"cost_account_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AdditionalExpenseStatus string

const (
	AdditionalExpenseStatusCreated        AdditionalExpenseStatus = "Kreiran"
	AdditionalExpenseStatusWaitingPayment AdditionalExpenseStatus = "Na čekanju"
	AdditionalExpenseStatusPaid           AdditionalExpenseStatus = "Plaćen"
)

type AdditionalExpensesTitles string

var (
	NetTitle                                 AdditionalExpensesTitles = "Neto"
	ObligationTaxTitle                       AdditionalExpensesTitles = "Porez"
	ObligationSubTaxTitle                    AdditionalExpensesTitles = "Prirez"
	LaborFundTitle                           AdditionalExpensesTitles = "Fond rada"
	ContributionForPIOTitle                  AdditionalExpensesTitles = "PIO"
	ContributionForPIOEmployeeTitle          AdditionalExpensesTitles = "PIO na teret zaposlenog"
	ContributionForPIOEmployerTitle          AdditionalExpensesTitles = "PIO na teret poslodavca"
	ContributionForUnemploymentTitle         AdditionalExpensesTitles = "Nezaposlenost"
	ContributionForUnemploymentEmployeeTitle AdditionalExpensesTitles = "Nezaposlenost na teret zaposlenog"
	ContributionForUnemploymentEmployerTitle AdditionalExpensesTitles = "Nezaposlenost na teret poslodavca"

	SupplierTitle = "Dobavljač"
)

type AdditionalExpenses struct {
	ID                   int                      `json:"id"`
	Title                AdditionalExpensesTitles `json:"title"`
	ObligationType       string                   `json:"obligation_type"`
	ObligationNumber     string                   `json:"obligation_number"`
	ObligationSupplierID int                      `json:"obligation_supplier_id"`
	AccountID            int                      `json:"account_id"`
	Price                float64                  `json:"price"`
	SubjectID            int                      `json:"subject_id"`
	BankAccount          string                   `json:"bank_account"`
	InvoiceID            int                      `json:"invoice_id"`
	OrganizationUnitID   int                      `json:"organization_unit_id"`
	Status               string                   `json:"status"`
	CreatedAt            time.Time                `json:"created_at"`
	UpdatedAt            time.Time                `json:"updated_at"`
}
