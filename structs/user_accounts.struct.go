package structs

type UserAccounts struct {
	Id             int    `json:"id"`
	RoleId         int    `json:"role_id"`
	Email          string `json:"email"`
	SecondaryEmail string `json:"secondary_email"`
	Phone          string `json:"phone"`
	Password       string `json:"password"`
	ResetPassword  string `json:"reset_password"`
	Pin            string `json:"pin"`
	Active         bool   `json:"active"`
	Token          string `json:"token"`
	RefreshToken   string `json:"refresh_token"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	VerifiedEmail  bool   `json:"verified_email"`
	VerifiedPhone  bool   `json:"verified_phone"`
	FolderId       int    `json:"folder_id"`
}

type UserAccountRoles struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
}
