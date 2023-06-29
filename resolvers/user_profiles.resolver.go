package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/graphql-go/graphql"
)

func UpdateRelatedUserAccount(userAccountId int, newData map[string]interface{}) map[string]interface{} {
	var projectRoot, _ = shared.GetProjectRoot()
	var allUserAccounts = shared.FetchByProperty("user_account", "", 0)
	var relatedUserAccount = shared.FindByProperty(allUserAccounts, "Id", userAccountId)

	// # Related User Account
	if len(relatedUserAccount) > 0 {
		var userAccountData = shared.WriteStructToInterface(relatedUserAccount[0])

		allUserAccounts = shared.FilterByProperty(allUserAccounts, "Id", userAccountData["id"])

		newData = shared.MergeMaps(userAccountData, newData, true)
	} else {
		newData["id"] = shared.GetRandomNumber()
		newData["password"] = "test1234"
		newData["pin"] = "1234"
		newData["active"] = true
	}

	var updatedData = append(allUserAccounts, newData)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_accounts.json"), updatedData)

	return newData
}

var UserProfilesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int

	id := params.Args["id"]
	active := params.Args["is_active"]
	organizationUnitId := params.Args["organization_unit_id"]
	jobPositionId := params.Args["job_position_id"]
	name := params.Args["name"]
	page := params.Args["page"]
	size := params.Args["size"]

	UserProfilesType := &structs.UserProfiles{}
	UserProfilesData, UserProfilesDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profiles.json", UserProfilesType)

	if UserProfilesDataErr != nil {
		fmt.Printf("Fetching User Profiles failed because of this error - %s.\n", UserProfilesDataErr)
	}

	// Fetch User Account data for each User Profile
	for _, item := range UserProfilesData {
		mergedItem := make(map[string]interface{})
		// # User Profile
		itemValue := reflect.ValueOf(item)

		if itemValue.Kind() == reflect.Ptr {
			itemValue = itemValue.Elem()
		}

		mergedItem["id"] = itemValue.FieldByName("Id").Interface()
		mergedItem["first_name"] = itemValue.FieldByName("FirstName").Interface()
		mergedItem["last_name"] = itemValue.FieldByName("LastName").Interface()
		mergedItem["date_of_birth"] = itemValue.FieldByName("DateOfBirth").Interface()
		mergedItem["created_at"] = itemValue.FieldByName("CreatedAt").Interface()
		mergedItem["updated_at"] = itemValue.FieldByName("UpdatedAt").Interface()

		// Filtering by name search
		if shared.IsString(name) && len(name.(string)) > 0 {
			UserProfileName := mergedItem["first_name"].(string) + mergedItem["last_name"].(string)

			if !shared.StringContains(UserProfileName, name.(string)) {
				continue
			}
		}

		// # Related User Account
		var relatedUserAccount = shared.FetchByProperty(
			"user_account",
			"Id",
			itemValue.FieldByName("UserAccountId").Interface(),
		)

		if len(relatedUserAccount) < 1 {
			// Filtering by Organization Unit or Job Position
			if (shared.IsInteger(organizationUnitId) && organizationUnitId != 0) || (shared.IsInteger(jobPositionId) && jobPositionId != 0) {
				continue
			}
		} else {
			relatedUserAccountValue := reflect.ValueOf(relatedUserAccount[0])

			if relatedUserAccountValue.Kind() == reflect.Ptr {
				relatedUserAccountValue = relatedUserAccountValue.Elem()
			}

			mergedItem["email"] = relatedUserAccountValue.FieldByName("Email").Interface()
			mergedItem["phone"] = relatedUserAccountValue.FieldByName("Phone").Interface()
			mergedItem["active"] = relatedUserAccountValue.FieldByName("Active").Interface()

			// # Related User Account Role
			var relatedRole = shared.FetchByProperty(
				"role",
				"Id",
				relatedUserAccountValue.FieldByName("RoleId").Interface(),
			)

			if len(relatedRole) < 1 {
				//fmt.Printf("Fetching relatedRole returned empty array. User Profile data - %s", item)
				// Filtering by Organization Unit or Job Position
				if (shared.IsInteger(organizationUnitId) && organizationUnitId != 0) || (shared.IsInteger(jobPositionId) && jobPositionId != 0) {
					continue
				}
			} else {
				relatedUserAccountRoleValue := reflect.ValueOf(relatedRole[0])

				if relatedUserAccountRoleValue.Kind() == reflect.Ptr {
					relatedUserAccountRoleValue = relatedUserAccountRoleValue.Elem()
				}

				mergedItem["role"] = relatedUserAccountRoleValue.FieldByName("Title").Interface()
			}

			// # Employee in Organization Unit data
			var EmployeeInOrganizationUnit = shared.FetchByProperty(
				"employees_in_organization_units",
				"UserAccountId",
				itemValue.FieldByName("UserAccountId").Interface(),
			)

			if len(EmployeeInOrganizationUnit) < 1 {
				//fmt.Printf("Fetching EmployeeInOrganizationUnit returned empty array. User Profile data - %s", item)
				// Filtering by Organization Unit or Job Position
				if (shared.IsInteger(organizationUnitId) && organizationUnitId != 0) || (shared.IsInteger(jobPositionId) && jobPositionId != 0) {
					continue
				}
			} else {
				EmployeeInOrganizationUnitValue := reflect.ValueOf(EmployeeInOrganizationUnit[0])

				if EmployeeInOrganizationUnitValue.Kind() == reflect.Ptr {
					EmployeeInOrganizationUnitValue = EmployeeInOrganizationUnitValue.Elem()
				}

				// # Job Position in Organization Unit data
				var JobPositionInOrganizationUnit = shared.FetchByProperty(
					"job_positions_in_organization_units",
					"Id",
					EmployeeInOrganizationUnitValue.FieldByName("PositionInOrganizationUnitId").Interface(),
				)

				if len(JobPositionInOrganizationUnit) < 1 {
					//fmt.Printf("Fetching JobPositionInOrganizationUnit returned empty array. User Profile data - %s", item)
					// Filtering by Organization Unit or Job Position
					if (shared.IsInteger(organizationUnitId) && organizationUnitId != 0) || (shared.IsInteger(jobPositionId) && jobPositionId != 0) {
						continue
					}
				} else {
					JobPositionInOrganizationUnitValue := reflect.ValueOf(JobPositionInOrganizationUnit[0])

					if JobPositionInOrganizationUnitValue.Kind() == reflect.Ptr {
						JobPositionInOrganizationUnitValue = JobPositionInOrganizationUnitValue.Elem()
					}

					// # Related Organization Unit
					var OrganizationUnit = shared.FetchByProperty(
						"organization_unit",
						"Id",
						JobPositionInOrganizationUnitValue.FieldByName("ParentOrganizationUnitId").Interface(),
					)

					if len(OrganizationUnit) < 1 {
						//fmt.Printf("Fetching OrganizationUnit returned empty array. User Profile data - %s", item)
						// Filtering by Organization Unit
						if shared.IsInteger(organizationUnitId) && organizationUnitId != 0 {
							continue
						}
					} else {
						OrganizationUnitValue := reflect.ValueOf(OrganizationUnit[0])

						if OrganizationUnitValue.Kind() == reflect.Ptr {
							OrganizationUnitValue = OrganizationUnitValue.Elem()
						}

						mergedItem["organization_unit"] = OrganizationUnitValue.FieldByName("Title").Interface()

						if shared.IsInteger(organizationUnitId) && organizationUnitId != 0 && OrganizationUnitValue.FieldByName("Id").Interface() != organizationUnitId {
							continue
						}
					}

					// # Related Job Position
					var JobPosition = shared.FetchByProperty(
						"job_positions",
						"Id",
						JobPositionInOrganizationUnitValue.FieldByName("JobPositionId").Interface(),
					)

					if len(JobPosition) < 1 {
						//fmt.Printf("Fetching JobPosition returned empty array. User Profile data - %s", item)
						// Filtering by Job Position
						if shared.IsInteger(jobPositionId) && jobPositionId != 0 {
							continue
						}
					} else {
						JobPositionValue := reflect.ValueOf(JobPosition[0])

						if JobPositionValue.Kind() == reflect.Ptr {
							JobPositionValue = JobPositionValue.Elem()
						}

						mergedItem["job_position"] = JobPositionValue.FieldByName("Title").Interface()
						mergedItem["is_judge"] = JobPositionValue.FieldByName("IsJudge").Interface()
						mergedItem["is_judge_president"] = JobPositionValue.FieldByName("IsJudgePresident").Interface()

						if shared.IsInteger(jobPositionId) && jobPositionId != 0 && JobPositionValue.FieldByName("Id").Interface() != jobPositionId {
							continue
						}
					}
				}
			}
		}

		items = append(items, mergedItem)
	}

	// Filtering by ID
	if shared.IsInteger(id) && id != 0 {
		items = shared.FindByProperty(items, "id", id)
	}
	// Filtering by User Account status
	if active == true || active == false {
		items = shared.FindByProperty(items, "active", active)
	}

	total = len(items)

	// Filtering by Pagination params
	if shared.IsInteger(page) && page != 0 && shared.IsInteger(size) && size != 0 {
		items = shared.Pagination(items, page.(int), size.(int))
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"total":   total,
		"items":   items,
	}, nil
}

var UserProfileBasicResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

	var item = shared.WriteStructToInterface(UserProfile[0])

	accountId = item["user_account_id"]

	var relatedUserAccount = shared.FetchByProperty(
		"user_account",
		"Id",
		accountId,
	)

	// # Related User Account
	if len(relatedUserAccount) > 0 {
		relatedUserAccountValue := reflect.ValueOf(relatedUserAccount[0])

		if relatedUserAccountValue.Kind() == reflect.Ptr {
			relatedUserAccountValue = relatedUserAccountValue.Elem()
		}

		item["email"] = relatedUserAccountValue.FieldByName("Email").Interface()
		item["phone"] = relatedUserAccountValue.FieldByName("Phone").Interface()

		// # Employee in Organization Unit data
		var EmployeeInOrganizationUnit = shared.FetchByProperty(
			"employees_in_organization_units",
			"UserAccountId",
			accountId,
		)

		if len(EmployeeInOrganizationUnit) > 0 {
			EmployeeInOrganizationUnitValue := reflect.ValueOf(EmployeeInOrganizationUnit[0])

			if EmployeeInOrganizationUnitValue.Kind() == reflect.Ptr {
				EmployeeInOrganizationUnitValue = EmployeeInOrganizationUnitValue.Elem()
			}

			// # Job Position in Organization Unit data
			var JobPositionInOrganizationUnit = shared.FetchByProperty(
				"job_positions_in_organization_units",
				"Id",
				EmployeeInOrganizationUnitValue.FieldByName("PositionInOrganizationUnitId").Interface(),
			)

			if len(JobPositionInOrganizationUnit) > 0 {
				JobPositionInOrganizationUnitValue := reflect.ValueOf(JobPositionInOrganizationUnit[0])

				if JobPositionInOrganizationUnitValue.Kind() == reflect.Ptr {
					JobPositionInOrganizationUnitValue = JobPositionInOrganizationUnitValue.Elem()
				}

				// # Related Organization Unit
				var OrganizationUnit = shared.FetchByProperty(
					"organization_unit",
					"Id",
					JobPositionInOrganizationUnitValue.FieldByName("ParentOrganizationUnitId").Interface(),
				)

				if len(OrganizationUnit) > 0 {
					OrganizationUnitValue := reflect.ValueOf(OrganizationUnit[0])

					if OrganizationUnitValue.Kind() == reflect.Ptr {
						OrganizationUnitValue = OrganizationUnitValue.Elem()
					}

					OrganizationUnitData := map[string]interface{}{
						"title": OrganizationUnitValue.FieldByName("Title").Interface(),
						"id":    OrganizationUnitValue.FieldByName("Id").Interface(),
					}

					item["organization_unit"] = OrganizationUnitData
				}

				// # Related Job Position
				var JobPosition = shared.FetchByProperty(
					"job_positions",
					"Id",
					JobPositionInOrganizationUnitValue.FieldByName("JobPositionId").Interface(),
				)

				if len(JobPosition) > 0 {
					JobPositionValue := reflect.ValueOf(JobPosition[0])

					if JobPositionValue.Kind() == reflect.Ptr {
						JobPositionValue = JobPositionValue.Elem()
					}

					JobPositionData := map[string]interface{}{
						"title": JobPositionValue.FieldByName("Title").Interface(),
						"id":    JobPositionValue.FieldByName("Id").Interface(),
					}

					item["job_position"] = JobPositionData
				}
			}
		}
	}

	var relatedEmployeeContract = shared.FetchByProperty(
		"employee_contract",
		"Id",
		profileId,
	)

	// # Related Employee Contract
	if len(relatedEmployeeContract) > 0 {
		relatedEmployeeContractValue := reflect.ValueOf(relatedEmployeeContract[0])

		if relatedEmployeeContractValue.Kind() == reflect.Ptr {
			relatedEmployeeContractValue = relatedEmployeeContractValue.Elem()
		}

		item["date_of_start"] = relatedEmployeeContractValue.FieldByName("DateOfStart").Interface()
		item["date_of_end"] = relatedEmployeeContractValue.FieldByName("DateOfEnd").Interface()

		// # Employee in Organization Unit data
		var relatedEmployeeContractType = shared.FetchByProperty(
			"contract_type",
			"Id",
			relatedEmployeeContractValue.FieldByName("ContractTypeId").Interface(),
		)

		if len(relatedEmployeeContractType) > 0 {
			relatedEmployeeContractTypeValue := reflect.ValueOf(relatedEmployeeContractType[0])

			if relatedEmployeeContractTypeValue.Kind() == reflect.Ptr {
				relatedEmployeeContractTypeValue = relatedEmployeeContractTypeValue.Elem()
			}

			ContractTypeData := map[string]interface{}{
				"title": relatedEmployeeContractTypeValue.FieldByName("Title").Interface(),
				"id":    relatedEmployeeContractTypeValue.FieldByName("Id").Interface(),
			}

			item["contract_type"] = ContractTypeData
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the item you asked for!",
		"item":    item,
	}, nil
}

var UserProfileBasicInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data map[string]interface{}
	var dataStruct structs.UserProfiles
	var newData map[string]interface{}
	var userAccountId int
	dataBytes, _ := json.Marshal(params.Args["data"])
	UserProfileBasicType := &structs.UserProfiles{}

	_ = json.Unmarshal(dataBytes, &data)
	_ = json.Unmarshal(dataBytes, &dataStruct)

	itemId := dataStruct.Id
	userProfileBasicData, userProfileBasicDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profiles.json", UserProfileBasicType)

	if userProfileBasicDataErr != nil {
		fmt.Printf("Fetching User Profile's Basic data failed because of this error - %s.\n", userProfileBasicDataErr)
	}

	// Edit existing User Profile and User Account
	if shared.IsInteger(itemId) && itemId != 0 {
		var editedUserProfile = shared.FindByProperty(userProfileBasicData, "Id", itemId)
		userProfileBasicData = shared.FilterByProperty(userProfileBasicData, "Id", itemId)

		if len(editedUserProfile) > 0 {
			var editedUserProfileData = shared.WriteStructToInterface(editedUserProfile[0])
			userAccountId = editedUserProfileData["user_account_id"].(int)

			newData = shared.MergeMaps(editedUserProfileData, data)
		}
	} else {
		// Create User Profile and User Account
		data["id"] = shared.GetRandomNumber()
		newData = data
		userAccountId = 0
	}

	var newUserAccountData = shared.WriteStructToInterface(structs.UserAccounts{
		Email: newData["email"].(string),
		Phone: newData["phone"].(string),
	})
	var updatedUserAccountData = UpdateRelatedUserAccount(userAccountId, newUserAccountData)

	newData["user_account_id"] = updatedUserAccountData["id"]

	var updatedData = append(userProfileBasicData, newData)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profiles.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    newData,
	}, nil
}

var UserProfileEducationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var educationGroups []interface{}

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

	var educationTypes = shared.FetchByProperty(
		"education_type",
		"",
		"",
	)

	for _, educationTypeItem := range educationTypes {
		var educationType = shared.WriteStructToInterface(educationTypeItem)
		var educationTypeItems []interface{}

		educationType["items"] = educationTypeItems
		educationGroups = append(educationGroups, educationType)
	}

	var relatedEducation = shared.FetchByProperty(
		"education",
		"UserProfileId",
		profileId,
	)

	// # Related Employee Education
	if len(relatedEducation) > 0 {
		for _, relatedEducationItem := range relatedEducation {
			var educationGroupItem = shared.WriteStructToInterface(relatedEducationItem)

			if educationGroupItem != nil && educationGroupItem["education_type_id"] != nil {
				educationGroups = shared.AppendByProperty(educationGroups, "id", educationGroupItem["education_type_id"], "items", educationGroupItem)
			}
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the item you asked for!",
		"items":   educationGroups,
	}, nil
}

var UserProfileEducationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.Education
	dataBytes, _ := json.Marshal(params.Args["data"])
	EducationType := &structs.Education{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	educationData, educationDataErr := shared.ReadJson("http://localhost:8080/mocked-data/educations.json", EducationType)

	if educationDataErr != nil {
		fmt.Printf("Fetching User Profile's education failed because of this error - %s.\n", educationDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		educationData = shared.FilterByProperty(educationData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(educationData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/educations.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var UserProfileEducationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	EducationType := &structs.Education{}
	educationData, educationDataErr := shared.ReadJson("http://localhost:8080/mocked-data/educations.json", EducationType)

	if educationDataErr != nil {
		fmt.Printf("Fetching User Profile's Education failed because of this error - %s.\n", educationDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		educationData = shared.FilterByProperty(educationData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/educations.json"), educationData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

var UserProfileExperienceResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

	var experienceItems = shared.FetchByProperty(
		"experience",
		"UserProfileId",
		profileId,
	)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the item you asked for!",
		"items":   experienceItems,
	}, nil
}

var UserProfileExperienceInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.Experience
	dataBytes, _ := json.Marshal(params.Args["data"])
	ExperienceType := &structs.Experience{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	ExperienceData, ExperienceDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profile_experiences.json", ExperienceType)

	if ExperienceDataErr != nil {
		fmt.Printf("Fetching User Profile's Experience failed because of this error - %s.\n", ExperienceDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		ExperienceData = shared.FilterByProperty(ExperienceData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(ExperienceData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_experiences.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var UserProfileExperienceDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	ExperienceType := &structs.Experience{}
	ExperienceData, ExperienceDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profile_experiences.json", ExperienceType)

	if ExperienceDataErr != nil {
		fmt.Printf("Fetching User Profile's Experience failed because of this error - %s.\n", ExperienceDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		ExperienceData = shared.FilterByProperty(ExperienceData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_experiences.json"), ExperienceData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

var UserProfileFamilyResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

	var familyItems = shared.FetchByProperty(
		"family",
		"UserProfileId",
		profileId,
	)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the item you asked for!",
		"items":   familyItems,
	}, nil
}

var UserProfileFamilyInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.Family
	dataBytes, _ := json.Marshal(params.Args["data"])
	FamilyType := &structs.Family{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	FamilyData, FamilyDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profile_family.json", FamilyType)

	if FamilyDataErr != nil {
		fmt.Printf("Fetching User Profile's Family failed because of this error - %s.\n", FamilyDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		FamilyData = shared.FilterByProperty(FamilyData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(FamilyData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_family.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var UserProfileFamilyDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	FamilyType := &structs.Family{}
	FamilyData, FamilyDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profile_family.json", FamilyType)

	if FamilyDataErr != nil {
		fmt.Printf("Fetching User Profile's Family failed because of this error - %s.\n", FamilyDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		FamilyData = shared.FilterByProperty(FamilyData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_family.json"), FamilyData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
