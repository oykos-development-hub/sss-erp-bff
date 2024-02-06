package config

type ContextKey string

const (
	HTTPResponseWriterKey ContextKey = "httpResponseWriter"
	Requestkey            ContextKey = "request"

	HTTPHeadersKey        ContextKey = "httpHeaders"
	TokenKey              ContextKey = "token"
	LoggedInAccountKey    ContextKey = "logged_in_account"
	LoggedInProfileKey    ContextKey = "logged_in_profile"
	OrganizationUnitIDKey ContextKey = "unit_id"
	ConfigKey             ContextKey = "config"

	ResolutionTypes                     string = "resolution_types"
	VacationTypeValueResolutionType     string = "vacation"
	EmploymentTerminationResolutionType string = "employment_termination"

	OfficeTypes    string = "office_types"
	EducationTypes string = "education_types"
)
