package dto

import (
	"bff/structs"
)

type LoginRequestMS struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetRequestMS struct {
	Email string `json:"email"`
}

type PinRequestMS struct {
	Pin string `json:"pin"`
}

type LoginResponseMS struct {
	Data struct {
		Token struct {
			Type  string `json:"type"`
			Token string `json:"token"`
		} `json:"token"`
		structs.UserAccounts `json:"user"`
	} `json:"data"`
}

type LoginResponse struct {
	Status               string                              `json:"status"`
	Message              string                              `json:"message"`
	Id                   int                                 `json:"id"`
	RoleId               structs.UserRole                    `json:"role_id"`
	FolderId             int                                 `json:"folder_id"`
	Email                string                              `json:"email"`
	Phone                string                              `json:"phone"`
	Token                string                              `json:"token"`
	CreatedAt            string                              `json:"created_at"`
	FirstName            string                              `json:"first_name"`
	LastName             string                              `json:"last_name"`
	BirthLastName        string                              `json:"birth_last_name"`
	Gender               string                              `json:"gender"`
	DateOfBecomingJudge  string                              `json:"date_of_becoming_judge"`
	Permissions          interface{}                         `json:"permissions"`
	Contract             *Contract                           `json:"contract"`
	Engagement           *structs.SettingsDropdown           `json:"engagement"`
	JobPosition          *structs.JobPositions               `json:"job_position"`
	OrganizationUnit     *structs.OrganizationUnits          `json:"organization_unit"`
	OrganizationUnitList []OrganizationUnitsOverviewResponse `json:"organization_units"`
	SupplierList         []*structs.Suppliers                `json:"suppliers"`
}

type RefreshTokenResponse struct {
	Data RefreshTokenData `json:"data"`
}

type RefreshTokenData struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}
