package dto

import "bff/structs"

type GetProcurementItemResponseMS struct {
	Data structs.PublicProcurementItem `json:"data"`
}

type GetProcurementItemListResponseMS struct {
	Data []*structs.PublicProcurementItem `json:"data"`
}

type GetProcurementItemListInputMS struct {
	Page                   *int    `json:"page"`
	Size                   *int    `json:"size"`
	PlanID                 *int    `json:"plan_id"`
	SortByTitle            *string `json:"sort_by_title"`
	SortBySerialNumber     *string `json:"sort_by_serial_number"`
	SortByDateOfPublishing *string `json:"sort_by_date_of_publishing"`
	SortByDateOfAwarding   *string `json:"sort_by_date_of_awarding"`
}

type ProcurementItemResponseItem struct {
	ID                int                               `json:"id"`
	BudgetIndent      DropdownBudgetIndent              `json:"budget_indent"`
	Plan              DropdownSimple                    `json:"plan_id"`
	IsOpenProcurement bool                              `json:"is_open_procurement"`
	TypeOfProcedure   string                            `json:""`
	Title             string                            `json:"title"`
	ArticleType       string                            `json:"article_type"`
	Status            structs.ProcurementStatus         `json:"status"`
	SerialNumber      *string                           `json:"serial_number"`
	DateOfPublishing  *string                           `json:"date_of_publishing"`
	DateOfAwarding    *string                           `json:"date_of_awarding"`
	FileID            *int                              `json:"file_id"`
	Articles          []*ProcurementArticleResponseItem `json:"articles"`
	ContractID        *int                              `json:"contract_id"`
	TotalGross        float64                           `json:"total_gross"`
	TotalNet          float64                           `json:"total_net"`
	CreatedAt         string                            `json:"created_at"`
	UpdatedAt         string                            `json:"updated_at"`
}

type Subtitles struct {
	PublicProcurement string `json:"public_procurement"`
	OrganizationUnit  string `json:"organization_unit"`
	Supplier          string `json:"supplier"`
}

type TableDataRow struct {
	ProcurementItem  string `json:"procurement_item"`
	KeyFeatures      string `json:"key_features"`
	ContractedAmount string `json:"contracted_amount"`
	AvailableAmount  string `json:"available_amount"`
	ConsumedAmount   string `json:"consumed_amount"`
}

type PdfData struct {
	Subtitles Subtitles      `json:"subtitles"`
	TableData []TableDataRow `json:"table_data"`
}
