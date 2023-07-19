package structs

type AccountItem struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	ParentId     int    `json:"parent_id"`
	SerialNumber string `json:"serial_number"`
}

type AccountItemNode struct {
	Id           int    `json:"id"`
	ParentId     int    `json:"parent_id"`
	SerialNumber string `json:"serial_number"`
	Title        string `json:"title"`
	Children     []*AccountItemNode
}

type AccountItemPath struct {
	Id           int    `json:"id"`
	ParentId     int    `json:"parent_id"`
	SerialNumber string `json:"serial_number"`
	Title        string `json:"title"`
	Path         string `json:"path"`
}
