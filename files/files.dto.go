package files

import (
	"bff/structs"
	"time"
)

type FileResponseData struct {
	Data    []FileResponseDTO `json:"data"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
}

type FileResponse struct {
	Data   []FileResponseDTO `json:"data"`
	Status string            `json:"status"`
}

type FileGetByIDResponse struct {
	Data    *SingleFileResponse `json:"data"`
	Message string              `json:"message"`
	Error   string              `json:"error"`
}

type SingleFileResponse struct {
	Data   *FileResponseDTO `json:"data"`
	Status string           `json:"status"`
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
	Files []int `json:"files"`
}

type errorResponse struct {
	Message string `json:"message"`
}
