package files

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"errors"
	"math"
	"net/http"
	"strconv"
)

func ReadArticlesPriceHandler(w http.ResponseWriter, r *http.Request) {
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

			var article ContractArticleResponseDTO
			var title, description string
			var price float32
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

			res, err := getProcurementArticlesList(&input)

			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			if len(res) == 0 {
				err = errors.New("article \"" + title + "\" is not valid.")
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			contractArticle, err := getProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
				ArticleID: &res[0].Id,
			})

			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			if len(contractArticle.Data) > 0 {
				article.ID = contractArticle.Data[0].Id
			}

			vatPercentage, _ := strconv.ParseFloat(res[0].VatPercentage, 32)
			vatFloat32 := float32(vatPercentage)
			article.ArticleID = res[0].Id
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

func ReadArticlesHandler(w http.ResponseWriter, r *http.Request) {
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

			article.PublicProcurementId = publicProcurementID

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

func ReadArticlesDonationHandler(w http.ResponseWriter, r *http.Request) {
	var response ProcurementArticleResponse

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
					if value == "" {
						break
					}

					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					article.NetPrice = float32(floatValue)

				}
			}
			articles = append(articles, article)
		}
	}
	response.Data = articles
	response.Status = "success"
	response.Message = "File was read successfuly"
	_ = MarshalAndWriteJSON(w, response)
}

func ReadArticlesSimpleProcurementHandler(w http.ResponseWriter, r *http.Request) {
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

			article.PublicProcurementId = publicProcurementID

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

func ReadExpireInventoriesHandler(w http.ResponseWriter, r *http.Request) {
	var response ExpireInventoriesResponse

	xlsFile, err := openExcelFile(r)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

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

			var dispatch structs.BasicInventoryAssessmentsTypesItem
			dispatch.Type = "financial"
			dispatch.Active = true

			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					inventoryID, _ := strconv.Atoi(value)
					dispatch.InventoryId = inventoryID
				case 1:
					continue
				case 2:
					floatValue, err := strconv.ParseFloat(value, 32)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
					dispatch.GrossPriceDifference = float32(floatValue)
				case 3:
					dispatch.DateOfAssessment = &value
				}
			}

			item, err := getInventoryItem(dispatch.InventoryId)
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			dispatch.DepreciationTypeId = item.DepreciationTypeId
			_, err = createAssessments(&dispatch)

			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
			}
		}
	}

	response.Status = "success"
	response.Message = "Inventory items was updates successfully"
	_ = MarshalAndWriteJSON(w, response)
}

func getProcurementArticlesList(input *dto.GetProcurementArticleListInputMS) ([]*structs.PublicProcurementArticle, error) {
	res := &dto.GetProcurementArticleListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ARTICLES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getProcurementContractArticlesList(input *dto.GetProcurementContractArticlesInput) (*dto.GetProcurementContractArticlesListResponseMS, error) {
	res := &dto.GetProcurementContractArticlesListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.CONTRACT_ARTICLE_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getInventoryItem(id int) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}
	_, err := shared.MakeAPIRequest("GET", config.INVENTORY_ITEM_ENDOPOINT+"/"+strconv.Itoa(id), nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func createAssessments(data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ASSESSMENTS_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
