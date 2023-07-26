package structs

type Permissions struct {
	ID       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Title    string `json:"title"`
	Route    string `json:"route"`
	Create   bool   `json:"create"`
	Read     bool   `json:"read"`
	Update   bool   `json:"update"`
	Delete   bool   `json:"delete"`
}
