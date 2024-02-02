package structs

type ProgramItem struct {
	ID          int    `json:"id"`
	ParentID    *int   `json:"parent_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
}
