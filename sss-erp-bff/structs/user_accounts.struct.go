package structs

type UserAccounts struct {
	Id             int     `json:"id"`
	RoleId         int     `json:"role_id"`
	Email          string  `json:"email"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	SecondaryEmail *string `json:"secondary_email,omitempty"`
	Phone          string  `json:"phone"`
	Password       string  `json:"password"`
	Pin            string  `json:"pin"`
	Active         bool    `json:"active"`
	VerifiedEmail  *bool   `json:"verified_email,omitempty"`
	VerifiedPhone  *bool   `json:"verified_phone,omitempty"`
	FolderId       *int    `json:"folder_id,omitempty"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type UserAccountRoles struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
}
