package structs

import (
	"fmt"
	"time"
)

type UserProfiles struct {
	ID                             int     `json:"id"`
	UserAccountID                  int     `json:"user_account_id"`
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
	OfficialPersonalID             string  `json:"official_personal_id"`
	OfficialPersonalDocumentNumber string  `json:"official_personal_document_number"`
	OfficialPersonalDocumentIssuer string  `json:"official_personal_document_issuer"`
	Gender                         string  `json:"gender"`
	SingleParent                   bool    `json:"single_parent"`
	HousingDone                    bool    `json:"housing_done"`
	HousingDescription             string  `json:"housing_description"`
	MaritalStatus                  string  `json:"marital_status"`
	DateOfTakingOath               *string `json:"date_of_taking_oath"`
	DateOfBecomingJudge            string  `json:"date_of_becoming_judge"`
	JudgeApplicationSubmissionDate *string `json:"judge_application_submission_date"`
	EngagementTypeID               *int    `json:"engagement_type_id,omitempty"`
	IsPresident                    bool    `json:"is_president"`
	IsJudge                        bool    `json:"is_judge"`
	SecondaryEmail                 string  `json:"secondary_email"`
	ActiveContract                 *bool   `json:"active_contract"`
	FileID                         int     `json:"file_id"`
	CreatedAt                      string  `json:"created_at"`
	UpdatedAt                      string  `json:"updated_at"`
}

type Revisor struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ID        int    `json:"id"`
}

func (u *UserProfiles) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *UserProfiles) GetAge() int {
	dateOfBirth, err := time.Parse(time.RFC3339, *u.DateOfBirth)
	if err != nil {
		fmt.Println("Error parsing date of birth:", err)
		return 0
	}

	currentDate := time.Now()

	age := currentDate.Sub(dateOfBirth)

	// Extract years, months, and days from the difference.
	years := age / (365 * 24 * time.Hour)
	//age -= years * 365 * 24 * time.Hour
	//months := age / (30 * 24 * time.Hour)
	//age -= months * 30 * 24 * time.Hour

	return int(years)
}
