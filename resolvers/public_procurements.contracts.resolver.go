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

var PublicProcurementContractsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.ProcurementContractResponseItem{}
	var total int

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]

	sortByDateOfExpiry := params.Args["sort_by_date_of_expiry"]
	sortByDateOfSigning := params.Args["sort_by_date_of_signing"]
	sortByGrossValue := params.Args["sort_by_gross_value"]
	sortBySerialNumber := params.Args["sort_by_serial_number"]
	procurement_id := params.Args["procurement_id"]
	supplier_id := params.Args["supplier_id"]
	year := params.Args["year"]

	input := dto.GetProcurementContractsInput{}
	if shared.IsInteger(page) && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if shared.IsInteger(size) && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}
	if shared.IsInteger(procurement_id) && procurement_id.(int) > 0 {
		procurementID := procurement_id.(int)
		input.ProcurementID = &procurementID
	}
	if shared.IsInteger(supplier_id) && supplier_id.(int) > 0 {
		supplierID := supplier_id.(int)
		input.SupplierID = &supplierID
	}
	if sortByDateOfExpiry != nil && sortByDateOfExpiry.(string) != "" {
		dateOfExpiratioin := sortByDateOfExpiry.(string)
		input.SortByDateOfExpiry = &dateOfExpiratioin
	}
	if sortByDateOfSigning != nil && sortByDateOfSigning.(string) != "" {
		value := sortByDateOfSigning.(string)
		input.SortByDateOfSigning = &value
	}
	if sortByGrossValue != nil && sortByGrossValue.(string) != "" {
		value := sortByGrossValue.(string)
		input.SortByGrossValue = &value
	}
	if sortBySerialNumber != nil && sortBySerialNumber.(string) != "" {
		value := sortBySerialNumber.(string)
		input.SortBySerialNumber = &value
	}
	if year != nil && year.(string) != "" {
		value := year.(string)
		input.Year = &value
	}

	if shared.IsInteger(id) && id.(int) > 0 {
		contract, err := getProcurementContract(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resItem, _ := buildProcurementContractResponseItem(contract)
		items = append(items, *resItem)
		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   items,
			Total:   1,
		}, nil
	} else {
		contractsRes, err := getProcurementContractsList(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		total = contractsRes.Total

		for _, contract := range contractsRes.Data {
			resItem, err := buildProcurementContractResponseItem(contract)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			items = append(items, *resItem)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   items,
			Total:   total,
		}, nil
	}

}

var PublicProcurementContractInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementContract
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateProcurementContract(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementContractResponseItem(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := createProcurementContract(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		articles, err := getProcurementArticlesList(&dto.GetProcurementArticleListInputMS{
			ItemID: &data.PublicProcurementId,
		})

		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, article := range articles {
			item := structs.PublicProcurementContractArticle{
				PublicProcurementArticleId:  article.Id,
				PublicProcurementContractId: res.Id,
				NetValue:                    0,
				GrossValue:                  0,
				VatPercentage:               article.VatPercentage,
				UsedArticles:                0,
			}

			_, err := createProcurementContractArticle(&item)

			if err != nil {
				return shared.HandleAPIError(err)
			}

		}

		item, err := buildProcurementContractResponseItem(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func buildProcurementContractResponseItem(item *structs.PublicProcurementContract) (*dto.ProcurementContractResponseItem, error) {
	publicProcurementItem, err := getProcurementItem(item.PublicProcurementId)
	if err != nil {
		return nil, err
	}

	supplier, err := getSupplier(item.SupplierId)
	if err != nil {
		return nil, err
	}

	var files []dto.FileDropdownSimple

	for _, id := range item.File {
		file, err := getFileByID(id)
		if err != nil {
			return nil, err
		}

		fileDropDown := dto.FileDropdownSimple{
			Id:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		}

		files = append(files, fileDropDown)
	}

	daysUntilExpiry, err := calculateDaysUntilExpiry(*item.DateOfExpiry)
	if err != nil {
		return nil, err
	}

	res := dto.ProcurementContractResponseItem{
		Id:                  item.Id,
		PublicProcurementId: item.PublicProcurementId,
		SupplierId:          item.SupplierId,
		DateOfSigning:       (string)(item.DateOfSigning),
		DateOfExpiry:        (*string)(item.DateOfExpiry),
		SerialNumber:        item.SerialNumber,
		NetValue:            item.NetValue,
		GrossValue:          item.GrossValue,
		File:                files,
		CreatedAt:           item.CreatedAt,
		UpdatedAt:           item.UpdatedAt,
		PublicProcurement: dto.DropdownSimple{
			Id:    publicProcurementItem.Id,
			Title: publicProcurementItem.Title,
		},
		Supplier: dto.DropdownSimple{
			Id:    supplier.Id,
			Title: supplier.Title,
		},
		DaysUntilExpiry: daysUntilExpiry,
	}

	return &res, nil
}

func calculateDaysUntilExpiry(expiryDateStr string) (int, error) {
	// Parse the expiry date
	expiryDate, err := time.Parse(time.RFC3339, expiryDateStr)
	if err != nil {
		return 0, err
	}

	// Get the current date (ignoring time of day)
	now := time.Now()
	current := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Calculate the difference in days
	duration := expiryDate.Sub(current)
	return int(duration.Hours() / 24), nil
}

var PublicProcurementContractDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteProcurementContract(itemId)
	if err != nil {
		fmt.Printf("Deleting procurement contract failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func createProcurementContract(resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error) {
	res := &dto.GetProcurementContractResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.CONTRACTS_ENDPOINT, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementContract(id int, resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error) {
	res := &dto.GetProcurementContractResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.CONTRACTS_ENDPOINT+"/"+strconv.Itoa(id), resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteProcurementContract(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.CONTRACTS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getProcurementContract(id int) (*structs.PublicProcurementContract, error) {
	res := &dto.GetProcurementContractResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.CONTRACTS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getProcurementContractsList(input *dto.GetProcurementContractsInput) (*dto.GetProcurementContractListResponseMS, error) {
	res := &dto.GetProcurementContractListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.CONTRACTS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getFileByID(id int) (*structs.File, error) {
	res := &dto.GetFileResponsePom{}
	_, err := shared.MakeAPIRequest("GET", config.FILES_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data.Data, nil
}

func getOrganizationUnitArticlesList(input dto.GetProcurementOrganizationUnitArticleListInputDTO) ([]dto.GetPublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleListResponse{}
	_, err := shared.MakeAPIRequest("GET", config.ORGANIZATION_UNIT_ARTICLE_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getOrganizationUnitArticleByID(id int) (*dto.GetPublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleResponse{}
	_, err := shared.MakeAPIRequest("GET", config.ORGANIZATION_UNIT_ARTICLE_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
