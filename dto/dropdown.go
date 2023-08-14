package dto

type DropdownSimple struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type DropdownProcurementArticle struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	VatPercentage string `json:"vat_percentage"`
	Description   string `json:"description"`
}
