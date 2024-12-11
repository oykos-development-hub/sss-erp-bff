package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"fmt"
	"strconv"
)

func (repo *MicroserviceRepository) GetAllInventoryDispatches(filter dto.InventoryDispatchFilter) (*dto.GetAllBasicInventoryDispatches, error) {
	res := &dto.GetAllBasicInventoryDispatches{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Dispatch, filter, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetMyInventoryDispatchesItems(filter *dto.DispatchInventoryItemFilter) ([]*structs.BasicInventoryDispatchItemsItem, error) {
	res := &dto.GetAllBasicInventoryDispatchItems{}

	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.DispatchItems, filter, res)

	if err != nil {
		fmt.Printf("Fetching Inventory items failed because of this error - %s.\n", err)
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateDispatchItem(ctx context.Context, item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error) {
	res := dto.GetBasicInventoryDispatch{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.Dispatch, item, &res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	if item.InventoryID != nil {
		for i := 0; i < len(item.InventoryID); i++ {
			itemDispatch := structs.BasicInventoryDispatchItemsItem{
				InventoryID: item.InventoryID[i],
				DispatchID:  res.Data.ID,
			}

			if item.Type != "revers" && item.Type != "return-revers" && item.Type != "created" {
				inventory, err := repo.GetInventoryItem(item.InventoryID[i])
				if err != nil {
					return nil, errors.Wrap(err, "repo get inventory item")
				}

				targetUserProfileID := 0
				OfficeID := 0

				if item.Type == "allocation" {
					targetUserProfileID = item.TargetUserProfileID
					OfficeID = item.OfficeID
				}
				if item.Type == "return" {
					page := 1
					size := 1000
					search := "Lager"
					organizationUnitID := ""
					if item.TargetOrganizationUnitID != 0 {
						organizationUnitID = strconv.Itoa(item.TargetOrganizationUnitID)
					} else {
						organizationUnitID = strconv.Itoa(item.SourceOrganizationUnitID)
					}

					input := dto.GetOfficesOfOrganizationInput{Page: &page, Size: &size, Search: &search, Value: &organizationUnitID}

					office, err := repo.GetOfficeDropdownSettings(&input)
					if err != nil {
						return nil, errors.Wrap(err, "repo get office dropdown settings")
					}
					if len(office.Data) > 0 {
						OfficeID = office.Data[0].ID
					}
				}

				inventory.TargetUserProfileID = targetUserProfileID
				inventory.OfficeID = OfficeID

				_, err = repo.UpdateInventoryItem(ctx, inventory.ID, inventory)
				if err != nil {
					return nil, errors.Wrap(err, "repo upate inventory item")
				}

			}

			if item.Type == "return-revers" {
				inventory, err := repo.GetInventoryItem(item.InventoryID[i])
				if err != nil {
					return nil, errors.Wrap(err, "repo get inventory item")
				}

				page := 1
				size := 1000
				search := "Lager"
				organizationUnitID := strconv.Itoa(item.SourceOrganizationUnitID)

				input := dto.GetOfficesOfOrganizationInput{Page: &page, Size: &size, Search: &search, Value: &organizationUnitID}

				office, err := repo.GetOfficeDropdownSettings(&input)
				if err != nil {
					return nil, errors.Wrap(err, "repo get office dropdown settings")
				}

				inventory.TargetUserProfileID = 0
				if len(office.Data) > 0 {
					inventory.OfficeID = office.Data[0].ID
				}

				_, err = repo.UpdateInventoryItem(ctx, inventory.ID, inventory)
				if err != nil {
					return nil, errors.Wrap(err, "repo update inventory item")
				}
			}

			_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.DispatchItems, itemDispatch, nil)
			if err != nil {
				return nil, errors.Wrap(err, "make api request")
			}
		}
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateDispatchItem(ctx context.Context, id int, item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error) {
	dispatch := dto.GetBasicInventoryDispatch{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Inventory.Dispatch+"/"+strconv.Itoa(id), item, &dispatch, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	if item.InventoryID != nil {
		filter := dto.DispatchInventoryItemFilter{
			DispatchID: &item.ID,
		}

		dispatchItems, _ := repo.GetMyInventoryDispatchesItems(&filter)

		for _, dispatch := range dispatchItems {
			_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Inventory.DispatchItems+"/"+strconv.Itoa(dispatch.ID), nil, nil)

			if err != nil {
				return nil, errors.Wrap(err, "make api request")
			}
		}

		for _, inventoryID := range item.InventoryID {
			itemDispatch := structs.BasicInventoryDispatchItemsItem{
				InventoryID: inventoryID,
				DispatchID:  dispatch.Data.ID,
			}

			_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.DispatchItems, itemDispatch, nil)
			if err != nil {
				return nil, errors.Wrap(err, "make api request")
			}
		}
	}

	return dispatch.Data, nil
}

func (repo *MicroserviceRepository) DeleteInventoryDispatch(ctx context.Context, id int) error {

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Inventory.Dispatch+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) CreateDispatchItemItem(item *structs.BasicInventoryDispatchItemsItem) (*structs.BasicInventoryDispatchItemsItem, error) {
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.DispatchItems, item, nil)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return item, err
}

func (repo *MicroserviceRepository) GetAllInventoryItemInOrgUnits(id int) ([]dto.GetAllItemsInOrgUnits, error) {
	res := dto.GetAllItemsInOrgUnitsMS{}

	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.ItemsInOrgUnit+"/"+strconv.Itoa(id), nil, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetDispatchItemByID(id int) (*structs.BasicInventoryDispatchItem, error) {
	res := dto.GetBasicInventoryDispatch{}

	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Dispatch+"/"+strconv.Itoa(id), nil, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateAssessments(ctx context.Context, data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.Assessments, data, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateAssessments(ctx context.Context, id int, data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Inventory.Assessments+"/"+strconv.Itoa(id), data, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteAssessment(ctx context.Context, id int) error {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Inventory.Assessments+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) CreateInventoryItem(ctx context.Context, item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.Item, item, &res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	if item.RealEstate != nil {
		item.RealEstate.ItemID = res.Data.ID
		_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.RealEstates, item.RealEstate, nil)
		if err != nil {
			return nil, errors.Wrap(err, "make api request")
		}
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateInventoryItem(ctx context.Context, id int, item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}
	res1 := dto.GetBasicInventoryInsertItem{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Inventory.Item+"/"+strconv.Itoa(id), item, &res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	if item.RealEstate != nil {
		item.RealEstate.ItemID = res.Data.ID
		if item.RealEstate.ID != 0 {
			_, err := makeAPIRequest("PUT", repo.Config.Microservices.Inventory.RealEstates+"/"+strconv.Itoa(item.RealEstate.ID), item.RealEstate, &res1)
			if err != nil {
				return nil, errors.Wrap(err, "make api request")
			}
		} else {
			_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.RealEstates, item.RealEstate, &res1)
			if err != nil {
				return nil, errors.Wrap(err, "make api request")
			}
		}
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetInventoryItem(id int) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Item+"/"+strconv.Itoa(id), nil, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}
func (repo *MicroserviceRepository) GetAllInventoryItem(filter dto.InventoryItemFilter) (*dto.GetAllBasicInventoryItem, error) {
	res := &dto.GetAllBasicInventoryItem{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Item, filter, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res, nil
}
func (repo *MicroserviceRepository) GetMyInventoryRealEstate(id int) (*structs.BasicInventoryRealEstatesItem, error) {
	res := &dto.GetMyInventoryRealEstateResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Base+"/item/"+strconv.Itoa(id)+"/real-estates", nil, res)

	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetMyInventoryAssessments(id int) ([]structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseArrayMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Assessments+"/"+strconv.Itoa(id)+"/item", nil, res)

	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok && apiErr.StatusCode != 404 {
			return nil, errors.Wrap(err, "make api request")
		}
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetDispatchItemByInventoryID(id int) ([]*structs.BasicInventoryDispatchItemsItem, error) {
	res1 := dto.GetAllBasicInventoryDispatchItems{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Base+"/item/"+strconv.Itoa(id)+"/dispatch-items", nil, &res1)

	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res1.Data, nil
}

func (repo *MicroserviceRepository) GetInventoryItemsByDispatch(dispatchID int) ([]*structs.BasicInventoryInsertItem, error) {
	res := dto.GetAllBasicInventoryItem{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Dispatch+"/"+strconv.Itoa(dispatchID)+"/items", nil, &res)

	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetInventoryRealEstatesList(input *dto.GetInventoryRealEstateListInputMS) (*dto.GetInventoryRealEstateListResponseMS, error) {
	res := &dto.GetInventoryRealEstateListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.RealEstates, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetInventoryRealEstate(id int) (*structs.BasicInventoryRealEstatesItem, error) {
	res := &dto.GetInventoryRealEstateResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.RealEstates+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetAllInventoryItemForReport(filter dto.ItemReportFilterDTO) ([]dto.ItemReportResponse, error) {
	res := &dto.GetAllItemsReportMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.ItemsReport, filter, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CheckInsertInventoryData(input []structs.BasicInventoryInsertItem) ([]structs.BasicInventoryInsertValidator, error) {
	inventoryMap := make(map[string]bool)
	serialMap := make(map[string]bool)
	var items []structs.BasicInventoryInsertValidator

	for _, item := range input {
		if serialMap[item.SerialNumber] && item.SerialNumber != "" {
			items = append(items, structs.BasicInventoryInsertValidator{
				Entity: "serial_number",
				Value:  item.SerialNumber,
			})
		}
		if inventoryMap[item.InventoryNumber] && item.InventoryNumber != "" {
			items = append(items, structs.BasicInventoryInsertValidator{
				Entity: "inventory_number",
				Value:  item.InventoryNumber,
			})
		}

		inventoryMap[item.InventoryNumber] = true
		serialMap[item.SerialNumber] = true
	}

	for _, item := range input {
		inventoryItem, err := repo.GetAllInventoryItem(dto.InventoryItemFilter{
			SerialNumber:       &item.SerialNumber,
			OrganizationUnitID: &item.OrganizationUnitID,
		})

		if err != nil {
			return nil, errors.Wrap(err, "repo get all inventory item")
		}

		if len(inventoryItem.Data) != 0 && inventoryItem.Data[0].ID != item.ID && item.SerialNumber != "" {
			items = append(items, structs.BasicInventoryInsertValidator{
				Entity: "serial_number",
				Value:  inventoryItem.Data[0].SerialNumber,
			})
		}

		inventoryItem, err = repo.GetAllInventoryItem(dto.InventoryItemFilter{
			InventoryNumber:    &item.InventoryNumber,
			OrganizationUnitID: &item.OrganizationUnitID,
		})

		if err != nil {
			return nil, errors.Wrap(err, "repo get all inventory item")
		}

		if len(inventoryItem.Data) != 0 && inventoryItem.Data[0].ID != item.ID && item.InventoryNumber != "" {
			items = append(items, structs.BasicInventoryInsertValidator{
				Entity: "inventory_number",
				Value:  inventoryItem.Data[0].InventoryNumber,
			})
		}
	}
	return items, nil
}

func (repo *MicroserviceRepository) CreateExcelInventoryItems(ctx context.Context, items []structs.ImportInventoryArticles) error {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.ExcelItems, items, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}
