package structs

type File struct {
	ID          int     `json:"id"`
	ParentID    *int    `json:"parent_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Size        int     `json:"size"`
	Type        *string `json:"type"`
}
