package structs

type CreateUser struct {
	UserProfileID                  int         `json:"id"`
	UserAccountID                  int         `json:"user_account_id"`
	FirstName                      string      `json:"first_name"`
	MiddleName                     string      `json:"middle_name"`
	LastName                       string      `json:"last_name"`
	BirthLastName                  string      `json:"birth_last_name"`
	FatherName                     string      `json:"father_name"`
	MotherName                     string      `json:"mother_name"`
	MotherBirthLastName            string      `json:"mother_birth_last_name"`
	DateOfBirth                    string      `json:"date_of_birth"`
	CountryOfBirth                 string      `json:"country_of_birth"`
	CityOfBirth                    string      `json:"city_of_birth"`
	Nationality                    string      `json:"nationality"`
	Citizenship                    string      `json:"citizenship"`
	Address                        string      `json:"address"`
	BankAccount                    string      `json:"bank_account"`
	BankName                       string      `json:"bank_name"`
	OfficialPersonalId             string      `json:"official_personal_id"`
	OfficialPersonalDocumentNumber string      `json:"official_personal_document_number"`
	OfficialPersonalDocumentIssuer string      `json:"official_personal_document_issuer"`
	Gender                         string      `json:"gender"`
	SingleParent                   bool        `json:"single_parent"`
	HousingDone                    bool        `json:"housing_done"`
	HousingDescription             string      `json:"housing_description"`
	RevisorRole                    bool        `json:"revisor_role"`
	MaritalStatus                  string      `json:"marital_status"`
	DateOfTakingOath               string      `json:"date_of_taking_oath"`
	DateOfBecomingJudge            string      `json:"date_of_becoming_judge"`
	EngagementTypeId               int         `json:"engagement_type_id"`
	NationalMinority               string      `json:"national_minority"`
	PositionInOrganizationUnitId   string      `json:"position_in_organization_unit_id"`
	Contracts                      []Contracts `json:"contracts"`
	Email                          string      `json:"email"`
	SecondaryEmail                 string      `json:"secondary_email"`
	Pin                            string      `json:"pin"`
	Phone                          string      `json:"phone"`
	Password                       string      `json:"password"`
	Role                           int         `json:"role"`
}
