package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/graphql-go/graphql"
)

func PopulateRevisionItemProperties(revisions []interface{}, filters ...int) []interface{} {
	var items []interface{}
	var id, organizationUnitId, revisorUserProfileId int

	switch len(filters) {
	case 1:
		id = filters[0]
	case 2:
		id = filters[0]
		organizationUnitId = filters[1]
	case 3:
		id = filters[0]
		organizationUnitId = filters[1]
		revisorUserProfileId = filters[2]
	}

	for _, item := range revisions {
		// # Revision
		itemValue := reflect.ValueOf(item)

		if itemValue.Kind() == reflect.Ptr {
			itemValue = itemValue.Elem()
		}

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}
		// Filtering by Organization Unit
		if shared.IsInteger(organizationUnitId) && organizationUnitId != 0 && organizationUnitId != mergedItem["revision_organization_unit_id"] {
			continue
		}
		// Filtering by Revisor User Profile
		if shared.IsInteger(revisorUserProfileId) && revisorUserProfileId != 0 && revisorUserProfileId != mergedItem["revisor_user_profile_id"] {
			continue
		}
		// # Related Revision Type
		var relatedRevisionType = shared.FetchByProperty(
			"revision_type",
			"Id",
			itemValue.FieldByName("RevisionTypeId").Interface(),
		)

		if len(relatedRevisionType) > 0 {
			relatedRevisionTypeValue := reflect.ValueOf(relatedRevisionType[0])

			if relatedRevisionTypeValue.Kind() == reflect.Ptr {
				relatedRevisionTypeValue = relatedRevisionTypeValue.Elem()
			}

			mergedItem["revision_type"] = map[string]interface{}{
				"title": relatedRevisionTypeValue.FieldByName("Title").Interface(),
				"id":    relatedRevisionTypeValue.FieldByName("Id").Interface(),
			}
		}
		// # Related Revisor User Profile
		var relatedRevisorUserProfile = shared.FetchByProperty(
			"user_profile",
			"Id",
			itemValue.FieldByName("RevisorUserProfileId").Interface(),
		)

		if len(relatedRevisorUserProfile) > 0 {
			var relatedRevisor = shared.WriteStructToInterface(relatedRevisorUserProfile[0])

			mergedItem["revisor_user_profile"] = map[string]interface{}{
				"title": relatedRevisor["first_name"].(string) + " " + relatedRevisor["last_name"].(string),
				"id":    relatedRevisor["id"],
			}
		}
		// # Related Organization Unit
		var relatedOrganizationUnit = shared.FetchByProperty(
			"organization_unit",
			"Id",
			itemValue.FieldByName("RevisionOrganizationUnitId").Interface(),
		)

		if len(relatedOrganizationUnit) > 0 {
			relatedOrganizationUnitValue := reflect.ValueOf(relatedOrganizationUnit[0])

			if relatedOrganizationUnitValue.Kind() == reflect.Ptr {
				relatedOrganizationUnitValue = relatedOrganizationUnitValue.Elem()
			}

			mergedItem["revision_organization_unit"] = map[string]interface{}{
				"title": relatedOrganizationUnitValue.FieldByName("Title").Interface(),
				"id":    relatedOrganizationUnitValue.FieldByName("Id").Interface(),
			}
		}
		// # Responsible User Profile
		var responsibleUserProfileId = itemValue.FieldByName("ResponsibleUserProfileId").Interface().(int)

		if responsibleUserProfileId > 0 {
			var responsibleUserProfile = shared.FetchByProperty(
				"user_profile",
				"Id",
				responsibleUserProfileId,
			)

			if len(responsibleUserProfile) > 0 {
				var responsibleUser = shared.WriteStructToInterface(responsibleUserProfile[0])

				mergedItem["responsible_user_profile"] = map[string]interface{}{
					"title": responsibleUser["first_name"].(string) + " " + responsibleUser["last_name"].(string),
					"id":    responsibleUser["id"],
				}
			}
		} else {
			mergedItem["responsible_user_profile"] = map[string]interface{}{
				"title": itemValue.FieldByName("ResponsibleUserProfile").Interface(),
				"id":    0,
			}
		}
		// # Implementation User Profile
		var implementationUserProfileId = itemValue.FieldByName("ImplementationUserProfileId").Interface().(int)

		if implementationUserProfileId > 0 {
			var implementationUserProfile = shared.FetchByProperty(
				"user_profile",
				"Id",
				implementationUserProfileId,
			)

			if len(implementationUserProfile) > 0 {
				var implementationUser = shared.WriteStructToInterface(implementationUserProfile[0])

				mergedItem["implementation_user_profile"] = map[string]interface{}{
					"title": implementationUser["first_name"].(string) + " " + implementationUser["last_name"].(string),
					"id":    implementationUser["id"],
				}
			}
		} else {
			mergedItem["implementation_user_profile"] = map[string]interface{}{
				"title": itemValue.FieldByName("ImplementationUserProfile").Interface(),
				"id":    0,
			}
		}

		items = append(items, mergedItem)
	}

	return items
}

var RevisionsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var revisors []interface{}
	var total int
	var id int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}
	var organizationUnitId int
	if params.Args["organization_unit_id"] == nil {
		organizationUnitId = 0
	} else {
		organizationUnitId = params.Args["organization_unit_id"].(int)
	}
	var revisorUserProfileId int
	if params.Args["revisor_user_profile_id"] == nil {
		revisorUserProfileId = 0
	} else {
		revisorUserProfileId = params.Args["revisor_user_profile_id"].(int)
	}
	page := params.Args["page"]
	size := params.Args["size"]

	RevisionsType := &structs.Revision{}
	RevisionsData, RevisionsDataErr := shared.ReadJson(shared.GetDataRoot()+"/revisions.json", RevisionsType)

	if RevisionsDataErr != nil {
		fmt.Printf("Fetching Revisions failed because of this error - %s.\n", RevisionsDataErr)
	}

	// Populate data for each Revision with Revision Type, Revision User, Responsible User, Implementation User, Organization Unit
	items = PopulateRevisionItemProperties(RevisionsData, id, organizationUnitId, revisorUserProfileId)

	// All Revisor User Profile
	var revisorUserProfiles = shared.FetchByProperty(
		"user_profile",
		"RevisorRole",
		true,
	)

	if len(revisorUserProfiles) > 0 {
		for _, userProfile := range revisorUserProfiles {
			var profile = shared.WriteStructToInterface(userProfile)
			var revisor = map[string]interface{}{}

			revisor["id"] = profile["id"]
			revisor["title"] = profile["first_name"].(string) + " " + profile["last_name"].(string)

			revisors = append(revisors, revisor)
		}
	}

	total = len(items)

	// Filtering by Pagination params
	if shared.IsInteger(page) && page != 0 && shared.IsInteger(size) && size != 0 {
		items = shared.Pagination(items, page.(int), size.(int))
	}

	return map[string]interface{}{
		"status":   "success",
		"message":  "Here's the list you asked for!",
		"total":    total,
		"revisors": revisors,
		"items":    items,
	}, nil
}

var RevisionResolver = func(params graphql.ResolveParams) (interface{}, error) {
	RevisionType := &structs.Revision{}
	RevisionData, RevisionDataErr := shared.ReadJson(shared.GetDataRoot()+"/revisions.json", RevisionType)

	var id int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	if RevisionDataErr != nil {
		fmt.Printf("Fetching Revisions failed because of this error - %s.\n", RevisionDataErr)
	}

	// Populate data for each Revision with Revision Type, Revision User, Responsible User, Implementation User, Organization Unit
	var items = PopulateRevisionItemProperties(RevisionData, id)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   items,
	}, nil
}

var RevisionInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.Revision
	dataBytes, _ := json.Marshal(params.Args["data"])
	RevisionType := &structs.Revision{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	revisionData, revisionDataErr := shared.ReadJson(shared.GetDataRoot()+"/revisions.json", RevisionType)

	if revisionDataErr != nil {
		fmt.Printf("Fetching Revision failed because of this error - %s.\n", revisionDataErr)
	}

	sliceData := []interface{}{data}
	// Populate data for each Revision with Revision Type, Revision User, Responsible User, Implementation User, Organization Unit
	var populatedData = PopulateRevisionItemProperties(sliceData, data.Id)

	if shared.IsInteger(itemId) && itemId != 0 {
		revisionData = shared.FilterByProperty(revisionData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(revisionData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/revisions.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedData,
	}, nil
}

var RevisionDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	RevisionType := &structs.Revision{}
	revisionData, revisionDataErr := shared.ReadJson(shared.GetDataRoot()+"/revisions.json", RevisionType)

	if revisionDataErr != nil {
		fmt.Printf("Fetching User Profile's Revision failed because of this error - %s.\n", revisionDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		revisionData = shared.FilterByProperty(revisionData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/revisions.json"), revisionData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
