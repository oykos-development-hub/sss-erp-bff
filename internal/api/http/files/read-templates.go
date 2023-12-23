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
				article.ID = contractArticle.Data[0].ID
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
