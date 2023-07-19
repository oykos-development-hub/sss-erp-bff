package shared

import (
	"bff/structs"
	"fmt"
	"reflect"
)

var FetchByProperty = func(entity string, property string, value interface{}, contain ...bool) []interface{} {
	var endpoint string
	var entityStruct interface{}
	var isValueString = reflect.TypeOf(value).Kind() == reflect.String

	switch entity {
	case "unit", "units", "organization_unit", "organization_units":
		endpoint = "organization_units.json"
		entityStruct = &structs.OrganizationUnits{}
	case "user_account", "user_accounts":
		endpoint = "user_accounts.json"
		entityStruct = &structs.UserAccounts{}
	case "profile", "profiles", "user_profile", "user_profiles":
		endpoint = "user_profiles.json"
		entityStruct = &structs.UserProfiles{}
	case "role", "roles", "user_account_role", "user_account_roles":
		endpoint = "user_account_roles.json"
		entityStruct = &structs.UserAccountRoles{}
	case "position", "positions", "job_position", "job_positions":
		endpoint = "job_positions.json"
		entityStruct = &structs.JobPositions{}
	case "job_positions_in_organization_units", "job_position_in_organization_unit", "job_position_in_organization_units":
		endpoint = "job_positions_in_organization_units.json"
		entityStruct = &structs.JobPositionsInOrganizationUnits{}
	case "employees_in_organization_units", "employee_in_organization_unit", "employee_in_organization_units":
		endpoint = "employees_in_organization_units.json"
		entityStruct = &structs.EmployeesInOrganizationUnits{}
	case "contract", "contracts", "employee_contract", "employee_contracts":
		endpoint = "contracts.json"
		entityStruct = &structs.Contracts{}
	case "contract_type", "contracts_types", "employee_contract_type", "employee_contracts_types":
		endpoint = "contract_types.json"
		entityStruct = &structs.ContractType{}
	case "education_type", "education_types":
		endpoint = "education_types.json"
		entityStruct = &structs.EducationType{}
	case "education", "educations":
		endpoint = "educations.json"
		entityStruct = &structs.Education{}
	case "experience", "experiences":
		endpoint = "user_profile_experiences.json"
		entityStruct = &structs.Experience{}
	case "family", "families":
		endpoint = "user_profile_family.json"
		entityStruct = &structs.Family{}
	case "foreigner", "foreigners":
		endpoint = "user_profile_foreigners.json"
		entityStruct = &structs.Foreigners{}
	case "salary_params", "salary_param":
		endpoint = "user_profile_salary_params.json"
		entityStruct = &structs.SalaryParams{}
	case "evaluation_type", "evaluation_types":
		endpoint = "user_profile_evaluation_types.json"
		entityStruct = &structs.EvaluationType{}
	case "evaluation", "evaluations":
		endpoint = "user_profile_evaluations.json"
		entityStruct = &structs.Evaluation{}
	case "vacation_type", "vacation_types":
		endpoint = "user_profile_vacation_types.json"
		entityStruct = &structs.AbsentType{}
	case "vacation", "vacations":
		endpoint = "user_profile_vacations.json"
		entityStruct = &structs.Absent{}
	case "relocation", "relocations":
		endpoint = "user_profile_relocations.json"
		entityStruct = &structs.Absent{}
	case "resolution_type", "resolution_types":
		endpoint = "user_profile_resolution_types.json"
		entityStruct = &structs.ResolutionType{}
	case "resolution", "resolutions":
		endpoint = "user_profile_resolutions.json"
		entityStruct = &structs.Resolution{}
	case "revision", "revisions":
		endpoint = "revisions.json"
		entityStruct = &structs.Revision{}
	case "job_tender", "job_tenders":
		endpoint = "job_tenders.json"
		entityStruct = &structs.JobTenders{}
	case "norm", "norms", "judge_norm", "judge_norms":
		endpoint = "judge_norms.json"
		entityStruct = &structs.JudgeNorms{}
	case "judge_resolution", "judge_resolutions":
		endpoint = "judge_resolutions.json"
		entityStruct = &structs.JudgeResolutions{}
	case "judge_resolution_item", "judge_resolution_items":
		endpoint = "judge_resolution_items.json"
		entityStruct = &structs.JudgeResolutionItems{}
	case "public_procurement_plan", "public_procurement_plans":
		endpoint = "public_procurement_plans.json"
		entityStruct = &structs.PublicProcurementPlan{}
	case "procurement", "procurements", "public_procurement_item", "public_procurement_items":
		endpoint = "public_procurement_items.json"
		entityStruct = &structs.PublicProcurementItem{}
	case "public_procurement_article", "public_procurement_articles":
		endpoint = "public_procurement_articles.json"
		entityStruct = &structs.PublicProcurementArticle{}
	case "public_procurement_limit", "public_procurement_limits":
		endpoint = "public_procurement_organization_unit_limits.json"
		entityStruct = &structs.PublicProcurementLimit{}
	case "public_procurement_organization_unit_article", "public_procurement_organization_unit_articles":
		endpoint = "public_procurement_organization_unit_articles.json"
		entityStruct = &structs.PublicProcurementOrganizationUnitArticle{}
	case "budget_indents", "budget_indent", "indents", "indent":
		endpoint = "budget_indents.json"
		entityStruct = &structs.BudgetIndent{}
	case "public_procurement_contract", "public_procurement_contracts":
		endpoint = "public_procurement_contracts.json"
		entityStruct = &structs.PublicProcurementContract{}
	case "public_procurement_contract_article", "public_procurement_contract_articles":
		endpoint = "public_procurement_contract_articles.json"
		entityStruct = &structs.PublicProcurementContractArticle{}
	case "suppliers", "supplier":
		endpoint = "suppliers.json"
		entityStruct = &structs.Suppliers{}
	case "basic_inventory_depreciation_types", "basic_inventory_depreciation_type":
		endpoint = "basic_inventory_depreciation_types.json"
		entityStruct = &structs.BasicInventoryDepreciationTypesItem{}
	case "basic_inventory_real_estates", "basic_inventory_real_estate":
		endpoint = "basic_inventory_real_estates.json"
		entityStruct = &structs.BasicInventoryRealEstatesItem{}
	case "offices_of_organization_units", "offices_of_organization_unit":
		endpoint = "offices_of_organization_units.json"
		entityStruct = &structs.OfficesOfOrganizationUnitItem{}
	case "inventory_class_type":
		endpoint = "settings_dropdown_options.json"
		entityStruct = &structs.SettingsDropdown{}
	case "order_procurement_article":
		endpoint = "order_procurement_article.json"
		entityStruct = &structs.OrderProcurementArticleItem{}
	case "accounts", "account":
		endpoint = "account.json"
		entityStruct = &structs.AccountItem{}
	case "activities":
		endpoint = "activities.json"
		entityStruct = &structs.ActivitiesItem{}
	}

	entityData, entityDataErr := ReadJson("http://localhost:8080/mocked-data/"+endpoint, entityStruct)

	if entityDataErr != nil {
		fmt.Printf("Fetching "+entity+" failed because of this error - %s.\n", entityDataErr)
	}

	if len(property) > 0 && ((isValueString && len(property) > 0) || (!isValueString && value != nil)) {
		return FindByProperty(entityData, property, value, contain...)
	}

	return entityData
}
