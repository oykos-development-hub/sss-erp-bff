package structs

type UserAccounts struct {
	ID             int     `json:"id"`
	RoleID         *int    `json:"role_id"`
	Role           Roles   `json:"role"`
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
	FolderID       *int    `json:"folder_id,omitempty"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type Roles struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Active       bool   `json:"active"`
}
