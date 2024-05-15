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

	TypeInvoice               string = "invoices"
	TypeContract              string = "contracts"
	TypeDecision              string = "decisions"
	TypeSalary                string = "salaries"
	TypeObligations           string = "obligations"
	TypePaymentOrder          string = "payment_orders"
	TypeEnforcedPayment       string = "enforced_payments"
	TypeReturnEnforcedPayment string = "return_enforced_payment"

	TitleInvoice               string = "Račun"
	TitleContract              string = "Ugovor"
	TitleDecision              string = "Rješenje"
	TitleSalary                string = "Zarada"
	TitleObligations           string = "Obaveza"
	TitlePaymentOrder          string = "Nalog za plaćanje"
	TitleEnforcedPayment       string = "Prinudna naplata"
	TitleReturnEnforcedPayment string = "Povraćaj prinudne naplate"
)
