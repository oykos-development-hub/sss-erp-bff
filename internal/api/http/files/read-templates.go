package files

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h *Handler) ReadArticlesPriceHandler(w http.ResponseWriter, r *http.Request) {
	var response ArticleResponse

	procurementID := r.FormValue("public_procurement_id")

	publicProcurementID, err := strconv.Atoi(procurementID)

	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	contractid := r.FormValue("contract_id")

	contractID, err := strconv.Atoi(contractid)

	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	var articles []ContractArticleResponseDTO

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			if rowindex == 0 {
				rowindex++
				continue
			}

			cols := rows.Columns()

			var (
				article            ContractArticleResponseDTO
				title, description string
				price              float32
			)

			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					title = value
				case 1:
					description = value
				case 2:
					if value == "" {
						break
					}

					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					price = float32(floatValue)
				}
			}

			if title == "" || description == "" || price == 0 {
				continue
			}

			input := dto.GetProcurementArticleListInputMS{
				Title:       &title,
				Description: &description,
				ItemID:      &publicProcurementID,
			}

			res, err := h.Repo.GetProcurementArticlesList(&input)

			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			if len(res) == 0 {
				err = errors.New("article \"" + title + "\" is not valid.")
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			contractArticle, err := h.Repo.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
				ArticleID: &res[0].ID,
			})

			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			if len(contractArticle.Data) > 0 {
				itemID := *contractArticle.Data[0]
				article.ID = itemID.ID
			}

			vatPercentage, _ := strconv.ParseFloat(res[0].VatPercentage, 32)
			vatFloat32 := float32(vatPercentage)
			article.ArticleID = res[0].ID
			grossValue := price + price*vatFloat32/100
			grossValue = float32(math.Round(float64(grossValue)*100) / 100)
			article.NetValue = &price
			article.GrossValue = &grossValue
			article.ContractID = contractID

			articles = append(articles, article)
		}
	}

	response.Status = "success"
	response.Message = "The file was read successfully"
	response.Data = articles

	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ReadArticlesHandler(w http.ResponseWriter, r *http.Request) {
	var response ProcurementArticleResponse

	procurementID := r.FormValue("public_procurement_id")

	publicProcurementID, err := strconv.Atoi(procurementID)

	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	var articles []structs.PublicProcurementArticle

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		if sheetName != "Stavke" {
			continue
		}

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			if rowindex == 0 {
				rowindex++
				continue
			}

			cols := rows.Columns()

			var article structs.PublicProcurementArticle
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					article.Title = value
				case 1:
					article.Description = value
				case 2:
					if value == "" {
						break
					}

					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					article.NetPrice = float32(floatValue)
				case 3:
					if value == "" {
						break
					}

					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}

					vatPercentage := 100 * floatValue / float64(article.NetPrice)
					round := math.Round(vatPercentage)

					valueVat := strconv.Itoa(int(round))

					article.VatPercentage = valueVat
				case 5:
					if value == "Materijalno knjigovodstvo" {
						article.VisibilityType = 2
					} else if value == "Osnovna sredstva" {
						article.VisibilityType = 3
					}
				}
			}

			article.PublicProcurementID = publicProcurementID

			if article.Title == "" || article.NetPrice == 0 || article.VatPercentage == "" {
				break
			}

			articles = append(articles, article)
		}

	}

	response.Data = articles
	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ReadArticlesInventoryHandler(w http.ResponseWriter, r *http.Request) {
	var response DonationArticleResponse

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	contractIDString := r.FormValue("contract_id")

	contractID, err := strconv.Atoi(contractIDString)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	contractArticles, err := h.Repo.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
		ContractID: &contractID})

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	var articlesData []structs.PublicProcurementArticle

	for _, article := range contractArticles.Data {
		articleData, err := h.Repo.GetProcurementArticle(article.PublicProcurementArticleID)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		articlesData = append(articlesData, *articleData)
	}

	var articles []structs.ReadArticlesDonation

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		if sheetName != "Stavke" {
			continue
		}

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			if rowindex == 0 {
				rowindex++
				continue
			}

			cols := rows.Columns()

			if len(cols) == 0 {
				break
			}
			var article structs.ReadArticlesDonation
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					if value == "" {
						break
					}
					article.Title = value
				case 1:
					if value == "" {
						break
					}
					article.SerialNumber = value
				case 2:

					article.Description = value
				}

			}
			if article.SerialNumber != "" {

				for _, articleData := range articlesData {
					if articleData.Title == article.Title {
						vatPercentageFloat, err := strconv.ParseFloat(articleData.VatPercentage, 32)

						if err != nil {
							handleError(w, err, http.StatusInternalServerError)
							return
						}

						article.GrossPrice = articleData.NetPrice + articleData.NetPrice*float32(vatPercentageFloat)/100
						article.ID = articleData.ID
						break
					}
				}

				articles = append(articles, article)
			}

		}
	}
	response.Data = articles
	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ReadArticlesDonationHandler(w http.ResponseWriter, r *http.Request) {
	var response DonationArticleResponse

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	var articles []structs.ReadArticlesDonation

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		if sheetName != "Stavke" {
			continue
		}

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			if rowindex == 0 {
				rowindex++
				continue
			}

			cols := rows.Columns()

			if len(cols) == 0 {
				break
			}
			var article structs.ReadArticlesDonation
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					if value == "" {
						break
					}
					article.Title = value
				case 1:
					if value == "" {
						break
					}

					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					article.GrossPrice = float32(floatValue)
				case 2:
					if value == "" {
						break
					}
					article.SerialNumber = value
				case 3:

					article.Description = value
				}

			}
			if article.SerialNumber != "" {
				articles = append(articles, article)
			}

		}
	}
	response.Data = articles
	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ReadExpireInventoriesHandler(w http.ResponseWriter, r *http.Request) {
	var response ExpireInventoriesResponse

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		if sheetName != "Stavke" && sheetName != "Sheet1" {
			break
		}

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			outerloop := true
			if rowindex == 0 {
				rowindex++
				continue
			}

			cols := rows.Columns()

			var dispatch structs.BasicInventoryAssessmentsTypesItem
			dispatch.Type = "financial"
			dispatch.Active = true
			var inventoryNumber string

			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					inventoryNumber = value

					inventory, err := h.Repo.GetAllInventoryItem(dto.InventoryItemFilter{InventoryNumber: &value})

					if err != nil || len(inventory.Data) == 0 {
						response.Message += "Inventarski broj " + value + " nije validan, ili ne postoji artikal sa tim brojem. "
						outerloop = false
						continue
					} else if inventoryNumber == "" {
						outerloop = false
						continue
					}

					dispatch.InventoryID = inventory.Data[0].ID

				case 4:
					floatValue, err := strconv.ParseFloat(value, 32)
					if err != nil {
						outerloop = false
					}
					dispatch.GrossPriceDifference = float32(floatValue)
				case 5:
					f, err := strconv.ParseFloat(value, 64)
					if err != nil {
						outerloop = false
					}
					estimatedDuration := int(f)
					if err != nil && estimatedDuration > 0 {
						outerloop = false
					}
					dispatch.EstimatedDuration = estimatedDuration
				case 6:
					residualPrice, err := strconv.ParseFloat(value, 32)
					if err != nil {
						continue
					}
					residualPriceFloat := float32(residualPrice)
					dispatch.ResidualPrice = &residualPriceFloat
				}
			}

			if inventoryNumber == "" && dispatch.InventoryID == 0 && dispatch.EstimatedDuration == 0 && dispatch.GrossPriceDifference == 0 && dispatch.ResidualPrice == nil {
				response.Status = "success"
				_ = MarshalAndWriteJSON(w, response)
				return
			}

			if !outerloop {
				continue
			}

			item, err := h.Repo.GetInventoryItem(dispatch.InventoryID)
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}
			now := time.Now()
			nowString := now.Format(config.ISO8601Format)
			dispatch.DateOfAssessment = &nowString
			dispatch.DepreciationTypeID = item.DepreciationTypeID

			response.Data = append(response.Data, dispatch)
			// item, err := h.Repo.GetInventoryItem(dispatch.InventoryId)
			// if err != nil {
			// 	handleError(w, err, http.StatusInternalServerError)
			// 	return
			// }
			// now := time.Now()
			// nowString := now.Format( config.ISO8601Format)
			// dispatch.DateOfAssessment = &nowString
			// dispatch.DepreciationTypeId = item.DepreciationTypeId
			// _, err = h.Repo.CreateAssessments(&dispatch)

			// if err != nil {
			// 	handleError(w, err, http.StatusInternalServerError)
			// }
		}
	}

	response.Status = "success"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ReadExpireImovableInventoriesHandler(w http.ResponseWriter, r *http.Request) {
	var response ExpireInventoriesResponse

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		if sheetName != "Stavke" && sheetName != "Sheet1" {
			break
		}

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			outerloop := true
			if rowindex == 0 {
				rowindex++
				continue
			}

			cols := rows.Columns()

			var dispatch structs.BasicInventoryAssessmentsTypesItem
			dispatch.Type = "financial"
			dispatch.Active = true
			var location string

			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					location = value

					if value == "" {
						continue
					}

					inventory, err := h.Repo.GetAllInventoryItem(dto.InventoryItemFilter{Location: &value})

					if err != nil || len(inventory.Data) == 0 {
						response.Message += "Lokacija " + value + " nije validna ili ne postoji sredstvo na toj lokaciji. "
						outerloop = false
						continue
					}

					dispatch.InventoryID = inventory.Data[0].ID

				case 4:
					floatValue, err := strconv.ParseFloat(value, 32)
					if err != nil {
						outerloop = false
					}
					dispatch.GrossPriceDifference = float32(floatValue)
				case 5:
					estimatedDuration, err := strconv.Atoi(value)
					if err != nil {
						outerloop = false
					}
					dispatch.EstimatedDuration = estimatedDuration
				case 6:
					residualPrice, err := strconv.ParseFloat(value, 32)
					if err != nil {
						continue
					}
					residualPriceFloat := float32(residualPrice)
					dispatch.ResidualPrice = &residualPriceFloat
				}
			}

			if location == "" && dispatch.InventoryID == 0 && dispatch.EstimatedDuration == 0 && dispatch.GrossPriceDifference == 0 && dispatch.ResidualPrice == nil {
				response.Status = "success"
				_ = MarshalAndWriteJSON(w, response)
				return
			}

			if !outerloop {
				continue
			}

			item, err := h.Repo.GetInventoryItem(dispatch.InventoryID)
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}
			now := time.Now()
			nowString := now.Format(config.ISO8601Format)
			dispatch.DateOfAssessment = &nowString
			dispatch.DepreciationTypeID = item.DepreciationTypeID
			_, err = h.Repo.CreateAssessments(r.Context(), &dispatch)

			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
			}
		}
	}

	response.Status = "success"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ReadArticlesSimpleProcurementHandler(w http.ResponseWriter, r *http.Request) {
	var response ProcurementArticleResponse

	procurementID := r.FormValue("public_procurement_id")

	publicProcurementID, err := strconv.Atoi(procurementID)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	var articles []structs.PublicProcurementArticle

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		if sheetName != "Stavke" {
			continue
		}

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			if rowindex == 0 {
				rowindex++
				continue
			}

			cols := rows.Columns()

			var article structs.PublicProcurementArticle
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					article.Title = value
				case 1:
					article.Description = value
				case 2:
					if value == "" {
						break
					}

					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					article.NetPrice = float32(floatValue)
				case 3:
					if value == "" {
						break
					}

					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return

					}

					vatPercentage := 100 * floatValue / float64(article.NetPrice)
					round := math.Round(vatPercentage)

					valueVat := strconv.Itoa(int(round))

					article.VatPercentage = valueVat
				case 4:
					if value == "" {
						break
					}

					amount, err := strconv.ParseFloat(value, 64)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}

					article.Amount = int(amount)
				case 6:
					if value == "Materijalno knjigovodstvo" {
						article.VisibilityType = 2
					} else if value == "Osnovna sredstva" {
						article.VisibilityType = 3
					}
				}
			}

			article.PublicProcurementID = publicProcurementID

			if article.Title == "" || article.NetPrice == 0 || article.VatPercentage == "" {
				break
			}

			articles = append(articles, article)
		}

	}

	response.Data = articles
	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ImportExcelOrgUnitInventoriesHandler(w http.ResponseWriter, r *http.Request) {
	var response ImportInventoriesResponse

	orgUnitIDstring := r.FormValue("organization_unit_id")

	orgUnitID, err := strconv.Atoi(orgUnitIDstring)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	var inventories []structs.BasicInventoryItem

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		if sheetName != "PS -1" {
			continue
		}

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			if rowindex < 9 {
				rowindex++
				continue
			}

			cols := rows.Columns()

			var inventory structs.BasicInventoryItem
			//var assessment structs.BasicInventoryAssessmentsTypesItem
			// var dispatch structs.BasicInventoryDispatchItem

			for cellIndex, cellValue := range cols {
				value := cellValue
				inventory.OrganizationUnitID = orgUnitID

				switch cellIndex {
				case 1:
					inventory.InventoryNumber = value
				case 2:
					input := dto.GetOfficesOfOrganizationInput{}

					input.Search = &value
					input.Value = &orgUnitIDstring

					office, err := h.Repo.GetOfficeDropdownSettings(&input)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					if len(office.Data) > 0 {
						inventory.OfficeID = office.Data[0].ID
					}
				case 3:
					inventory.Title = value
				case 5:
					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					inventory.GrossPrice = float32(floatValue)
				/*case 14:
				if value != "" {
					formattedDate, err := ConvertDateFormat(value)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					assessment.DateOfAssessment = &formattedDate
				}*/
				/*case 15:
				floatValue, err := strconv.ParseFloat(value, 32)

				if err != nil {
					handleError(w, err, http.StatusInternalServerError)
					return
				}
				assessment.GrossPriceDifference = float32(floatValue)*/
				/*case 16:
				floatValue, err := strconv.ParseFloat(value, 32)

				if err != nil {
					handleError(w, err, http.StatusInternalServerError)
					return
				}
				val := float32(floatValue)
				assessment.ResidualPrice = &val*/
				case 26:
					input := dto.GetSettingsInput{}

					input.Entity = "inventory_class_type"
					input.Search = &value

					class, err := h.Repo.GetDropdownSettings(&input)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					if len(class.Data) > 0 {
						inventory.ClassTypeID = class.Data[0].ID
					}
				case 29:
					if value != "" {
						floatValue, err := strconv.ParseFloat(value, 64)
						if err != nil {
							handleError(w, err, http.StatusInternalServerError)
							return
						}
						inventory.DateOfPurchase = ExcelDateToTimeString(floatValue)
					}
				case 30:
					input := dto.GetSettingsInput{}

					input.Entity = "deprecation_types"
					input.Search = &value

					deprecation, err := h.Repo.GetDropdownSettings(&input)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					if len(deprecation.Data) > 0 {
						//assessment.DepreciationTypeID = deprecation.Data[0].ID
						inventory.DepreciationTypeID = deprecation.Data[0].ID
					}
				}

			}

			inventories = append(inventories, inventory)
		}

	}

	response.Data = inventories
	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ImportUserProfileVacationsHandler(w http.ResponseWriter, r *http.Request) {
	var response ImportUserProfileVacationsResponse

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	sheetMap := xlsFile.GetSheetMap()

	var userProfileVacations []ImportUserProfileVacation

	users, err := h.Repo.GetUserProfiles(&dto.GetUserProfilesInput{})

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	usersMap := make(map[int]string)

	for _, user := range users {
		usersMap[user.ID] = user.FirstName + " " + user.LastName
	}

	for _, sheetName := range sheetMap {

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			if rowindex == 0 {
				rowindex++
				continue
			}

			cols := rows.Columns()

			var item ImportUserProfileVacation
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					userID, err := strconv.Atoi(value)

					if err != nil && value != "" {
						responseMessage := ValidationResponse{
							Column:  0,
							Row:     rowindex,
							Message: "ID korisnika nije validno unijet!",
						}
						response.Validation = append(response.Validation, responseMessage)
					} else if value != "" {
						_, exists := usersMap[userID]
						if !exists {
							responseMessage := ValidationResponse{
								Column:  0,
								Row:     rowindex,
								Message: "ID korisnika ne postoji!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							item.UserProfileID = userID
						}
					}
				case 1:
					userName, exists := usersMap[item.UserProfileID]
					if !exists || value != userName {
						responseMessage := ValidationResponse{
							Column:  1,
							Row:     rowindex,
							Message: "Korisnik i ID se ne podudaraju!",
						}
						response.Validation = append(response.Validation, responseMessage)
					}

				case 3:
					if value == "" {
						break
					}

					NumberOfDays, err := strconv.ParseFloat(value, 64)

					if err != nil {
						responseMessage := ValidationResponse{
							Column:  3,
							Row:     rowindex,
							Message: "Broj dana nije validno unijet!",
						}
						response.Validation = append(response.Validation, responseMessage)
					}

					item.NumberOfDays = int(NumberOfDays)
				}
			}
			userProfileVacations = append(userProfileVacations, item)
		}

	}

	response.Data = userProfileVacations
	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ImportUserExpirienceHandler(w http.ResponseWriter, r *http.Request) {
	var response ImportPS1Inventories

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	sheetMap := xlsFile.GetSheetMap()

	var userProfileExpiriences []structs.Experience

	users, err := h.Repo.GetUserProfiles(&dto.GetUserProfilesInput{})

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	usersMap := make(map[int]string)

	for _, user := range users {
		usersMap[user.ID] = user.FirstName + " " + user.LastName
	}

	organizationUnits, err := h.Repo.GetOrganizationUnits(&dto.GetOrganizationUnitsInput{})

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	OUMap := make(map[string]int)

	for _, OU := range organizationUnits.Data {
		OUMap[OU.Title] = OU.ID
	}

	for _, sheetName := range sheetMap {

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			if rowindex == 0 {
				rowindex++
				continue
			}

			rowindex++

			cols := rows.Columns()

			var item structs.Experience
			item.YearsOfInsuredExperience = -1
			item.MonthsOfInsuredExperience = -1
			item.DaysOfInsuredExperience = -1
			var dateOfStart time.Time
			var dateOfEnd time.Time
			var err error
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					userID, err := strconv.Atoi(value)

					if err != nil && value != "" {
						responseMessage := ValidationResponse{
							Column:  0,
							Row:     rowindex,
							Message: "ID korisnika nije validno unijet!",
						}
						response.Validation = append(response.Validation, responseMessage)
					} else if value != "" {
						_, exists := usersMap[userID]
						if !exists {
							responseMessage := ValidationResponse{
								Column:  0,
								Row:     rowindex,
								Message: "ID korisnika ne postoji!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							item.UserProfileID = userID
						}
					}
				case 1:
					userName, exists := usersMap[item.UserProfileID]
					if !exists || value != userName {
						responseMessage := ValidationResponse{
							Column:  1,
							Row:     rowindex,
							Message: "Korisnik i ID se ne podudaraju!",
						}
						response.Validation = append(response.Validation, responseMessage)
					}
				case 2:
					if value != "Da" && value != "Ne" {
						responseMessage := ValidationResponse{
							Column:  2,
							Row:     rowindex,
							Message: "Nevalidna vrijednost! Dozvoljene vrijednosti su \"Da\" i \"Ne\"!",
						}
						response.Validation = append(response.Validation, responseMessage)
					}

					if value == "Da" {
						item.Relevant = true
					} else {
						item.Relevant = false
					}
				case 3:
					if !item.Relevant {
						item.OrganizationUnit = value
					} else {
						id, exists := OUMap[value]
						if !exists {
							responseMessage := ValidationResponse{
								Column:  3,
								Row:     rowindex,
								Message: "Organizaciona jednica nije validna!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							item.OrganizationUnitID = id
						}
					}
				case 4:
					dateOfStart, err = parseDate(value)
					if err != nil {
						responseMessage := ValidationResponse{
							Column:  4,
							Row:     rowindex,
							Message: "Pocetak radnog odnosa nije validan!",
						}
						response.Validation = append(response.Validation, responseMessage)
					} else {
						item.DateOfStart = dateOfStart.Format(config.ISO8601Format)
					}
				case 5:
					dateOfEnd, err = parseDate(value)
					if err != nil {
						responseMessage := ValidationResponse{
							Column:  5,
							Row:     rowindex,
							Message: "Kraj radnog odnosa nije validan!",
						}
						response.Validation = append(response.Validation, responseMessage)
					} else {
						if dateOfEnd.Before(dateOfStart) {
							responseMessage := ValidationResponse{
								Column:  5,
								Row:     rowindex,
								Message: "Kraj radnog odnosa mora biti posle početka radnog odnosa!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							item.DateOfEnd = dateOfEnd.Format(config.ISO8601Format)
						}
					}
				case 6:
					years, err := strconv.Atoi(value)

					if err != nil && value != "" {
						responseMessage := ValidationResponse{
							Column:  6,
							Row:     rowindex,
							Message: "Godine prijavljenog staza nijesu validno unijete!",
						}
						response.Validation = append(response.Validation, responseMessage)
					} else if value != "" {
						item.YearsOfInsuredExperience = years
					}
				case 7:
					months, err := strconv.Atoi(value)

					if err != nil && value != "" {
						responseMessage := ValidationResponse{
							Column:  7,
							Row:     rowindex,
							Message: "Mjeseci prijavljenog staza nijesu validno unijete!",
						}
						response.Validation = append(response.Validation, responseMessage)
					} else if value != "" {
						item.MonthsOfInsuredExperience = months
					}
				case 8:
					days, err := strconv.Atoi(value)

					if err != nil && value != "" {
						responseMessage := ValidationResponse{
							Column:  8,
							Row:     rowindex,
							Message: "Dani prijavljenog staza nijesu validno unijeti!",
						}
						response.Validation = append(response.Validation, responseMessage)
					} else if value != "" {
						item.DaysOfInsuredExperience = days
					}

				}
			}

			yearsDiff := dateOfEnd.Year() - dateOfStart.Year()
			monthsDiff := int(dateOfEnd.Month()) - int(dateOfStart.Month())

			if monthsDiff < 0 {
				monthsDiff += 12
				yearsDiff--
			}

			daysDiff := int(dateOfEnd.Day()) - int(dateOfStart.Day())
			if daysDiff < 0 {
				monthsDiff--
				daysDiff += 30
				if monthsDiff < 0 {
					yearsDiff--
					monthsDiff += 12
				}
			}

			if item.YearsOfInsuredExperience == -1 || item.MonthsOfInsuredExperience == -1 || item.DaysOfInsuredExperience == -1 {
				item.YearsOfInsuredExperience = yearsDiff
				item.MonthsOfInsuredExperience = monthsDiff
				item.DaysOfInsuredExperience = daysDiff
			}

			item.YearsOfExperience = yearsDiff
			item.MonthsOfExperience = monthsDiff
			item.DaysOfExperience = daysDiff

			if (item.OrganizationUnit != "" || item.OrganizationUnitID != 0) && item.DateOfStart != "" && item.DateOfEnd != "" {
				userProfileExpiriences = append(userProfileExpiriences, item)
			}
		}
	}

	if len(response.Validation) == 0 {
		response.Data = userProfileExpiriences
	}

	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ImportSalariesHandler(w http.ResponseWriter, r *http.Request) {
	var response ImportSalaries

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	organizationUnitIDStr := r.FormValue("organization_unit_id")

	organizationUnitID, err := strconv.Atoi(organizationUnitIDStr)

	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0
		var typeOfInput string

		for rows.Next() {
			rowindex++

			cols := rows.Columns()

			additionalSalaryExpense := structs.SalaryAdditionalExpense{}
			//var err error
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					switch typeOfInput {
					case "Doprinosi":
						additionalSalaryExpense.Type = "contributions"
					case "Porezi":
						additionalSalaryExpense.Type = "taxes"
					case "Obustave":
						additionalSalaryExpense.Type = "suspensions"
					case "Prirezi":
						additionalSalaryExpense.Type = "subtaxes"
					case "Banka":
						additionalSalaryExpense.Type = "banks"
					}

					additionalSalaryExpense.Title = value

					if value == "Doprinosi" || value == "Porezi" || value == "Prirezi" || value == "Obustave" || value == "Banka" {
						typeOfInput = value
					}

					if strings.Contains(value, "Ukupno") || strings.Contains(value, "Za isplatu") {
						typeOfInput = ""
						additionalSalaryExpense.Type = ""
					}

				case 3:
					if value != "" && value != "Broj zaposlenih" {
						numberOfEmployees, err := strconv.Atoi(value)
						if err != nil {
							responseMessage := ValidationResponse{
								Column:  7,
								Row:     rowindex,
								Message: "Broj zaposlenih nije ispravno unijet!",
							}
							response.Validation = append(response.Validation, responseMessage)
						}
						response.NumberOfEmployees = numberOfEmployees

					}
				case 7:
					price, err := strconv.ParseFloat(value, 32)

					if err != nil && value != "" && additionalSalaryExpense.Type != "" && additionalSalaryExpense.Title != "" {
						responseMessage := ValidationResponse{
							Column:  7,
							Row:     rowindex,
							Message: "Iznos nije ispravno unijet!",
						}
						response.Validation = append(response.Validation, responseMessage)
					} else {
						additionalSalaryExpense.Amount = price
					}

				}
			}
			if additionalSalaryExpense.Type != "" && additionalSalaryExpense.Title != "" && additionalSalaryExpense.Amount > 0 {
				additionalSalaryExpense.OrganizationUnitID = organizationUnitID
				additionalSalaryExpense.Status = "Kreiran"
				response.Data = append(response.Data, additionalSalaryExpense)
			}
		}
	}

	/*	if len(response.Validation) == 0 {
			response.Data = item
		}
	*/
	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ImportSuspensionsHandler(w http.ResponseWriter, r *http.Request) {
	var response ImportSalaries

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	organizationUnitIDStr := r.FormValue("organization_unit_id")

	organizationUnitID, err := strconv.Atoi(organizationUnitIDStr)

	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0
		var title string

		for rows.Next() {
			rowindex++
			if rowindex < 2 {
				continue
			}

			cols := rows.Columns()

			additionalSalaryExpense := structs.SalaryAdditionalExpense{}
			//var err error
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					if rowindex == 2 && value != "" {
						title = value
					}
				case 1:
					if value != "" && value != "Zaposleni" {
						additionalSalaryExpense.Type = "suspensions"
					}
				case 5:
					additionalSalaryExpense.BankAccount = value
				case 4:
					additionalSalaryExpense.IdentificatorNumber = value
				case 11:
					noThousands := strings.ReplaceAll(value, ".", "")

					normalized := strings.ReplaceAll(noThousands, ",", "")

					price, err := strconv.ParseFloat(normalized, 32)

					price = price / 100

					if err != nil && value != "" && additionalSalaryExpense.Type != "" && additionalSalaryExpense.Title != "" && value != "Iznos rate" {
						responseMessage := ValidationResponse{
							Column:  11,
							Row:     rowindex,
							Message: "Iznos nije ispravno unijet!",
						}
						response.Validation = append(response.Validation, responseMessage)
					} else {
						additionalSalaryExpense.Amount = price
					}

				}
			}
			if additionalSalaryExpense.Type != "" && rowindex > 10 {
				additionalSalaryExpense.Title = title
				additionalSalaryExpense.OrganizationUnitID = organizationUnitID
				additionalSalaryExpense.Status = "Kreiran"
				response.Data = append(response.Data, additionalSalaryExpense)
			}
		}
	}

	/*	if len(response.Validation) == 0 {
			response.Data = item
		}
	*/
	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func (h *Handler) ImportSAPHandler(w http.ResponseWriter, r *http.Request) {
	var response ImportSAP

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {

		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		rowindex := 0

		for rows.Next() {
			rowindex++
			if rowindex < 2 {
				continue
			}

			cols := rows.Columns()

			paymentOrder := structs.PaymentOrder{}
			var SAPID string
			var DateOfSAP time.Time
			//var err error
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					SAPID = value
				case 6:
					if value != "" {
						date, err := parseDate(value)

						if err != nil {
							responseMessage := ValidationResponse{
								Column:  6,
								Row:     rowindex,
								Message: "Datum nije ispravno unijet!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							DateOfSAP = date
						}
					}
				case 14:
					if value != "" {
						idOfStatement, err := strconv.Atoi(value)
						if err != nil {
							responseMessage := ValidationResponse{
								Column:  14,
								Row:     rowindex,
								Message: "ID naloga nije ispravno unijet!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							item, err := h.Repo.GetPaymentOrderByIDOfStatement(idOfStatement)

							if err != nil || item.ID == 0 {
								responseMessage := ValidationResponse{
									Column:  14,
									Row:     rowindex,
									Message: "Nalog sa zadatim ID-em ne postoji!",
								}
								response.Validation = append(response.Validation, responseMessage)
							} else {
								paymentOrder = *item
							}
						}
					}
				}
			}

			paymentOrder.SAPID = &SAPID
			paymentOrder.DateOfSAP = &DateOfSAP
			if paymentOrder.SAPID != nil && paymentOrder.DateOfSAP != nil && paymentOrder.ID != 0 {
				response.Data = append(response.Data, paymentOrder)
			}
		}
	}

	if len(response.Validation) == 0 {
		for _, item := range response.Data {
			_, err := h.Repo.UpdatePaymentOrder(r.Context(), &item)

			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}
		}
	}

	response.Message = "File was read successfuly"
	response.Status = "success"
	_ = MarshalAndWriteJSON(w, response)
}
