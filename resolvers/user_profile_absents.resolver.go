package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
)

var UserProfileAbsentResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var absentItems []interface{}
	var absentSummary = map[string]interface{}{
		"current_available_days": 0,
		"past_available_days":    0,
		"used_days":              0,
	}

	profileId := params.Args["user_profile_id"]
	accountId := params.Args["user_account_id"]

	if !shared.IsInteger(profileId) && !shared.IsInteger(accountId) {
		return map[string]interface{}{
			"status":  "error",
			"message": "Argument 'user_profile_id' must not be empty!",
			"item":    nil,
		}, nil
	}

	UserProfilesType := &structs.UserProfiles{}
	UserProfilesData, UserProfilesDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profiles.json", UserProfilesType)

	if UserProfilesDataErr != nil {
		fmt.Printf("Fetching User Profiles failed because of this error - %s.\n", UserProfilesDataErr)
	}

	var UserProfile = shared.FindByProperty(UserProfilesData, "Id", profileId)

	if UserProfile == nil || UserProfile[0] == nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "User Profile not found for provided 'user_profile_id'!",
			"item":    nil,
		}, nil
	}

	var vacationTypes = shared.FetchByProperty(
		"vacation_type",
		"",
		"",
	)
	var relatedVacations = shared.FetchByProperty(
		"vacations",
		"UserProfileId",
		profileId,
	)
	// # Related Employee Vacations
	if len(relatedVacations) > 0 {
		for _, relatedAbsentItem := range relatedVacations {
			var relatedAbsentItemData = shared.WriteStructToInterface(relatedAbsentItem)
			var relatedAbsentType = shared.FindByProperty(vacationTypes, "Id", relatedAbsentItemData["vacation_type_id"])

			if relatedAbsentType != nil && len(relatedAbsentType) > 0 {
				var relatedAbsentTypeData = shared.WriteStructToInterface(relatedAbsentType[0])

				relatedAbsentItemData["vacation_type"] = map[string]interface{}{
					"title": relatedAbsentTypeData["title"],
					"id":    relatedAbsentTypeData["id"],
				}
			}

			absentItems = append(absentItems, relatedAbsentItemData)
		}
	}

	var relatedRelocations = shared.FetchByProperty(
		"relocations",
		"UserProfileId",
		profileId,
	)
	// # Related Employee Relocations
	if len(relatedRelocations) > 0 {
		for _, relatedAbsentItem := range relatedRelocations {
			var relatedAbsentItemData = shared.WriteStructToInterface(relatedAbsentItem)

			absentItems = append(absentItems, relatedAbsentItemData)
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here are the items you asked for!",
		"summary": absentSummary,
		"items":   absentItems,
	}, nil
}

var UserProfileAbsentInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var absentDataItems []interface{}
	var absentItemEndpoint string
	var absentItemType interface{}
	var data structs.Absent
	dataBytes, _ := json.Marshal(params.Args["data"])
	VacationType := &structs.Vacation{}
	RelocationType := &structs.Relocation{}

	json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	vacationTypeId := data.VacationTypeId

	if shared.IsInteger(vacationTypeId) && vacationTypeId != 0 {
		absentItemType = VacationType
		absentItemEndpoint = "user_profile_vacations"
	} else {
		absentItemType = RelocationType
		absentItemEndpoint = "user_profile_relocations"
	}

	absentData, absentDataErr := shared.ReadJson("http://localhost:8080/mocked-data/"+absentItemEndpoint+".json", absentItemType)

	if absentDataErr != nil {
		fmt.Printf("Fetching User Profile's absent items failed because of this error - %s.\n", absentDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		absentDataItems = shared.FilterByProperty(absentData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
		absentDataItems = absentData
	}

	var updatedData = append(absentDataItems, data)

	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/"+absentItemEndpoint+".json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var UserProfileAbsentDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var absentDataItems []interface{}
	var absentItemEndpoint string
	var absentItemType interface{}
	itemId := params.Args["id"]
	vacationTypeId := params.Args["vacation_type_id"]
	VacationType := &structs.Vacation{}
	RelocationType := &structs.Relocation{}
	fmt.Printf("\n vacationTypeId %s\n", vacationTypeId)
	if shared.IsInteger(vacationTypeId) && vacationTypeId != 0 {
		absentItemType = VacationType
		absentItemEndpoint = "user_profile_vacations"
	} else {
		absentItemType = RelocationType
		absentItemEndpoint = "user_profile_relocations"
	}

	absentData, absentDataErr := shared.ReadJson("http://localhost:8080/mocked-data/"+absentItemEndpoint+".json", absentItemType)

	if absentDataErr != nil {
		fmt.Printf("Fetching User Profile's absent items failed because of this error - %s.\n", absentDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		absentDataItems = shared.FilterByProperty(absentData, "Id", itemId)
	} else {
		absentDataItems = absentData
	}

	fmt.Printf("\n absentDataItems %s\n", absentDataItems)
	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/"+absentItemEndpoint+".json"), absentDataItems)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
