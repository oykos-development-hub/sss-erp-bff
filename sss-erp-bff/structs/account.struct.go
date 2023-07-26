package structs

type AccountItem struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	ParentId     int    `json:"parent_id"`
	SerialNumber string `json:"serial_number"`
}

type AccountItemNode struct {
	Id                 int    `json:"id"`
	ParentId           int    `json:"parent_id"`
	SerialNumber       string `json:"serial_number"`
	Title              string `json:"title"`
	ValueCurrentYear   int    `json:"value_current_year"`
	ValueNextYear      int    `json:"value_next_year"`
	ValueAfterNextYear int    `json:"value_after_next_year"`
	Children           []*AccountItemNode
}

type AccountItemPath struct {
	Id           int    `json:"id"`
	ParentId     int    `json:"parent_id"`
	SerialNumber string `json:"serial_number"`
	Title        string `json:"title"`
	Path         string `json:"path"`
}
