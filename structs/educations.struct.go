package structs

type Education struct {
	Id                  int     `json:"id"`
	Title               string  `json:"title"`
	TypeId              int     `json:"type_id"`
	UserProfileId       int     `json:"user_profile_id"`
	Description         string  `json:"description"`
	DateOfCertification *string `json:"date_of_certification"`
	Price               float32 `json:"price"`
	DateOfStart         *string `json:"date_of_start"`
	DateOfEnd           *string `json:"date_of_end"`
	AcademicTitle       string  `json:"academic_title"`
	ExpertiseLevel      string  `json:"expertise_level"`
	CertificateIssuer   string  `json:"certificate_issuer"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
	FileId              int     `json:"file_id"`
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
