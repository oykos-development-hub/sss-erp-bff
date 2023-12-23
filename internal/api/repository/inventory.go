package repository

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"fmt"
	"strconv"
)

func (repo *MicroserviceRepository) GetAllInventoryDispatches(filter dto.InventoryDispatchFilter) (*dto.GetAllBasicInventoryDispatches, error) {
	res := &dto.GetAllBasicInventoryDispatches{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Dispatch, filter, &res)
	if err != nil {
		return nil, err
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

func (repo *MicroserviceRepository) CreateDispatchItem(item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error) {
	res := dto.GetBasicInventoryDispatch{}

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.Dispatch, item, &res)
	if err != nil {
		return nil, err
	}

	if item.InventoryID != nil {
		for i := 0; i < len(item.InventoryID); i++ {
			itemDispatch := structs.BasicInventoryDispatchItemsItem{
				InventoryID: item.InventoryID[i],
				DispatchID:  res.Data.ID,
			}

			if item.Type != "revers" {
				inventory, err := repo.GetInventoryItem(item.InventoryID[i])
				if err != nil {
					return nil, err
				}

				targetOrganizationUnitID := 0
				targetUserProfileID := 0
				OfficeID := 0

				if item.Type == "allocation" {
					targetUserProfileID = item.TargetUserProfileID
					OfficeID = item.OfficeID
				}
				if item.Type == "return" {
					inventory.TargetOrganizationUnitID = targetOrganizationUnitID
				}

				inventory.TargetUserProfileID = targetUserProfileID
				inventory.OfficeID = OfficeID

				_, err = repo.UpdateInventoryItem(inventory.ID, inventory)
				if err != nil {
					return nil, err
				}

			}
			_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.DispatchItems, itemDispatch, nil)
			if err != nil {
				return nil, err
			}
		}
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateDispatchItem(id int, item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error) {
	dispatch := dto.GetBasicInventoryDispatch{}

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Inventory.Dispatch+"/"+strconv.Itoa(id), item, &dispatch)
	if err != nil {
		return nil, err
	}

	if item.InventoryID != nil {
		filter := dto.DispatchInventoryItemFilter{
			DispatchID: &item.ID,
		}

		dispatchItems, _ := repo.GetMyInventoryDispatchesItems(&filter)

		for _, dispatch := range dispatchItems {
			_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Inventory.DispatchItems+"/"+strconv.Itoa(dispatch.ID), nil, nil)

			if err != nil {
				return nil, err
			}
		}

		for _, inventoryID := range item.InventoryID {
			itemDispatch := structs.BasicInventoryDispatchItemsItem{
				InventoryID: inventoryID,
				DispatchID:  dispatch.Data.ID,
			}

			_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.DispatchItems, itemDispatch, nil)
			if err != nil {
				return nil, err
			}
		}
	}

	return dispatch.Data, nil
}

func (repo *MicroserviceRepository) DeleteInventoryDispatch(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Inventory.Dispatch+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetDispatchItemByID(id int) (*structs.BasicInventoryDispatchItem, error) {
	res := dto.GetBasicInventoryDispatch{}

	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Dispatch+"/"+strconv.Itoa(id), nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateAssessments(data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.Assessments, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateAssessments(id int, data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Inventory.Assessments+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteAssessment(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Inventory.Assessments+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateInventoryItem(item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.Item, item, &res)
	if err != nil {
		return nil, err
	}

	if item.RealEstate != nil {
		item.RealEstate.ItemID = res.Data.ID
		_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.RealEstates, item.RealEstate, nil)
		if err != nil {
			return nil, err
		}
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateInventoryItem(id int, item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}
	res1 := dto.GetBasicInventoryInsertItem{}

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Inventory.Item+"/"+strconv.Itoa(id), item, &res)
	if err != nil {
		return nil, err
	}

	if item.RealEstate != nil {
		item.RealEstate.ItemID = res.Data.ID
		if item.RealEstate.ID != 0 {
			_, err := makeAPIRequest("PUT", repo.Config.Microservices.Inventory.RealEstates+"/"+strconv.Itoa(item.RealEstate.ID), item.RealEstate, &res1)
			if err != nil {
				return nil, err
			}
		} else {
			_, err := makeAPIRequest("POST", repo.Config.Microservices.Inventory.RealEstates, item.RealEstate, &res1)
			if err != nil {
				return nil, err
			}
		}
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetInventoryItem(id int) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Item+"/"+strconv.Itoa(id), nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetAllInventoryItem(filter dto.InventoryItemFilter) (*dto.GetAllBasicInventoryItem, error) {
	res := &dto.GetAllBasicInventoryItem{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Item, filter, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetMyInventoryRealEstate(id int) (*structs.BasicInventoryRealEstatesItem, error) {
	res := &dto.GetMyInventoryRealEstateResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Base+"/item/"+strconv.Itoa(id)+"/real-estates", nil, res)

	if err != nil {
		fmt.Printf("Fetching Real Estate failed because of this error - %s.\n", err)
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetMyInventoryAssessments(id int) ([]structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseArrayMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Assessments+"/"+strconv.Itoa(id)+"/item", nil, res)

	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok && apiErr.StatusCode != 404 {
			fmt.Printf("Fetching Assessments failed because of this error - %s.\n", err)
			return nil, err
		}
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetDispatchItemByInventoryID(id int) ([]*structs.BasicInventoryDispatchItemsItem, error) {
	res1 := dto.GetAllBasicInventoryDispatchItems{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Base+"/item/"+strconv.Itoa(id)+"/dispatch-items", nil, &res1)

	if err != nil {
		return nil, err
	}

	return res1.Data, nil
}

func (repo *MicroserviceRepository) GetInventoryRealEstatesList(input *dto.GetInventoryRealEstateListInputMS) (*dto.GetInventoryRealEstateListResponseMS, error) {
	res := &dto.GetInventoryRealEstateListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.RealEstates, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetInventoryRealEstate(id int) (*structs.BasicInventoryRealEstatesItem, error) {
	res := &dto.GetInventoryRealEstateResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.RealEstates+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CheckInsertInventoryData(input []structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, bool, int, error) {
	inventoryMap := make(map[string]bool)
	serialMap := make(map[string]bool)

	for _, item := range input {
		if serialMap[item.SerialNumber] {
			return &item, false, 1, nil
		}
		if inventoryMap[item.InventoryNumber] {
			return &item, false, 2, nil
		}

		inventoryMap[item.InventoryNumber] = true
		serialMap[item.SerialNumber] = true
	}

	for _, item := range input {
		inventoryItem, err := repo.GetAllInventoryItem(dto.InventoryItemFilter{
			SerialNumber: &item.SerialNumber,
		})

		if err != nil {
			return nil, false, 0, err
		}

		if len(inventoryItem.Data) != 0 {
			return &item, false, 1, nil
		}

		inventoryItem, err = repo.GetAllInventoryItem(dto.InventoryItemFilter{
			InventoryNumber: &item.InventoryNumber,
		})

		if err != nil {
			return nil, false, 0, err
		}

		if len(inventoryItem.Data) != 0 {
			return &item, false, 2, nil
		}
	}
	return nil, true, 0, nil
}
