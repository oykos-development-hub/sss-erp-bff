package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) PublicProcurementContractsOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.ProcurementContractResponseItem{}
	var total int

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]

	sortByDateOfExpiry := params.Args["sort_by_date_of_expiry"]
	sortByDateOfSigning := params.Args["sort_by_date_of_signing"]
	sortByGrossValue := params.Args["sort_by_gross_value"]
	sortBySerialNumber := params.Args["sort_by_serial_number"]
	procurementID := params.Args["procurement_id"]
	supplierID := params.Args["supplier_id"]
	year := params.Args["year"]

	input := dto.GetProcurementContractsInput{}
	if page != nil && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if size != nil && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}
	if procurementID != nil && procurementID.(int) > 0 {
		procurementID := procurementID.(int)
		input.ProcurementID = &procurementID
	}
	if supplierID != nil && supplierID.(int) > 0 {
		supplierID := supplierID.(int)
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

	if id != nil && id.(int) > 0 {
		contract, err := r.Repo.GetProcurementContract(id.(int))
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		resItem, _ := buildProcurementContractResponseItem(r.Repo, contract)
		items = append(items, *resItem)
		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   items,
			Total:   1,
		}, nil
	}
	contractsRes, err := r.Repo.GetProcurementContractsList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	total = contractsRes.Total

	for _, contract := range contractsRes.Data {
		resItem, err := buildProcurementContractResponseItem(r.Repo, contract)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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

func (r *Resolver) PublicProcurementContractInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementContract
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		res, err := r.Repo.UpdateProcurementContract(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildProcurementContractResponseItem(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := r.Repo.CreateProcurementContract(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		articles, err := r.Repo.GetProcurementArticlesList(&dto.GetProcurementArticleListInputMS{
			ItemID: &data.PublicProcurementID,
		})

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, article := range articles {
			item := structs.PublicProcurementContractArticle{
				PublicProcurementArticleID:  article.ID,
				PublicProcurementContractID: res.ID,
				NetValue:                    0,
				GrossValue:                  0,
				VatPercentage:               article.VatPercentage,
				UsedArticles:                0,
			}

			_, err := r.Repo.CreateProcurementContractArticle(&item)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

		}

		item, err := buildProcurementContractResponseItem(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		contractArticles, err := r.Repo.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
			ContractID: &item.ID,
		})

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		isAccounting := false
		isInventory := false

		for _, article := range contractArticles.Data {
			procurementArticle, err := r.Repo.GetProcurementArticle(article.PublicProcurementArticleID)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			if procurementArticle.VisibilityType == 3 {
				isInventory = true
			} else if procurementArticle.VisibilityType == 2 {
				isAccounting = true
			}

			if isAccounting && isInventory {
				break
			}

		}

		if isAccounting {
			loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

			targetUsers, _ := r.Repo.GetUsersByPermission(config.AccountingContract, config.OperationRead)
			/*if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}*/

			for _, user := range targetUsers {
				if loggedInUser.ID != user.ID {
					dataJSON, _ := json.Marshal(item)
					_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
						Content:     "Potpisan je novi ugovor",
						Module:      "Javne nabavke",
						FromUserID:  loggedInUser.ID,
						ToUserID:    user.ID,
						FromContent: "Službenik za javne nabavke",
						IsRead:      false,
						Data:        dataJSON,
						Path:        fmt.Sprintf("/accounting/contracts/%d/contract-details", item.ID),
					})
					if err != nil {
						_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
						return errors.HandleAPPError(err)
					}
				}
			}
		}

		if isInventory {
			loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

			targetUsers, _ := r.Repo.GetUsersByPermission(config.InventoryMovableItems, config.OperationRead)
			/*if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}*/

			searchParam := "Sekretarijat"

			targetOrganizationUnit, _ := r.Repo.GetOrganizationUnits(
				&dto.GetOrganizationUnitsInput{
					ParentID: nil,
					Search:   &searchParam,
				},
			)

			/*if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}*/

			var targetOrganziationUnitID int

			for _, OU := range targetOrganizationUnit.Data {
				if OU.ID < 5 {
					targetOrganziationUnitID = OU.ID
				}
			}

			if targetOrganziationUnitID != 0 {

				employees, _ := GetEmployeesOfOrganizationUnit(r.Repo, targetOrganziationUnitID)
				/*if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}*/

				for _, user := range targetUsers {
					for _, employee := range employees {
						if employee.UserAccountID == user.ID && loggedInUser.ID != user.ID {
							dataJSON, _ := json.Marshal(item)
							_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
								Content:     "Potpisan je novi ugovor",
								Module:      "Javne nabavke",
								FromUserID:  loggedInUser.ID,
								ToUserID:    user.ID,
								FromContent: "Službenik za javne nabavke",
								IsRead:      false,
								Data:        dataJSON,
								Path:        fmt.Sprintf("/procurements/contracts/%d/details", item.ID),
							})
							if err != nil {
								_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
								return errors.HandleAPPError(err)
							}
						}
					}
				}
			}
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func buildProcurementContractResponseItem(r repository.MicroserviceRepositoryInterface, item *structs.PublicProcurementContract) (*dto.ProcurementContractResponseItem, error) {
	publicProcurementItem, err := r.GetProcurementItem(item.PublicProcurementID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get procurement item")
	}

	supplier, _ := r.GetSupplier(item.SupplierID)
	/*if err != nil {
		return nil, errors.Wrap(err, "repo get supplier")
	}*/

	var files []dto.FileDropdownSimple

	for _, id := range item.File {
		file, _ := r.GetFileByID(id)
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}*/

		if file != nil {

			fileDropDown := dto.FileDropdownSimple{
				ID:   file.ID,
				Name: file.Name,
				Type: *file.Type,
			}

			files = append(files, fileDropDown)
		}
	}

	daysUntilExpiry, err := calculateDaysUntilExpiry(*item.DateOfExpiry)
	if err != nil {
		return nil, errors.Wrap(err, "calculate days until expiry")
	}

	res := dto.ProcurementContractResponseItem{
		ID:                  item.ID,
		PublicProcurementID: item.PublicProcurementID,
		SupplierID:          item.SupplierID,
		DateOfSigning:       (string)(item.DateOfSigning),
		DateOfExpiry:        (*string)(item.DateOfExpiry),
		SerialNumber:        item.SerialNumber,
		NetValue:            item.NetValue,
		GrossValue:          item.GrossValue,
		File:                files,
		CreatedAt:           item.CreatedAt,
		UpdatedAt:           item.UpdatedAt,
		PublicProcurement: dto.DropdownSimple{
			ID:    publicProcurementItem.ID,
			Title: publicProcurementItem.Title,
		},
		DaysUntilExpiry: daysUntilExpiry,
	}

	if supplier != nil {
		res.Supplier = dto.DropdownSimple{
			ID:    supplier.ID,
			Title: supplier.Title,
		}
	}

	return &res, nil
}

func calculateDaysUntilExpiry(expiryDateStr string) (int, error) {
	// Parse the expiry date
	expiryDate, err := time.Parse(time.RFC3339, expiryDateStr)
	if err != nil {
		return 0, errors.Wrap(err, "time parse")
	}

	// Get the current date (ignoring time of day)
	now := time.Now()
	current := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Calculate the difference in days
	duration := expiryDate.Sub(current)
	return int(duration.Hours() / 24), nil
}

func (r *Resolver) PublicProcurementContractDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteProcurementContract(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
