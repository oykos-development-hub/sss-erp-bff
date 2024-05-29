package repository

import (
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
	CreateAbsent(absent *structs.Absent) (*structs.Absent, error)
	CreateAbsentType(absent *structs.AbsentType) (*structs.AbsentType, error)
	CreateAccountItemList(accountItemList []structs.AccountItem) ([]*structs.AccountItem, error)
	CreateAssessments(data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error)
	CreateDispatchItem(item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error)
	CreateDispatchItemItem(item *structs.BasicInventoryDispatchItemsItem) (*structs.BasicInventoryDispatchItemsItem, error)
	CreateDropdownSettings(data *structs.SettingsDropdown) (*structs.SettingsDropdown, error)
	CreateEmployeeContract(contract *structs.Contracts) (*structs.Contracts, error)
	CreateEmployeeEducation(education *structs.Education) (*structs.Education, error)
	CreateEmployeeEvaluation(evaluation *structs.Evaluation) (*structs.Evaluation, error)
	CreateEmployeeFamilyMember(familyMember *structs.Family) (*structs.Family, error)
	CreateEmployeeForeigner(foreigner *structs.Foreigners) (*structs.Foreigners, error)
	CreateEmployeeSalaryParams(salaries *structs.SalaryParams) (*structs.SalaryParams, error)
	CreateEmployeesInOrganizationUnits(data *structs.EmployeesInOrganizationUnits) (*structs.EmployeesInOrganizationUnits, error)
	CreateExperience(contract *structs.Experience) (*structs.Experience, error)
	CreateInventoryItem(item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error)
	CreateJobPositions(data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error)
	CreateJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error)
	CreateJobTender(jobTender *structs.JobTenders) (*structs.JobTenders, error)
	CreateJobTenderApplication(jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error)
	CreateJobTenderType(jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error)
	CreateJudgeNorm(norm *structs.JudgeNorms) (*structs.JudgeNorms, error)
	CreateJudgeResolutionItems(item *structs.JudgeResolutionItems) (*structs.JudgeResolutionItems, error)
	CreateJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitItem) (*dto.JudgeResolutionsOrganizationUnitItem, error)
	CreateJudgeResolutions(resolution *structs.JudgeResolutions) (*structs.JudgeResolutions, error)
	CreateMovementArticle(input dto.MovementArticle) (*dto.MovementArticle, error)
	CreateMovements(input structs.OrderAssetMovementItem) (*structs.Movement, error)
	CreateNotification(notification *structs.Notifications) (*structs.Notifications, error)
	CreateOrderListItem(orderListItem *structs.OrderListItem) (*structs.OrderListItem, error)
	CreateOrderListProcurementArticles(orderListID int, data structs.OrderListInsertItem) error
	CreateOrderProcurementArticle(orderProcurementArticleItem *structs.OrderProcurementArticleItem) (*structs.OrderProcurementArticleItem, error)
	CreateOrganizationUnits(data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error)
	CreateProcurementArticle(article *structs.PublicProcurementArticle) (*structs.PublicProcurementArticle, error)
	CreateProcurementContract(resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error)
	CreateProcurementContractArticle(article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error)
	CreateProcurementContractArticleOverage(articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error)
	CreateProcurementItem(item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error)
	CreateProcurementOUArticle(article *structs.PublicProcurementOrganizationUnitArticle) (*structs.PublicProcurementOrganizationUnitArticle, error)
	CreateProcurementOULimit(limit *structs.PublicProcurementLimit) (*structs.PublicProcurementLimit, error)
	CreateProcurementPlan(resolution *structs.PublicProcurementPlan) (*structs.PublicProcurementPlan, error)
	CreateResolution(resolution *structs.Resolution) (*structs.Resolution, error)
	CreateRevision(revision *structs.Revision) (*structs.Revision, error)
	CreateRevisionOrgUnit(plan *dto.RevisionOrgUnit) error
	CreateRevisionPlan(plan *dto.RevisionPlanItem) (*dto.RevisionPlanItem, error)
	CreateRevisionRevisor(plan *dto.RevisionRevisor) error
	CreateRevisionTips(plan *structs.RevisionTips) (*structs.RevisionTips, error)
	CreateRevisions(plan *structs.Revisions) (*structs.Revisions, error)
	CreateRole(data structs.Roles) (*structs.Roles, error)
	CreateStock(input dto.MovementArticle) error
	CreateSupplier(supplier *structs.Suppliers) (*structs.Suppliers, error)
	CreateSystematization(data *structs.Systematization) (*structs.Systematization, error)
	CreateUserAccount(user structs.UserAccounts) (*structs.UserAccounts, error)
	CreateUserProfile(user structs.UserProfiles) (*structs.UserProfiles, error)
	DeactivateUserAccount(userID int) (*structs.UserAccounts, error)
	DeleteAbsent(id int) error
	DeleteAbsentType(id int) error
	DeleteAccount(id int) error
	DeleteAssessment(id int) error
	DeleteDropdownSettings(id int) error
	DeleteEmployeeContract(id int) error
	DeleteEmployeeEducation(id int) error
	DeleteEmployeeFamilyMember(id int) error
	DeleteEmployeeInOrganizationUnit(jobPositionInOrganizationUnitID int) error
	DeleteEmployeeInOrganizationUnitByID(jobPositionInOrganizationUnitID int) error
	DeleteEvaluation(id int) error
	DeleteExperience(id int) error
	DeleteFile(id int) error
	DeleteForeigner(id int) error
	DeleteInventoryDispatch(id int) error
	DeleteJJudgeResolutionOrganizationUnit(id int) error
	DeleteJobPositions(id int) error
	DeleteJobPositionsInOrganizationUnits(id int) error
	DeleteJobTender(id int) error
	DeleteJobTenderApplication(id int) error
	DeleteJobTenderType(id int) error
	DeleteJudgeNorm(id int) error
	DeleteJudgeResolution(id int) error
	DeleteMovement(id int) error
	DeleteNotification(notificationID int) error
	DeleteOrderList(id int) error
	DeleteOrderProcurementArticle(id int) error
	DeleteOrganizationUnits(id int) error
	DeleteProcurementArticle(id int) error
	DeleteProcurementContract(id int) error
	DeleteProcurementContractArticleOverage(id int) error
	DeleteProcurementItem(id int) error
	DeleteProcurementPlan(id int) error
	DeleteResolution(id int) error
	DeleteRevision(id int) error
	DeleteRevisionOrgUnit(id int) error
	DeleteRevisionPlan(id int) error
	DeleteRevisionRevisor(id int) error
	DeleteRevisionTips(id int) error
	DeleteRevisions(id int) error
	DeleteSalaryParams(id int) error
	DeleteSupplier(id int) error
	DeleteSystematization(id int) error
	DeleteUserAccount(id int) error
	DeleteUserProfile(id int) error
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
	SendOrderListToFinance(id int) error
	SyncPermissions(roleID int, input []*structs.RolePermission) ([]structs.RolePermission, error)
	UpdateAbsent(id int, absent *structs.Absent) (*structs.Absent, error)
	UpdateAbsentType(id int, absent *structs.AbsentType) (*structs.AbsentType, error)
	UpdateAccountItem(id int, accountItem *structs.AccountItem) (*structs.AccountItem, error)
	UpdateAssessments(id int, data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error)
	UpdateDispatchItem(id int, item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error)
	UpdateDropdownSettings(id int, data *structs.SettingsDropdown) (*structs.SettingsDropdown, error)
	UpdateEmployeeContract(id int, contract *structs.Contracts) (*structs.Contracts, error)
	UpdateEmployeeEducation(id int, education *structs.Education) (*structs.Education, error)
	UpdateEmployeeEvaluation(id int, evaluation *structs.Evaluation) (*structs.Evaluation, error)
	UpdateEmployeeFamilyMember(id int, education *structs.Family) (*structs.Family, error)
	UpdateEmployeeForeigner(id int, foreigner *structs.Foreigners) (*structs.Foreigners, error)
	UpdateEmployeeSalaryParams(id int, salaries *structs.SalaryParams) (*structs.SalaryParams, error)
	UpdateExperience(id int, contract *structs.Experience) (*structs.Experience, error)
	UpdateInventoryItem(id int, item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error)
	UpdateJobPositions(id int, data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error)
	UpdateJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error)
	UpdateJobTender(id int, jobTender *structs.JobTenders) (*structs.JobTenders, error)
	UpdateJobTenderApplication(id int, jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error)
	UpdateJobTenderType(id int, jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error)
	UpdateJudgeNorm(id int, norm *structs.JudgeNorms) (*structs.JudgeNorms, error)
	UpdateJudgeResolutionItems(id int, item *structs.JudgeResolutionItems) (*structs.JudgeResolutionItems, error)
	UpdateJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitItem) (*dto.JudgeResolutionsOrganizationUnitItem, error)
	UpdateJudgeResolutions(id int, resolution *structs.JudgeResolutions) (*structs.JudgeResolutions, error)
	UpdateMovements(input structs.OrderAssetMovementItem) (*structs.Movement, error)
	UpdateNotification(notificationID int, notification *structs.Notifications) error
	UpdateOrderListItem(id int, orderListItem *structs.OrderListItem) (*structs.OrderListItem, error)
	UpdateOrderProcurementArticle(item *structs.OrderProcurementArticleItem) (*structs.OrderProcurementArticleItem, error)
	UpdateOrganizationUnits(id int, data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error)
	UpdateProcurementArticle(id int, article *structs.PublicProcurementArticle) (*structs.PublicProcurementArticle, error)
	UpdateProcurementContract(id int, resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error)
	UpdateProcurementContractArticle(id int, article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error)
	UpdateProcurementContractArticleOverage(id int, articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error)
	UpdateProcurementItem(id int, item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error)
	UpdateProcurementOUArticle(id int, article *structs.PublicProcurementOrganizationUnitArticle) (*structs.PublicProcurementOrganizationUnitArticle, error)
	UpdateProcurementOULimit(id int, limit *structs.PublicProcurementLimit) (*structs.PublicProcurementLimit, error)
	UpdateProcurementPlan(id int, resolution *structs.PublicProcurementPlan) (*structs.PublicProcurementPlan, error)
	UpdateResolution(id int, resolution *structs.Resolution) (*structs.Resolution, error)
	UpdateRevision(id int, revision *structs.Revision) (*structs.Revision, error)
	UpdateRevisionPlan(id int, plan *dto.RevisionPlanItem) (*dto.RevisionPlanItem, error)
	UpdateRevisionTips(id int, plan *structs.RevisionTips) (*structs.RevisionTips, error)
	UpdateRevisions(id int, plan *structs.Revisions) (*structs.Revisions, error)
	UpdateRole(id int, data structs.Roles) (*structs.Roles, error)
	UpdateStock(input structs.StockArticle) error
	UpdateSupplier(id int, supplier *structs.Suppliers) (*structs.Suppliers, error)
	UpdateSystematization(id int, data *structs.Systematization) (*structs.Systematization, error)
	UpdateUserAccount(userID int, user structs.UserAccounts) (*structs.UserAccounts, error)
	UpdateUserProfile(userID int, user structs.UserProfiles) (*structs.UserProfiles, error)
	ValidateMail(input *dto.ResetPasswordVerify) (*dto.ResetPasswordVerifyResponseMS, error)
	ValidatePin(pin string, headers map[string]string) error

	GetLatestVersionOfAccounts() (int, error)

	CreateBudget(budget *structs.Budget) (*structs.Budget, error)
	UpdateBudget(budget *structs.Budget) (*structs.Budget, error)
	GetBudget(id int) (*structs.Budget, error)
	GetBudgetList(input *dto.GetBudgetListInputMS) ([]structs.Budget, error)
	DeleteBudget(id int) error

	CreateBudgetRequest(budget *structs.BudgetRequest) (*structs.BudgetRequest, error)
	UpdateBudgetRequest(budget *structs.BudgetRequest) (*structs.BudgetRequest, error)
	GetBudgetRequest(id int) (*structs.BudgetRequest, error)
	GetBudgetRequestList(input *dto.GetBudgetRequestListInputMS) ([]structs.BudgetRequest, error)
	GetOneBudgetRequest(input *dto.GetBudgetRequestListInputMS) (*structs.BudgetRequest, error)

	UpdateFinancialBudget(financialBudget *structs.FinancialBudget) (*structs.FinancialBudget, error)
	CreateFinancialBudget(financialBudget *structs.FinancialBudget) (*structs.FinancialBudget, error)
	GetFinancialBudgetByBudgetID(budgetID int) (*structs.FinancialBudget, error)
	GetFinancialBudgetByID(id int) (*structs.FinancialBudget, error)
	GetFilledFinancialBudgetList(requestID int) ([]structs.FilledFinanceBudget, error)

	GetBudgetLimits(budgetID int) ([]structs.FinancialBudgetLimit, error)
	GetBudgetUnitLimit(budgetID, unitID int) (int, error)
	CreateBudgetLimit(budgetLimit *structs.FinancialBudgetLimit) (*structs.FinancialBudgetLimit, error)
	UpdateBudgetLimit(budgetLimit *structs.FinancialBudgetLimit) (*structs.FinancialBudgetLimit, error)
	DeleteBudgetLimit(id int) error

	UpdateNonFinancialBudget(id int, program *structs.NonFinancialBudgetItem) (*structs.NonFinancialBudgetItem, error)
	CreateNonFinancialBudget(budget *structs.NonFinancialBudgetItem) (*structs.NonFinancialBudgetItem, error)
	FillFinancialBudget(budget *structs.FilledFinanceBudget) (*structs.FilledFinanceBudget, error)
	UpdateFilledFinancialBudget(id int, budget *structs.FilledFinanceBudget) (*structs.FilledFinanceBudget, error)
	FillActualFinancialBudget(id int, actual decimal.Decimal) (*structs.FilledFinanceBudget, error)
	GetFinancialFilledSummary(budgetID int, reqType structs.RequestType) ([]structs.FilledFinanceBudget, error)
	DeleteFilledFinancialBudgetData(id int) error
	DeleteNonFinancialBudget(id int) error
	GetNonFinancialBudget(id int) (*structs.NonFinancialBudgetItem, error)
	GetNonFinancialBudgetList(input *dto.GetNonFinancialBudgetListInputMS) ([]structs.NonFinancialBudgetItem, error)
	GetNonFinancialBudgetByRequestID(requestID int) (structs.NonFinancialBudgetItem, error)

	UpdateNonFinancialGoal(id int, activity *structs.NonFinancialGoalItem) (*structs.NonFinancialGoalItem, error)
	CreateNonFinancialGoal(goal *structs.NonFinancialGoalItem) (*structs.NonFinancialGoalItem, error)
	DeleteNonFinancialGoal(id int) error
	GetNonFinancialGoal(id int) (*structs.NonFinancialGoalItem, error)
	GetNonFinancialGoalList(input *dto.GetNonFinancialGoalListInputMS) ([]structs.NonFinancialGoalItem, error)

	UpdateNonFinancialGoalIndicator(id int, goalIndicator *structs.NonFinancialGoalIndicatorItem) (*structs.NonFinancialGoalIndicatorItem, error)
	CreateNonFinancialGoalIndicator(goal *structs.NonFinancialGoalIndicatorItem) (*structs.NonFinancialGoalIndicatorItem, error)
	DeleteNonFinancialGoalIndicator(id int) error
	GetNonFinancialGoalIndicator(id int) (*structs.NonFinancialGoalIndicatorItem, error)
	GetNonFinancialGoalIndicatorList(input *dto.GetNonFinancialGoalIndicatorListInputMS) ([]structs.NonFinancialGoalIndicatorItem, error)

	UpdateProgram(id int, program *structs.ProgramItem) (*structs.ProgramItem, error)
	CreateProgram(program *structs.ProgramItem) (*structs.ProgramItem, error)
	DeleteProgram(id int) error
	GetProgram(id int) (*structs.ProgramItem, error)
	GetProgramList(input *dto.GetFinanceProgramListInputMS) ([]structs.ProgramItem, error)

	UpdateActivity(id int, activity *structs.ActivitiesItem) (*structs.ActivitiesItem, error)
	CreateActivity(activity *structs.ActivitiesItem) (*structs.ActivitiesItem, error)
	DeleteActivity(id int) error
	GetActivity(id int) (*structs.ActivitiesItem, error)
	GetActivityList(input *dto.GetFinanceActivityListInputMS) ([]structs.ActivitiesItem, error)
	GetActivityByUnit(organizationUnitID int) (*structs.ActivitiesItem, error)

	GetSpendingDynamic(budgetID, unitID int, input *dto.GetSpendingDynamicHistoryInput) ([]dto.SpendingDynamicDTO, error)
	GetSpendingDynamicHistory(budgetID, unitID int) ([]dto.SpendingDynamicHistoryDTO, error)
	GetSpendingDynamicActual(budgetID, unitID, accountID int) (decimal.NullDecimal, error)
	CreateSpendingDynamic(ctx context.Context, spendingDynamic []structs.SpendingDynamicInsert) ([]dto.SpendingDynamicDTO, error)

	CreateSpendingRelease(ctx context.Context, spendingRelease *structs.SpendingReleaseInsert) (*structs.SpendingRelease, error)

	CreateCurrentBudget(ctx context.Context, currentBudget *structs.CurrentBudget) (*structs.CurrentBudget, error)

	CreateInvoice(item *structs.Invoice) (*structs.Invoice, error)
	UpdateInvoice(item *structs.Invoice) (*structs.Invoice, error)
	GetInvoice(id int) (*structs.Invoice, error)
	GetInvoiceList(input *dto.GetInvoiceListInputMS) ([]structs.Invoice, int, error)
	GetInvoiceArticleList(id int) ([]structs.InvoiceArticles, error)
	GetAdditionalExpenses(input *dto.AdditionalExpensesListInputMS) ([]structs.AdditionalExpenses, int, error)
	DeleteInvoice(id int) error
	DeleteInvoiceArticle(id int) error
	CreateInvoiceArticle(article *structs.InvoiceArticles) (*structs.InvoiceArticles, error)
	UpdateInvoiceArticle(item *structs.InvoiceArticles) (*structs.InvoiceArticles, error)

	GetTaxAuthorityCodebookByID(id int) (*structs.TaxAuthorityCodebook, error)
	GetTaxAuthorityCodebooks(input dto.TaxAuthorityCodebookFilter) (*dto.GetTaxAuthorityCodebooksResponseMS, error)
	CreateTaxAuthorityCodebook(item *structs.TaxAuthorityCodebook) (*structs.TaxAuthorityCodebook, error)
	UpdateTaxAuthorityCodebook(id int, data *structs.TaxAuthorityCodebook) (*structs.TaxAuthorityCodebook, error)
	DeactivateTaxAuthorityCodebook(id int, active bool) error
	DeleteTaxAuthorityCodebook(id int) error

	CreateFee(item *structs.Fee) (*structs.Fee, error)
	GetFee(id int) (*structs.Fee, error)
	GetFeeList(input *dto.GetFeeListInputMS) ([]structs.Fee, int, error)
	DeleteFee(id int) error
	UpdateFee(item *structs.Fee) (*structs.Fee, error)

	CreateFeePayment(item *structs.FeePayment) (*structs.FeePayment, error)
	GetFeePayment(id int) (*structs.FeePayment, error)
	GetFeePaymentList(input *dto.GetFeePaymentListInputMS) ([]structs.FeePayment, int, error)
	DeleteFeePayment(id int) error
	UpdateFeePayment(item *structs.FeePayment) (*structs.FeePayment, error)

	CreateFine(item *structs.Fine) (*structs.Fine, error)
	GetFine(id int) (*structs.Fine, error)
	GetFineList(input *dto.GetFineListInputMS) ([]structs.Fine, int, error)
	DeleteFine(id int) error
	UpdateFine(item *structs.Fine) (*structs.Fine, error)

	CreateFinePayment(item *structs.FinePayment) (*structs.FinePayment, error)
	GetFinePayment(id int) (*structs.FinePayment, error)
	GetFinePaymentList(input *dto.GetFinePaymentListInputMS) ([]structs.FinePayment, int, error)
	DeleteFinePayment(id int) error
	UpdateFinePayment(item *structs.FinePayment) (*structs.FinePayment, error)

	CreateProcedureCost(item *structs.ProcedureCost) (*structs.ProcedureCost, error)
	GetProcedureCost(id int) (*structs.ProcedureCost, error)
	GetProcedureCostList(input *dto.GetProcedureCostListInputMS) ([]structs.ProcedureCost, int, error)
	DeleteProcedureCost(id int) error
	UpdateProcedureCost(item *structs.ProcedureCost) (*structs.ProcedureCost, error)

	CreateProcedureCostPayment(item *structs.ProcedureCostPayment) (*structs.ProcedureCostPayment, error)
	GetProcedureCostPayment(id int) (*structs.ProcedureCostPayment, error)
	GetProcedureCostPaymentList(input *dto.GetProcedureCostPaymentListInputMS) ([]structs.ProcedureCostPayment, int, error)
	DeleteProcedureCostPayment(id int) error
	UpdateProcedureCostPayment(item *structs.ProcedureCostPayment) (*structs.ProcedureCostPayment, error)

	CreateFlatRate(item *structs.FlatRate) (*structs.FlatRate, error)
	GetFlatRate(id int) (*structs.FlatRate, error)
	GetFlatRateList(input *dto.GetFlatRateListInputMS) ([]structs.FlatRate, int, error)
	DeleteFlatRate(id int) error
	UpdateFlatRate(item *structs.FlatRate) (*structs.FlatRate, error)

	CreateFlatRatePayment(item *structs.FlatRatePayment) (*structs.FlatRatePayment, error)
	GetFlatRatePayment(id int) (*structs.FlatRatePayment, error)
	GetFlatRatePaymentList(input *dto.GetFlatRatePaymentListInputMS) ([]structs.FlatRatePayment, int, error)
	DeleteFlatRatePayment(id int) error
	UpdateFlatRatePayment(item *structs.FlatRatePayment) (*structs.FlatRatePayment, error)

	CreatePropBenConf(item *structs.PropBenConf) (*structs.PropBenConf, error)
	GetPropBenConf(id int) (*structs.PropBenConf, error)
	GetPropBenConfList(input *dto.GetPropBenConfListInputMS) ([]structs.PropBenConf, int, error)
	DeletePropBenConf(id int) error
	UpdatePropBenConf(item *structs.PropBenConf) (*structs.PropBenConf, error)

	CreatePropBenConfPayment(item *structs.PropBenConfPayment) (*structs.PropBenConfPayment, error)
	GetPropBenConfPayment(id int) (*structs.PropBenConfPayment, error)
	GetPropBenConfPaymentList(input *dto.GetPropBenConfPaymentListInputMS) ([]structs.PropBenConfPayment, int, error)
	DeletePropBenConfPayment(id int) error
	UpdatePropBenConfPayment(item *structs.PropBenConfPayment) (*structs.PropBenConfPayment, error)

	CreateFixedDeposit(item *structs.FixedDeposit) (*structs.FixedDeposit, error)
	DeleteFixedDeposit(id int) error
	UpdateFixedDeposit(item *structs.FixedDeposit) (*structs.FixedDeposit, error)
	GetFixedDepositByID(id int) (*structs.FixedDeposit, error)
	GetFixedDepositList(input dto.FixedDepositFilter) ([]structs.FixedDeposit, int, error)

	CreateFixedDepositWill(item *structs.FixedDepositWill) (*structs.FixedDepositWill, error)
	DeleteFixedDepositWill(id int) error
	UpdateFixedDepositWill(item *structs.FixedDepositWill) (*structs.FixedDepositWill, error)
	GetFixedDepositWillByID(id int) (*structs.FixedDepositWill, error)
	GetFixedDepositWillList(input dto.FixedDepositWillFilter) ([]structs.FixedDepositWill, int, error)

	CreateFixedDepositItem(item *structs.FixedDepositItem) error
	UpdateFixedDepositItem(item *structs.FixedDepositItem) error
	DeleteFixedDepositItem(id int) error
	CreateFixedDepositDispatch(item *structs.FixedDepositDispatch) error
	UpdateFixedDepositDispatch(item *structs.FixedDepositDispatch) error
	DeleteFixedDepositDispatch(id int) error
	CreateFixedDepositJudge(item *structs.FixedDepositJudge) error
	UpdateFixedDepositJudge(item *structs.FixedDepositJudge) error
	DeleteFixedDepositJudge(id int) error
	CreateFixedDepositWillDispatch(item *structs.FixedDepositWillDispatch) error
	UpdateFixedDepositWillDispatch(item *structs.FixedDepositWillDispatch) error
	DeleteFixedDepositWillDispatch(id int) error

	CreateSalary(item *structs.Salary) (*structs.Salary, error)
	DeleteSalary(id int) error
	UpdateSalary(item *structs.Salary) (*structs.Salary, error)
	GetSalaryByID(id int) (*structs.Salary, error)
	GetSalaryList(input dto.SalaryFilter) ([]structs.Salary, int, error)

	CreateDepositPayment(item *structs.DepositPayment) (*structs.DepositPayment, error)
	DeleteDepositPayment(id int) error
	UpdateDepositPayment(item *structs.DepositPayment) (*structs.DepositPayment, error)
	GetDepositPaymentByID(id int) (*structs.DepositPayment, error)
	GetDepositPaymentList(input dto.DepositPaymentFilter) ([]structs.DepositPayment, int, error)
	GetInitialState(input dto.DepositInitialStateFilter) ([]structs.DepositPayment, error)
	GetDepositPaymentCaseNumber(caseNumber string, bankAccount string) (*structs.DepositPayment, error)
	GetCaseNumber(organizationUnitID int, bankAccount string) ([]structs.DepositPayment, error)

	CreateDepositPaymentOrder(item *structs.DepositPaymentOrder) (*structs.DepositPaymentOrder, error)
	DeleteDepositPaymentOrder(id int) error
	UpdateDepositPaymentOrder(item *structs.DepositPaymentOrder) (*structs.DepositPaymentOrder, error)
	GetDepositPaymentOrderByID(id int) (*structs.DepositPaymentOrder, error)
	GetDepositPaymentOrderList(input dto.DepositPaymentOrderFilter) ([]structs.DepositPaymentOrder, int, error)
	GetDepositPaymentAdditionalExpenses(input *dto.DepositPaymentAdditionalExpensesListInputMS) ([]structs.DepositPaymentAdditionalExpenses, int, error)
	PayDepositPaymentOrder(input structs.DepositPaymentOrder) error

	CreatePaymentOrder(item *structs.PaymentOrder) (*structs.PaymentOrder, error)
	DeletePaymentOrder(id int) error
	UpdatePaymentOrder(item *structs.PaymentOrder) (*structs.PaymentOrder, error)
	GetPaymentOrderByID(id int) (*structs.PaymentOrder, error)
	GetPaymentOrderList(input dto.PaymentOrderFilter) ([]structs.PaymentOrder, int, error)
	GetAllObligations(input dto.ObligationsFilter) ([]dto.Obligation, int, error)
	PayPaymentOrder(input structs.PaymentOrder) error
	CancelPaymentOrder(id int) error

	CreateEnforcedPayment(item *structs.EnforcedPayment) (*structs.EnforcedPayment, error)
	UpdateEnforcedPayment(item *structs.EnforcedPayment) (*structs.EnforcedPayment, error)
	GetEnforcedPaymentByID(id int) (*structs.EnforcedPayment, error)
	GetEnforcedPaymentList(input dto.EnforcedPaymentFilter) ([]structs.EnforcedPayment, int, error)
	ReturnEnforcedPayment(input structs.EnforcedPayment) error

	GetAllObligationsForAccounting(input dto.ObligationsFilter) ([]dto.ObligationForAccounting, int, error)
	GetAllPaymentOrdersForAccounting(input dto.ObligationsFilter) ([]dto.PaymentOrdersForAccounting, int, error)
	GetAllEnforcedPaymentsForAccounting(input dto.ObligationsFilter) ([]dto.PaymentOrdersForAccounting, int, error)
	GetAllReturnedEnforcedPaymentsForAccounting(input dto.ObligationsFilter) ([]dto.PaymentOrdersForAccounting, int, error)
	BuildAccountingOrderForObligations(data structs.AccountingOrderForObligationsData) (*dto.AccountingOrderForObligations, error)

	UpdateModelsOfAccounting(item *structs.ModelsOfAccounting) (*structs.ModelsOfAccounting, error)
	GetModelsOfAccountingByID(id int) (*structs.ModelsOfAccounting, error)
	GetModelsOfAccountingList(input dto.ModelsOfAccountingFilter) ([]structs.ModelsOfAccounting, int, error)

	CreateAccountingEntry(item *structs.AccountingEntry) (*structs.AccountingEntry, error)
	DeleteAccountingEntry(id int) error
	GetAccountingEntryByID(id int) (*structs.AccountingEntry, error)
	GetAccountingEntryList(input dto.AccountingEntryFilter) ([]structs.AccountingEntry, int, error)

	GetAnalyticalCard(input dto.AnalyticalCardFilter) ([]structs.AnalyticalCard, error)
}
