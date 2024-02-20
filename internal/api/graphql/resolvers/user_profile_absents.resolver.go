package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfileVacationResolver(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"].(int)

	resolutions, err := r.Repo.GetEmployeeResolutions(userProfileID, nil)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	vacationResItemList, err := buildVacationResponseItemList(r.Repo, resolutions)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   vacationResItemList,
	}, nil
}

func buildVacationResponseItemList(r repository.MicroserviceRepositoryInterface, items []*structs.Resolution) (resItemList []*dto.Vacation, err error) {
	for _, item := range items {
		resItem, err := buildVacationResItem(r, item)
		if err != nil {
			return nil, err
		}
		if resItem != nil {
			resItemList = append(resItemList, resItem)
		}
	}
	return
}

func (r *Resolver) UserProfileVacationResolutionInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Vacation
	var inputData structs.Resolution
	vacationTypeValue := config.VacationTypeValueResolutionType
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	vacationType, err := r.Repo.GetDropdownSettings(&dto.GetSettingsInput{Value: &vacationTypeValue, Entity: config.ResolutionTypes})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	dateOfEnd := time.Date(data.Year, time.December, 31, 23, 59, 59, 999999999, time.UTC).Format("2006-01-02T15:04:05Z")
	inputData.ResolutionTypeID = vacationType.Data[0].ID
	inputData.DateOfStart = time.Date(data.Year, time.January, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05Z")
	inputData.DateOfEnd = &dateOfEnd
	inputData.ID = data.ID
	inputData.FileID = data.FileID
	inputData.UserProfileID = data.UserProfileID
	inputData.Value = strconv.Itoa(data.NumberOfDays)
	inputData.ResolutionPurpose = data.ResolutionPurpose

	if inputData.ID != 0 {
		resolution, err := r.Repo.UpdateResolution(inputData.ID, &inputData)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resolutionResItem, err := buildVacationResItem(r.Repo, resolution)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You updated this item!"
	} else {
		resolution, err := r.Repo.CreateResolution(&inputData)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resolutionResItem, err := buildVacationResItem(r.Repo, resolution)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) UserProfileVacationResolutionsInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.VacationArray

	vacationTypeValue := config.VacationTypeValueResolutionType
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.Response{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	vacationType, err := r.Repo.GetDropdownSettings(&dto.GetSettingsInput{Value: &vacationTypeValue, Entity: config.ResolutionTypes})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	var vacations []dto.Vacation

	for _, vacation := range data.Data {
		var inputData structs.Resolution
		dateOfEnd := time.Date(data.Year, time.December, 31, 23, 59, 59, 999999999, time.UTC).Format("2006-01-02T15:04:05Z")
		inputData.ResolutionTypeID = vacationType.Data[0].ID
		inputData.DateOfStart = time.Date(data.Year, time.January, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05Z")
		inputData.DateOfEnd = &dateOfEnd
		inputData.UserProfileID = vacation.UserProfileID
		inputData.Value = strconv.Itoa(vacation.NumberOfDays)

		resolution, err := r.Repo.CreateResolution(&inputData)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resolutionResItem, err := buildVacationResItem(r.Repo, resolution)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		vacations = append(vacations, *resolutionResItem)
	}

	response.Data = vacations

	return response, nil
}

func buildVacationResItem(r repository.MicroserviceRepositoryInterface, item *structs.Resolution) (*dto.Vacation, error) {
	userProfile, err := r.GetUserProfileByID(item.UserProfileID)
	if err != nil {
		return nil, err
	}
	resolutionType, err := r.GetDropdownSettingByID(item.ResolutionTypeID)
	if err != nil {
		return nil, err
	}

	if resolutionType.Value != config.VacationTypeValueResolutionType {
		return nil, nil
	}

	dataOfStart, _ := time.Parse("2006-01-02T15:04:05Z", item.DateOfStart)
	numberOfDays, _ := strconv.Atoi(item.Value)

	var file dto.FileDropdownSimple

	if item.FileID > 0 {
		res, err := r.GetFileByID(item.FileID)

		if err != nil {
			return nil, err
		}

		file.ID = res.ID
		file.Name = res.Name
		file.Type = *res.Type
	}
	return &dto.Vacation{
		ID:                item.ID,
		ResolutionPurpose: item.ResolutionPurpose,
		UserProfile: dto.DropdownSimple{
			ID:    userProfile.ID,
			Title: userProfile.GetFullName(),
		},
		ResolutionType: dto.DropdownSimple{
			ID:    resolutionType.ID,
			Title: resolutionType.Title,
		},
		Year:         dataOfStart.Year(),
		NumberOfDays: numberOfDays,
		FileID:       item.FileID,
		File:         file,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}, nil
}

func (r *Resolver) UserProfileAbsentResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		absentSummary dto.AbsentsSummary
	//	usedDays      int
	)

	profileID := params.Args["user_profile_id"].(int)

	// year ago
	sumDaysOfCurrentYear, availableDaysOfCurrentYear, availableDaysOfPreviousYear, err := GetNumberOfCurrentAndPreviousYearAvailableDays(r.Repo, profileID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	// get all absents in a period of current year
	currentYear := time.Now().Year()
	startOfYear := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := time.Date(currentYear, time.December, 31, 23, 59, 59, 999999999, time.UTC)
	absents, err := r.Repo.GetEmployeeAbsents(profileID, &dto.EmployeeAbsentsInput{From: &startOfYear, To: &endOfYear})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, absent := range absents {
		absentType, err := r.Repo.GetAbsentTypeByID(absent.AbsentTypeID)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		absent.AbsentType = *absentType

		/*if absentType.AccountingDaysOff {
			daysTakenBeforeJuly, daysTakenAfterJuly := getTakenVacationDaysBeforeAndAfterJuly(absent.DateOfStart, absent.DateOfEnd)

			// Subtract vacation days taken before July from available previous year days
			if daysTakenBeforeJuly > 0 {
				if availableDaysOfPreviousYear >= daysTakenBeforeJuly {
					availableDaysOfPreviousYear -= daysTakenBeforeJuly
				} else {
					// if available days from previous year are not enough, we should use current year too
					availableDaysOfCurrentYear -= daysTakenBeforeJuly - availableDaysOfPreviousYear
					availableDaysOfPreviousYear = 0
				}
			}
			// Subtract vacation days taken after July from available current year vacation days
			if daysTakenAfterJuly > 0 {
				availableDaysOfCurrentYear -= daysTakenAfterJuly
			}

			usedDays += (daysTakenBeforeJuly + daysTakenAfterJuly)
		}*/
	}

	allAbsents, _ := r.Repo.GetEmployeeAbsents(profileID, nil)
	for _, absent := range allAbsents {
		if absent.TargetOrganizationUnitID != nil {
			organizationUnit, err := r.Repo.GetOrganizationUnitByID(*absent.TargetOrganizationUnitID)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			absent.TargetOrganizationUnit = organizationUnit
		}

		absentType, err := r.Repo.GetAbsentTypeByID(absent.AbsentTypeID)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		if absent.FileID > 0 {
			res, err := r.Repo.GetFileByID(absent.FileID)

			if err != nil {
				return nil, err
			}

			absent.File.ID = res.ID
			absent.File.Name = res.Name
			absent.File.Type = *res.Type
		}
		absent.AbsentType = *absentType
	}

	absentSummary.CurrentAvailableDays = availableDaysOfCurrentYear
	absentSummary.PastAvailableDays = availableDaysOfPreviousYear
	absentSummary.UsedDays = sumDaysOfCurrentYear - availableDaysOfCurrentYear

	return dto.Response{
		Status:  "success",
		Message: "Here's the items you asked for!",
		Summary: absentSummary,
		Items:   allAbsents,
	}, nil
}

/*
	func getTakenVacationDaysBeforeAndAfterJuly(startDate string, endDate string) (int, int) {
		// Parse the date string
		start, _ := time.Parse(time.RFC3339, startDate)
		end, _ := time.Parse(time.RFC3339, endDate)

		currentYear := time.Now().Year()

		workingDaysBeforeJuly := 0
		workingDaysAfterJuly := 0

		july := time.Date(currentYear, time.July, 1, 0, 0, 0, 0, start.Location())

		for !start.After(end) {
			if start.Year() == currentYear && start.Weekday() != time.Saturday && start.Weekday() != time.Sunday {
				if start.After(july) {
					workingDaysAfterJuly++
				} else {
					workingDaysBeforeJuly++
				}
			}
			start = start.AddDate(0, 0, 1)
		}

		return workingDaysBeforeJuly, workingDaysAfterJuly
	}
*/
func GetNumberOfCurrentAndPreviousYearAvailableDays(r repository.MicroserviceRepositoryInterface, profileID int) (int, int, int, error) {
	currentYear := time.Now().Year()
	startDatePreviousYear := time.Date(currentYear-1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDateCurrentYear := time.Date(currentYear, time.December, 31, 23, 59, 59, 0, time.UTC)

	vacationDays := 0
	pastVacationDays := 0
	resolutions, err := r.GetEmployeeResolutions(profileID, &dto.EmployeeResolutionListInput{From: &startDatePreviousYear, To: &endDateCurrentYear})
	if err != nil {
		fmt.Println("error getting employee resolution - " + err.Error())
	}

	for _, resolution := range resolutions {
		resolutionType, err := r.GetDropdownSettingByID(resolution.ResolutionTypeID)
		if err != nil {
			return 0, 0, 0, err
		}
		resolution.ResolutionType = resolutionType
	}

	for _, resolution := range resolutions {
		if resolution.ResolutionType.Value != config.VacationTypeValueResolutionType {
			continue
		}

		start, _ := time.Parse(time.RFC3339, resolution.DateOfStart)

		if start.Year() == time.Now().Year() {
			totalDays, _ := strconv.Atoi(resolution.Value)
			vacationDays += totalDays
		} else if start.Year() == time.Now().Year()-1 {
			totalDays, _ := strconv.Atoi(resolution.Value)
			pastVacationDays += totalDays
		}
	}

	usedDays, err := calculateUsedDays(r, profileID)
	if err != nil {
		return 0, 0, 0, err
	}

	// pastVacationDays -= usedDaysPreviousYear
	pastVacationDays -= usedDays
	currentYearVacationDays := vacationDays

	if pastVacationDays < 0 {
		vacationDays += pastVacationDays
		pastVacationDays = 0
	}

	return currentYearVacationDays, vacationDays, pastVacationDays, nil
}

func calculateUsedDays(r repository.MicroserviceRepositoryInterface, profileID int) (int, error) {
	usedDays := 0

	currentYear := time.Now().Year()
	startOfYear := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := time.Date(currentYear, time.December, 31, 23, 59, 59, 0, time.UTC)

	absents, err := r.GetEmployeeAbsents(profileID, &dto.EmployeeAbsentsInput{
		From: &startOfYear,
		To:   &endOfYear,
	})
	if err != nil {
		return 0, err
	}

	for _, absent := range absents {
		start, _ := time.Parse(time.RFC3339, absent.DateOfStart)
		end, _ := time.Parse(time.RFC3339, absent.DateOfEnd)

		absentType, err := r.GetAbsentTypeByID(absent.AbsentTypeID)
		if err != nil {
			return 0, err
		}

		if absentType.AccountingDaysOff {
			usedDays += countWorkingDaysBetweenDates(start, end)
		}
	}

	return usedDays, nil
}

/*
	func calculateUsedDaysOfPreviousYear(profileID int) (int, error) {
		currentYear := time.Now().Year()
		startDatePreviousYear := time.Date(currentYear-1, time.January, 1, 0, 0, 0, 0, time.UTC)
		endDatePreviousYear := time.Date(currentYear-1, time.December, 31, 23, 59, 59, 0, time.UTC)

		// Initialize usedDays variable
		usedDaysPreviousYear := 0

		// Get all absents of the employee in the previous year
		absents, err := r.Repo.GetEmployeeAbsents(profileID, &dto.EmployeeAbsentsInput{From: &startDatePreviousYear, To: &endDatePreviousYear})
		if err != nil {
			return 0, err
		}

		for _, absent := range absents {
			start, _ := time.Parse(time.RFC3339, absent.DateOfStart)
			end, _ := time.Parse(time.RFC3339, absent.DateOfEnd)

			absentType, err := r.Repo.GetAbsentTypeByID(absent.AbsentTypeID)
			if err != nil {
				return 0, err
			}

			if absentType.AccountingDaysOff {
				usedDaysPreviousYear += countWorkingDaysBetweenDates(start, end)
			}
		}

		return usedDaysPreviousYear, nil
	}
*/
func countWorkingDaysBetweenDates(start, end time.Time) int {
	daysCount := 0

	for !start.After(end) {
		if start.Weekday() != time.Saturday && start.Weekday() != time.Sunday {
			daysCount++
		}
		start = start.AddDate(0, 0, 1)
	}

	return daysCount
}

func (r *Resolver) UserProfileAbsentInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var data structs.Absent
	var item *structs.Absent

	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.ErrorResponse("Bad request: user profile absent data"), nil
	}

	absentType, err := r.Repo.GetAbsentTypeByID(data.AbsentTypeID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	if absentType.Title == "Godišnji odmor" {
		_, currYearDays, pastYearDays, err := GetNumberOfCurrentAndPreviousYearAvailableDays(r.Repo, data.UserProfileID)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		dateOfStart, err := time.Parse("2006-01-02T15:04:05.000Z", data.DateOfStart)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		dateOfEnd, err := time.Parse("2006-01-02T15:04:05.000Z", data.DateOfEnd)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		newUsedData := countWorkingDaysBetweenDates(dateOfStart, dateOfEnd)

		if newUsedData > (currYearDays + pastYearDays) {
			err = fmt.Errorf("limit is reached")
			return errors.HandleAPIError(err)
		}
	}

	if data.ID != 0 {
		item, err = r.Repo.UpdateAbsent(data.ID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
	} else {
		item, err = r.Repo.CreateAbsent(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
	}

	resItem, err := buildAbsentResponseItem(r.Repo, *item)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	response.Item = resItem

	return response, nil
}

func buildAbsentResponseItem(r repository.MicroserviceRepositoryInterface, absent structs.Absent) (*structs.Absent, error) {
	absentType, err := r.GetAbsentTypeByID(absent.AbsentTypeID)
	if err != nil {
		return nil, err
	}
	absent.AbsentType = *absentType

	if absent.TargetOrganizationUnitID != nil {
		organizationUnit, err := r.GetOrganizationUnitByID(*absent.TargetOrganizationUnitID)
		if err != nil {
			return nil, err
		}
		absent.TargetOrganizationUnit = organizationUnit
	}

	if absent.FileID > 0 {
		res, err := r.GetFileByID(absent.FileID)

		if err != nil {
			return nil, err
		}

		absent.File.ID = res.ID
		absent.File.Name = res.Name
		absent.File.Type = *res.Type
	}

	return &absent, nil
}

func (r *Resolver) UserProfileAbsentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteAbsent(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) TerminateEmployment(params graphql.ResolveParams) (interface{}, error) {
	userID := params.Args["user_profile_id"].(int)

	user, err := r.Repo.GetUserProfileByID(userID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	userResponse, err := buildUserProfileOverviewResponse(r.Repo, user)

	if err != nil {
		return errors.HandleAPIError(err)
	}
	active := true

	if userResponse.IsJudge {
		input := dto.GetJudgeResolutionListInputMS{
			Active: &active,
		}

		resolution, err := r.Repo.GetJudgeResolutionList(&input)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		filter := dto.JudgeResolutionsOrganizationUnitInput{
			UserProfileID: &userID,
			ResolutionID:  &resolution.Data[0].ID,
		}

		judge, _, err := r.Repo.GetJudgeResolutionOrganizationUnit(&filter)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		err = r.Repo.DeleteJJudgeResolutionOrganizationUnit(judge[0].ID)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		contract, err := r.Repo.GetEmployeeContracts(userID, nil)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		now := time.Now()
		format := "2006-01-02T00:00:00Z"
		dateOfEnd := now.Format(format)
		dateOfStart, _ := time.Parse(format, *contract[0].DateOfStart)
		yearsDiff := now.Year() - dateOfStart.Year()
		monthsDiff := int(now.Month()) - int(dateOfStart.Month())

		if monthsDiff < 0 {
			monthsDiff += 12
			yearsDiff--
		}

		totalMonths := (yearsDiff * 12) + monthsDiff

		experience := structs.Experience{
			UserProfileID:             userID,
			OrganizationUnitID:        judge[0].OrganizationUnitID,
			Relevant:                  true,
			DateOfStart:               *contract[0].DateOfStart,
			DateOfEnd:                 dateOfEnd,
			AmountOfExperience:        totalMonths,
			AmountOfInsuredExperience: totalMonths,
		}
		_, err = r.Repo.CreateExperience(&experience)
		if err != nil {
			return errors.HandleAPIError(err)
		}

	} else {
		contract, err := r.Repo.GetEmployeeContracts(userID, nil)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		employeeInOrgUnit, err := r.Repo.GetEmployeesInOrganizationUnitsByProfileID(userID)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		if employeeInOrgUnit != nil {
			err = r.Repo.DeleteEmployeeInOrganizationUnitByID(employeeInOrgUnit.ID)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			now := time.Now()
			format := "2006-01-02T00:00:00Z"
			dateOfEnd := now.Format(format)
			dateOfStart, _ := time.Parse(format, *contract[0].DateOfStart)
			yearsDiff := now.Year() - dateOfStart.Year()
			monthsDiff := int(now.Month()) - int(dateOfStart.Month())

			if monthsDiff < 0 {
				monthsDiff += 12
				yearsDiff--
			}

			totalMonths := (yearsDiff * 12) + monthsDiff
			experience := structs.Experience{
				UserProfileID:             userID,
				OrganizationUnitID:        contract[0].OrganizationUnitID,
				Relevant:                  true,
				DateOfStart:               *contract[0].DateOfStart,
				DateOfEnd:                 dateOfEnd,
				AmountOfExperience:        totalMonths,
				AmountOfInsuredExperience: totalMonths,
			}
			_, err = r.Repo.CreateExperience(&experience)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			contract[0].JobPositionInOrganizationUnitID = 0
			contract[0].OrganizationUnitDepartmentID = nil
			contract[0].Active = false
			_, err = r.Repo.UpdateEmployeeContract(contract[0].ID, contract[0])
			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	}

	terminateResolutionTypeValue := config.EmploymentTerminationResolutionType

	terminateResolutionType, err := r.Repo.GetDropdownSettings(&dto.GetSettingsInput{Value: &terminateResolutionTypeValue, Entity: config.ResolutionTypes})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	now := time.Now()
	fileID := params.Args["file_id"].(int)

	_, err = r.Repo.CreateResolution(&structs.Resolution{
		UserProfileID:    userID,
		ResolutionTypeID: terminateResolutionType.Data[0].ID,
		IsAffect:         true,
		DateOfStart:      now.Format("2006-01-02T00:00:00Z"),
		FileID:           fileID,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	active = false
	user.ActiveContract = &active
	_, err = r.Repo.UpdateUserProfile(user.ID, *user)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	_, err = r.Repo.DeactivateUserAccount(user.UserAccountID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deactivated this user!",
	}, nil
}

type ReportType int

const (
	AnnualLeaveReport ReportType = 1
)

func (r *Resolver) GetVacationReportData(params graphql.ResolveParams) (interface{}, error) {
	organizationUnitID := params.Args["organization_unit_id"].(int)
	reportType := ReportType(params.Args["type"].(int))
	employeeID, _ := params.Args["employee_id"].(int)

	if reportType == AnnualLeaveReport {
		data, err := getDataForUsedAnnualLeaveDaysForEmployees(r.Repo, organizationUnitID, employeeID)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   data,
		}, nil
	}
	return errors.HandleAPIError(fmt.Errorf("unsupported type %d", reportType))
}

func getDataForUsedAnnualLeaveDaysForEmployees(repo repository.MicroserviceRepositoryInterface, organizationUnitID int, employeeID int) ([]dto.VacationReportResItem, error) {
	var employees []*structs.UserProfiles
	var err error

	organizationUnit, err := repo.GetOrganizationUnitByID(organizationUnitID)
	if err != nil {
		return nil, err
	}

	if employeeID > 0 {
		employee, err := repo.GetUserProfileByID(employeeID)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	} else {
		employees, err = GetEmployeesOfOrganizationUnit(repo, organizationUnitID)
		if err != nil {
			return nil, err
		}
	}

	var data []dto.VacationReportResItem
	for _, employee := range employees {
		sumDaysOfCurrentYear, availableDaysOfCurrentYear, _, err := GetNumberOfCurrentAndPreviousYearAvailableDays(repo, employee.ID)
		if err != nil {
			return nil, err
		}

		resItem := dto.VacationReportResItem{
			FullName:         employee.GetFullName(),
			OrganizationUnit: organizationUnit.Title,
			TotalDays:        sumDaysOfCurrentYear,
			UsedDays:         sumDaysOfCurrentYear - availableDaysOfCurrentYear,
			LeftDays:         availableDaysOfCurrentYear,
		}

		data = append(data, resItem)
	}

	return data, nil
}
