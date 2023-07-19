package structs

type ProgramItem struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Title    string `json:"title"`
}
