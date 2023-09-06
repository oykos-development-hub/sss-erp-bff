package structs

type Contracts struct {
	Id                              int      `json:"id"`
	Title                           string   `json:"title"`
	ContractTypeId                  int      `json:"contract_type_id"`
	OrganizationUnitID              int      `json:"organization_unit_id"`
	OrganizationUnitDepartmentID    *int     `json:"organization_unit_department_id"`
	JobPositionInOrganizationUnitID int      `json:"job_position_in_organization_unit_id"`
	UserProfileId                   int      `json:"user_profile_id"`
	Abbreviation                    string   `json:"abbreviation"`
	Description                     string   `json:"description"`
	Active                          bool     `json:"active"`
	SerialNumber                    string   `json:"serial_number"`
	NetSalary                       string   `json:"net_salary"`
	GrossSalary                     string   `json:"gross_salary"`
	BankAccount                     string   `json:"bank_account"`
	BankName                        string   `json:"bank_name"`
	DateOfSignature                 JSONDate `json:"date_of_signature"`
	DateOfEligibility               JSONDate `json:"date_of_eligibility"`
	DateOfStart                     JSONDate `json:"date_of_start"`
	DateOfEnd                       JSONDate `json:"date_of_end"`
	CreatedAt                       string   `json:"created_at"`
	UpdatedAt                       string   `json:"updated_at"`
	FileId                          int      `json:"file_id"`
}

type ContractType struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
}
