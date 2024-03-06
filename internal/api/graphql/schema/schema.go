package schema

import (
	"bff/config"
	"bff/internal/api/graphql/fields"
	"bff/internal/api/graphql/resolvers"
	"bff/internal/api/repository"
	"bff/internal/api/websockets/notifications"

	"github.com/graphql-go/graphql"
)

func SetupGraphQLSchema(notificationService *notifications.Websockets, repo repository.MicroserviceRepositoryInterface, cfg *config.Config) (*graphql.Schema, error) {
	resolvers := resolvers.NewResolver(cfg, notificationService, repo)
	fields := fields.NewFields(resolvers)

	mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"role_Insert":                                     fields.RoleInsertField(),
			"permissions_Update":                              fields.PermissionsUpdate(),
			"userAccount_Insert":                              fields.UserAccountInsertField(),
			"userAccount_Delete":                              fields.UserAccountDeleteField(),
			"settingsDropdown_Insert":                         fields.SettingsDropdownInsertField(),
			"settingsDropdown_Delete":                         fields.SettingsDropdownDeleteField(),
			"organizationUnits_Insert":                        fields.OrganizationUnitInsertField(),
			"organizationUnits_Order":                         fields.OrganizationUnitOrderField(),
			"organizationUnits_Delete":                        fields.OrganizationUnitDeleteField(),
			"jobPositions_Insert":                             fields.JobPositionInsertField(),
			"jobPositions_Delete":                             fields.JobPositionDeleteField(),
			"jobPositionInOrganizationUnit_Insert":            fields.JobPositionInOrganizationUnitInsertField(),
			"jobPositionInOrganizationUnit_Delete":            fields.JobPositionInOrganizationUnitDeleteField(),
			"jobTenderTypes_Insert":                           fields.JobTenderTypeInsertField(),
			"jobTenderTypes_Delete":                           fields.JobTenderTypeDeleteField(),
			"jobTenders_Insert":                               fields.JobTenderInsertField(),
			"jobTenders_Delete":                               fields.JobTenderDeleteField(),
			"jobTender_Applications_Insert":                   fields.JobTenderApplicationsInsertField(),
			"jobTender_Applications_Delete":                   fields.JobTenderApplicationsDeleteField(),
			"systematizations_Insert":                         fields.SystematizationInsertField(),
			"systematizations_Delete":                         fields.SystematizationDeleteField(),
			"userProfile_Basic_Insert":                        fields.UserProfileBasicInsertField(),
			"userProfile_Update":                              fields.UserProfileUpdateField(),
			"userProfile_Contract_Insert":                     fields.UserProfileContractInsertField(),
			"userProfile_Contract_Delete":                     fields.UserProfileContractDeleteField(),
			"userProfile_Education_Insert":                    fields.UserProfileEducationInsertField(),
			"userProfile_Education_Delete":                    fields.UserProfileEducationDeleteField(),
			"userProfile_Experience_Insert":                   fields.UserProfileExperienceInsertField(),
			"userProfile_Experience_Delete":                   fields.UserProfileExperienceDeleteField(),
			"userProfile_Family_Insert":                       fields.UserProfileFamilyInsertField(),
			"userProfile_Family_Delete":                       fields.UserProfileFamilyDeleteField(),
			"userProfile_Foreigner_Insert":                    fields.UserProfileForeignerInsertField(),
			"userProfile_Foreigner_Delete":                    fields.UserProfileForeignerDeleteField(),
			"userProfile_SalaryParams_Insert":                 fields.UserProfileSalaryParamsInsertField(),
			"userProfile_SalaryParams_Delete":                 fields.UserProfileSalaryParamsDeleteField(),
			"userProfile_Evaluation_Insert":                   fields.UserProfileEvaluationInsertField(),
			"userProfile_Evaluation_Delete":                   fields.UserProfileEvaluationDeleteField(),
			"absentType_Insert":                               fields.AbsentTypeInsertField(),
			"absentType_Delete":                               fields.AbsentTypeDeleteField(),
			"userProfile_Absent_Insert":                       fields.UserProfileAbsentInsertField(),
			"userProfile_Vacation_Insert":                     fields.UserProfileVacationInsertField(),
			"userProfile_Vacations_Insert":                    fields.UserProfileVacationsInsertField(),
			"userProfile_Absent_Delete":                       fields.UserProfileAbsentDeleteField(),
			"userProfile_Resolution_Insert":                   fields.UserProfileResolutionInsertField(),
			"userProfile_Resolution_Delete":                   fields.UserProfileResolutionDeleteField(),
			"revisions_Insert":                                fields.RevisionInsertField(),
			"revisions_Delete":                                fields.RevisionDeleteField(),
			"revision_plans_Insert":                           fields.RevisionPlansInsertField(),
			"revision_plans_Delete":                           fields.RevisionPlansDelete(),
			"revision_Insert":                                 fields.RevisionInsert(),
			"revision_Delete":                                 fields.RevisionDelete(),
			"revision_tips_Insert":                            fields.RevisionTipsInsert(),
			"revision_tips_Delete":                            fields.RevisionTipsDelete(),
			"judgeNorms_Insert":                               fields.JudgeNormsInsertField(),
			"judgeNorms_Delete":                               fields.JudgeNormsDeleteField(),
			"judgeResolutions_Insert":                         fields.JudgeResolutionsInsertField(),
			"judgeResolutions_Delete":                         fields.JudgeResolutionsDeleteField(),
			"publicProcurementPlan_Insert":                    fields.PublicProcurementPlanInsertField(),
			"publicProcurementPlan_Delete":                    fields.PublicProcurementPlanDeleteField(),
			"publicProcurementPlanItem_Insert":                fields.PublicProcurementPlanItemInsertField(),
			"publicProcurementPlanItem_Delete":                fields.PublicProcurementPlanItemDeleteField(),
			"publicProcurementPlanItemLimit_Insert":           fields.PublicProcurementPlanItemLimitInsertField(),
			"publicProcurementPlanItemArticle_Insert":         fields.PublicProcurementPlanItemArticleInsertField(),
			"publicProcurementPlanItemArticle_Delete":         fields.PublicProcurementPlanItemArticleDeleteField(),
			"publicProcurementOrganizationUnitArticle_Insert": fields.PublicProcurementOrganizationUnitArticleInsertField(),
			"publicProcurementSendPlanOnRevision_Update":      fields.PublicProcurementSendPlanOnRevision(),
			"publicProcurementContracts_Insert":               fields.PublicProcurementContractsInsertField(),
			"publicProcurementContracts_Delete":               fields.PublicProcurementContractsDeleteField(),
			"publicProcurementContractArticle_Insert":         fields.PublicProcurementContractArticleInsertField(),
			"publicProcurementContractArticleOverage_Insert":  fields.PublicProcurementContractArticleOverageInsertField(),
			"publicProcurementContractArticleOverage_Delete":  fields.PublicProcurementContractArticleOverageDeleteField(),
			"suppliers_Insert":                                fields.SuppliersInsertField(),
			"suppliers_Delete":                                fields.SuppliersDeleteField(),
			"basicInventory_Insert":                           fields.BasicInventoryInsertField(),
			"basicInventory_Deactivate":                       fields.BasicInventoryDeactivateField(),
			"basicInventoryAssessments_Insert":                fields.BasicInventoryAssessmentsInsertField(),
			"basicEXCLInventoryAssessments_Insert":            fields.BasicEXCLInventoryAssessmentsInsertField(),
			"basicInventoryAssessments_Delete":                fields.BasicInventoryAssessmentsDeleteField(),
			"basicInventoryDispatch_Insert":                   fields.BasicInventoryDispatchInsertField(),
			"basicInventoryDispatch_Delete":                   fields.BasicInventoryDispatchDeleteField(),
			"basicInventoryDispatch_Accept":                   fields.BasicInventoryDispatchAcceptField(),
			"officesOfOrganizationUnits_Insert":               fields.OfficesOfOrganizationUnitInsertField(),
			"officesOfOrganizationUnits_Delete":               fields.OfficesOfOrganizationUnitDeleteField(),
			"orderList_Insert":                                fields.OrderListInsertField(),
			"orderList_Receive":                               fields.OrderListReceiveField(),
			"orderList_Delete":                                fields.OrderListDeleteField(),
			"orderListReceive_Delete":                         fields.OrderListReceiveDeleteField(),
			"movement_Insert":                                 fields.MovementInsertField(),
			"movement_Delete":                                 fields.MovementDeleteField(),
			"activities_Delete":                               fields.ActivitiesDeleteField(),
			"account_Delete":                                  fields.AccountDeleteField(),
			"program_Delete":                                  fields.ProgramDeleteField(),
			"budget_Delete":                                   fields.BudgetDeleteField(),
			"activities_Insert":                               fields.ActivitiesInsertField(),
			"account_Insert":                                  fields.AccountInsertField(),
			"program_Insert":                                  fields.ProgramInsertField(),
			"budget_Insert":                                   fields.BudgetInsertField(),
			"budget_Send":                                     fields.BudgetSendField(),
			"budget_SendOnReview":                             fields.BudgetSendOnReviewField(),
			"accountBudgetActivity_Insert":                    fields.AccountBudgetActivityInsertField(),
			"nonFinancialBudget_Insert":                       fields.NonFinancialBudgetInsertField(),
			"nonFinancialGoalIndicator_Insert":                fields.NonFinacialGoalIndicatorInsertField(),
			"nonFinacialBudgetGoal_Insert":                    fields.NonFinacialBudgetGoalInsertField(),
			"financialBudget_Fill":                            fields.FinancialBudgetFillField(),
			"financialBudgetVersion_Update":                   fields.FinancialBudgetVersionUpdateField(),
			"invoice_Insert":                                  fields.InvoiceInsertField(),
			"invoice_Delete":                                  fields.InvoiceDeleteField(),
			"fine_Insert":                                     fields.FineInsertField(),
			"fine_Delete":                                     fields.FineDeleteField(),
			"finePayment_Insert":                              fields.FinePaymentInsertField(),
			"finePayment_Delete":                              fields.FinePaymentDeleteField(),
		},
	})
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"role_Overview":                fields.RoleOverviewField(),
			"role_Details":                 fields.RoleDetailsField(),
			"permissionsForRole":           fields.PermissionsForRoleField(),
			"login":                        fields.LoginField(),
			"logout":                       fields.LogoutField(),
			"refresh":                      fields.RefreshField(),
			"pin":                          fields.PinField(),
			"userAccount_Overview":         fields.UserAccountField(),
			"userAccount_ForgotPassword":   fields.UserForgotPassword(),
			"userAccount_ValidateEmail":    fields.UserValidateMail(),
			"userAccount_ResetPassword":    fields.UserResetPassword(),
			"settingsDropdown_Overview":    fields.SettingsDropdownField(),
			"organizationUnits":            fields.OrganizationUnitsField(),
			"jobPositions":                 fields.JobPositionsField(),
			"jobPositionsOrganizationUnit": fields.JobPositionsOrganizationUnitField(),
			"jobPositionsAvailableInOrganizationUnit": fields.JobPositionAvailableInOrganizationUnitField(),
			"jobTenderTypes":                                             fields.JobTenderTypesField(),
			"jobTenders_Overview":                                        fields.JobTendersOverviewField(),
			"jobTender_Applications":                                     fields.JobTenderApplicationsField(),
			"systematizations_Overview":                                  fields.SystematizationsOverviewField(),
			"systematization_Details":                                    fields.SystematizationDetailsField(),
			"userProfiles_Overview":                                      fields.UserProfilesOverviewField(),
			"userProfile_Contracts":                                      fields.UserProfileContractsField(),
			"userProfile_Basic":                                          fields.UserProfileBasicField(),
			"userProfile_Education":                                      fields.UserProfileEducationField(),
			"userProfile_Experience":                                     fields.UserProfileExperienceField(),
			"userProfile_Family":                                         fields.UserProfileFamilyField(),
			"userProfile_Foreigner":                                      fields.UserProfileForeignerField(),
			"userProfile_SalaryParams":                                   fields.UserProfileSalaryParamsField(),
			"userProfile_Evaluation":                                     fields.UserProfileEvaluationField(),
			"judgeEvaluation_Report":                                     fields.ReportJudgeEvaluation(),
			"userProfile_Absent":                                         fields.UserProfileAbsentField(),
			"userProfile_Vacation":                                       fields.UserProfileVacationField(),
			"terminateEmployment":                                        fields.TerminateEmployment(),
			"absentType":                                                 fields.AbsentTypeField(),
			"userProfile_Resolution":                                     fields.UserProfileResolutionField(),
			"revisions_Overview":                                         fields.RevisionsOverviewField(),
			"revisions_Details":                                          fields.RevisionDetailsField(),
			"revision_plans_Overview":                                    fields.RevisionPlansOverview(),
			"revision_plans_Details":                                     fields.RevisionPlansDetails(),
			"revision_Overview":                                          fields.RevisionOverview(),
			"revision_Details":                                           fields.RevisionDetails(),
			"revision_tips_Overview":                                     fields.RevisionTipsOverview(),
			"revision_tips_Details":                                      fields.RevisionTipsDetails(),
			"judges_Overview":                                            fields.JudgesOverviewField(),
			"judgeResolutions_Overview":                                  fields.JudgeResolutionsOverviewField(),
			"judgeResolutions_Active":                                    fields.JudgeResolutionsActiveField(),
			"checkJudgeAndPresidentIsAvailable":                          fields.CheckJudgeAndPresidentIsAvailableField(),
			"organizationUintCalculateEmployeeStats":                     fields.OrganizationUintCalculateEmployeeStatsField(),
			"vacation_Report":                                            fields.ReportVacations(),
			"publicProcurementPlans_Overview":                            fields.PublicProcurementPlansOverviewField(),
			"publicProcurementPlan_Details":                              fields.PublicProcurementPlanDetailsField(),
			"publicProcurementPlanItem_Details":                          fields.PublicProcurementPlanItemDetailsField(),
			"publicProcurementPlanItem_PDF":                              fields.PublicProcurementPlanItemPDFField(),
			"publicProcurementPlan_PDF":                                  fields.PublicProcurementPlanPDFField(),
			"publicProcurementPlanItem_Limits":                           fields.PublicProcurementPlanItemLimitsField(),
			"publicProcurementOrganizationUnitArticles_Overview":         fields.PublicProcurementOrganizationUnitArticlesOverviewField(),
			"publicProcurementOrganizationUnitArticles_Details":          fields.PublicProcurementOrganizationUnitArticlesDetailsField(),
			"publicProcurementContracts_Overview":                        fields.PublicProcurementContractsOverviewField(),
			"publicProcurementContractArticles_Overview":                 fields.PublicProcurementContractArticlesOverviewField(),
			"publicProcurementContractArticlesOrganizationUnit_Overview": fields.PublicProcurementContractOrganizationUnitArticlesOverviewField(),
			"suppliers_Overview":                                         fields.SuppliersOverviewField(),
			"basicInventory_Overview":                                    fields.BasicInventoryOverviewField(),
			"ReportValueClassInventory_PDF":                              fields.ReportValueClassInventoryField(),
			"reportInventoryList_PDF":                                    fields.ReportInventoryList(),
			"basicInventory_Details":                                     fields.BasicInventoryDetailsField(),
			"basicInventoryRealEstates_Overview":                         fields.BasicInventoryRealEstatesOverviewField(),
			"officesOfOrganizationUnits_Overview":                        fields.OfficesOfOrganizationUnitOverviewField(),
			"basicInventoryDispatch_Overview":                            fields.BasicInventoryDispatchOverviewField(),
			"orderList_Overview":                                         fields.OrderListOverviewField(),
			"orderProcurementAvailableList_Overview":                     fields.OrderProcurementAvailableField(),
			"stock_Overview":                                             fields.StockOverviewFiled(),
			"movement_Overview":                                          fields.MovementOverviewField(),
			"movement_Details":                                           fields.MovementDetailsField(),
			"movementArticles_Overview":                                  fields.MovementArticlesField(),
			"recipientUsers_Overview":                                    fields.RecipientUsersField(),
			"overallSpending_Report":                                     fields.OverallSpendingField(),
			"account_Overview":                                           fields.AccountOverviewField(),
			"accountBudgetActivity_Overview":                             fields.AccountBudgetActivityOverviewField(),
			"budget_Overview":                                            fields.BudgetOverviewField(),
			"filledFinancialBudget_Overview":                             fields.FilledFinancialBudgetOverview(),
			"financialBudget_Details":                                    fields.FinancialBudgetDetails(),
			"programs_Overview":                                          fields.ProgramOverviewField(),
			"nonFinancialBudget_Overview":                                fields.NonFinancialBudgetOverviewType(),
			"activities_Overview":                                        fields.ActivitiesOverviewField(),
			"invoice_Overview":                                           fields.InvoiceOverviewField(),
			"fine_Overview":                                              fields.FineOverviewField(),
			"finePayment_Overview":                                       fields.FinePaymentOverviewField(),
		},
	})
	schemaConfig := graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
