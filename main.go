package main

import (
	"bff/config"
	"bff/fields"
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func extractTokenFromHeader(headerValue string) string {
	if headerValue == "" {
		return ""
	}
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
		ctx := context.WithValue(r.Context(), config.TokenKey, token)
		r = r.WithContext(ctx)
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func errorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a buffer to capture the response
		buf := &bytes.Buffer{}
		responseWriter := httptest.NewRecorder()
		// Replace the original response writer with the recorder
		defer func() {
			// Check for errors in the response
			if responseWriter.Code >= http.StatusBadRequest {
				// Handle the error by logging or returning a custom error message
				log.Println("HTTP error:", responseWriter.Code, buf.String())
			}
			// Copy the response from the recorder to the original writer
			for key, values := range responseWriter.Header() {
				w.Header()[key] = values
			}
			w.WriteHeader(responseWriter.Code)
			_, _ = buf.WriteTo(w)
		}()
		// Replace the response writer with the buffer
		responseWriter.Body = buf
		// Pass the modified response writer to the next handler
		next.ServeHTTP(responseWriter, r)
	})
}

func main() {
	// Open the log file for writing
	err := os.MkdirAll("./log", 0777)
	if err != nil {
		fmt.Println("Failed to create a log directory:", err)
		return
	}
	logFile, err := os.OpenFile("./log/sss-erp-bff.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	// Set the log output to the log file
	log.SetOutput(logFile)
	// Redirect standard error to the log file
	os.Stderr = logFile
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
			"jobTenderTypes_Insert":                           fields.JobTenderTypeInsertField,
			"jobTenderTypes_Delete":                           fields.JobTenderTypeDeleteField,
			"jobTenders_Insert":                               fields.JobTenderInsertField,
			"jobTenders_Delete":                               fields.JobTenderDeleteField,
			"jobTender_Applications_Insert":                   fields.JobTenderApplicationsInsertField,
			"jobTender_Applications_Delete":                   fields.JobTenderApplicationsDeleteField,
			"systematizations_Insert":                         fields.SystematizationInsertField,
			"systematizations_Delete":                         fields.SystematizationDeleteField,
			"userProfile_Basic_Insert":                        fields.UserProfileBasicInsertField,
			"userProfile_Update":                              fields.UserProfileUpdateField,
			"userProfile_Contract_Insert":                     fields.UserProfileContractInsertField,
			"userProfile_Contract_Delete":                     fields.UserProfileContractDeleteField,
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
			"absentType_Insert":                               fields.AbsentTypeInsertField,
			"absentType_Delete":                               fields.AbsentTypeDeleteField,
			"userProfile_Absent_Insert":                       fields.UserProfileAbsentInsertField,
			"userProfile_Vacation_Insert":                     fields.UserProfileVacationInsertField,
			"userProfile_Absent_Delete":                       fields.UserProfileAbsentDeleteField,
			"userProfile_Resolution_Insert":                   fields.UserProfileResolutionInsertField,
			"userProfile_Resolution_Delete":                   fields.UserProfileResolutionDeleteField,
			"revisions_Insert":                                fields.RevisionInsertField,
			"revisions_Delete":                                fields.RevisionDeleteField,
			"judgeNorms_Insert":                               fields.JudgeNormsInsertField,
			"judgeNorms_Delete":                               fields.JudgeNormsDeleteField,
			"judgeResolutions_Insert":                         fields.JudgeResolutionsInsertField,
			"judgeResolutions_Delete":                         fields.JudgeResolutionsDeleteField,
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
			"officesOfOrganizationUnits_Insert":               fields.OfficesOfOrganizationUnitInsertField,
			"officesOfOrganizationUnits_Delete":               fields.OfficesOfOrganizationUnitDeleteField,
			"orderList_Insert":                                fields.OrderListInsertField,
			"orderList_Receive":                               fields.OrderListReceiveField,
			"orderList_Movement":                              fields.OrderListAssetMovementField,
			"orderList_Delete":                                fields.OrderListDeleteField,
			"orderListReceive_Delete":                         fields.OrderListReceiveDeleteField,
			"orderListAssetMovement_Delete":                   fields.OrderListAssetMovementDeleteField,
			"activities_Delete":                               fields.ActivitiesDeleteField,
			"account_Delete":                                  fields.AccountDeleteField,
			"program_Delete":                                  fields.ProgramDeleteField,
			"budget_Delete":                                   fields.BudgetDeleteField,
			"activities_Insert":                               fields.ActivitiesInsertField,
			"account_Insert":                                  fields.AccountInsertField,
			"program_Insert":                                  fields.ProgramInsertField,
			"budget_Insert":                                   fields.BudgetInsertField,
			"budget_Send":                                     fields.BudgetSendField,
			"accountBudgetActivity_Insert":                    fields.AccountBudgetActivityInsertField,
		},
	})
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"login":                        fields.LoginField,
			"logout":                       fields.LogoutField,
			"refresh":                      fields.RefreshField,
			"pin":                          fields.PinField,
			"userAccount_Overview":         fields.UserAccountField,
			"settingsDropdown_Overview":    fields.SettingsDropdownField,
			"organizationUnits":            fields.OrganizationUnitsField,
			"jobPositions":                 fields.JobPositionsField,
			"jobPositionsOrganizationUnit": fields.JobPositionsOrganizationUnitField,
			"jobPositionsAvailableInOrganizationUnit": fields.JobPositionAvailableInOrganizationUnitField,
			"jobTenderTypes":                                     fields.JobTenderTypesField,
			"jobTenders_Overview":                                fields.JobTendersOverviewField,
			"jobTender_Applications":                             fields.JobTenderApplicationsField,
			"systematizations_Overview":                          fields.SystematizationsOverviewField,
			"systematization_Details":                            fields.SystematizationDetailsField,
			"userProfiles_Overview":                              fields.UserProfilesOverviewField,
			"userProfile_Contracts":                              fields.UserProfileContractsField,
			"userProfile_Basic":                                  fields.UserProfileBasicField,
			"userProfile_Education":                              fields.UserProfileEducationField,
			"userProfile_Experience":                             fields.UserProfileExperienceField,
			"userProfile_Family":                                 fields.UserProfileFamilyField,
			"userProfile_Foreigner":                              fields.UserProfileForeignerField,
			"userProfile_SalaryParams":                           fields.UserProfileSalaryParamsField,
			"userProfile_Evaluation":                             fields.UserProfileEvaluationField,
			"userProfile_Absent":                                 fields.UserProfileAbsentField,
			"userProfile_Vacation":                               fields.UserProfileVacationField,
			"absentType":                                         fields.AbsentTypeField,
			"userProfile_Resolution":                             fields.UserProfileResolutionField,
			"revisions_Overview":                                 fields.RevisionsOverviewField,
			"revision_Details":                                   fields.RevisionDetailsField,
			"judges_Overview":                                    fields.JudgesOverviewField,
			"judgeResolutions_Overview":                          fields.JudgeResolutionsOverviewField,
			"organizationUintCalculateEmployeeStats":             fields.OrganizationUintCalculateEmployeeStatsField,
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
			"basicInventoryRealEstates_Overview":                 fields.BasicInventoryRealEstatesOverviewField,
			"officesOfOrganizationUnits_Overview":                fields.OfficesOfOrganizationUnitOverviewField,
			"basicInventoryDispatch_Overview":                    fields.BasicInventoryDispatchOverviewField,
			"orderList_Overview":                                 fields.OrderListOverviewField,
			"orderProcurementAvailableList_Overview":             fields.OrderProcurementAvailableField,
			"recipientUsers_Overview":                            fields.RecipientUsersField,
			"account_Overview":                                   fields.AccountOverviewField,
			"accountBudgetActivity_Overview":                     fields.AccountBudgetActivityOverviewField,
			"activities_Overview":                                fields.ActivitiesOverviewField,
			"budget_Overview":                                    fields.BudgetOverviewField,
			"programs_Overview":                                  fields.ProgramOverviewField,
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
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Create a new HTTP handler function to serve the JSON files
	fs := http.FileServer(http.Dir("mocked-data"))
	http.Handle("/mocked-data/", http.StripPrefix("/mocked-data", fs))
	// Create a CORS-enabled handler
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:3002",
			"http://localhost:3003",
			"http://localhost:3004",
			"http://localhost:3005",
			"https://localhost:3000",
			"https://localhost:3001",
			"https://localhost:3002",
			"https://localhost:3003",
			"https://localhost:3004",
			"https://localhost:3005",
			config.HR_FRONTEND,
			config.PROCUREMENTS_FRONTEND,
			config.ACCOUNTING_FRONTEND,
			config.INVENTORY_FRONTEND,
			config.FINANCE_FRONTEND,
			config.CORE_FRONTEND,
		}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)
	// Insert the custom middleware handler
	graphqlHandler := errorHandlerMiddleware(extractTokenMiddleware(corsHandler(addResponseWriterToContext(RequestContextMiddleware(h)))))
	// Start your HTTP server with the CORS-enabled handler
	http.Handle("/", graphqlHandler)
	_ = http.ListenAndServe(":8080", nil)
}

func RequestContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), config.Requestkey, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func addResponseWriterToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), config.HttpResponseWriterKey, w)
		// Retrieve the Authorization header value from the request
		authHeader := r.Header.Get("Authorization")
		// Add the bearer token as a header in the context
		headers := map[string]string{
			"Authorization": authHeader,
		}
		ctx = context.WithValue(ctx, config.HttpHeadersKey, headers)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
