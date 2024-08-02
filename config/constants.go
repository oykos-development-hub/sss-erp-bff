package config

type ContextKey string
type Module string
type PermissionPath string
type PermissionOperations string

const (
	HTTPResponseWriterKey ContextKey = "httpResponseWriter"
	Requestkey            ContextKey = "request"

	HTTPHeadersKey        ContextKey = "httpHeaders"
	TokenKey              ContextKey = "token"
	LoggedInAccountKey    ContextKey = "logged_in_account"
	LoggedInProfileKey    ContextKey = "logged_in_profile"
	OrganizationUnitIDKey ContextKey = "unit_id"
	ConfigKey             ContextKey = "config"

	ISO8601Format     string = "2006-01-02T00:00:00Z"
	DefaultDateString string = "0001-01-01T00:00:00Z"

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

	TitleInvoice               string = "Računi"
	TitleContract              string = "Ugovori"
	TitleDecision              string = "Rješenja"
	TitleSalary                string = "Zarade"
	TitleObligations           string = "Obaveze"
	TitlePaymentOrder          string = "Nalozi za plaćanje"
	TitleEnforcedPayment       string = "Prinudna naplate"
	TitleReturnEnforcedPayment string = "Povraćaji prinudne naplate"

	ModuleCore         Module = "CORE"
	ModuleHR           Module = "HR"
	ModuleInventory    Module = "INVENTORY"
	ModuleAccounting   Module = "ACCOUNTING"
	ModuleProcurements Module = "PROCUREMENTS"
	ModuleFinance      Module = "FINANCE"
	ModuleBFF          Module = "BFF"

	PublicProcurementPlan PermissionPath = "/procurements/plans"
	AccountingContract    PermissionPath = "/accounting/contracts"
	InventoryMovableItems PermissionPath = "/inventory/movable-inventory"
	FinanceBudget         PermissionPath = "/finance/budget"

	OperationCreate     PermissionOperations = "can_create"
	OperationUpdate     PermissionOperations = "can_update"
	OperationRead       PermissionOperations = "can_read"
	OperationDelete     PermissionOperations = "can_delete"
	OperationFullAccess PermissionOperations = "full_access"
)
