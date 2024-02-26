package files

import (
	"bff/internal/api/dto"
	"bff/structs"
	"errors"
	"math"
	"net/http"
	"strconv"
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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

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
						if err != nil {
							handleError(w, err, http.StatusInternalServerError)
							return
						}
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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}
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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}
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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

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
			nowString := now.Format("2006-01-02T00:00:00Z")
			dispatch.DateOfAssessment = &nowString
			dispatch.DepreciationTypeID = item.DepreciationTypeID

			response.Data = append(response.Data, dispatch)
			// item, err := h.Repo.GetInventoryItem(dispatch.InventoryId)
			// if err != nil {
			// 	handleError(w, err, http.StatusInternalServerError)
			// 	return
			// }
			// now := time.Now()
			// nowString := now.Format("2006-01-02T00:00:00Z")
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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

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
			nowString := now.Format("2006-01-02T00:00:00Z")
			dispatch.DateOfAssessment = &nowString
			dispatch.DepreciationTypeID = item.DepreciationTypeID
			_, err = h.Repo.CreateAssessments(&dispatch)

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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

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
						if err != nil {
							handleError(w, err, http.StatusInternalServerError)
							return
						}
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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			var inventory structs.BasicInventoryItem
			var assessment structs.BasicInventoryAssessmentsTypesItem
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
				case 14:
					if value != "" {
						formattedDate, err := ConvertDateFormat(value)
						if err != nil {
							handleError(w, err, http.StatusInternalServerError)
							return
						}
						assessment.DateOfAssessment = &formattedDate
					}
				case 15:
					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					assessment.GrossPriceDifference = float32(floatValue)
				case 16:
					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					val := float32(floatValue)
					assessment.ResidualPrice = &val
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
						assessment.DepreciationTypeID = deprecation.Data[0].ID
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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			var item ImportUserProfileVacation
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					if value == "" {
						break
					}

					id, err := strconv.ParseFloat(value, 64)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}

					item.UserProfileID = int(id)
				case 3:
					if value == "" {
						break
					}

					NumberOfDays, err := strconv.ParseFloat(value, 64)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
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

func (h *Handler) ImportExcelPS1(w http.ResponseWriter, r *http.Request) {
	var response ImportPS1Inventories
	organizationUnitIDStr := r.FormValue("organization_unit_id")

	organizationUnitID, err := strconv.Atoi(organizationUnitIDStr)

	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	var articles []ImportInventoryArticles

	sheetMap := xlsFile.GetSheetMap()
	classTypes, err := h.Repo.GetDropdownSettings(&dto.GetSettingsInput{Entity: "inventory_class_type"})

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	mapOfClassTypes := make(map[string]int)

	for _, obj := range classTypes.Data {
		mapOfClassTypes[obj.Abbreviation] = obj.ID
	}

	deprecationTypes, err := h.Repo.GetDropdownSettings(&dto.GetSettingsInput{Entity: "deprecation_types"})

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	mapOfDeprecationTypes := make(map[string]int)

	for _, obj := range deprecationTypes.Data {
		mapOfDeprecationTypes[obj.Title] = obj.ID
	}

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
			rowindex++
			if rowindex <= 9 {
				continue
			}

			if rowindex > 100 {
				break
			}

			cols := rows.Columns()
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			var article ImportInventoryArticles
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 1:
					article.Article.InventoryNumber = value
				case 3:
					article.Article.Title = value
				case 5:
					price, err := strconv.ParseFloat(value, 32)

					if value != "" && err != nil {
						responseMessage := ValidationResponse{
							Column:  5,
							Row:     rowindex,
							Message: "Cijena nije validno unijeta!",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						article.Article.GrossPrice = float32(price)
					}
				case 14:
					DateOfAssessment, err := parseDate(value)

					if value != "" && err != nil {
						responseMessage := ValidationResponse{
							Column:  29,
							Row:     rowindex,
							Message: "Datum amortizacije nije validno unijet!",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						dateOfAssessment := DateOfAssessment.Format("2006-01-02T15:04:05Z")
						article.Amortization.DateOfAssessment = &dateOfAssessment
					}
				case 15:
					grossPriceNew, err := strconv.ParseFloat(value, 32)

					if value != "" && err != nil {
						responseMessage := ValidationResponse{
							Column:  15,
							Row:     rowindex,
							Message: "Nova cijena nije validno unijeta!",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						article.Amortization.GrossPriceNew = float32(grossPriceNew)
					}
				case 16:
					residualPrice, err := strconv.ParseFloat(value, 32)
					if value != "" && err != nil {
						responseMessage := ValidationResponse{
							Column:  16,
							Row:     rowindex,
							Message: "Rezidualna cijena nije validno unijeta!",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						residualPriceFloat32 := float32(residualPrice)
						article.Amortization.ResidualPrice = &residualPriceFloat32
					}
				case 24:
					estimatedDuration, err := strconv.Atoi(value)
					if value != "" && err != nil {
						responseMessage := ValidationResponse{
							Column:  24,
							Row:     rowindex,
							Message: "Vijek trajanja nije validno unijet!",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						article.Amortization.EstimatedDuration = estimatedDuration
					}
				case 25:
					article.Article.Description = value
				case 26:
					if _, exists := mapOfClassTypes[value]; !exists && value != "" {
						responseMessage := ValidationResponse{
							Column:  26,
							Row:     rowindex,
							Message: "Klasa sredstva " + value + " nije validna.",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						article.Article.ClassTypeID = mapOfClassTypes[value]
					}
				case 29:
					dateOfPurchase, err := parseDate(value)

					if value != "" && err != nil {
						responseMessage := ValidationResponse{
							Column:  29,
							Row:     rowindex,
							Message: "Datum nabavke nije validno unijet!",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						article.Article.DateOfPurchase = dateOfPurchase.Format("2006-01-02T15:04:05Z")
					}
				case 30:
					if _, exists := mapOfDeprecationTypes[value]; !exists && value != "" {
						responseMessage := ValidationResponse{
							Column:  30,
							Row:     rowindex,
							Message: "Amortizaciona grupa " + value + " nije validna.",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						article.Article.DepreciationTypeID = mapOfDeprecationTypes[value]
						article.Amortization.DepreciationTypeID = mapOfDeprecationTypes[value]
					}
				}

			}
			articles = append(articles, article)
		}
	}

	if len(response.Data) == 0 {
		for _, article := range articles {
			article.Article.OrganizationUnitID = organizationUnitID
			article.Article.Type = "movable"
			article.Article.Active = true
			newArticle, err := h.Repo.CreateInventoryItem(&article.Article)
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			if article.Amortization.EstimatedDuration > 0 {
				article.Amortization.InventoryID = newArticle.ID
				article.Amortization.Type = "financial"
				article.Amortization.Active = true
				_, err = h.Repo.CreateAssessments(&article.Amortization)
				if err != nil {
					handleError(w, err, http.StatusInternalServerError)
					return
				}
			}
		}
	}

	response.Status = "success"
	response.Message = "The file was read successfully"

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
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			var item structs.Experience
			var dateOfStart time.Time
			var dateOfEnd time.Time
			var err error
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					userID, err := strconv.Atoi(value)

					if err != nil {
						responseMessage := ValidationResponse{
							Column:  0,
							Row:     rowindex,
							Message: "ID korisnika nije validno unijet!",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						_, exists := usersMap[userID]
						if !exists {
							responseMessage := ValidationResponse{
								Column:  0,
								Row:     rowindex,
								Message: "ID korisnika ne postoji!",
							}
							response.Data = append(response.Data, responseMessage)
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
						response.Data = append(response.Data, responseMessage)
					}
				case 2:
					if value != "Da" && value != "Ne" {
						responseMessage := ValidationResponse{
							Column:  2,
							Row:     rowindex,
							Message: "Nevalidna vrijednost!",
						}
						response.Data = append(response.Data, responseMessage)
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
							response.Data = append(response.Data, responseMessage)
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
						response.Data = append(response.Data, responseMessage)
					} else {
						item.DateOfStart = dateOfStart.Format("2006-01-02T00:00:00Z")
					}
				case 5:
					dateOfEnd, err = parseDate(value)
					if err != nil {
						responseMessage := ValidationResponse{
							Column:  5,
							Row:     rowindex,
							Message: "Kraj radnog odnosa nije validan!",
						}
						response.Data = append(response.Data, responseMessage)
					} else {
						item.DateOfEnd = dateOfEnd.Format("2006-01-02T00:00:00Z")

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

						item.YearsOfExperience = yearsDiff
						item.YearsOfInsuredExperience = yearsDiff
						item.MonthsOfExperience = monthsDiff
						item.MonthsOfInsuredExperience = monthsDiff
						item.DaysOfExperience = daysDiff
						item.DaysOfInsuredExperience = daysDiff
					}
				}
			}
			userProfileExpiriences = append(userProfileExpiriences, item)
		}
	}

	if len(response.Data) == 0 {
		for _, item := range userProfileExpiriences {
			_, err := h.Repo.CreateExperience(&item)

			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}
		}
	}

	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}
