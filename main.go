package main

import (
	"bff/fields"
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func extractTokenFromHeader(headerValue string) string {
	// Assuming the Authorization header follows the "Bearer <token>" format
	split := strings.Split(headerValue, " ")
	if len(split) == 2 && split[0] == "Bearer" {
		return split[1]
	}
	return "" // Return an empty token if the header format is invalid or empty
}

func extractTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		// Extract the token value from the header
		token := extractTokenFromHeader(authHeader)
		// Store the token value in the request context
		ctx := context.WithValue(r.Context(), "token", token)
		r = r.WithContext(ctx)
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"userAccount_Insert":                              fields.UserAccountInsertField,
			"userAccount_Delete":                              fields.UserAccountDeleteField,
			"settingsDropdown_Insert":                         fields.SettingsDropdownInsertField,
			"settingsDropdown_Delete":                         fields.SettingsDropdownDeleteField,
			"organizationUnits_Insert":                        fields.OrganizationUnitInsertField,
			"organizationUnits_Delete":                        fields.OrganizationUnitDeleteField,
			"jobPositions_Insert":                             fields.JobPositionInsertField,
			"jobPositions_Delete":                             fields.JobPositionDeleteField,
			"jobPositionInOrganizationUnit_Insert":            fields.JobPositionInOrganizationUnitInsertField,
			"jobPositionInOrganizationUnit_Delete":            fields.JobPositionInOrganizationUnitDeleteField,
			"employeeInOrganizationUnit_Insert":               fields.EmployeeInOrganizationUnitInsertField,
			"employeeInOrganizationUnit_Delete":               fields.EmployeeInOrganizationUnitDeleteField,
			"jobTenderTypes_Insert":                           fields.JobTenderTypeInsertField,
			"jobTenderTypes_Delete":                           fields.JobTenderTypeDeleteField,
			"jobTenders_Insert":                               fields.JobTenderInsertField,
			"jobTenders_Delete":                               fields.JobTenderDeleteField,
			"jobTender_Applications_Insert":                   fields.JobTenderApplicationsInsertField,
			"jobTender_Applications_Delete":                   fields.JobTenderApplicationsDeleteField,
			"systematizations_Insert":                         fields.SystematizationInsertField,
			"systematizations_Delete":                         fields.SystematizationDeleteField,
			"userProfile_Basic_Insert":                        fields.UserProfileBasicInsertField,
			"userProfile_Education_Insert":                    fields.UserProfileEducationInsertField,
			"userProfile_Education_Delete":                    fields.UserProfileEducationDeleteField,
			"userProfile_Experience_Insert":                   fields.UserProfileExperienceInsertField,
			"userProfile_Experience_Delete":                   fields.UserProfileExperienceDeleteField,
			"userProfile_Family_Insert":                       fields.UserProfileFamilyInsertField,
			"userProfile_Family_Delete":                       fields.UserProfileFamilyDeleteField,
			"userProfile_Foreigner_Insert":                    fields.UserProfileForeignerInsertField,
			"userProfile_Foreigner_Delete":                    fields.UserProfileForeignerDeleteField,
			"userProfile_SalaryParams_Insert":                 fields.UserProfileSalaryParamsInsertField,
			"userProfile_SalaryParams_Delete":                 fields.UserProfileSalaryParamsDeleteField,
			"userProfile_Evaluation_Insert":                   fields.UserProfileEvaluationInsertField,
			"userProfile_Evaluation_Delete":                   fields.UserProfileEvaluationDeleteField,
			"userProfile_Absent_Insert":                       fields.UserProfileAbsentInsertField,
			"userProfile_Absent_Delete":                       fields.UserProfileAbsentDeleteField,
			"userProfile_Resolution_Insert":                   fields.UserProfileResolutionInsertField,
			"userProfile_Resolution_Delete":                   fields.UserProfileResolutionDeleteField,
			"revisions_Insert":                                fields.RevisionInsertField,
			"revisions_Delete":                                fields.RevisionDeleteField,
			"judgeNorms_Insert":                               fields.JudgeNormsInsertField,
			"judgeNorms_Delete":                               fields.JudgeNormsDeleteField,
			"publicProcurementPlan_Insert":                    fields.PublicProcurementPlanInsertField,
			"publicProcurementPlan_Delete":                    fields.PublicProcurementPlanDeleteField,
			"publicProcurementPlanItem_Insert":                fields.PublicProcurementPlanItemInsertField,
			"publicProcurementPlanItem_Delete":                fields.PublicProcurementPlanItemDeleteField,
			"publicProcurementPlanItemLimit_Insert":           fields.PublicProcurementPlanItemLimitInsertField,
			"publicProcurementPlanItemArticle_Insert":         fields.PublicProcurementPlanItemArticleInsertField,
			"publicProcurementPlanItemArticle_Delete":         fields.PublicProcurementPlanItemArticleDeleteField,
			"publicProcurementOrganizationUnitArticle_Insert": fields.PublicProcurementOrganizationUnitArticleInsertField,
			"publicProcurementContracts_Insert":               fields.PublicProcurementContractsInsertField,
			"publicProcurementContracts_Delete":               fields.PublicProcurementContractsDeleteField,
			"publicProcurementContractArticle_Insert":         fields.PublicProcurementContractArticleInsertField,
			"suppliers_Insert":                                fields.SuppliersInsertField,
			"suppliers_Delete":                                fields.SuppliersDeleteField,
			"basicInventory_Insert":                           fields.BasicInventoryInsertField,
			"basicInventoryAssessments_Insert":                fields.BasicInventoryAssessmentsInsertField,
			"basicInventoryAssessments_Delete":                fields.BasicInventoryAssessmentsDeleteField,
			"basicInventoryDispatch_Insert":                   fields.BasicInventoryDispatchInsertField,
			"basicInventoryDispatch_Delete":                   fields.BasicInventoryDispatchDeleteField,
			"basicInventoryDispatch_Accept":                   fields.BasicInventoryDispatchAcceptField,
			"basicInventoryDepreciationTypes_Insert":          fields.BasicInventoryDepreciationTypesInsertField,
			"basicInventoryDepreciationTypes_Delete":          fields.BasicInventoryDepreciationTypesDeleteField,
			"officesOfOrganizationUnits_Insert":               fields.OfficesOfOrganizationUnitInsertField,
			"officesOfOrganizationUnits_Delete":               fields.OfficesOfOrganizationUnitDeleteField,
		},
	})
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"login":                                              fields.LoginField,
			"userAccount_Overview":                               fields.UserAccountField,
			"settingsDropdown_Overview":                          fields.SettingsDropdownField,
			"organizationUnits":                                  fields.OrganizationUnitsField,
			"jobPositions":                                       fields.JobPositionsField,
			"jobTenderTypes":                                     fields.JobTenderTypesField,
			"jobTenders_Overview":                                fields.JobTendersOverviewField,
			"jobTender_Details":                                  fields.JobTenderDetailsField,
			"jobTender_Applications":                             fields.JobTenderApplicationsField,
			"systematizations_Overview":                          fields.SystematizationsOverviewField,
			"systematization_Details":                            fields.SystematizationDetailsField,
			"userProfiles_Overview":                              fields.UserProfilesOverviewField,
			"userProfile_Basic":                                  fields.UserProfileBasicField,
			"userProfile_Education":                              fields.UserProfileEducationField,
			"userProfile_Experience":                             fields.UserProfileExperienceField,
			"userProfile_Family":                                 fields.UserProfileFamilyField,
			"userProfile_Foreigner":                              fields.UserProfileForeignerField,
			"userProfile_SalaryParams":                           fields.UserProfileSalaryParamsField,
			"userProfile_Evaluation":                             fields.UserProfileEvaluationField,
			"userProfile_Absent":                                 fields.UserProfileAbsentField,
			"userProfile_Resolution":                             fields.UserProfileResolutionField,
			"revisions_Overview":                                 fields.RevisionsOverviewField,
			"revision_Details":                                   fields.RevisionDetailsField,
			"judges_Overview":                                    fields.JudgesOverviewField,
			"judgeResolutions_Overview":                          fields.JudgeResolutionsOverviewField,
			"judgeResolution_Details":                            fields.JudgeResolutionDetailsField,
			"publicProcurementPlans_Overview":                    fields.PublicProcurementPlansOverviewField,
			"publicProcurementPlan_Details":                      fields.PublicProcurementPlanDetailsField,
			"publicProcurementPlanItem_Details":                  fields.PublicProcurementPlanItemDetailsField,
			"publicProcurementPlanItem_Limits":                   fields.PublicProcurementPlanItemLimitsField,
			"publicProcurementOrganizationUnitArticles_Overview": fields.PublicProcurementOrganizationUnitArticlesOverviewField,
			"publicProcurementOrganizationUnitArticles_Details":  fields.PublicProcurementOrganizationUnitArticlesDetailsField,
			"publicProcurementContracts_Overview":                fields.PublicProcurementContractsOverviewField,
			"publicProcurementContractArticles_Overview":         fields.PublicProcurementContractArticlesOverviewField,
			"suppliers_Overview":                                 fields.SuppliersOverviewField,
			"basicInventory_Overview":                            fields.BasicInventoryOverviewField,
			"basicInventory_Details":                             fields.BasicInventoryDetailsField,
			"basicInventoryDepreciationTypes_Overview":           fields.BasicInventoryDepreciationTypesOverviewField,
			"basicInventoryRealEstates_Overview":                 fields.BasicInventoryRealEstatesOverviewField,
			"officesOfOrganizationUnits_Overview":                fields.OfficesOfOrganizationUnitOverviewField,
			"basicInventoryDispatch_Overview":                    fields.BasicInventoryDispatchOverviewField,
		},
	})
	schemaConfig := graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		panic(err)
	}
	// Create a GraphQL HTTP handler
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// Create a new HTTP handler function to serve the JSON files
	fs := http.FileServer(http.Dir("mocked-data"))
	http.Handle("/mocked-data/", http.StripPrefix("/mocked-data", fs))
	// Create a CORS-enabled handler
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	// Insert the custom middleware handler
	graphqlHandler := extractTokenMiddleware(corsHandler(h))
	// Start your HTTP server with the CORS-enabled handler
	http.Handle("/", graphqlHandler)
	http.ListenAndServe(":8080", nil)
}
