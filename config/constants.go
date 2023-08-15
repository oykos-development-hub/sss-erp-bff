package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ContextKey string

const (
	HttpResponseWriterKey ContextKey = "httpResponseWriter"
	HttpHeadersKey        ContextKey = "httpHeaders"
	EducationTypes        string     = "education_types"
)

var (
	LOGGED_IN_USER_ENDPOINT string

	HR_MS_BASE_URL                               string
	CORE_MS_BASE_URL                             string
	LOGIN_ENDPOINT                               string
	PIN_ENDPOINT                                 string
	USER_ACCOUNTS_ENDPOINT                       string
	ROLES_ENDPOINT                               string
	CONTRACT_TYPES_ENDPOINT                      string
	SETTINGS_ENDPOINT                            string
	EVALUATIONS                                  string
	EVALUATION_TYPES_ENDPOINT                    string
	FOREIGNERS                                   string
	SALARIES                                     string
	ORGANIZATION_UNITS_ENDPOINT                  string
	JOB_POSITIONS_ENDPOINT                       string
	JOB_POSITIONS_IN_ORGANIZATION_UNITS_ENDPOINT string
	EDUCATION_TYPES_ENDPOINT                     string
	USER_PROFILES_ENDPOINT                       string
	EMPLOYEE_CONTRACTS                           string
	EMPLOYEE_EDUCATIONS                          string
	EMPLOYEE_EXPERIENCES                         string
	EMPLOYEES_IN_ORGANIZATION_UNITS_ENDPOINT     string
	RESOLUTIONS_ENDPOINT                         string
	SYSTEMATIZATIONS_ENDPOINT                    string
	EMPLOYEE_FAMILY_MEMBERS                      string
	JUDGE_NORM_ENDPOINT                          string
	ABSENT_TYPE                                  string
	EMPLOYEE_ABSENTS                             string
	REVISIONS_ENDPOINT                           string
	JOB_TENDERS_ENDPOINT                         string
	JOB_TENDER_TYPES_ENDPOINT                    string
	JOB_TENDER_APPLICATIONS_ENDPOINT             string
	JUDGE_RESOLUTIONS_ENDPOINT                   string
	JUDGE_RESOLUTION_ITEMS_ENDPOINT              string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	HR_MS_BASE_URL = os.Getenv("HR_MS_BASE_URL")
	CORE_MS_BASE_URL = os.Getenv("CORE_MS_BASE_URL")
	LOGIN_ENDPOINT = CORE_MS_BASE_URL + "/users/login"
	ROLES_ENDPOINT = CORE_MS_BASE_URL + "/roles"
	PIN_ENDPOINT = CORE_MS_BASE_URL + "/users/validate-pin"
	USER_ACCOUNTS_ENDPOINT = CORE_MS_BASE_URL + "/users"
	USER_PROFILES_ENDPOINT = HR_MS_BASE_URL + "/user-profiles"
	EMPLOYEE_CONTRACTS = HR_MS_BASE_URL + "/employee-contracts"
	EMPLOYEE_EDUCATIONS = HR_MS_BASE_URL + "/employee-educations"
	EMPLOYEE_EXPERIENCES = HR_MS_BASE_URL + "/employee-experiences"
	CONTRACT_TYPES_ENDPOINT = HR_MS_BASE_URL + "/contract-types"
	SETTINGS_ENDPOINT = CORE_MS_BASE_URL + "/settings"
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
	RESOLUTIONS_ENDPOINT = HR_MS_BASE_URL + "/employee-resolutions"
	EMPLOYEE_FAMILY_MEMBERS = HR_MS_BASE_URL + "/employee-family-members"
	JUDGE_NORM_ENDPOINT = HR_MS_BASE_URL + "/user-norms"
	ABSENT_TYPE = HR_MS_BASE_URL + "/absent-types"
	EMPLOYEE_ABSENTS = HR_MS_BASE_URL + "/employee-absents"
	REVISIONS_ENDPOINT = HR_MS_BASE_URL + "/revisions-of-organization-units"
	JOB_TENDERS_ENDPOINT = HR_MS_BASE_URL + "/tenders-in-organization-units"
	JOB_TENDER_TYPES_ENDPOINT = HR_MS_BASE_URL + "/tender-types"
	JOB_TENDER_APPLICATIONS_ENDPOINT = HR_MS_BASE_URL + "/tender-applications-in-organization-units"
	JUDGE_RESOLUTIONS_ENDPOINT = HR_MS_BASE_URL + "/judge-number-resolutions"
	JUDGE_RESOLUTION_ITEMS_ENDPOINT = HR_MS_BASE_URL + "/judge-number-resolution-organization-units"
	LOGGED_IN_USER_ENDPOINT = CORE_MS_BASE_URL + "/logged-in-user"

}
