package dto

import (
	"bff/structs"
)

type GetUserProfileResponseMS struct {
	Data structs.UserProfiles `json:"data"`
}

type GetUserProfileContractResponseMS struct {
	Data structs.Contracts `json:"data"`
}

type ExperienceItemResponseMS struct {
	Data structs.Experience `json:"data"`
}

type GetEmployeeExperienceListResponseMS struct {
	Data []*structs.Experience `json:"data"`
}

type GetUserProfilesInput struct {
	IsRevisor *bool `json:"is_revisor"`
	AccountID *int  `json:"account_id"`
}

type GetEmployeeContracts struct {
	Active *bool `json:"active"`
}

type GetUserProfileListResponseMS struct {
	Data  []*structs.UserProfiles `json:"data"`
	Total int                     `json:"total"`
}

type GetEmployeeContractListResponseMS struct {
	Data []*structs.Contracts `json:"data"`
}

type MutateUserProfileActiveContract struct {
	Contract *structs.Contracts `json:"contract"`
}

type GetEmployeeEducationResponseMS struct {
	Data structs.Education `json:"data"`
}

type GetEducationTypesResponseMS struct {
	Data []structs.EducationType `json:"data"`
}

type GetEmployeeEducationListResponseMS struct {
	Data []structs.Education `json:"data"`
}

type GetEmployeeEvaluationListResponseMS struct {
	Data []*structs.Evaluation `json:"data"`
}

type GetEvaluationResponse struct {
	Data *structs.Evaluation `json:"data"`
}

type GetEmployeeForeignersListResponseMS struct {
	Data []*structs.Foreigners `json:"data"`
}

type GetEmployeeForeignersResponseMS struct {
	Data *structs.Foreigners `json:"data"`
}

type GetEmployeeFamilyMemberResponseMS struct {
	Data structs.Family `json:"data"`
}

type GetEmployeeFamilyMemberListResponseMS struct {
	Data []*structs.Family `json:"data"`
}

type GetEmployeeSalaryParamsListResponseMS struct {
	Data []*structs.SalaryParams `json:"data"`
}

type GetEmployeeSalaryParamsResponseMS struct {
	Data *structs.SalaryParams `json:"data"`
}

type UserProfileBasicResponse struct {
	ID                            int                        `json:"id"`
	FirstName                     string                     `json:"first_name"`
	LastName                      string                     `json:"last_name"`
	DateOfBirth                   string                     `json:"date_of_birth"`
	BirthLastName                 string                     `json:"birth_last_name"`
	CountryOfBirth                string                     `json:"country_of_birth"`
	CityOfBirth                   string                     `json:"city_of_birth"`
	Nationality                   string                     `json:"nationality"`
	Citizenship                   string                     `json:"citizenship"`
	Address                       string                     `json:"address"`
	FatherName                    string                     `json:"father_name"`
	MotherName                    string                     `json:"mother_name"`
	MotherBirthLastName           string                     `json:"mother_birth_last_name"`
	BankAccount                   string                     `json:"bank_account"`
	BankName                      string                     `json:"bank_name"`
	OfficialPersonalID            string                     `json:"official_personal_id"`
	OfficialPersonalDocNumber     string                     `json:"official_personal_document_number"`
	OfficialPersonalDocIssuer     string                     `json:"official_personal_document_issuer"`
	Gender                        string                     `json:"gender"`
	SingleParent                  bool                       `json:"single_parent"`
	HousingDone                   bool                       `json:"housing_done"`
	RevisorRole                   bool                       `json:"revisor_role"`
	HousingDescription            string                     `json:"housing_description"`
	MaritalStatus                 string                     `json:"marital_status"`
	DateOfTakingOath              string                     `json:"date_of_taking_oath"`
	DateOfBecomingJudge           string                     `json:"date_of_becoming_judge"`
	Email                         string                     `json:"email"`
	Phone                         string                     `json:"phone"`
	OrganizationUnit              *structs.OrganizationUnits `json:"organization_unit"`
	JobPosition                   *structs.JobPositions      `json:"job_position"`
	Contract                      *Contract                  `json:"contracts"`
	JobPositionInOrganizationUnit int                        `json:"position_in_organization_unit_id"`
	CreatedAt                     string                     `json:"created_at"`
	UpdatedAt                     string                     `json:"updated_at"`
	NationalMinority              string                     `json:"national_minority"`
	PrivateEmail                  string                     `json:"private_email"`
	PIN                           string                     `json:"pin"`
}

type UserProfileOverviewResponse struct {
	ID               int                      `json:"id"`
	FirstName        string                   `json:"first_name"`
	LastName         string                   `json:"last_name"`
	DateOfBirth      string                   `json:"date_of_birth"`
	Email            string                   `json:"email"`
	Phone            string                   `json:"phone"`
	Active           bool                     `json:"active"`
	IsJudge          bool                     `json:"is_judge"`
	IsJudgePresident bool                     `json:"is_judge_president"`
	Role             structs.SettingsDropdown `json:"role"`
	OrganizationUnit structs.SettingsDropdown `json:"organization_unit"`
	JobPosition      structs.SettingsDropdown `json:"job_position"`
	CreatedAt        string                   `json:"created_at"`
	UpdatedAt        string                   `json:"updated_at"`
}

type EducationResponseItem struct {
	ID           int          `json:"id"`
	Abbreviation string       `json:"abbreviation"`
	Title        string       `json:"title"`
	Value        string       `json:"value"`
	Educations   []*Education `json:"educations"`
}

type Education struct {
	Id                  int              `json:"id"`
	Title               string           `json:"title"`
	TypeId              int              `json:"type_id"`
	SubType             DropdownSimple   `json:"sub_type"`
	UserProfileId       int              `json:"user_profile_id"`
	Description         string           `json:"description"`
	DateOfCertification structs.JSONDate `json:"date_of_certification"`
	Price               int              `json:"price"`
	DateOfStart         structs.JSONDate `json:"date_of_start"`
	DateOfEnd           structs.JSONDate `json:"date_of_end"`
	AcademicTitle       string           `json:"academic_title"`
	ExpertiseLevel      string           `json:"expertise_level"`
	CertificateIssuer   string           `json:"certificate_issuer"`
	CreatedAt           string           `json:"created_at"`
	UpdatedAt           string           `json:"updated_at"`
	FileId              int              `json:"file_id"`
}

type EducationInput struct {
	UserProfileID int  `json:"user_profile_id"`
	TypeID        *int `json:"type_id"`
	SubTypeID     *int `json:"sub_type_id"`
}
