package structs

type UserProfiles struct {
	Id                             int     `json:"id"`
	UserAccountId                  int     `json:"user_account_id"`
	FirstName                      string  `json:"first_name"`
	MiddleName                     string  `json:"middle_name"`
	LastName                       string  `json:"last_name"`
	BirthLastName                  string  `json:"birth_last_name"`
	FatherName                     string  `json:"father_name"`
	MotherName                     string  `json:"mother_name"`
	MotherBirthLastName            string  `json:"mother_birth_last_name"`
	DateOfBirth                    *string `json:"date_of_birth"`
	CountryOfBirth                 string  `json:"country_of_birth"`
	CityOfBirth                    string  `json:"city_of_birth"`
	Nationality                    string  `json:"nationality"`
	NationalMinority               string  `json:"national_minority"`
	Citizenship                    string  `json:"citizenship"`
	Address                        string  `json:"address"`
	BankAccount                    string  `json:"bank_account"`
	BankName                       string  `json:"bank_name"`
	PersonalID                     *string `json:"personal_id"`
	OfficialPersonalId             string  `json:"official_personal_id"`
	OfficialPersonalDocumentNumber string  `json:"official_personal_document_number"`
	OfficialPersonalDocumentIssuer string  `json:"official_personal_document_issuer"`
	Gender                         string  `json:"gender"`
	SingleParent                   bool    `json:"single_parent"`
	HousingDone                    bool    `json:"housing_done"`
	HousingDescription             string  `json:"housing_description"`
	MaritalStatus                  string  `json:"marital_status"`
	DateOfTakingOath               *string `json:"date_of_taking_oath"`
	DateOfBecomingJudge            string  `json:"date_of_becoming_judge"`
	EngagementTypeId               *int    `json:"engagement_type_id,omitempty"`
	RevisorRole                    bool    `json:"revisor_role"`
	SecondaryEmail                 string  `json:"secondary_email"`
	ActiveContract                 *bool   `db:"active_contract"`
	CreatedAt                      string  `json:"created_at"`
	UpdatedAt                      string  `json:"updated_at"`
}

func (u *UserProfiles) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
