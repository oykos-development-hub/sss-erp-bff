package structs

type UserRole int

const (
	UserRoleAdmin                         UserRole = 1
	UserRoleManagerOJ                     UserRole = 2
	UserRoleOfficialForPublicProcurements UserRole = 3
)

type UserPermission string

const (
	PermissionManageOrganizationUnits UserPermission = "manage_organization_units"
	// Add other permissions as needed
)

// This function returns the permissions for a given user role
func (user UserAccounts) GetPermissionsForRole() []UserPermission {
	var permissions []UserPermission
	switch user.RoleID {
	case UserRoleAdmin:
		permissions = append(permissions, PermissionManageOrganizationUnits)
	case UserRoleOfficialForPublicProcurements:
		permissions = append(permissions, PermissionManageOrganizationUnits)
	case UserRoleManagerOJ:

	}
	return permissions
}

func (user UserAccounts) HasPermission(requiredPermission UserPermission) bool {
	permissions := user.GetPermissionsForRole()
	for _, p := range permissions {
		if p == requiredPermission {
			return true
		}
	}
	return false
}

type UserAccounts struct {
	ID             int      `json:"id"`
	RoleID         UserRole `json:"role_id"`
	Email          string   `json:"email"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	SecondaryEmail *string  `json:"secondary_email,omitempty"`
	Phone          string   `json:"phone"`
	Password       string   `json:"password"`
	Pin            string   `json:"pin"`
	Active         bool     `json:"active"`
	VerifiedEmail  *bool    `json:"verified_email,omitempty"`
	VerifiedPhone  *bool    `json:"verified_phone,omitempty"`
	FolderID       *int     `json:"folder_id,omitempty"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

type Roles struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Active       bool   `json:"active"`
}
