package structs

type PublicProcurementPlan struct {
	Id               int       `json:"id"`
	Year             string    `json:"year"`
	Title            string    `json:"title"`
	Active           bool      `json:"active"`
	SerialNumber     *string   `json:"serial_number"`
	IsPreBudget      bool      `json:"is_pre_budget"`
	DateOfPublishing *JSONDate `json:"date_of_publishing"`
	DateOfClosing    *JSONDate `json:"date_of_closing"`
	PreBudgetId      *int      `json:"pre_budget_id"`
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
	FileId           *int      `json:"file_id"`
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
	Id                  int     `json:"id"`
	BudgetIndentId      int     `json:"budget_indent_id"`
	PublicProcurementId int     `json:"public_procurement_id"`
	Title               string  `json:"title"`
	Description         string  `json:"description"`
	NetPrice            float32 `json:"net_price"`
	VatPercentage       string  `json:"vat_percentage"`
	Manufacturer        *string `json:"manufacturer"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
}

type PublicProcurementLimit struct {
	Id                  int `json:"id"`
	PublicProcurementId int `json:"public_procurement_id"`
	OrganizationUnitId  int `json:"organization_unit_id"`
	Limit               int `json:"limit"`
}

type PublicProcurementOrganizationUnitArticle struct {
	Id                         int    `json:"id"`
	PublicProcurementArticleId int    `json:"public_procurement_article_id"`
	OrganizationUnitId         int    `json:"organization_unit_id"`
	Amount                     int    `json:"amount"`
	Status                     string `json:"status"`
	IsRejected                 bool   `json:"is_rejected"`
	RejectedDescription        string `json:"rejected_description"`
	CreatedAt                  string `json:"created_at"`
	UpdatedAt                  string `json:"updated_at"`
}

type PublicProcurementContract struct {
	Id                  int       `json:"id"`
	PublicProcurementId int       `json:"public_procurement_id" validate:"required"`
	SupplierId          int       `json:"supplier_id" validate:"required"`
	SerialNumber        string    `json:"serial_number" validate:"required"`
	DateOfSigning       JSONDate  `json:"date_of_signing" validate:"required"`
	DateOfExpiry        *JSONDate `json:"date_of_expiry"`
	NetValue            float32   `json:"net_value" validate:"required"`
	GrossValue          float32   `json:"gross_value" validate:"required"`
	CreatedAt           string    `json:"created_at"`
	UpdatedAt           string    `json:"updated_at"`
	FileId              *int      `json:"file_id"`
}

type PublicProcurementContractArticle struct {
	Id                          int     `json:"id"`
	PublicProcurementArticleId  int     `json:"public_procurement_article_id"`
	PublicProcurementContractId int     `json:"public_procurement_contract_id"`
	Amount                      int     `json:"amount"`
	NetValue                    float32 `json:"net_value"`
	GrossValue                  float32 `json:"gross_value"`
	VatPercentage               string  `json:"vat_percentage"`
	CreatedAt                   string  `json:"created_at"`
	UpdatedAt                   string  `json:"updated_at"`
}
