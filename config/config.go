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

	defaultCoreApi         = "http://localhost:4000"
	defaultHRApi           = "http://localhost:4100"
	defaultProcurementsApi = "http://localhost:4200"
	defaultInventoryApi    = "http://localhost:4300"
	defaultAccountingApi   = "http://localhost:4400"
	defaultFileApi         = "http://localhost:4500"
	defaultFinanceApi      = "http://localhost:4600"

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
	BASE                      string
	PLANS                     string
	ITEMS                     string
	ARTICLES                  string
	CONTRACTS                 string
	OU_LIMITS                 string
	CONTRACT_ARTICLE          string
	CONTRACT_ARTICLE_OVERAGE  string
	ORGANIZATION_UNIT_ARTICLE string
}

type CoreMS struct {
	BASE            string
	LOGIN           string
	LOGOUT          string
	REFRESH         string
	PIN             string
	USER_ACCOUNTS   string
	ROLES           string
	PERMISSIONS     string
	SETTINGS        string
	SUPPLIERS       string
	NOTIFICATIONS   string
	LOGGED_IN_USER  string
	ACCOUNT         string
	FORGOT_PASSWORD string
	VALIDATE_MAIL   string
	RESET_PASSWORD  string
}

type HrMS struct {
	BASE                                  string
	EVALUATIONS                           string
	EVALUATION_TYPES                      string
	FOREIGNERS                            string
	SALARIES                              string
	ORGANIZATION_UNITS                    string
	JOB_POSITIONS                         string
	JOB_POSITIONS_IN_ORGANIZATION_UNITS   string
	EDUCATION_TYPES                       string
	USER_PROFILES                         string
	EMPLOYEE_CONTRACTS                    string
	EMPLOYEE_EDUCATIONS                   string
	EMPLOYEE_EXPERIENCES                  string
	EMPLOYEES_IN_ORGANIZATION_UNITS       string
	EMPLOYEES_IN_ORGANIZATION_UNITS_BY_ID string
	RESOLUTIONS                           string
	SYSTEMATIZATIONS                      string
	EMPLOYEE_FAMILY_MEMBERS               string
	JUDGE_NORM                            string
	JUDGES                                string
	ABSENT_TYPE                           string
	EMPLOYEE_ABSENTS                      string
	REVISIONS                             string
	JOB_TENDERS                           string
	JOB_TENDER_TYPES                      string
	JOB_TENDER_APPLICATIONS               string
	JUDGE_RESOLUTIONS                     string
	JUDGE_RESOLUTION_ITEMS                string
	REVISION_PLAN                         string
	REVISION                              string
	REVISION_TIPS                         string
	REVISORS                              string
	REVISION_REVISORS                     string
	REVISION_ORG_UNIT                     string
}

type AccountingMS struct {
	Base                       string
	ORDER_LISTS                string
	ORDER_PROCUREMENT_ARTICLES string
	MOVEMENTS                  string
	MOVEMENT_REPORT            string
	MOVEMENT_ARTICLES          string
	STOCK                      string
}

type InventoryMS struct {
	Base           string
	ITEM           string
	REAL_ESTATES   string
	ASSESSMENTS    string
	DISPATCH       string
	DISPATCH_ITEMS string
}

type FilesMS struct {
	Base                  string
	FILES                 string
	FILES_MULTIPLE_DELETE string
	FILES_DOWNLOAD        string
	FILES_OVERVIEW        string
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
	valueBool, err := strconv.ParseBool(value)
	if err != nil {
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

	hrBase := getEnvString("HR_MS_BASE_URL", defaultHRApi)
	coreBase := getEnvString("CORE_MS_BASE_URL", defaultCoreApi)
	procurementsBase := getEnvString("PROCUREMENT_MS_BASE_URL", defaultProcurementsApi)
	accountingBase := getEnvString("ACCOUNTING_MS_BASE_URL", defaultAccountingApi)
	inventoryBase := getEnvString("BASIC_INVENTORY_MS_BASE_URL", defaultInventoryApi)
	filesBase := getEnvString("FILE_MS_BASE_URL", defaultFileApi)

	return &Config{
		BaseAppDir: baseAppDir,
		AppPort:    getEnvInt("APP_PORT", defaultAppPort),
		Microservices: MicroservicesConfig{
			HR: HrMS{
				BASE:                                  hrBase,
				USER_PROFILES:                         hrBase + "/user-profiles",
				EMPLOYEE_CONTRACTS:                    hrBase + "/employee-contracts",
				EMPLOYEE_EDUCATIONS:                   hrBase + "/employee-educations",
				EMPLOYEE_EXPERIENCES:                  hrBase + "/employee-experiences",
				EVALUATION_TYPES:                      hrBase + "/evaluation-types",
				EVALUATIONS:                           hrBase + "/evaluations",
				FOREIGNERS:                            hrBase + "/foreigners",
				SALARIES:                              hrBase + "/salaries",
				EDUCATION_TYPES:                       hrBase + "/education-types",
				ORGANIZATION_UNITS:                    hrBase + "/organization-units",
				SYSTEMATIZATIONS:                      hrBase + "/systematizations",
				JOB_POSITIONS:                         hrBase + "/job-positions",
				JOB_POSITIONS_IN_ORGANIZATION_UNITS:   hrBase + "/job-positions-in-organization-units",
				EMPLOYEES_IN_ORGANIZATION_UNITS:       hrBase + "/employees-in-organization-units",
				EMPLOYEES_IN_ORGANIZATION_UNITS_BY_ID: hrBase + "/employees-in-organization-units-by-id",
				RESOLUTIONS:                           hrBase + "/employee-resolutions",
				EMPLOYEE_FAMILY_MEMBERS:               hrBase + "/employee-family-members",
				JUDGE_NORM:                            hrBase + "/user-norms",
				JUDGES:                                hrBase + "/judges",
				ABSENT_TYPE:                           hrBase + "/absent-types",
				EMPLOYEE_ABSENTS:                      hrBase + "/employee-absents",
				REVISIONS:                             hrBase + "/revisions-of-organization-units",
				JOB_TENDERS:                           hrBase + "/tenders-in-organization-units",
				JOB_TENDER_TYPES:                      hrBase + "/tender-types",
				JOB_TENDER_APPLICATIONS:               hrBase + "/tender-applications-in-organization-units",
				JUDGE_RESOLUTIONS:                     hrBase + "/judge-number-resolutions",
				JUDGE_RESOLUTION_ITEMS:                hrBase + "/judge-number-resolution-organization-units",
				REVISION_PLAN:                         hrBase + "/plans",
				REVISION:                              hrBase + "/revisions",
				REVISION_TIPS:                         hrBase + "/revision-tips",
				REVISORS:                              hrBase + "/get-revisors",
				REVISION_REVISORS:                     hrBase + "/revision-revisors",
				REVISION_ORG_UNIT:                     hrBase + "/revisions-in-organization-units",
			},
			Core: CoreMS{
				BASE:            coreBase,
				LOGIN:           coreBase + "/users/login",
				LOGOUT:          coreBase + "/users/logout",
				REFRESH:         coreBase + "/refresh",
				ROLES:           coreBase + "/roles",
				PERMISSIONS:     coreBase + "/permissions",
				PIN:             coreBase + "/users/validate-pin",
				USER_ACCOUNTS:   coreBase + "/users",
				SETTINGS:        coreBase + "/settings",
				SUPPLIERS:       coreBase + "/suppliers",
				NOTIFICATIONS:   coreBase + "/notifications",
				LOGGED_IN_USER:  coreBase + "/logged-in-user",
				ACCOUNT:         coreBase + "/accounts",
				FORGOT_PASSWORD: coreBase + "/users/password/forgot",
				VALIDATE_MAIL:   coreBase + "/users/password/validate-email",
				RESET_PASSWORD:  coreBase + "/users/password/reset",
			},
			Procurements: ProcurementMS{
				BASE:                      procurementsBase,
				PLANS:                     procurementsBase + "/plans",
				ITEMS:                     procurementsBase + "/items",
				ARTICLES:                  procurementsBase + "/articles",
				CONTRACTS:                 procurementsBase + "/contracts",
				CONTRACT_ARTICLE:          procurementsBase + "/contract-articles",
				OU_LIMITS:                 procurementsBase + "/organization-unit-plan-limits",
				ORGANIZATION_UNIT_ARTICLE: procurementsBase + "/organization-unit-articles",
				CONTRACT_ARTICLE_OVERAGE:  procurementsBase + "/contract-article-overages",
			},
			Inventory: InventoryMS{
				Base:           inventoryBase,
				ITEM:           inventoryBase + "/items",
				REAL_ESTATES:   inventoryBase + "/real-estates",
				ASSESSMENTS:    inventoryBase + "/assessments",
				DISPATCH:       inventoryBase + "/dispatches",
				DISPATCH_ITEMS: inventoryBase + "/dispatch-items",
			},
			Files: FilesMS{
				Base:                  filesBase,
				FILES:                 filesBase + "/files",
				FILES_DOWNLOAD:        filesBase + "/file-overview",
				FILES_OVERVIEW:        filesBase + "/download",
				FILES_MULTIPLE_DELETE: filesBase + "/files/batch-delete",
			},
			Accounting: AccountingMS{
				Base:                       accountingBase,
				ORDER_LISTS:                accountingBase + "/order-lists",
				ORDER_PROCUREMENT_ARTICLES: accountingBase + "/order-procurement-articles",
				MOVEMENTS:                  accountingBase + "/movements",
				MOVEMENT_REPORT:            accountingBase + "/movements-report",
				MOVEMENT_ARTICLES:          accountingBase + "/movement-articles",
				STOCK:                      accountingBase + "/stocks",
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
