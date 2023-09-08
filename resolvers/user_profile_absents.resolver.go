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

var UserProfileAbsentResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		absentSummary dto.AbsentsSummary
		usedDays      int
	)

	profileId := params.Args["user_profile_id"].(int)

	absents, err := getEmployeeAbsents(profileId, nil)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	currentAvailableDays, previousYearAvailableDays, err := getNumberOfCurrentAndPreviousYearAvailableDays(profileId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, absent := range absents {
		if absent.TargetOrganizationUnitID != 0 {
			organizationUnit, err := getOrganizationUnitById(absent.TargetOrganizationUnitID)
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

		if absentType.AccountingDaysOff {
			daysTakenBeforeJuly, daysTakenAfterJuly := getTakenVacationDaysBeforeAndAfterJuly(absent.DateOfStart, absent.DateOfEnd)

			// Subtract vacation days taken before July from available previous year days
			if daysTakenBeforeJuly > 0 {
				if previousYearAvailableDays >= daysTakenBeforeJuly {
					previousYearAvailableDays -= daysTakenBeforeJuly
				} else {
					// if available days from previous year are not enough, we should use current year too
					currentAvailableDays -= daysTakenBeforeJuly - previousYearAvailableDays
					previousYearAvailableDays = 0
				}
			}
			// Subtract vacation days taken after July from available current year vacation days
			if daysTakenAfterJuly > 0 {
				currentAvailableDays -= daysTakenAfterJuly
			}

			usedDays += (daysTakenBeforeJuly + daysTakenAfterJuly)
		}
	}

	absentSummary.CurrentAvailableDays = currentAvailableDays
	absentSummary.PastAvailableDays = previousYearAvailableDays
	absentSummary.UsedDays = usedDays

	return dto.Response{
		Status:  "success",
		Message: "Here's the items you asked for!",
		Summary: absentSummary,
		Items:   absents,
	}, nil
}

func buildAbsentResponseItem(absent structs.Absent) (*structs.Absent, error) {
	absentType, err := getAbsentTypeById(absent.AbsentTypeId)
	if err != nil {
		return nil, err
	}
	absent.AbsentType = *absentType

	if absent.TargetOrganizationUnitID != 0 {
		organizationUnit, err := getOrganizationUnitById(absent.TargetOrganizationUnitID)
		if err != nil {
			return nil, err
		}
		absent.TargetOrganizationUnit = organizationUnit
	}

	return &absent, nil
}

func getTakenVacationDaysBeforeAndAfterJuly(start time.Time, end time.Time) (int, int) {
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
	vacationDays := 0
	pastVacationDays := 0
	resolutions, err := getEmployeeResolutions(profileID)
	if err != nil {
		fmt.Println("error hydrating resolution types - " + err.Error())
	}
	for _, resolution := range resolutions {
		resolutionType, err := getDropdownSettingById(resolution.ResolutionTypeId)
		if err != nil {
			return 0, 0, err
		}
		resolution.ResolutionType = &structs.SettingsDropdown{Id: resolutionType.Id, Title: resolutionType.Title}
	}

	for _, resolution := range resolutions {
		if resolution.DateOfStart.Year() != time.Now().Year() {
			continue
		}
		if resolution.ResolutionType.Value == "vacation" {
			vacationDays += getNumberOfWorkingDays(resolution.DateOfStart, resolution.DateOfEnd)
		} else if resolution.ResolutionType.Value == "vacation_past" {
			pastVacationDays += getNumberOfWorkingDays(resolution.DateOfStart, resolution.DateOfEnd)
		}
	}
	return vacationDays, pastVacationDays, nil
}

func getNumberOfWorkingDays(start time.Time, end time.Time) int {
	workingDays := 0

	for !start.After(end) {
		if start.Weekday() != time.Saturday && start.Weekday() != time.Sunday {
			workingDays++
		}
		start = start.AddDate(0, 0, 1)
	}

	return workingDays
}

var UserProfileAbsentInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var data structs.Absent

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
		item, err := updateAbsent(data.Id, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := createAbsent(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
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
