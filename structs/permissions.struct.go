package structs

type Permissions struct {
	ID       int    `json:"id"`
	ParentID *int   `json:"parent_id"`
	Title    string `json:"title"`
	Route    string `json:"route"`
	Create   bool   `json:"create"`
	Read     bool   `json:"read"`
	Update   bool   `json:"update"`
	Delete   bool   `json:"delete"`
}

type RolePermission struct {
	PermissionID int  `json:"permission_id"`
	RoleID       int  `json:"role_id"`
	Create       bool `json:"create"`
	Read         bool `json:"read"`
	Update       bool `json:"update"`
	Delete       bool `json:"delete"`
}
