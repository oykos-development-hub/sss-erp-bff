package structs

type Family struct {
	Id                   int    `json:"id"`
	UserProfileId        int    `json:"user_profile_id"`
	FirstName            string `json:"first_name"`
	MiddleName           string `json:"middle_name"`
	LastName             string `json:"last_name"`
	BirthLastName        string `json:"birth_last_name"`
	FatherName           string `json:"father_name"`
	MotherName           string `json:"mother_name"`
	MotherBirthLastName  string `json:"mother_birth_last_name"`
	DateOfBirth          string `json:"date_of_birth"`
	CountryOfBirth       string `json:"country_of_birth"`
	CityOfBirth          string `json:"city_of_birth"`
	Nationality          string `json:"nationality"`
	Citizenship          string `json:"citizenship"`
	Address              string `json:"address"`
	OfficialPersonalId   string `json:"official_personal_id"`
	Gender               string `json:"gender"`
	EmployeeRelationship string `json:"employee_relationship"`
	InsuranceCoverage    string `json:"insurance_coverage"`
	HandicappedPerson    bool   `json:"handicapped_person"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}
