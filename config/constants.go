package config

import (
	"log"
	"os"
	"path"
	"strconv"

	"github.com/joho/godotenv"
)

type ContextKey string

const (
	HttpResponseWriterKey ContextKey = "httpResponseWriter"
	Requestkey            ContextKey = "request"

	HttpHeadersKey        ContextKey = "httpHeaders"
	TokenKey              ContextKey = "token"
	LoggedInAccountKey    ContextKey = "logged_in_account"
	LoggedInProfileKey    ContextKey = "logged_in_profile"
	OrganizationUnitIDKey ContextKey = "unit_id"

	ResolutionTypes string = "resolution_types"
	OfficeTypes     string = "office_types"
	EducationTypes  string = "education_types"
)

var (
	DEBUG bool

	REPORT_DIR_PATH string
	REPORT_DIR_NAME string
	BASE_APP_DIR    string

	CORE_FRONTEND         string
	HR_FRONTEND           string
	PROCUREMENTS_FRONTEND string
	ACCOUNTING_FRONTEND   string
	FINANCE_FRONTEND      string
	INVENTORY_FRONTEND    string

	HR_MS_BASE_URL              string
	CORE_MS_BASE_URL            string
	PROCUREMENT_MS_BASE_URL     string
	BASIC_INVENTORY_MS_BASE_URL string
	ACCOUNTING_MS_BASE_URL      string
	FILE_MS_BASE_URL            string

	LOGIN_ENDPOINT          string
	LOGOUT_ENDPOINT         string
	REFRESH_ENDPOINT        string
	PIN_ENDPOINT            string
	USER_ACCOUNTS_ENDPOINT  string
	ROLES_ENDPOINT          string
	SETTINGS_ENDPOINT       string
	SUPPLIERS_ENDPOINT      string
	LOGGED_IN_USER_ENDPOINT string
	ACCOUNT_ENDPOINT        string

	EVALUATIONS                                    string
	EVALUATION_TYPES_ENDPOINT                      string
	FOREIGNERS                                     string
	SALARIES                                       string
	ORGANIZATION_UNITS_ENDPOINT                    string
	JOB_POSITIONS_ENDPOINT                         string
	JOB_POSITIONS_IN_ORGANIZATION_UNITS_ENDPOINT   string
	EDUCATION_TYPES_ENDPOINT                       string
	USER_PROFILES_ENDPOINT                         string
	EMPLOYEE_CONTRACTS                             string
	EMPLOYEE_EDUCATIONS                            string
	EMPLOYEE_EXPERIENCES                           string
	EMPLOYEES_IN_ORGANIZATION_UNITS_ENDPOINT       string
	EMPLOYEES_IN_ORGANIZATION_UNITS_BY_ID_ENDPOINT string
	RESOLUTIONS_ENDPOINT                           string
	SYSTEMATIZATIONS_ENDPOINT                      string
	EMPLOYEE_FAMILY_MEMBERS                        string
	JUDGE_NORM_ENDPOINT                            string
	JUDGES                                         string
	ABSENT_TYPE                                    string
	EMPLOYEE_ABSENTS                               string
	REVISIONS_ENDPOINT                             string
	JOB_TENDERS_ENDPOINT                           string
	JOB_TENDER_TYPES_ENDPOINT                      string
	JOB_TENDER_APPLICATIONS_ENDPOINT               string
	JUDGE_RESOLUTIONS_ENDPOINT                     string
	JUDGE_RESOLUTION_ITEMS_ENDPOINT                string
	REVISION_PLAN_ENDPOINT                         string
	REVISION_ENDPOINT                              string
	REVISION_TIPS_ENDPOINT                         string
	REVISORS_ENDPOINT                              string
	REVISION_REVISORS_ENDPOINT                     string
	REVISION_ORG_UNIT_ENDPOINT                     string

	PLANS_ENDPOINT                     string
	ITEMS_ENDPOINT                     string
	FORGOT_PASSWORD                    string
	ARTICLES_ENDPOINT                  string
	CONTRACTS_ENDPOINT                 string
	OU_LIMITS_ENDPOINT                 string
	CONTRACT_ARTICLE_ENDPOINT          string
	CONTRACT_ARTICLE_OVERAGE_ENDPOINT  string
	ORGANIZATION_UNIT_ARTICLE_ENDPOINT string

	INVENTORY_ITEM_ENDOPOINT           string
	REAL_ESTATES_ENDPOINT              string
	ASSESSMENTS_ENDPOINT               string
	INVENTORY_DISPATCH_ENDOPOINT       string
	INVENTORY_DISPATCH_ITEMS_ENDOPOINT string

	ORDER_LISTS_ENDPOINT                string
	ORDER_PROCUREMENT_ARTICLES_ENDPOINT string

	FILES_ENDPOINT string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	BASE_APP_DIR = os.Getenv("BASE_APP_DIR")

	debugValue, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatal("Error parsing debug config")
	}
	DEBUG = debugValue

	REPORT_DIR_PATH = os.Getenv("REPORT_DIR")

	REPORT_DIR_NAME = path.Base(REPORT_DIR_PATH)

	CORE_FRONTEND = os.Getenv("CORE_FRONTEND_URL")
	HR_FRONTEND = os.Getenv("HR_FRONTEND_URL")
	PROCUREMENTS_FRONTEND = os.Getenv("PROCUREMENTS_FRONTEND_URL")
	ACCOUNTING_FRONTEND = os.Getenv("ACCOUNTING_FRONTEND_URL")
	FINANCE_FRONTEND = os.Getenv("FINANCE_FRONTEND_URL")
	INVENTORY_FRONTEND = os.Getenv("INVENTORY_FRONTEND_URL")

	HR_MS_BASE_URL = os.Getenv("HR_MS_BASE_URL")
	CORE_MS_BASE_URL = os.Getenv("CORE_MS_BASE_URL")
	PROCUREMENT_MS_BASE_URL = os.Getenv("PROCUREMENT_MS_BASE_URL")
	BASIC_INVENTORY_MS_BASE_URL = os.Getenv("BASIC_INVENTORY_MS_BASE_URL")
	ACCOUNTING_MS_BASE_URL = os.Getenv("ACCOUNTING_MS_BASE_URL")
	FILE_MS_BASE_URL = os.Getenv("FILE_MS_BASE_URL")

	// CORE MS endpoints
	LOGIN_ENDPOINT = CORE_MS_BASE_URL + "/users/login"
	LOGOUT_ENDPOINT = CORE_MS_BASE_URL + "/users/logout"
	REFRESH_ENDPOINT = CORE_MS_BASE_URL + "/refresh"
	ROLES_ENDPOINT = CORE_MS_BASE_URL + "/roles"
	PIN_ENDPOINT = CORE_MS_BASE_URL + "/users/validate-pin"
	USER_ACCOUNTS_ENDPOINT = CORE_MS_BASE_URL + "/users"
	SETTINGS_ENDPOINT = CORE_MS_BASE_URL + "/settings"
	SUPPLIERS_ENDPOINT = CORE_MS_BASE_URL + "/suppliers"
	LOGGED_IN_USER_ENDPOINT = CORE_MS_BASE_URL + "/logged-in-user"
	ACCOUNT_ENDPOINT = CORE_MS_BASE_URL + "/accounts"
	FORGOT_PASSWORD = CORE_MS_BASE_URL + "/v2/users/password/forgot"

	// HR MS endpoints
	USER_PROFILES_ENDPOINT = HR_MS_BASE_URL + "/user-profiles"
	EMPLOYEE_CONTRACTS = HR_MS_BASE_URL + "/employee-contracts"
	EMPLOYEE_EDUCATIONS = HR_MS_BASE_URL + "/employee-educations"
	EMPLOYEE_EXPERIENCES = HR_MS_BASE_URL + "/employee-experiences"
	EVALUATION_TYPES_ENDPOINT = HR_MS_BASE_URL + "/evaluation-types"
	EVALUATIONS = HR_MS_BASE_URL + "/evaluations"
	FOREIGNERS = HR_MS_BASE_URL + "/foreigners"
	SALARIES = HR_MS_BASE_URL + "/salaries"
	EDUCATION_TYPES_ENDPOINT = HR_MS_BASE_URL + "/education-types"
	ORGANIZATION_UNITS_ENDPOINT = HR_MS_BASE_URL + "/organization-units"
	SYSTEMATIZATIONS_ENDPOINT = HR_MS_BASE_URL + "/systematizations"
	JOB_POSITIONS_ENDPOINT = HR_MS_BASE_URL + "/job-positions"
	JOB_POSITIONS_IN_ORGANIZATION_UNITS_ENDPOINT = HR_MS_BASE_URL + "/job-positions-in-organization-units"
	EMPLOYEES_IN_ORGANIZATION_UNITS_ENDPOINT = HR_MS_BASE_URL + "/employees-in-organization-units"
	EMPLOYEES_IN_ORGANIZATION_UNITS_BY_ID_ENDPOINT = HR_MS_BASE_URL + "/employees-in-organization-units-by-id"
	RESOLUTIONS_ENDPOINT = HR_MS_BASE_URL + "/employee-resolutions"
	EMPLOYEE_FAMILY_MEMBERS = HR_MS_BASE_URL + "/employee-family-members"
	JUDGE_NORM_ENDPOINT = HR_MS_BASE_URL + "/user-norms"
	JUDGES = HR_MS_BASE_URL + "/judges"
	ABSENT_TYPE = HR_MS_BASE_URL + "/absent-types"
	EMPLOYEE_ABSENTS = HR_MS_BASE_URL + "/employee-absents"
	REVISIONS_ENDPOINT = HR_MS_BASE_URL + "/revisions-of-organization-units"
	JOB_TENDERS_ENDPOINT = HR_MS_BASE_URL + "/tenders-in-organization-units"
	JOB_TENDER_TYPES_ENDPOINT = HR_MS_BASE_URL + "/tender-types"
	JOB_TENDER_APPLICATIONS_ENDPOINT = HR_MS_BASE_URL + "/tender-applications-in-organization-units"
	JUDGE_RESOLUTIONS_ENDPOINT = HR_MS_BASE_URL + "/judge-number-resolutions"
	JUDGE_RESOLUTION_ITEMS_ENDPOINT = HR_MS_BASE_URL + "/judge-number-resolution-organization-units"
	REVISION_PLAN_ENDPOINT = HR_MS_BASE_URL + "/plans"
	REVISION_ENDPOINT = HR_MS_BASE_URL + "/revisions"
	REVISION_TIPS_ENDPOINT = HR_MS_BASE_URL + "/revision-tips"
	REVISORS_ENDPOINT = HR_MS_BASE_URL + "/get-revisors"
	REVISION_REVISORS_ENDPOINT = HR_MS_BASE_URL + "/revision-revisors"
	REVISION_ORG_UNIT_ENDPOINT = HR_MS_BASE_URL + "/revisions-in-organization-units"

	// public procurement endpoints
	PLANS_ENDPOINT = PROCUREMENT_MS_BASE_URL + "/plans"
	ITEMS_ENDPOINT = PROCUREMENT_MS_BASE_URL + "/items"
	ARTICLES_ENDPOINT = PROCUREMENT_MS_BASE_URL + "/articles"
	CONTRACTS_ENDPOINT = PROCUREMENT_MS_BASE_URL + "/contracts"
	CONTRACT_ARTICLE_ENDPOINT = PROCUREMENT_MS_BASE_URL + "/contract-articles"
	OU_LIMITS_ENDPOINT = PROCUREMENT_MS_BASE_URL + "/organization-unit-plan-limits"
	ORGANIZATION_UNIT_ARTICLE_ENDPOINT = PROCUREMENT_MS_BASE_URL + "/organization-unit-articles"
	CONTRACT_ARTICLE_OVERAGE_ENDPOINT = PROCUREMENT_MS_BASE_URL + "/contract-article-overages"

	// basic inventory endpoints
	INVENTORY_ITEM_ENDOPOINT = BASIC_INVENTORY_MS_BASE_URL + "/items"
	REAL_ESTATES_ENDPOINT = BASIC_INVENTORY_MS_BASE_URL + "/real-estates"
	ASSESSMENTS_ENDPOINT = BASIC_INVENTORY_MS_BASE_URL + "/assessments"
	INVENTORY_DISPATCH_ENDOPOINT = BASIC_INVENTORY_MS_BASE_URL + "/dispatches"
	INVENTORY_DISPATCH_ITEMS_ENDOPOINT = BASIC_INVENTORY_MS_BASE_URL + "/dispatch-items"

	// accounting endpoints
	ORDER_LISTS_ENDPOINT = ACCOUNTING_MS_BASE_URL + "/order-lists"
	ORDER_PROCUREMENT_ARTICLES_ENDPOINT = ACCOUNTING_MS_BASE_URL + "/order-procurement-articles"

	FILES_ENDPOINT = FILE_MS_BASE_URL + "/files"
}
