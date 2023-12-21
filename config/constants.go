package config

type ContextKey string

const (
	HttpResponseWriterKey ContextKey = "httpResponseWriter"
	Requestkey            ContextKey = "request"

	HttpHeadersKey        ContextKey = "httpHeaders"
	TokenKey              ContextKey = "token"
	LoggedInAccountKey    ContextKey = "logged_in_account"
	LoggedInProfileKey    ContextKey = "logged_in_profile"
	OrganizationUnitIDKey ContextKey = "unit_id"
	ConfigKey             ContextKey = "config"

	ResolutionTypes string = "resolution_types"
	OfficeTypes     string = "office_types"
	EducationTypes  string = "education_types"
)
