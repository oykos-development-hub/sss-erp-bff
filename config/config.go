package config

import (
	"bff/log"
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	defaultAppPort = 8080

	defaultCoreAPI         = "http://localhost:4000/api"
	defaultHRAPI           = "http://localhost:4100/api"
	defaultProcurementsAPI = "http://localhost:4200/api"
	defaultInventoryAPI    = "http://localhost:4300/api"
	defaultAccountingAPI   = "http://localhost:4400/api"
	defaultFileAPI         = "http://localhost:4500/api"
	defaultFinanceAPI      = "http://localhost:4600/api"

	defaultCoreFE         = "http://localhost:3000"
	defaultHRFE           = "http://localhost:3001"
	defaultProcurementsFE = "http://localhost:3002"
	defaultInventoryFE    = "http://localhost:3003"
	defaultAccountingFE   = "http://localhost:3004"
	defaultFinanceFE      = "http://localhost:3005"

	defaultIsDebug = false
)

type Config struct {
	BaseAppDir    string
	AppPort       int
	Microservices MicroservicesConfig
	Frontend      FrontendConfig
	Debug         bool
}

type FrontendConfig struct {
	Core         string
	HR           string
	Procurements string
	Inventory    string
	Accounting   string
	Finance      string
}

type MicroservicesConfig struct {
	Core         CoreMS
	HR           HrMS
	Procurements ProcurementMS
	Accounting   AccountingMS
	Inventory    InventoryMS
	Files        FilesMS
	Finance      FinanceMS
}

type FinanceMS struct {
	Base                                     string
	Budget                                   string
	BudgetRequest                            string
	SpendingDynamicGet                       string
	SpendingDynamicGetHistory                string
	SpendingDynamicActual                    string
	SpendingDynamicInsert                    string
	SpendingReleaseInsert                    string
	SpendingReleaseDelete                    string
	SpendingReleaseOverview                  string
	SpendingReleaseList                      string
	SpendingReleaseRequest                   string
	AcceptSpendingReleaseRequest             string
	CurrentBudget                            string
	CurrentBudgetUnitList                    string
	Program                                  string
	Activity                                 string
	FinancialBudget                          string
	FilledFinancialBudget                    string
	FinancialBudgetLimit                     string
	NonFinancialBudget                       string
	NonFinancialGoal                         string
	NonFinancialGoalIndicator                string
	Invoice                                  string
	InvoiceArticle                           string
	AdditionalExpenses                       string
	TaxAuthorityCodebook                     string
	DeactivateTaxAuthorityCodebook           string
	Fee                                      string
	FeePayment                               string
	Fine                                     string
	FinePayment                              string
	ProcedureCost                            string
	ProcedureCostPayment                     string
	FlatRate                                 string
	FlatRatePayment                          string
	PropBenConf                              string
	PropBenConfPayment                       string
	FixedDeposit                             string
	FixedDepositItem                         string
	FixedDepositJudge                        string
	FixedDepositDispatch                     string
	FixedDepositWill                         string
	FixedDepositWillDispatch                 string
	Salary                                   string
	DepositPayment                           string
	DepositPaymentCaseNumber                 string
	GetDepositPaymentCaseNumber              string
	GetInitialState                          string
	DepositPaymentOrder                      string
	DepositPaymentAdditionalExpenses         string
	PayDepositPaymentOrder                   string
	PaymentOrder                             string
	GetPaymentOrderByIDOfStatement           string
	PayPaymentOrder                          string
	CancelPaymentOrder                       string
	GetObligation                            string
	GetObligationsForAccounting              string
	GetPaymentOrdersForAccounting            string
	GetEnforcedPaymentsForAccounting         string
	GetReturnedEnforcedPaymentsForAccounting string
	EnforcedPayment                          string
	ReturnEnforcedPayment                    string
	ModelsOfAccounting                       string
	BuildAccountingOrderForObligations       string
	AccountingEntry                          string
	AnalyticalCard                           string
	InternalReallocation                     string
	ExternalReallocation                     string
	AcceptOUExternalReallocation             string
	RejectOUExternalReallocation             string
	AcceptSSSExternalReallocation            string
	RejectSSSExternalReallocation            string
	GetCurrentBudgetByOrganizationUnit       string
	Logs                                     string
	ErrorLogs                                string
}

type ProcurementMS struct {
	Base                    string
	Plans                   string
	Items                   string
	Articles                string
	Contracts               string
	OULimits                string
	ContractArticle         string
	ContractArticleOverage  string
	OrganizationUnitArticle string
	Logs                    string
	ErrorLogs               string
}

type CoreMS struct {
	Base                string
	Login               string
	Logout              string
	Refresh             string
	Pin                 string
	UserAccounts        string
	Roles               string
	Permissions         string
	GetUserByPermission string
	Settings            string
	Suppliers           string
	Notifications       string
	LoggedInUser        string
	Account             string
	ForgotPassword      string
	ValidateMail        string
	ResetPassword       string
	Logs                string
	Templates           string
	TemplateItems       string
	ErrorLogs           string
	BffErrorLogs        string
	CustomerSupport     string
	ListOfParameters    string
}

type HrMS struct {
	Base                            string
	Evaluations                     string
	EvaluationTypes                 string
	Foreigners                      string
	Salaries                        string
	OrganizationUnits               string
	JobPositions                    string
	JobPositionInOrganizationUnits  string
	EducationTypes                  string
	UserProfiles                    string
	EmployeeContracts               string
	EmployeeEducations              string
	EmployeeExperiences             string
	EmployeesInOrganizationUnits    string
	EmployeesInOrganizationUnitByID string
	Resolutions                     string
	Systematization                 string
	EmployeeFamilyMembers           string
	JudgeNorm                       string
	Judges                          string
	AbsentType                      string
	EmployeeAbsents                 string
	Revisions                       string
	JobTenders                      string
	JobTenderTypes                  string
	JobTenderApplications           string
	JudgeResolutions                string
	JudgeResolutionItems            string
	RevisionPlan                    string
	Revision                        string
	RevisionTips                    string
	RevisionTipImplementations      string
	Revisors                        string
	RevisionRevisors                string
	RevisionOrgUnit                 string
	Logs                            string
	ErrorLogs                       string
}

type AccountingMS struct {
	Base                     string
	OrderLists               string
	OrderListSendToFinance   string
	OrderProcurementArticles string
	Movements                string
	MovementReport           string
	MovementArticles         string
	Stock                    string
	StockReport              string
	Logs                     string
	ErrorLogs                string
	StockOrderArticle        string
}

type InventoryMS struct {
	Base           string
	Item           string
	RealEstates    string
	Assessments    string
	Dispatch       string
	DispatchItems  string
	ItemsInOrgUnit string
	ItemsReport    string
	Logs           string
	ErrorLogs      string
}

type FilesMS struct {
	Base                string
	Files               string
	FilesMultipleDelete string
	FilesDownload       string
	FilesOverview       string
	ErrorLogs           string
}

func getEnvString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	if valueBool, err := strconv.ParseBool(value); err == nil {
		return valueBool
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	if valInt, err := strconv.Atoi(value); err == nil {
		return valInt
	}
	return defaultValue
}

func LoadDefaultConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Logger.Fatal("Error loading .env file")
	}

	baseAppDir, exists := os.LookupEnv("BASE_APP_DIR")
	if !exists {
		return nil, errors.New("BASE_APP_DIR environment variable is required but not set")
	}

	hrBase := getEnvString("HR_MS_BASE_URL", defaultHRAPI)
	coreBase := getEnvString("CORE_MS_BASE_URL", defaultCoreAPI)
	procurementsBase := getEnvString("PROCUREMENT_MS_BASE_URL", defaultProcurementsAPI)
	accountingBase := getEnvString("ACCOUNTING_MS_BASE_URL", defaultAccountingAPI)
	inventoryBase := getEnvString("BASIC_INVENTORY_MS_BASE_URL", defaultInventoryAPI)
	filesBase := getEnvString("FILE_MS_BASE_URL", defaultFileAPI)
	financeBase := getEnvString("FINANCE_MS_BASE_URL", defaultFinanceAPI)

	return &Config{
		BaseAppDir: baseAppDir,
		AppPort:    getEnvInt("APP_PORT", defaultAppPort),
		Microservices: MicroservicesConfig{
			HR: HrMS{
				Base:                            hrBase,
				UserProfiles:                    hrBase + "/user-profiles",
				EmployeeContracts:               hrBase + "/employee-contracts",
				EmployeeEducations:              hrBase + "/employee-educations",
				EmployeeExperiences:             hrBase + "/employee-experiences",
				EvaluationTypes:                 hrBase + "/evaluation-types",
				Evaluations:                     hrBase + "/evaluations",
				Foreigners:                      hrBase + "/foreigners",
				Salaries:                        hrBase + "/salaries",
				EducationTypes:                  hrBase + "/education-types",
				OrganizationUnits:               hrBase + "/organization-units",
				Systematization:                 hrBase + "/systematizations",
				JobPositions:                    hrBase + "/job-positions",
				JobPositionInOrganizationUnits:  hrBase + "/job-positions-in-organization-units",
				EmployeesInOrganizationUnits:    hrBase + "/employees-in-organization-units",
				EmployeesInOrganizationUnitByID: hrBase + "/employees-in-organization-units-by-id",
				Resolutions:                     hrBase + "/employee-resolutions",
				EmployeeFamilyMembers:           hrBase + "/employee-family-members",
				JudgeNorm:                       hrBase + "/user-norms",
				Judges:                          hrBase + "/judges",
				AbsentType:                      hrBase + "/absent-types",
				EmployeeAbsents:                 hrBase + "/employee-absents",
				Revisions:                       hrBase + "/revisions-of-organization-units",
				JobTenders:                      hrBase + "/tenders-in-organization-units",
				JobTenderTypes:                  hrBase + "/tender-types",
				JobTenderApplications:           hrBase + "/tender-applications-in-organization-units",
				JudgeResolutions:                hrBase + "/judge-number-resolutions",
				JudgeResolutionItems:            hrBase + "/judge-number-resolution-organization-units",
				RevisionPlan:                    hrBase + "/plans",
				Revision:                        hrBase + "/revisions",
				RevisionTips:                    hrBase + "/revision-tips",
				RevisionTipImplementations:      hrBase + "/revision-tip-implementations",
				Revisors:                        hrBase + "/get-revisors",
				RevisionRevisors:                hrBase + "/revision-revisors",
				RevisionOrgUnit:                 hrBase + "/revisions-in-organization-units",
				Logs:                            hrBase + "/logs",
				ErrorLogs:                       hrBase + "/error-logs",
			},
			Core: CoreMS{
				Base:                coreBase,
				Login:               coreBase + "/users/login",
				Logout:              coreBase + "/users/logout",
				Refresh:             coreBase + "/refresh",
				Roles:               coreBase + "/roles",
				Permissions:         coreBase + "/permissions",
				GetUserByPermission: coreBase + "/get-users-by-permission",
				Pin:                 coreBase + "/users/validate-pin",
				UserAccounts:        coreBase + "/users",
				Settings:            coreBase + "/settings",
				Suppliers:           coreBase + "/suppliers",
				Notifications:       coreBase + "/notifications",
				LoggedInUser:        coreBase + "/logged-in-user",
				Account:             coreBase + "/accounts",
				ForgotPassword:      coreBase + "/users/password/forgot",
				ValidateMail:        coreBase + "/users/password/validate-email",
				ResetPassword:       coreBase + "/users/password/reset",
				Logs:                coreBase + "/logs",
				Templates:           coreBase + "/templates",
				TemplateItems:       coreBase + "/template-items",
				ErrorLogs:           coreBase + "/error-logs",
				BffErrorLogs:        coreBase + "/bff-error-logs",
				CustomerSupport:     coreBase + "/customer-supports",
				ListOfParameters:    coreBase + "/list-of-parameters",
			},
			Procurements: ProcurementMS{
				Base:                    procurementsBase,
				Plans:                   procurementsBase + "/plans",
				Items:                   procurementsBase + "/items",
				Articles:                procurementsBase + "/articles",
				Contracts:               procurementsBase + "/contracts",
				ContractArticle:         procurementsBase + "/contract-articles",
				OULimits:                procurementsBase + "/organization-unit-plan-limits",
				OrganizationUnitArticle: procurementsBase + "/organization-unit-articles",
				ContractArticleOverage:  procurementsBase + "/contract-article-overages",
				Logs:                    procurementsBase + "/logs",
				ErrorLogs:               procurementsBase + "/error-logs",
			},
			Inventory: InventoryMS{
				Base:           inventoryBase,
				Item:           inventoryBase + "/items",
				RealEstates:    inventoryBase + "/real-estates",
				Assessments:    inventoryBase + "/assessments",
				Dispatch:       inventoryBase + "/dispatches",
				DispatchItems:  inventoryBase + "/dispatch-items",
				ItemsInOrgUnit: inventoryBase + "/items-in-organization-unit",
				ItemsReport:    inventoryBase + "/items-for-item-list-report",
				Logs:           inventoryBase + "/logs",
				ErrorLogs:      inventoryBase + "/error-logs",
			},
			Files: FilesMS{
				Base:                filesBase,
				Files:               filesBase + "/files",
				FilesDownload:       filesBase + "/file-overview",
				FilesOverview:       filesBase + "/download",
				FilesMultipleDelete: filesBase + "/files/batch-delete",
				ErrorLogs:           filesBase + "/error-logs",
			},
			Accounting: AccountingMS{
				Base:                     accountingBase,
				OrderLists:               accountingBase + "/order-lists",
				OrderListSendToFinance:   accountingBase + "/order-list-send-to-finance",
				OrderProcurementArticles: accountingBase + "/order-procurement-articles",
				Movements:                accountingBase + "/movements",
				MovementReport:           accountingBase + "/movements-report",
				MovementArticles:         accountingBase + "/movement-articles",
				Stock:                    accountingBase + "/stocks",
				StockReport:              accountingBase + "/get-all-stocks",
				Logs:                     accountingBase + "/logs",
				ErrorLogs:                accountingBase + "/error-logs",
				StockOrderArticle:        accountingBase + "/stock-order-articles",
			},
			Finance: FinanceMS{
				Base:                                     financeBase,
				Budget:                                   financeBase + "/budgets",
				BudgetRequest:                            financeBase + "/budget-requests",
				SpendingDynamicGet:                       financeBase + "/budgets/%d/units/%d/spending-dynamics",
				SpendingDynamicGetHistory:                financeBase + "/budgets/%d/units/%d/spending-dynamics/history",
				SpendingDynamicActual:                    financeBase + "/budgets/%d/units/%d/accounts/%d/actual",
				SpendingDynamicInsert:                    financeBase + "/budgets/%d/units/%d/spending-dynamics",
				SpendingReleaseInsert:                    financeBase + "/budgets/%d/units/%d/spending-releases",
				SpendingReleaseDelete:                    financeBase + "/spending-releases",
				SpendingReleaseOverview:                  financeBase + "/spending-releases/overview",
				SpendingReleaseList:                      financeBase + "/spending-releases",
				SpendingReleaseRequest:                   financeBase + "/spending-release-requests",
				AcceptSpendingReleaseRequest:             financeBase + "/accept-spending-release-request",
				CurrentBudget:                            financeBase + "/current-budgets",
				CurrentBudgetUnitList:                    financeBase + "/current-budgets/units",
				FinancialBudget:                          financeBase + "/financial-budgets",
				FilledFinancialBudget:                    financeBase + "/filled-financial-budgets",
				FinancialBudgetLimit:                     financeBase + "/financial-budget-limits",
				Program:                                  financeBase + "/programs",
				Activity:                                 financeBase + "/activities",
				NonFinancialGoal:                         financeBase + "/non-financial-budget-goals",
				NonFinancialBudget:                       financeBase + "/non-financial-budgets",
				NonFinancialGoalIndicator:                financeBase + "/goal-indicators",
				Invoice:                                  financeBase + "/invoices",
				InvoiceArticle:                           financeBase + "/articles",
				AdditionalExpenses:                       financeBase + "/additional-expenses",
				TaxAuthorityCodebook:                     financeBase + "/tax-authority-codebooks",
				DeactivateTaxAuthorityCodebook:           financeBase + "/tax-authority-codebook-deactivate",
				Fee:                                      financeBase + "/fees",
				FeePayment:                               financeBase + "/fee-payments",
				Fine:                                     financeBase + "/fines",
				FinePayment:                              financeBase + "/fine-payments",
				ProcedureCost:                            financeBase + "/procedure-costs",
				ProcedureCostPayment:                     financeBase + "/procedure-cost-payments",
				FlatRate:                                 financeBase + "/flat-rates",
				FlatRatePayment:                          financeBase + "/flat-rate-payments",
				PropBenConf:                              financeBase + "/property-benefits-confiscations",
				PropBenConfPayment:                       financeBase + "/property-benefits-confiscation-payments",
				FixedDeposit:                             financeBase + "/fixed-deposits",
				FixedDepositItem:                         financeBase + "/fixed-deposit-items",
				FixedDepositDispatch:                     financeBase + "/fixed-deposit-dispatches",
				FixedDepositJudge:                        financeBase + "/fixed-deposit-judges",
				FixedDepositWill:                         financeBase + "/fixed-deposit-wills",
				FixedDepositWillDispatch:                 financeBase + "/fixed-deposit-will-dispatches",
				Salary:                                   financeBase + "/salaries",
				GetInitialState:                          financeBase + "/get-initial-state",
				DepositPayment:                           financeBase + "/deposit-payments",
				DepositPaymentCaseNumber:                 financeBase + "/deposit-payments-case-number",
				GetDepositPaymentCaseNumber:              financeBase + "/get-case-number",
				DepositPaymentOrder:                      financeBase + "/deposit-payment-orders",
				PayDepositPaymentOrder:                   financeBase + "/pay-deposit-payment-order",
				DepositPaymentAdditionalExpenses:         financeBase + "/deposit-additional-expenses",
				PaymentOrder:                             financeBase + "/payment-orders",
				GetPaymentOrderByIDOfStatement:           financeBase + "/payment-orders-id-of-statement",
				PayPaymentOrder:                          financeBase + "/pay-payment-order",
				CancelPaymentOrder:                       financeBase + "/cancel-payment-order",
				GetObligation:                            financeBase + "/get-all-obligations",
				GetObligationsForAccounting:              financeBase + "/get-obligations-for-accounting",
				GetPaymentOrdersForAccounting:            financeBase + "/get-payment-orders-for-accounting",
				GetEnforcedPaymentsForAccounting:         financeBase + "/get-enforced-payments-for-accounting",
				GetReturnedEnforcedPaymentsForAccounting: financeBase + "/get-returned-enforced-payments-for-accounting",
				EnforcedPayment:                          financeBase + "/enforced-payments",
				ReturnEnforcedPayment:                    financeBase + "/return-enforced-payment",
				ModelsOfAccounting:                       financeBase + "/models-of-accountings",
				BuildAccountingOrderForObligations:       financeBase + "/build-accounting-order-for-obligations",
				AccountingEntry:                          financeBase + "/accounting-entries",
				AnalyticalCard:                           financeBase + "/analytical-card",
				InternalReallocation:                     financeBase + "/internal-reallocations",
				ExternalReallocation:                     financeBase + "/external-reallocations",
				AcceptOUExternalReallocation:             financeBase + "/accept-ou-external-reallocations",
				RejectOUExternalReallocation:             financeBase + "/reject-ou-external-reallocations",
				AcceptSSSExternalReallocation:            financeBase + "/accept-sss-external-reallocations",
				RejectSSSExternalReallocation:            financeBase + "/reject-sss-external-reallocations",
				GetCurrentBudgetByOrganizationUnit:       financeBase + "/get-actual-current-budget",
				Logs:                                     financeBase + "/logs",
				ErrorLogs:                                financeBase + "/error-logs",
			},
		},
		Frontend: FrontendConfig{
			Core:         getEnvString("CORE_FRONTEND_URL", defaultCoreFE),
			HR:           getEnvString("HR_FRONTEND_URL", defaultHRFE),
			Procurements: getEnvString("PROCUREMENTS_FRONTEND_URL", defaultProcurementsFE),
			Inventory:    getEnvString("INVENTORY_FRONTEND_URL", defaultInventoryFE),
			Accounting:   getEnvString("ACCOUNTING_FRONTEND_URL", defaultAccountingFE),
			Finance:      getEnvString("FINANCE_FRONTEND_URL", defaultFinanceFE),
		},
		Debug: getEnvBool("DEBUG", defaultIsDebug),
	}, nil
}
