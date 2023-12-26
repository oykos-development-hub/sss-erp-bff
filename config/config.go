package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	defaultAppPort = 8080

	defaultCoreAPI         = "http://localhost:4000"
	defaultHRAPI           = "http://localhost:4100"
	defaultProcurementsAPI = "http://localhost:4200"
	defaultInventoryAPI    = "http://localhost:4300"
	defaultAccountingAPI   = "http://localhost:4400"
	defaultFileAPI         = "http://localhost:4500"
	defaultFinanceAPI      = "http://localhost:4600"

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
	Finance      string
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
}

type CoreMS struct {
	Base           string
	Login          string
	Logout         string
	Refresh        string
	Pin            string
	UserAccounts   string
	Roles          string
	Permissions    string
	Settings       string
	Suppliers      string
	Notifications  string
	LoggedInUser   string
	Account        string
	ForgotPassword string
	ValidateMail   string
	ResetPassword  string
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
	Revisors                        string
	RevisionRevisors                string
	RevisionOrgUnit                 string
}

type AccountingMS struct {
	Base                     string
	OrderLists               string
	OrderProcurementArticles string
	Movements                string
	MovementReport           string
	MovementArticles         string
	Stock                    string
}

type InventoryMS struct {
	Base           string
	Item           string
	RealEstates    string
	Assessments    string
	Dispatch       string
	DispatchItems  string
	ItemsInOrgUnit string
}

type FilesMS struct {
	Base                string
	Files               string
	FilesMultipleDelete string
	FilesDownload       string
	FilesOverview       string
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
		log.Fatal("Error loading .env file")
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
				Revisors:                        hrBase + "/get-revisors",
				RevisionRevisors:                hrBase + "/revision-revisors",
				RevisionOrgUnit:                 hrBase + "/revisions-in-organization-units",
			},
			Core: CoreMS{
				Base:           coreBase,
				Login:          coreBase + "/users/login",
				Logout:         coreBase + "/users/logout",
				Refresh:        coreBase + "/refresh",
				Roles:          coreBase + "/roles",
				Permissions:    coreBase + "/permissions",
				Pin:            coreBase + "/users/validate-pin",
				UserAccounts:   coreBase + "/users",
				Settings:       coreBase + "/settings",
				Suppliers:      coreBase + "/suppliers",
				Notifications:  coreBase + "/notifications",
				LoggedInUser:   coreBase + "/logged-in-user",
				Account:        coreBase + "/accounts",
				ForgotPassword: coreBase + "/users/password/forgot",
				ValidateMail:   coreBase + "/users/password/validate-email",
				ResetPassword:  coreBase + "/users/password/reset",
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
			},
			Inventory: InventoryMS{
				Base:           inventoryBase,
				Item:           inventoryBase + "/items",
				RealEstates:    inventoryBase + "/real-estates",
				Assessments:    inventoryBase + "/assessments",
				Dispatch:       inventoryBase + "/dispatches",
				DispatchItems:  inventoryBase + "/dispatch-items",
				ItemsInOrgUnit: inventoryBase + "/items-in-organization-unit",
			},
			Files: FilesMS{
				Base:                filesBase,
				Files:               filesBase + "/files",
				FilesDownload:       filesBase + "/file-overview",
				FilesOverview:       filesBase + "/download",
				FilesMultipleDelete: filesBase + "/files/batch-delete",
			},
			Accounting: AccountingMS{
				Base:                     accountingBase,
				OrderLists:               accountingBase + "/order-lists",
				OrderProcurementArticles: accountingBase + "/order-procurement-articles",
				Movements:                accountingBase + "/movements",
				MovementReport:           accountingBase + "/movements-report",
				MovementArticles:         accountingBase + "/movement-articles",
				Stock:                    accountingBase + "/stocks",
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
