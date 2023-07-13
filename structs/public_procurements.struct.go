package structs

type PublicProcurementPlan struct {
	Id               int    `json:"id"`
	Year             string `json:"year"`
	Title            string `json:"title"`
	Active           bool   `json:"active"`
	SerialNumber     string `json:"serial_number"`
	IsPreBudget      bool   `json:"is_pre_budget"`
	DateOfPublishing string `json:"date_of_publishing"`
	DateOfClosing    string `json:"date_of_closing"`
	PreBudgetId      int    `json:"pre_budget_id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	FileId           int    `json:"file_id"`
}

type PublicProcurementItem struct {
	Id                int       `json:"id"`
	BudgetIndentId    int       `json:"budget_indent_id" validate:"required"`
	PlanId            int       `json:"plan_id" validate:"required"`
	IsOpenProcurement bool      `json:"is_open_procurement" validate:"required"`
	Title             string    `json:"title" validate:"required"`
	ArticleType       string    `json:"article_type" validate:"required"`
	Status            *string   `json:"status"`
	SerialNumber      *string   `json:"serial_number"`
	DateOfPublishing  *JSONDate `json:"date_of_publishing"`
	DateOfAwarding    *JSONDate `json:"date_of_awarding"`
	CreatedAt         string    `json:"created_at"`
	UpdatedAt         string    `json:"updated_at"`
	FileId            *int      `json:"file_id"`
}

type PublicProcurementArticle struct {
	Id                  int    `json:"id"`
	BudgetIndentId      int    `json:"budget_indent_id"`
	PublicProcurementId int    `json:"public_procurement_id"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	NetPrice            string `json:"net_price"`
	VatPercentage       string `json:"vat_percentage"`
}

type PublicProcurementLimit struct {
	Id                  int    `json:"id"`
	PublicProcurementId int    `json:"public_procurement_id"`
	OrganizationUnitId  int    `json:"organization_unit_id"`
	Limit               string `json:"limit"`
}

type PublicProcurementOrganizationUnitArticle struct {
	Id                         int    `json:"id"`
	PublicProcurementArticleId int    `json:"public_procurement_article_id"`
	OrganizationUnitId         int    `json:"organization_unit_id"`
	Amount                     int    `json:"amount"`
	Status                     string `json:"status"`
	IsRejected                 bool   `json:"is_rejected"`
	RejectedDescription        string `json:"rejected_description"`
}

type PublicProcurementContract struct {
	Id                  int    `json:"id"`
	PublicProcurementId int    `json:"public_procurement_id"`
	SupplierId          int    `json:"supplier_id"`
	SerialNumber        string `json:"serial_number"`
	DateOfSigning       string `json:"date_of_signing"`
	DateOfExpiry        string `json:"date_of_expiry"`
	NetValue            string `json:"net_value"`
	GrossValue          string `json:"gross_value"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	FileId              int    `json:"file_id"`
}

type PublicProcurementContractArticle struct {
	Id                          int    `json:"id"`
	PublicProcurementArticleId  int    `json:"public_procurement_article_id"`
	PublicProcurementContractId int    `json:"public_procurement_contract_id"`
	Amount                      string `json:"amount"`
	NetValue                    string `json:"net_value"`
	GrossValue                  string `json:"gross_value"`
	VatPercentage               string `json:"vat_percentage"`
}
