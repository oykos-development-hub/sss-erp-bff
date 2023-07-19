package structs

type Contracts struct {
	Id                int               `json:"id"`
	Title             *string           `json:"title,omitempty"`
	ContractTypeId    int               `json:"contract_type_id"`
	ContractType      *SettingsDropdown `json:"contract_type,omitempty"`
	UserProfileId     int               `json:"user_profile_id"`
	Abbreviation      *string           `json:"abbreviation,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Active            bool              `json:"active"`
	SerialNumber      *string           `json:"serial_number,omitempty"`
	NetSalary         *string           `json:"net_salary,omitempty"`
	GrossSalary       *string           `json:"gross_salary,omitempty"`
	BankAccount       *string           `json:"bank_account,omitempty"`
	BankName          *string           `json:"bank_name,omitempty"`
	DateOfSignature   JSONDate          `json:"date_of_signature,omitempty"`
	DateOfEligibility JSONDate          `json:"date_of_eligibility,omitempty"`
	DateOfStart       JSONDate          `json:"date_of_start"`
	DateOfEnd         JSONDate          `json:"date_of_end,omitempty"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	FileId            *int              `json:"file_id,omitempty"`
}

type ContractType struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
}
