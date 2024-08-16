package files

import (
	"bff/structs"
	"time"
)

type FileResponseData struct {
	Data    []FileResponseDTO `json:"data"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
	Status  string            `json:"status"`
}

type FileResponse struct {
	Data   []FileResponseDTO `json:"data"`
	Status string            `json:"status"`
}

type SingleFileResponse struct {
	Data    *FileResponseDTO `json:"data"`
	Message string           `json:"message"`
	Status  string           `json:"status"`
}

type FileResponseDTO struct {
	ID          int       `json:"id"`
	ParentID    *int      `json:"parent_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Size        int64     `json:"size"`
	Type        *string   `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ArticleResponse struct {
	Data    []ContractArticleResponseDTO `json:"data"`
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Error   string                       `json:"error"`
}

type ContractArticleResponseDTO struct {
	ID           int      `json:"id"`
	ArticleID    int      `json:"public_procurement_article_id"`
	ContractID   int      `json:"public_procurement_contract_id"`
	NetValue     *float32 `json:"net_value"`
	GrossValue   *float32 `json:"gross_value"`
	UsedArticles int      `json:"used_articles"`
}

type ProcurementArticleResponse struct {
	Data    []structs.PublicProcurementArticle `json:"data"`
	Status  string                             `json:"status"`
	Message string                             `json:"message"`
	Error   string                             `json:"error"`
}

type MultipleDeleteFiles struct {
	Files  []int  `json:"files"`
	Status string `json:"status"`
}

type DonationArticleResponse struct {
	Data    []structs.ReadArticlesDonation `json:"data"`
	Status  string                         `json:"status"`
	Message string                         `json:"message"`
	Error   string                         `json:"error"`
}

type ExpireInventoriesResponse struct {
	Status  string                                       `json:"status"`
	Message string                                       `json:"message"`
	Error   string                                       `json:"error"`
	Data    []structs.BasicInventoryAssessmentsTypesItem `json:"data"`
}

type errorResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ImportInventoriesResponse struct {
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Error   string                       `json:"error"`
	Data    []structs.BasicInventoryItem `json:"data"`
}

type ValidationResponse struct {
	Column  int    `json:"column"`
	Row     int    `json:"row"`
	Message string `json:"message"`
}

type ImportPS1Inventories struct {
	Status     string               `json:"status"`
	Message    string               `json:"message"`
	Error      string               `json:"error"`
	Validation []ValidationResponse `json:"validation"`
	Data       []structs.Experience `json:"data"`
}

type ImportSalaries struct {
	Status            string                            `json:"status"`
	Message           string                            `json:"message"`
	Error             string                            `json:"error"`
	Validation        []ValidationResponse              `json:"validation"`
	NumberOfEmployees int                               `json:"number_of_employees"`
	Data              []structs.SalaryAdditionalExpense `json:"data"`
}

type ImportSAP struct {
	Status     string                 `json:"status"`
	Message    string                 `json:"message"`
	Error      string                 `json:"error"`
	Validation []ValidationResponse   `json:"validation"`
	Data       []structs.PaymentOrder `json:"data"`
}

type ImportUserProfileVacation struct {
	UserProfileID int `json:"user_profile_id"`
	NumberOfDays  int `json:"number_of_days"`
}

type ImportUserProfileVacationsResponse struct {
	Status     string                      `json:"status"`
	Message    string                      `json:"message"`
	Error      string                      `json:"error"`
	Validation []ValidationResponse        `json:"validation"`
	Data       []ImportUserProfileVacation `json:"data"`
}

type ImportInventoryArticles struct {
	Article            structs.BasicInventoryInsertItem           `json:"article"`
	FirstAmortization  structs.BasicInventoryAssessmentsTypesItem `json:"first_amortization"`
	SecondAmortization structs.BasicInventoryAssessmentsTypesItem `json:"second_amortization"`
	Dispatch           structs.BasicInventoryDispatchItem         `json:"dispatch"`
	ReversDispatch     structs.BasicInventoryDispatchItem         `json:"revers_dispatch"`
	DispatchItem       structs.BasicInventoryDispatchItemsItem    `json:"dispatch_item"`
	ReversDispatchItem structs.BasicInventoryDispatchItemsItem    `json:"revers_dispatch_item"`
}
