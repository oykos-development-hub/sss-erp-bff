package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"net/http"
)

type MicroserviceRepositoryInterface interface {
	AddOnStock(stock []structs.StockArticle, article structs.OrderProcurementArticleItem, organizationUnitID int) error
	AuthenticateUser(r *http.Request) (*structs.UserAccounts, error)
	CheckInsertInventoryData(input []structs.BasicInventoryInsertItem) ([]structs.BasicInventoryInsertValidator, error)
	CreateAbsent(absent *structs.Absent) (*structs.Absent, error)
	CreateAbsentType(absent *structs.AbsentType) (*structs.AbsentType, error)
	CreateAccountItem(accountItem *structs.AccountItem) (*structs.AccountItem, error)
	CreateAssessments(data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error)
	CreateDispatchItem(item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error)
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
	GetDispatchItemByID(id int) (*structs.BasicInventoryDispatchItem, error)
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
	GetJudgeNormListByEmployee(userProfileID int) ([]structs.JudgeNorms, error)
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
}
