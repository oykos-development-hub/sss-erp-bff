package structs

type Education struct {
	ID                  int     `json:"id"`
	Title               string  `json:"title"`
	TypeID              int     `json:"type_id"`
	UserProfileID       int     `json:"user_profile_id"`
	Description         string  `json:"description"`
	DateOfCertification *string `json:"date_of_certification"`
	Price               float64 `json:"price"`
	DateOfStart         *string `json:"date_of_start"`
	DateOfEnd           *string `json:"date_of_end"`
	AcademicTitle       string  `json:"academic_title"`
	ExpertiseLevel      string  `json:"expertise_level"`
	Score               *string `json:"score"`
	CertificateIssuer   string  `json:"certificate_issuer"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
	FileIDs             []int   `json:"file_ids"`
}

type EducationType struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	Value        string `json:"value"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
}
