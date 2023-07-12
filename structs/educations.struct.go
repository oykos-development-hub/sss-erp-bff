package structs

type Education struct {
	Id                  int      `json:"id"`
	Title               string   `json:"title"`
	EducationTypeId     int      `json:"education_type_id"`
	UserProfileId       int      `json:"user_profile_id"`
	Description         string   `json:"description"`
	DateOfCertification JSONDate `json:"date_of_certification"`
	Price               int      `json:"price"`
	DateOfStart         JSONDate `json:"date_of_start"`
	DateOfEnd           JSONDate `json:"date_of_end"`
	AcademicTitle       string   `json:"academic_title"`
	ExpertiseLevel      string   `json:"expertise_level"`
	CertificateIssuer   string   `json:"certificate_issuer"`
	CreatedAt           string   `json:"created_at"`
	UpdatedAt           string   `json:"updated_at"`
	FileId              int      `json:"file_id"`
}

type EducationType struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	Value        string `json:"value"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
}
