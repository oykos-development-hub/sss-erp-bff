package structs

type AccountItem struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	ParentID     *int   `json:"parent_id"`
	SerialNumber string `json:"serial_number"`
	Version      int    `json:"version"`
}

type AccountItemNode struct {
	ID                 int    `json:"id"`
	ParentID           int    `json:"parent_id"`
	SerialNumber       string `json:"serial_number"`
	Title              string `json:"title"`
	ValueCurrentYear   int    `json:"value_current_year"`
	ValueNextYear      int    `json:"value_next_year"`
	ValueAfterNextYear int    `json:"value_after_next_year"`
	Children           []*AccountItemNode
}

type AccountItemPath struct {
	ID           int    `json:"id"`
	ParentID     int    `json:"parent_id"`
	SerialNumber string `json:"serial_number"`
	Title        string `json:"title"`
	Path         string `json:"path"`
}
