package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"net/http"

	"github.com/shopspring/decimal"
)

type MicroserviceRepositoryInterface interface {
	AddOnStock(stock []structs.StockArticle, article structs.OrderProcurementArticleItem, organizationUnitID int) error
	AuthenticateUser(r *http.Request) (*structs.UserAccounts, error)
	CheckInsertInventoryData(input []structs.BasicInventoryInsertItem) ([]structs.BasicInventoryInsertValidator, error)
	CreateAbsent(ctx context.Context, absent *structs.Absent) (*structs.Absent, error)
	CreateAbsentType(absent *structs.AbsentType) (*structs.AbsentType, error)
	CreateAccountItemList(ctx context.Context, accountItemList []structs.AccountItem) ([]*structs.AccountItem, error)
	CreateAssessments(ctx context.Context, data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error)
	CreateDispatchItem(ctx context.Context, item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error)
	CreateDispatchItemItem(item *structs.BasicInventoryDispatchItemsItem) (*structs.BasicInventoryDispatchItemsItem, error)
	CreateDropdownSettings(data *structs.SettingsDropdown) (*structs.SettingsDropdown, error)
	CreateEmployeeContract(ctx context.Context, contract *structs.Contracts) (*structs.Contracts, error)
	CreateEmployeeEducation(education *structs.Education) (*structs.Education, error)
	CreateEmployeeEvaluation(ctx context.Context, evaluation *structs.Evaluation) (*structs.Evaluation, error)
	CreateEmployeeFamilyMember(familyMember *structs.Family) (*structs.Family, error)
	CreateEmployeeForeigner(foreigner *structs.Foreigners) (*structs.Foreigners, error)
	CreateEmployeeSalaryParams(ctx context.Context, salaries *structs.SalaryParams) (*structs.SalaryParams, error)
	CreateEmployeesInOrganizationUnits(data *structs.EmployeesInOrganizationUnits) (*structs.EmployeesInOrganizationUnits, error)
	CreateExperience(contract *structs.Experience) (*structs.Experience, error)
	CreateInventoryItem(ctx context.Context, item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error)
	CreateJobPositions(ctx context.Context, data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error)
	CreateJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error)
	CreateJobTender(ctx context.Context, jobTender *structs.JobTenders) (*structs.JobTenders, error)
	CreateJobTenderApplication(ctx context.Context, jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error)
	CreateJobTenderType(jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error)
	CreateJudgeNorm(ctx context.Context, norm *structs.JudgeNorms) (*structs.JudgeNorms, error)
	CreateJudgeResolutionItems(item *structs.JudgeResolutionItems) (*structs.JudgeResolutionItems, error)
	CreateJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitItem) (*dto.JudgeResolutionsOrganizationUnitItem, error)
	CreateJudgeResolutions(ctx context.Context, resolution *structs.JudgeResolutions) (*structs.JudgeResolutions, error)
	CreateMovementArticle(input dto.MovementArticle) (*dto.MovementArticle, error)
	CreateMovements(ctx context.Context, input structs.OrderAssetMovementItem) (*structs.Movement, error)
	CreateNotification(notification *structs.Notifications) (*structs.Notifications, error)
	CreateOrderListItem(ctx context.Context, orderListItem *structs.OrderListItem) (*structs.OrderListItem, error)
	CreateOrderListProcurementArticles(orderListID int, data structs.OrderListInsertItem) error
	CreateOrderProcurementArticle(orderProcurementArticleItem *structs.OrderProcurementArticleItem) (*structs.OrderProcurementArticleItem, error)
	CreateOrganizationUnits(ctx context.Context, data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error)
	CreateProcurementArticle(article *structs.PublicProcurementArticle) (*structs.PublicProcurementArticle, error)
	CreateProcurementContract(ctx context.Context, resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error)
	CreateProcurementContractArticle(article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error)
	CreateProcurementContractArticleOverage(articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error)
	CreateProcurementItem(ctx context.Context, item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error)
	CreateProcurementOUArticle(article *structs.PublicProcurementOrganizationUnitArticle) (*structs.PublicProcurementOrganizationUnitArticle, error)
	CreateProcurementOULimit(limit *structs.PublicProcurementLimit) (*structs.PublicProcurementLimit, error)
	CreateProcurementPlan(ctx context.Context, resolution *structs.PublicProcurementPlan) (*structs.PublicProcurementPlan, error)
	CreateResolution(ctx context.Context, resolution *structs.Resolution) (*structs.Resolution, error)
	CreateRevision(ctx context.Context, revision *structs.Revision) (*structs.Revision, error)
	CreateRevisionOrgUnit(plan *dto.RevisionOrgUnit) error
	CreateRevisionPlan(ctx context.Context, plan *dto.RevisionPlanItem) (*dto.RevisionPlanItem, error)
	CreateRevisionRevisor(plan *dto.RevisionRevisor) error
	CreateRevisionTips(ctx context.Context, plan *structs.RevisionTips) (*structs.RevisionTips, error)
	CreateRevisions(ctx context.Context, plan *structs.Revisions) (*structs.Revisions, error)
	CreateRole(ctx context.Context, data structs.Roles) (*structs.Roles, error)
	CreateStock(input dto.MovementArticle) error
	CreateSupplier(supplier *structs.Suppliers) (*structs.Suppliers, error)
	CreateSystematization(ctx context.Context, data *structs.Systematization) (*structs.Systematization, error)
	CreateUserAccount(ctx context.Context, user structs.UserAccounts) (*structs.UserAccounts, error)
	CreateUserProfile(ctx context.Context, user structs.UserProfiles) (*structs.UserProfiles, error)
	DeactivateUserAccount(ctx context.Context, userID int) (*structs.UserAccounts, error)
	DeleteAbsent(ctx context.Context, id int) error
	DeleteAbsentType(id int) error
	DeleteAccount(ctx context.Context, id int) error
	DeleteAssessment(ctx context.Context, id int) error
	DeleteDropdownSettings(id int) error
	DeleteEmployeeContract(ctx context.Context, id int) error
	DeleteEmployeeEducation(id int) error
	DeleteEmployeeFamilyMember(id int) error
	DeleteEmployeeInOrganizationUnit(jobPositionInOrganizationUnitID int) error
	DeleteEmployeeInOrganizationUnitByID(jobPositionInOrganizationUnitID int) error
	DeleteEvaluation(ctx context.Context, id int) error
	DeleteExperience(id int) error
	DeleteFile(id int) error
	DeleteForeigner(id int) error
	DeleteInventoryDispatch(ctx context.Context, id int) error
	DeleteJJudgeResolutionOrganizationUnit(id int) error
	DeleteJobPositions(ctx context.Context, id int) error
	DeleteJobPositionsInOrganizationUnits(id int) error
	DeleteJobTender(ctx context.Context, id int) error
	DeleteJobTenderApplication(id int) error
	DeleteJobTenderType(id int) error
	DeleteJudgeNorm(ctx context.Context, id int) error
	DeleteJudgeResolution(ctx context.Context, id int) error
	DeleteMovement(ctx context.Context, id int) error
	DeleteNotification(notificationID int) error
	DeleteOrderList(ctx context.Context, id int) error
	DeleteOrderProcurementArticle(id int) error
	DeleteOrganizationUnits(ctx context.Context, id int) error
	DeleteProcurementArticle(id int) error
	DeleteProcurementContract(ctx context.Context, id int) error
	DeleteProcurementContractArticleOverage(id int) error
	DeleteProcurementItem(ctx context.Context, id int) error
	DeleteProcurementPlan(ctx context.Context, id int) error
	DeleteResolution(ctx context.Context, id int) error
	DeleteRevision(ctx context.Context, id int) error
	DeleteRevisionOrgUnit(id int) error
	DeleteRevisionPlan(ctx context.Context, id int) error
	DeleteRevisionRevisor(id int) error
	DeleteRevisionTips(ctx context.Context, id int) error
	DeleteRevisions(ctx context.Context, id int) error
	DeleteSalaryParams(ctx context.Context, id int) error
	DeleteSupplier(id int) error
	DeleteSystematization(ctx context.Context, id int) error
	DeleteUserAccount(ctx context.Context, id int) error
	DeleteUserProfile(ctx context.Context, id int) error
	FetchNotifications(userID int) ([]*structs.Notifications, error)
	ForgotPassword(email string) error
	GetAbsentByID(absentID int) (*structs.Absent, error)
	GetAbsentTypeByID(absentTypeID int) (*structs.AbsentType, error)
	GetAbsentTypes() (*dto.GetAbsentTypeListResponseMS, error)
	GetAccountItemByID(id int) (*structs.AccountItem, error)
	GetAccountItems(filters *dto.GetAccountsFilter) (*dto.GetAccountItemListResponseMS, error)
	GetAllInventoryDispatches(filter dto.InventoryDispatchFilter) (*dto.GetAllBasicInventoryDispatches, error)
	GetAllInventoryItem(filter dto.InventoryItemFilter) (*dto.GetAllBasicInventoryItem, error)
	GetAllInventoryItemInOrgUnits(id int) ([]dto.GetAllItemsInOrgUnits, error)
	GetAllInventoryItemForReport(filter dto.ItemReportFilterDTO) ([]dto.ItemReportResponse, error)
	GetDispatchItemByID(id int) (*structs.BasicInventoryDispatchItem, error)
	GetInventoryItemsByDispatch(dispatchID int) ([]*structs.BasicInventoryInsertItem, error)
	GetDispatchItemByInventoryID(id int) ([]*structs.BasicInventoryDispatchItemsItem, error)
	GetDropdownSettingByID(id int) (*structs.SettingsDropdown, error)
	GetDropdownSettings(input *dto.GetSettingsInput) (*dto.GetDropdownTypesResponseMS, error)
	GetEmployeeAbsents(userProfileID int, input *dto.EmployeeAbsentsInput) ([]*structs.Absent, error)
	GetEmployeeContracts(employeeID int, input *dto.GetEmployeeContracts) ([]*structs.Contracts, error)
	GetEmployeeEducations(input dto.EducationInput) ([]structs.Education, error)
	GetEmployeeEvaluations(userProfileID int) ([]*structs.Evaluation, error)
	GetEmployeeExperiences(employeeID int) ([]*structs.Experience, error)
	GetEmployeeFamilyMembers(employeeID int) ([]*structs.Family, error)
	GetEmployeeForeigners(userProfileID int) ([]*structs.Foreigners, error)
	GetEmployeeResolution(id int) (*structs.Resolution, error)
	GetEmployeeResolutions(employeeID int, input *dto.EmployeeResolutionListInput) ([]*structs.Resolution, error)
	GetEmployeeSalaryParams(userProfileID int) ([]*structs.SalaryParams, error)
	GetEmployeesInOrganizationUnitList(input *dto.GetEmployeesInOrganizationUnitInput) ([]*structs.EmployeesInOrganizationUnits, error)
	GetEmployeesInOrganizationUnitsByProfileID(profileID int) (*structs.EmployeesInOrganizationUnits, error)
	GetEvaluation(evaulationID int) (*structs.Evaluation, error)
	GetEvaluationList(input *dto.GetEvaluationListInputMS) ([]*structs.Evaluation, error)
	GetFileByID(id int) (*structs.File, error)
	GetInventoryItem(id int) (*structs.BasicInventoryInsertItem, error)
	GetInventoryRealEstate(id int) (*structs.BasicInventoryRealEstatesItem, error)
	GetInventoryRealEstatesList(input *dto.GetInventoryRealEstateListInputMS) (*dto.GetInventoryRealEstateListResponseMS, error)
	GetJobPositionByID(id int) (*structs.JobPositions, error)
	GetJobPositions(input *dto.GetJobPositionsInput) (*dto.GetJobPositionsResponseMS, error)
	GetJobPositionsInOrganizationUnits(input *dto.GetJobPositionInOrganizationUnitsInput) (*dto.GetJobPositionsInOrganizationUnitsResponseMS, error)
	GetJobPositionsInOrganizationUnitsByID(id int) (*structs.JobPositionsInOrganizationUnits, error)
	GetJobTender(id int) (*structs.JobTenders, error)
	GetJobTenderList() ([]*structs.JobTenders, error)
	GetJudgeNormListByEmployee(userProfileID int, input dto.GetUserNormFulfilmentListInput) ([]structs.JudgeNorms, error)
	GetJudgeResolution(id int) (*structs.JudgeResolutions, error)
	GetJudgeResolutionItemsList(input *dto.GetJudgeResolutionItemListInputMS) ([]*structs.JudgeResolutionItems, error)
	GetJudgeResolutionList(input *dto.GetJudgeResolutionListInputMS) (*dto.GetJudgeResolutionListResponseMS, error)
	GetJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitInput) ([]dto.JudgeResolutionsOrganizationUnitItem, int, error)
	GetLoggedInUser(token string) (*structs.UserAccounts, error)
	GetMovementArticleList(filter dto.OveralSpendingFilter) ([]dto.ArticleReport, error)
	GetMovementArticles(id int) ([]dto.MovementArticle, error)
	GetMovementByID(id int) (*structs.Movement, error)
	GetMovements(input *dto.MovementFilter) ([]structs.Movement, *int, error)
	GetMyInventoryAssessments(id int) ([]structs.BasicInventoryAssessmentsTypesItem, error)
	GetMyInventoryDispatchesItems(filter *dto.DispatchInventoryItemFilter) ([]*structs.BasicInventoryDispatchItemsItem, error)
	GetMyInventoryRealEstate(id int) (*structs.BasicInventoryRealEstatesItem, error)
	GetNotification(id int) (*structs.Notifications, error)
	GetOfficeDropdownSettings(input *dto.GetOfficesOfOrganizationInput) (*dto.GetDropdownTypesResponseMS, error)
	GetOrderListByID(id int) (*structs.OrderListItem, error)
	GetOrderLists(input *dto.GetOrderListInput) (*dto.GetOrderListsResponseMS, error)
	GetOrderProcurementArticleByID(id int) (*structs.OrderProcurementArticleItem, error)
	GetOrderProcurementArticles(input *dto.GetOrderProcurementArticleInput) (*dto.GetOrderProcurementArticlesResponseMS, error)
	GetOrganizationUnitArticleByID(id int) (*dto.GetPublicProcurementOrganizationUnitArticle, error)
	GetOrganizationUnitArticlesList(input dto.GetProcurementOrganizationUnitArticleListInputDTO) ([]dto.GetPublicProcurementOrganizationUnitArticle, error)
	GetOrganizationUnitByID(id int) (*structs.OrganizationUnits, error)
	GetOrganizationUnitIDByUserProfile(id int) (*int, error)
	GetOrganizationUnits(input *dto.GetOrganizationUnitsInput) (*dto.GetOrganizationUnitsResponseMS, error)
	GetPermissionList(roleID int) ([]structs.Permissions, error)
	GetProcurementArticle(id int) (*structs.PublicProcurementArticle, error)
	GetProcurementArticlesList(input *dto.GetProcurementArticleListInputMS) ([]*structs.PublicProcurementArticle, error)
	GetProcurementContract(id int) (*structs.PublicProcurementContract, error)
	GetProcurementContractArticleOverageList(input *dto.GetProcurementContractArticleOverageInput) ([]*structs.PublicProcurementContractArticleOverage, error)
	GetProcurementContractArticlesList(input *dto.GetProcurementContractArticlesInput) (*dto.GetProcurementContractArticlesListResponseMS, error)
	GetProcurementContractsList(input *dto.GetProcurementContractsInput) (*dto.GetProcurementContractListResponseMS, error)
	GetProcurementItem(id int) (*structs.PublicProcurementItem, error)
	GetProcurementItemList(input *dto.GetProcurementItemListInputMS) ([]*structs.PublicProcurementItem, error)
	GetProcurementOUArticleList(input *dto.GetProcurementOrganizationUnitArticleListInputDTO) ([]*structs.PublicProcurementOrganizationUnitArticle, error)
	GetProcurementOULimitList(input *dto.GetProcurementOULimitListInputMS) ([]*structs.PublicProcurementLimit, error)
	GetProcurementPlan(id int) (*structs.PublicProcurementPlan, error)
	GetProcurementPlanList(input *dto.GetProcurementPlansInput) ([]*structs.PublicProcurementPlan, error)
	GetRevisionByID(id int) (*structs.Revision, error)
	GetRevisionList(input *dto.GetRevisionsInput) (*dto.GetRevisionListResponseMS, error)
	GetRevisionOrgUnitList(input *dto.RevisionOrgUnitFilter) ([]*dto.RevisionOrgUnit, error)
	GetRevisionPlanByID(id int) (*dto.RevisionPlanItem, error)
	GetRevisionPlanList(input *dto.GetPlansInput) (*dto.GetRevisionPlanResponseMS, error)
	GetRevisionRevisorList(input *dto.RevisionRevisorFilter) ([]*dto.RevisionRevisor, error)
	GetRevisionTipByID(id int) (*structs.RevisionTips, error)
	GetRevisionTipsList(input *dto.GetRevisionTipFilter) (*dto.GetRevisionTipsResponseMS, error)
	GetRevisionsByID(id int) (*structs.Revisions, error)
	GetRevisionsList(input *dto.GetRevisionFilter) (*dto.GetRevisionsResponseMS, error)
	GetRevisors() ([]*structs.Revisor, error)
	GetRole(id structs.UserRole) (*structs.Roles, error)
	GetRoleList() ([]structs.Roles, error)
	GetStock(input *dto.StockFilter) ([]structs.StockArticle, *int, error)
	GetStockByID(id int) (*structs.StockArticle, error)
	GetSupplier(id int) (*structs.Suppliers, error)
	GetSupplierList(input *dto.GetSupplierInputMS) (*dto.GetSupplierListResponseMS, error)
	GetSystematizationByID(id int) (*structs.Systematization, error)
	GetSystematizations(input *dto.GetSystematizationsInput) (*dto.GetSystematizationsResponseMS, error)
	GetTenderApplication(id int) (*structs.JobTenderApplications, error)
	GetTenderApplicationList(input *dto.GetJobTenderApplicationsInput) (*dto.GetJobTenderApplicationListResponseMS, error)
	GetTenderType(id int) (*structs.JobTenderTypes, error)
	GetTenderTypeList(input *dto.GetJobTenderTypeInputMS) ([]*structs.JobTenderTypes, error)
	GetUserAccountByID(id int) (*structs.UserAccounts, error)
	GetUserAccounts(input *dto.GetUserAccountListInput) (*dto.GetUserAccountListResponseMS, error)
	GetUserProfileByID(id int) (*structs.UserProfiles, error)
	GetUserProfileByUserAccountID(accountID int) (*structs.UserProfiles, error)
	GetUserProfiles(input *dto.GetUserProfilesInput) ([]*structs.UserProfiles, error)
	LoginUser(email string, password string) (*dto.LoginResponseMS, []*http.Cookie, error)
	Logout(token string) error
	MarkNotificationRead(notificationID int) error
	RefreshToken(cookie *http.Cookie) (*dto.RefreshTokenResponse, []*http.Cookie, error)
	ResetPassword(input *dto.ResetPassword) error
	SendOrderListToFinance(ctx context.Context, id int) error
	SyncPermissions(roleID int, input []*structs.RolePermission) ([]structs.RolePermission, error)
	UpdateAbsent(ctx context.Context, id int, absent *structs.Absent) (*structs.Absent, error)
	UpdateAbsentType(id int, absent *structs.AbsentType) (*structs.AbsentType, error)
	UpdateAccountItem(ctx context.Context, id int, accountItem *structs.AccountItem) (*structs.AccountItem, error)
	UpdateAssessments(ctx context.Context, id int, data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error)
	UpdateDispatchItem(ctx context.Context, id int, item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error)
	UpdateDropdownSettings(id int, data *structs.SettingsDropdown) (*structs.SettingsDropdown, error)
	UpdateEmployeeContract(ctx context.Context, id int, contract *structs.Contracts) (*structs.Contracts, error)
	UpdateEmployeeEducation(id int, education *structs.Education) (*structs.Education, error)
	UpdateEmployeeEvaluation(ctx context.Context, id int, evaluation *structs.Evaluation) (*structs.Evaluation, error)
	UpdateEmployeeFamilyMember(id int, education *structs.Family) (*structs.Family, error)
	UpdateEmployeeForeigner(id int, foreigner *structs.Foreigners) (*structs.Foreigners, error)
	UpdateEmployeeSalaryParams(ctx context.Context, id int, salaries *structs.SalaryParams) (*structs.SalaryParams, error)
	UpdateExperience(id int, contract *structs.Experience) (*structs.Experience, error)
	UpdateInventoryItem(ctx context.Context, id int, item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error)
	UpdateJobPositions(ctx context.Context, id int, data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error)
	UpdateJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error)
	UpdateJobTender(ctx context.Context, id int, jobTender *structs.JobTenders) (*structs.JobTenders, error)
	UpdateJobTenderApplication(ctx context.Context, id int, jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error)
	UpdateJobTenderType(id int, jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error)
	UpdateJudgeNorm(ctx context.Context, id int, norm *structs.JudgeNorms) (*structs.JudgeNorms, error)
	UpdateJudgeResolutionItems(id int, item *structs.JudgeResolutionItems) (*structs.JudgeResolutionItems, error)
	UpdateJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitItem) (*dto.JudgeResolutionsOrganizationUnitItem, error)
	UpdateJudgeResolutions(ctx context.Context, id int, resolution *structs.JudgeResolutions) (*structs.JudgeResolutions, error)
	UpdateMovements(ctx context.Context, input structs.OrderAssetMovementItem) (*structs.Movement, error)
	UpdateNotification(notificationID int, notification *structs.Notifications) error
	UpdateOrderListItem(ctx context.Context, id int, orderListItem *structs.OrderListItem) (*structs.OrderListItem, error)
	UpdateOrderProcurementArticle(item *structs.OrderProcurementArticleItem) (*structs.OrderProcurementArticleItem, error)
	UpdateOrganizationUnits(ctx context.Context, id int, data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error)
	UpdateProcurementArticle(id int, article *structs.PublicProcurementArticle) (*structs.PublicProcurementArticle, error)
	UpdateProcurementContract(ctx context.Context, id int, resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error)
	UpdateProcurementContractArticle(id int, article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error)
	UpdateProcurementContractArticleOverage(id int, articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error)
	UpdateProcurementItem(ctx context.Context, id int, item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error)
	UpdateProcurementOUArticle(id int, article *structs.PublicProcurementOrganizationUnitArticle) (*structs.PublicProcurementOrganizationUnitArticle, error)
	UpdateProcurementOULimit(id int, limit *structs.PublicProcurementLimit) (*structs.PublicProcurementLimit, error)
	UpdateProcurementPlan(ctx context.Context, id int, resolution *structs.PublicProcurementPlan) (*structs.PublicProcurementPlan, error)
	UpdateResolution(ctx context.Context, id int, resolution *structs.Resolution) (*structs.Resolution, error)
	UpdateRevision(ctx context.Context, id int, revision *structs.Revision) (*structs.Revision, error)
	UpdateRevisionPlan(ctx context.Context, id int, plan *dto.RevisionPlanItem) (*dto.RevisionPlanItem, error)
	UpdateRevisionTips(ctx context.Context, id int, plan *structs.RevisionTips) (*structs.RevisionTips, error)
	UpdateRevisions(ctx context.Context, id int, plan *structs.Revisions) (*structs.Revisions, error)
	UpdateRole(ctx context.Context, id int, data structs.Roles) (*structs.Roles, error)
	UpdateStock(input structs.StockArticle) error
	UpdateSupplier(id int, supplier *structs.Suppliers) (*structs.Suppliers, error)
	UpdateSystematization(ctx context.Context, id int, data *structs.Systematization) (*structs.Systematization, error)
	UpdateUserAccount(ctx context.Context, userID int, user structs.UserAccounts) (*structs.UserAccounts, error)
	UpdateUserProfile(ctx context.Context, userID int, user structs.UserProfiles) (*structs.UserProfiles, error)
	ValidateMail(input *dto.ResetPasswordVerify) (*dto.ResetPasswordVerifyResponseMS, error)
	ValidatePin(pin string, headers map[string]string) error

	GetLatestVersionOfAccounts() (int, error)

	CreateBudget(ctx context.Context, budget *structs.Budget) (*structs.Budget, error)
	UpdateBudget(ctx context.Context, budget *structs.Budget) (*structs.Budget, error)
	GetBudget(id int) (*structs.Budget, error)
	GetBudgetList(input *dto.GetBudgetListInputMS) ([]structs.Budget, error)
	DeleteBudget(ctx context.Context, id int) error

	GetCurrentBudgetByOrganizationUnit(organizationUnitID int) ([]structs.CurrentBudget, error)

	CreateBudgetRequest(ctx context.Context, budget *structs.BudgetRequest) (*structs.BudgetRequest, error)
	UpdateBudgetRequest(ctx context.Context, budget *structs.BudgetRequest) (*structs.BudgetRequest, error)
	GetBudgetRequest(id int) (*structs.BudgetRequest, error)
	GetBudgetRequestList(input *dto.GetBudgetRequestListInputMS) ([]structs.BudgetRequest, error)
	GetOneBudgetRequest(input *dto.GetBudgetRequestListInputMS) (*structs.BudgetRequest, error)

	UpdateFinancialBudget(ctx context.Context, financialBudget *structs.FinancialBudget) (*structs.FinancialBudget, error)
	CreateFinancialBudget(ctx context.Context, financialBudget *structs.FinancialBudget) (*structs.FinancialBudget, error)
	GetFinancialBudgetByBudgetID(budgetID int) (*structs.FinancialBudget, error)
	GetFinancialBudgetByID(id int) (*structs.FinancialBudget, error)
	GetFilledFinancialBudgetList(input *dto.FilledFinancialBudgetInputMS) ([]structs.FilledFinanceBudget, error)

	GetBudgetLimits(budgetID int) ([]structs.FinancialBudgetLimit, error)
	GetBudgetUnitLimit(budgetID, unitID int) (int, error)
	CreateBudgetLimit(budgetLimit *structs.FinancialBudgetLimit) (*structs.FinancialBudgetLimit, error)
	UpdateBudgetLimit(budgetLimit *structs.FinancialBudgetLimit) (*structs.FinancialBudgetLimit, error)
	DeleteBudgetLimit(id int) error

	UpdateNonFinancialBudget(ctx context.Context, id int, program *structs.NonFinancialBudgetItem) (*structs.NonFinancialBudgetItem, error)
	CreateNonFinancialBudget(ctx context.Context, budget *structs.NonFinancialBudgetItem) (*structs.NonFinancialBudgetItem, error)
	FillFinancialBudget(ctx context.Context, budget *structs.FilledFinanceBudget) (*structs.FilledFinanceBudget, error)
	UpdateFilledFinancialBudget(ctx context.Context, id int, budget *structs.FilledFinanceBudget) (*structs.FilledFinanceBudget, error)
	FillActualFinancialBudget(ctx context.Context, id int, actual decimal.Decimal, requestID int) (*structs.FilledFinanceBudget, error)
	GetFinancialFilledSummary(budgetID int, reqType structs.RequestType) ([]structs.FilledFinanceBudget, error)
	DeleteFilledFinancialBudgetData(ctx context.Context, id int) error
	DeleteNonFinancialBudget(ctx context.Context, id int) error
	GetNonFinancialBudget(id int) (*structs.NonFinancialBudgetItem, error)
	GetNonFinancialBudgetList(input *dto.GetNonFinancialBudgetListInputMS) ([]structs.NonFinancialBudgetItem, error)
	GetNonFinancialBudgetByRequestID(requestID int) (structs.NonFinancialBudgetItem, error)

	UpdateNonFinancialGoal(ctx context.Context, id int, activity *structs.NonFinancialGoalItem) (*structs.NonFinancialGoalItem, error)
	CreateNonFinancialGoal(ctx context.Context, goal *structs.NonFinancialGoalItem) (*structs.NonFinancialGoalItem, error)
	DeleteNonFinancialGoal(ctx context.Context, id int) error
	GetNonFinancialGoal(id int) (*structs.NonFinancialGoalItem, error)
	GetNonFinancialGoalList(input *dto.GetNonFinancialGoalListInputMS) ([]structs.NonFinancialGoalItem, error)

	UpdateNonFinancialGoalIndicator(ctx context.Context, id int, goalIndicator *structs.NonFinancialGoalIndicatorItem) (*structs.NonFinancialGoalIndicatorItem, error)
	CreateNonFinancialGoalIndicator(ctx context.Context, goal *structs.NonFinancialGoalIndicatorItem) (*structs.NonFinancialGoalIndicatorItem, error)
	DeleteNonFinancialGoalIndicator(ctx context.Context, id int) error
	GetNonFinancialGoalIndicator(id int) (*structs.NonFinancialGoalIndicatorItem, error)
	GetNonFinancialGoalIndicatorList(input *dto.GetNonFinancialGoalIndicatorListInputMS) ([]structs.NonFinancialGoalIndicatorItem, error)

	UpdateProgram(ctx context.Context, id int, program *structs.ProgramItem) (*structs.ProgramItem, error)
	CreateProgram(ctx context.Context, program *structs.ProgramItem) (*structs.ProgramItem, error)
	DeleteProgram(ctx context.Context, id int) error
	GetProgram(id int) (*structs.ProgramItem, error)
	GetProgramList(input *dto.GetFinanceProgramListInputMS) ([]structs.ProgramItem, error)

	UpdateActivity(ctx context.Context, id int, activity *structs.ActivitiesItem) (*structs.ActivitiesItem, error)
	CreateActivity(ctx context.Context, activity *structs.ActivitiesItem) (*structs.ActivitiesItem, error)
	DeleteActivity(ctx context.Context, id int) error
	GetActivity(id int) (*structs.ActivitiesItem, error)
	GetActivityList(input *dto.GetFinanceActivityListInputMS) ([]structs.ActivitiesItem, error)
	GetActivityByUnit(organizationUnitID int) (*structs.ActivitiesItem, error)

	GetSpendingDynamic(budgetID, unitID int, input *dto.GetSpendingDynamicHistoryInput) ([]dto.SpendingDynamicDTO, error)
	GetSpendingDynamicHistory(budgetID, unitID int) ([]dto.SpendingDynamicHistoryDTO, error)
	GetSpendingDynamicActual(budgetID, unitID, accountID int) (decimal.NullDecimal, error)
	CreateSpendingDynamic(ctx context.Context, budgetID, unitID int, spendingDynamic []structs.SpendingDynamicInsert) ([]dto.SpendingDynamicDTO, error)

	CreateSpendingRelease(ctx context.Context, spendingReleaseList []structs.SpendingReleaseInsert, budgetID, unitID int) ([]structs.SpendingRelease, error)
	GetSpendingReleaseOverview(ctx context.Context, input *dto.SpendingReleaseOverviewFilterDTO) ([]dto.SpendingReleaseOverviewItem, error)
	GetSpendingReleaseList(ctx context.Context, input *dto.GetSpendingReleaseListInput) ([]structs.SpendingRelease, error)
	DeleteSpendingRelease(ctx context.Context, input *dto.DeleteSpendingReleaseInput) error

	CreateCurrentBudget(ctx context.Context, currentBudget *structs.CurrentBudget) (*structs.CurrentBudget, error)
	GetCurrentBudgetUnitList(ctx context.Context) ([]int, error)

	CreateInvoice(ctx context.Context, item *structs.Invoice) (*structs.Invoice, error)
	UpdateInvoice(ctx context.Context, item *structs.Invoice) (*structs.Invoice, error)
	GetInvoice(id int) (*structs.Invoice, error)
	GetInvoiceList(input *dto.GetInvoiceListInputMS) ([]structs.Invoice, int, error)
	GetInvoiceArticleList(id int) ([]structs.InvoiceArticles, error)
	GetAdditionalExpenses(input *dto.AdditionalExpensesListInputMS) ([]structs.AdditionalExpenses, int, error)
	DeleteInvoice(ctx context.Context, id int) error
	DeleteInvoiceArticle(id int) error
	CreateInvoiceArticle(article *structs.InvoiceArticles) (*structs.InvoiceArticles, error)
	UpdateInvoiceArticle(item *structs.InvoiceArticles) (*structs.InvoiceArticles, error)

	GetTaxAuthorityCodebookByID(id int) (*structs.TaxAuthorityCodebook, error)
	GetTaxAuthorityCodebooks(input dto.TaxAuthorityCodebookFilter) (*dto.GetTaxAuthorityCodebooksResponseMS, error)
	CreateTaxAuthorityCodebook(ctx context.Context, item *structs.TaxAuthorityCodebook) (*structs.TaxAuthorityCodebook, error)
	UpdateTaxAuthorityCodebook(ctx context.Context, id int, data *structs.TaxAuthorityCodebook) (*structs.TaxAuthorityCodebook, error)
	DeactivateTaxAuthorityCodebook(ctx context.Context, id int, active bool) error
	DeleteTaxAuthorityCodebook(ctx context.Context, id int) error

	CreateFee(ctx context.Context, item *structs.Fee) (*structs.Fee, error)
	GetFee(id int) (*structs.Fee, error)
	GetFeeList(input *dto.GetFeeListInputMS) ([]structs.Fee, int, error)
	DeleteFee(ctx context.Context, id int) error
	UpdateFee(ctx context.Context, item *structs.Fee) (*structs.Fee, error)

	CreateFeePayment(ctx context.Context, item *structs.FeePayment) (*structs.FeePayment, error)
	GetFeePayment(id int) (*structs.FeePayment, error)
	GetFeePaymentList(input *dto.GetFeePaymentListInputMS) ([]structs.FeePayment, int, error)
	DeleteFeePayment(ctx context.Context, id int) error
	UpdateFeePayment(ctx context.Context, item *structs.FeePayment) (*structs.FeePayment, error)

	CreateFine(ctx context.Context, item *structs.Fine) (*structs.Fine, error)
	GetFine(id int) (*structs.Fine, error)
	GetFineList(input *dto.GetFineListInputMS) ([]structs.Fine, int, error)
	DeleteFine(ctx context.Context, id int) error
	UpdateFine(ctx context.Context, item *structs.Fine) (*structs.Fine, error)

	CreateFinePayment(ctx context.Context, item *structs.FinePayment) (*structs.FinePayment, error)
	GetFinePayment(id int) (*structs.FinePayment, error)
	GetFinePaymentList(input *dto.GetFinePaymentListInputMS) ([]structs.FinePayment, int, error)
	DeleteFinePayment(ctx context.Context, id int) error
	UpdateFinePayment(ctx context.Context, item *structs.FinePayment) (*structs.FinePayment, error)

	CreateProcedureCost(ctx context.Context, item *structs.ProcedureCost) (*structs.ProcedureCost, error)
	GetProcedureCost(id int) (*structs.ProcedureCost, error)
	GetProcedureCostList(input *dto.GetProcedureCostListInputMS) ([]structs.ProcedureCost, int, error)
	DeleteProcedureCost(ctx context.Context, id int) error
	UpdateProcedureCost(ctx context.Context, item *structs.ProcedureCost) (*structs.ProcedureCost, error)

	CreateProcedureCostPayment(ctx context.Context, item *structs.ProcedureCostPayment) (*structs.ProcedureCostPayment, error)
	GetProcedureCostPayment(id int) (*structs.ProcedureCostPayment, error)
	GetProcedureCostPaymentList(input *dto.GetProcedureCostPaymentListInputMS) ([]structs.ProcedureCostPayment, int, error)
	DeleteProcedureCostPayment(ctx context.Context, id int) error
	UpdateProcedureCostPayment(ctx context.Context, item *structs.ProcedureCostPayment) (*structs.ProcedureCostPayment, error)

	CreateFlatRate(ctx context.Context, item *structs.FlatRate) (*structs.FlatRate, error)
	GetFlatRate(id int) (*structs.FlatRate, error)
	GetFlatRateList(input *dto.GetFlatRateListInputMS) ([]structs.FlatRate, int, error)
	DeleteFlatRate(ctx context.Context, id int) error
	UpdateFlatRate(ctx context.Context, item *structs.FlatRate) (*structs.FlatRate, error)

	CreateFlatRatePayment(ctx context.Context, item *structs.FlatRatePayment) (*structs.FlatRatePayment, error)
	GetFlatRatePayment(id int) (*structs.FlatRatePayment, error)
	GetFlatRatePaymentList(input *dto.GetFlatRatePaymentListInputMS) ([]structs.FlatRatePayment, int, error)
	DeleteFlatRatePayment(ctx context.Context, id int) error
	UpdateFlatRatePayment(ctx context.Context, item *structs.FlatRatePayment) (*structs.FlatRatePayment, error)

	CreatePropBenConf(ctx context.Context, item *structs.PropBenConf) (*structs.PropBenConf, error)
	GetPropBenConf(id int) (*structs.PropBenConf, error)
	GetPropBenConfList(input *dto.GetPropBenConfListInputMS) ([]structs.PropBenConf, int, error)
	DeletePropBenConf(ctx context.Context, id int) error
	UpdatePropBenConf(ctx context.Context, item *structs.PropBenConf) (*structs.PropBenConf, error)

	CreatePropBenConfPayment(ctx context.Context, item *structs.PropBenConfPayment) (*structs.PropBenConfPayment, error)
	GetPropBenConfPayment(id int) (*structs.PropBenConfPayment, error)
	GetPropBenConfPaymentList(input *dto.GetPropBenConfPaymentListInputMS) ([]structs.PropBenConfPayment, int, error)
	DeletePropBenConfPayment(ctx context.Context, id int) error
	UpdatePropBenConfPayment(ctx context.Context, item *structs.PropBenConfPayment) (*structs.PropBenConfPayment, error)

	CreateFixedDeposit(ctx context.Context, item *structs.FixedDeposit) (*structs.FixedDeposit, error)
	DeleteFixedDeposit(ctx context.Context, id int) error
	UpdateFixedDeposit(ctx context.Context, item *structs.FixedDeposit) (*structs.FixedDeposit, error)
	GetFixedDepositByID(id int) (*structs.FixedDeposit, error)
	GetFixedDepositList(input dto.FixedDepositFilter) ([]structs.FixedDeposit, int, error)

	CreateFixedDepositWill(ctx context.Context, item *structs.FixedDepositWill) (*structs.FixedDepositWill, error)
	DeleteFixedDepositWill(ctx context.Context, id int) error
	UpdateFixedDepositWill(ctx context.Context, item *structs.FixedDepositWill) (*structs.FixedDepositWill, error)
	GetFixedDepositWillByID(id int) (*structs.FixedDepositWill, error)
	GetFixedDepositWillList(input dto.FixedDepositWillFilter) ([]structs.FixedDepositWill, int, error)

	CreateFixedDepositItem(ctx context.Context, item *structs.FixedDepositItem) error
	UpdateFixedDepositItem(ctx context.Context, item *structs.FixedDepositItem) error
	DeleteFixedDepositItem(ctx context.Context, id int) error
	CreateFixedDepositDispatch(ctx context.Context, item *structs.FixedDepositDispatch) error
	UpdateFixedDepositDispatch(ctx context.Context, item *structs.FixedDepositDispatch) error
	DeleteFixedDepositDispatch(ctx context.Context, id int) error
	CreateFixedDepositJudge(item *structs.FixedDepositJudge) error
	UpdateFixedDepositJudge(item *structs.FixedDepositJudge) error
	DeleteFixedDepositJudge(id int) error
	CreateFixedDepositWillDispatch(ctx context.Context, item *structs.FixedDepositWillDispatch) error
	UpdateFixedDepositWillDispatch(ctx context.Context, item *structs.FixedDepositWillDispatch) error
	DeleteFixedDepositWillDispatch(ctx context.Context, id int) error

	CreateSalary(ctx context.Context, item *structs.Salary) (*structs.Salary, error)
	DeleteSalary(ctx context.Context, id int) error
	UpdateSalary(ctx context.Context, item *structs.Salary) (*structs.Salary, error)
	GetSalaryByID(id int) (*structs.Salary, error)
	GetSalaryList(input dto.SalaryFilter) ([]structs.Salary, int, error)

	CreateDepositPayment(ctx context.Context, item *structs.DepositPayment) (*structs.DepositPayment, error)
	DeleteDepositPayment(ctx context.Context, id int) error
	UpdateDepositPayment(ctx context.Context, item *structs.DepositPayment) (*structs.DepositPayment, error)
	GetDepositPaymentByID(id int) (*structs.DepositPayment, error)
	GetDepositPaymentList(input dto.DepositPaymentFilter) ([]structs.DepositPayment, int, error)
	GetInitialState(input dto.DepositInitialStateFilter) ([]structs.DepositPayment, error)
	GetDepositPaymentCaseNumber(caseNumber string, bankAccount string) (*structs.DepositPayment, error)
	GetCaseNumber(organizationUnitID int, bankAccount string) ([]structs.DepositPayment, error)

	CreateDepositPaymentOrder(ctx context.Context, item *structs.DepositPaymentOrder) (*structs.DepositPaymentOrder, error)
	DeleteDepositPaymentOrder(ctx context.Context, id int) error
	UpdateDepositPaymentOrder(ctx context.Context, item *structs.DepositPaymentOrder) (*structs.DepositPaymentOrder, error)
	GetDepositPaymentOrderByID(id int) (*structs.DepositPaymentOrder, error)
	GetDepositPaymentOrderList(input dto.DepositPaymentOrderFilter) ([]structs.DepositPaymentOrder, int, error)
	GetDepositPaymentAdditionalExpenses(input *dto.DepositPaymentAdditionalExpensesListInputMS) ([]structs.DepositPaymentAdditionalExpenses, int, error)
	PayDepositPaymentOrder(ctx context.Context, input structs.DepositPaymentOrder) error

	CreatePaymentOrder(ctx context.Context, item *structs.PaymentOrder) (*structs.PaymentOrder, error)
	DeletePaymentOrder(ctx context.Context, id int) error
	UpdatePaymentOrder(ctx context.Context, item *structs.PaymentOrder) (*structs.PaymentOrder, error)
	GetPaymentOrderByID(id int) (*structs.PaymentOrder, error)
	GetPaymentOrderList(input dto.PaymentOrderFilter) ([]structs.PaymentOrder, int, error)
	GetAllObligations(input dto.ObligationsFilter) ([]dto.Obligation, int, error)
	PayPaymentOrder(ctx context.Context, input structs.PaymentOrder) error
	CancelPaymentOrder(ctx context.Context, id int) error

	CreateEnforcedPayment(ctx context.Context, item *structs.EnforcedPayment) (*structs.EnforcedPayment, error)
	UpdateEnforcedPayment(ctx context.Context, item *structs.EnforcedPayment) (*structs.EnforcedPayment, error)
	GetEnforcedPaymentByID(id int) (*structs.EnforcedPayment, error)
	GetEnforcedPaymentList(input dto.EnforcedPaymentFilter) ([]structs.EnforcedPayment, int, error)
	ReturnEnforcedPayment(ctx context.Context, input structs.EnforcedPayment) error

	GetAllObligationsForAccounting(input dto.ObligationsFilter) ([]dto.ObligationForAccounting, int, error)
	GetAllPaymentOrdersForAccounting(input dto.ObligationsFilter) ([]dto.PaymentOrdersForAccounting, int, error)
	GetAllEnforcedPaymentsForAccounting(input dto.ObligationsFilter) ([]dto.PaymentOrdersForAccounting, int, error)
	GetAllReturnedEnforcedPaymentsForAccounting(input dto.ObligationsFilter) ([]dto.PaymentOrdersForAccounting, int, error)
	BuildAccountingOrderForObligations(data structs.AccountingOrderForObligationsData) (*dto.AccountingOrderForObligations, error)

	UpdateModelsOfAccounting(ctx context.Context, item *structs.ModelsOfAccounting) (*structs.ModelsOfAccounting, error)
	GetModelsOfAccountingByID(id int) (*structs.ModelsOfAccounting, error)
	GetModelsOfAccountingList(input dto.ModelsOfAccountingFilter) ([]structs.ModelsOfAccounting, int, error)

	CreateAccountingEntry(ctx context.Context, item *structs.AccountingEntry) (*structs.AccountingEntry, error)
	DeleteAccountingEntry(ctx context.Context, id int) error
	GetAccountingEntryByID(id int) (*structs.AccountingEntry, error)
	GetAccountingEntryList(input dto.AccountingEntryFilter) ([]structs.AccountingEntry, int, error)

	GetAnalyticalCard(input dto.AnalyticalCardFilter) ([]structs.AnalyticalCard, error)

	CreateInternalReallocation(ctx context.Context, item *structs.InternalReallocation) (*structs.InternalReallocation, error)
	DeleteInternalReallocation(ctx context.Context, id int) error
	GetInternalReallocationByID(id int) (*structs.InternalReallocation, error)
	GetInternalReallocationList(input dto.InternalReallocationFilter) ([]structs.InternalReallocation, int, error)

	CreateExternalReallocation(ctx context.Context, item *structs.ExternalReallocation) (*structs.ExternalReallocation, error)
	DeleteExternalReallocation(ctx context.Context, id int) error
	GetExternalReallocationByID(id int) (*structs.ExternalReallocation, error)
	GetExternalReallocationList(input dto.ExternalReallocationFilter) ([]structs.ExternalReallocation, int, error)
	AcceptOUExternalReallocation(ctx context.Context, item *structs.ExternalReallocation) (*structs.ExternalReallocation, error)
	RejectOUExternalReallocation(ctx context.Context, id int) error
	AcceptSSSExternalReallocation(ctx context.Context, id int) error
	RejectSSSExternalReallocation(ctx context.Context, id int) error

	GetLog(entity config.Module, id int) (*structs.Logs, error)
	GetLogs(filter dto.LogFilterDTO) ([]structs.Logs, int, error)

	CreateTemplate(ctx context.Context, item *structs.Template) error
	DeleteTemplate(ctx context.Context, id int) error
	UpdateTemplate(ctx context.Context, item *structs.Template) error
	UpdateTemplateItem(ctx context.Context, item *structs.Template) error
	GetTemplateByID(id int) (*structs.Template, error)
	GetTemplateList(input dto.TemplateFilter) ([]structs.Template, int, error)
}
