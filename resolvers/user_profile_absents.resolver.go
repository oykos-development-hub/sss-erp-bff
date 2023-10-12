package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

const (
	VacationTypeValue string = "vacation"
)

var UserProfileVacationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	userProfileId := params.Args["user_profile_id"].(int)

	resolutions, err := getEmployeeResolutions(userProfileId, nil)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	vacationResItemList, err := buildVacationResponseItemList(resolutions)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   vacationResItemList,
	}, nil
}

func buildVacationResponseItemList(items []*structs.Resolution) (resItemList []*dto.Vacation, err error) {
	for _, item := range items {
		resItem, err := buildVacationResItem(item)
		if err != nil {
			return nil, err
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

var UserProfileVacationResolutionInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Vacation
	var inputData structs.Resolution
	vacationTypeValue := VacationTypeValue
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	vacationType, err := getDropdownSettings(&dto.GetSettingsInput{Value: &vacationTypeValue, Entity: config.ResolutionTypes})
	if err != nil {
		return shared.HandleAPIError(err)
	}
	inputData.ResolutionTypeId = vacationType.Data[0].Id
	inputData.DateOfStart = time.Date(data.Year, time.January, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05Z")
	inputData.DateOfEnd = time.Date(data.Year, time.December, 31, 23, 59, 59, 999999999, time.UTC).Format("2006-01-02T15:04:05Z")
	inputData.Id = data.Id
	inputData.FileId = data.FileId
	inputData.UserProfileId = data.UserProfileId
	inputData.Value = strconv.Itoa(data.NumberOfDays)
	inputData.ResolutionPurpose = data.ResolutionPurpose

	if shared.IsInteger(inputData.Id) && inputData.Id != 0 {
		resolution, err := updateResolution(inputData.Id, &inputData)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resolutionResItem, err := buildVacationResItem(resolution)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You updated this item!"
	} else {
		resolution, err := createResolution(&inputData)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resolutionResItem, err := buildVacationResItem(resolution)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You created this item!"
	}

	return response, nil
}

func buildVacationResItem(item *structs.Resolution) (*dto.Vacation, error) {
	userProfile, err := getUserProfileById(item.UserProfileId)
	if err != nil {
		return nil, err
	}
	resolutionType, err := getDropdownSettingById(item.ResolutionTypeId)
	if err != nil {
		return nil, err
	}

	dataOfStart, _ := time.Parse("2006-01-02T15:04:05Z", item.DateOfStart)
	numberOfDays, _ := strconv.Atoi(item.Value)

	return &dto.Vacation{
		Id:                item.Id,
		ResolutionPurpose: item.ResolutionPurpose,
		UserProfile: dto.DropdownSimple{
			Id:    userProfile.Id,
			Title: userProfile.GetFullName(),
		},
		ResolutionType: dto.DropdownSimple{
			Id:    resolutionType.Id,
			Title: resolutionType.Title,
		},
		Year:         dataOfStart.Year(),
		NumberOfDays: numberOfDays,
		FileId:       item.FileId,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}, nil
}

var UserProfileAbsentResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		absentSummary dto.AbsentsSummary
		usedDays      int
	)

	profileID := params.Args["user_profile_id"].(int)

	// year ago
	availableDaysOfCurrentYear, availableDaysOfPreviousYear, err := getNumberOfCurrentAndPreviousYearAvailableDays(profileID)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	// get all absents in a period of current year
	currentYear := time.Now().Year()
	startOfYear := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := time.Date(currentYear, time.December, 31, 23, 59, 59, 999999999, time.UTC)
	absents, err := getEmployeeAbsents(profileID, &dto.EmployeeAbsentsInput{From: &startOfYear, To: &endOfYear})
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, absent := range absents {
		absentType, err := getAbsentTypeById(absent.AbsentTypeId)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		absent.AbsentType = *absentType

		if absentType.AccountingDaysOff {
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
		}
	}

	allAbsents, _ := getEmployeeAbsents(profileID, nil)
	for _, absent := range allAbsents {
		if absent.TargetOrganizationUnitID != nil {
			organizationUnit, err := getOrganizationUnitById(*absent.TargetOrganizationUnitID)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			absent.TargetOrganizationUnit = organizationUnit
		}

		absentType, err := getAbsentTypeById(absent.AbsentTypeId)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		absent.AbsentType = *absentType
	}

	absentSummary.CurrentAvailableDays = availableDaysOfCurrentYear
	absentSummary.PastAvailableDays = availableDaysOfPreviousYear
	absentSummary.UsedDays = usedDays

	return dto.Response{
		Status:  "success",
		Message: "Here's the items you asked for!",
		Summary: absentSummary,
		Items:   allAbsents,
	}, nil
}

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

func getNumberOfCurrentAndPreviousYearAvailableDays(profileID int) (int, int, error) {
	currentYear := time.Now().Year()
	startDatePreviousYear := time.Date(currentYear-1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDateCurrentYear := time.Date(currentYear, time.December, 31, 23, 59, 59, 0, time.UTC)

	vacationDays := 0
	pastVacationDays := 0
	resolutions, err := getEmployeeResolutions(profileID, &dto.EmployeeResolutionListInput{From: &startDatePreviousYear, To: &endDateCurrentYear})
	if err != nil {
		fmt.Println("error getting employee resolution - " + err.Error())
	}

	for _, resolution := range resolutions {
		resolutionType, err := getDropdownSettingById(resolution.ResolutionTypeId)
		if err != nil {
			return 0, 0, err
		}
		resolution.ResolutionType = resolutionType
	}

	for _, resolution := range resolutions {
		if resolution.ResolutionType.Value != VacationTypeValue {
			continue
		}

		start, _ := time.Parse(time.RFC3339, resolution.DateOfStart)

		if start.Year() == time.Now().Year() {
			totalDays, _ := strconv.Atoi(resolution.Value)
			vacationDays += totalDays
		}
		if start.Year() == time.Now().Year()-1 {
			totalDays, _ := strconv.Atoi(resolution.Value)
			pastVacationDays += totalDays
		}
	}

	usedDaysPreviousYear, err := calculateUsedDaysOfPreviousYear(profileID)
	if err != nil {
		return 0, 0, err
	}

	pastVacationDays -= usedDaysPreviousYear

	return vacationDays, pastVacationDays, nil
}

func calculateUsedDaysOfPreviousYear(profileID int) (int, error) {
	currentYear := time.Now().Year()
	startDatePreviousYear := time.Date(currentYear-1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDatePreviousYear := time.Date(currentYear-1, time.December, 31, 23, 59, 59, 0, time.UTC)

	// Initialize usedDays variable
	usedDaysPreviousYear := 0

	// Get all absents of the employee in the previous year
	absents, err := getEmployeeAbsents(profileID, &dto.EmployeeAbsentsInput{From: &startDatePreviousYear, To: &endDatePreviousYear})
	if err != nil {
		return 0, err
	}

	for _, absent := range absents {
		start, _ := time.Parse(time.RFC3339, absent.DateOfStart)
		end, _ := time.Parse(time.RFC3339, absent.DateOfEnd)

		absentType, err := getAbsentTypeById(absent.AbsentTypeId)
		if err != nil {
			return 0, err
		}

		if absentType.AccountingDaysOff {
			usedDaysPreviousYear += countWorkingDaysBetweenDates(start, end)
		}
	}

	return usedDaysPreviousYear, nil
}

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

var UserProfileAbsentInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
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
		return shared.ErrorResponse("Bad request: user profile absent data"), nil
	}

	if shared.IsInteger(data.Id) && data.Id != 0 {
		item, err = updateAbsent(data.Id, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
	} else {
		item, err = createAbsent(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
	}

	resItem, err := buildAbsentResponseItem(*item)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	response.Item = resItem

	return response, nil
}

func buildAbsentResponseItem(absent structs.Absent) (*structs.Absent, error) {
	absentType, err := getAbsentTypeById(absent.AbsentTypeId)
	if err != nil {
		return nil, err
	}
	absent.AbsentType = *absentType

	if absent.TargetOrganizationUnitID != nil {
		organizationUnit, err := getOrganizationUnitById(*absent.TargetOrganizationUnitID)
		if err != nil {
			return nil, err
		}
		absent.TargetOrganizationUnit = organizationUnit
	}

	return &absent, nil
}

var UserProfileAbsentDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteAbsent(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

var TerminateEmployment = func(params graphql.ResolveParams) (interface{}, error) {
	userID := params.Args["user_profile_id"].(int)

	user, err := getUserProfileById(userID)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	userResponse, err := buildUserProfileOverviewResponse(user)

	if err != nil {
		return shared.HandleAPIError(err)
	}
	active := true

	if userResponse.IsJudge {
		input := dto.GetJudgeResolutionListInputMS{
			Active: &active,
		}

		resolution, err := getJudgeResolutionList(&input)

		if err != nil {
			return shared.HandleAPIError(err)
		}

		filter := dto.JudgeResolutionsOrganizationUnitInput{
			UserProfileId: &userID,
			ResolutionId:  &resolution.Data[0].Id,
		}

		judge, _, err := getJudgeResolutionOrganizationUnit(&filter)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		err = deleteJJudgeResolutionOrganizationUnit(judge[0].Id)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		contract, err := getEmployeeContracts(userID, nil)
		if err != nil {
			return shared.HandleAPIError(err)
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
			UserProfileId:             userID,
			OrganizationUnitId:        judge[0].OrganizationUnitId,
			Relevant:                  true,
			DateOfStart:               *contract[0].DateOfStart,
			DateOfEnd:                 dateOfEnd,
			AmountOfExperience:        totalMonths,
			AmountOfInsuredExperience: totalMonths,
		}
		_, err = createExperience(&experience)
		if err != nil {
			return shared.HandleAPIError(err)
		}

	} else {
		contract, err := getEmployeeContracts(userID, nil)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		employeeInOrgUnit, err := getEmployeesInOrganizationUnitsByProfileId(userID)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		if employeeInOrgUnit != nil {
			err = deleteEmployeeInOrganizationUnitByID(employeeInOrgUnit.Id)
			if err != nil {
				return shared.HandleAPIError(err)
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
				UserProfileId:             userID,
				OrganizationUnitId:        contract[0].OrganizationUnitID,
				Relevant:                  true,
				DateOfStart:               *contract[0].DateOfStart,
				DateOfEnd:                 dateOfEnd,
				AmountOfExperience:        totalMonths,
				AmountOfInsuredExperience: totalMonths,
			}
			_, err = createExperience(&experience)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			contract[0].JobPositionInOrganizationUnitID = 0
			contract[0].OrganizationUnitDepartmentID = nil
			contract[0].Active = false
			_, err = updateEmployeeContract(contract[0].Id, contract[0])
			if err != nil {
				return shared.HandleAPIError(err)
			}
		}
	}

	active = false
	user.ActiveContract = &active
	_, err = updateUserProfile(user.Id, *user)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	_, err = DeactivateUserAccount(user.UserAccountId)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deactivated this user!",
	}, nil
}

func createAbsent(absent *structs.Absent) (*structs.Absent, error) {
	res := &dto.GetAbsentResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.EMPLOYEE_ABSENTS, absent, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateAbsent(id int, absent *structs.Absent) (*structs.Absent, error) {
	res := &dto.GetAbsentResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.EMPLOYEE_ABSENTS+"/"+strconv.Itoa(id), absent, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteAbsent(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.EMPLOYEE_ABSENTS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getEmployeeAbsents(userProfileID int, input *dto.EmployeeAbsentsInput) ([]*structs.Absent, error) {
	res := &dto.GetAbsentListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(userProfileID)+"/absents", input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getAbsentById(absentID int) (*structs.Absent, error) {
	res := &dto.GetAbsentResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.EMPLOYEE_ABSENTS+"/"+strconv.Itoa(absentID), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
