package files

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func ReadArticlesPriceHandler(w http.ResponseWriter, r *http.Request) {
	var response ArticleResponse

	procurementID := r.FormValue("public_procurement_id")

	publicProcurementID, err := strconv.Atoi(procurementID)

	if err != nil {
		response.Message = "You must provide valid public_procurement_id"
		response.Error = err.Error()
		response.Status = "failed"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contractid := r.FormValue("contract_id")

	contractID, err := strconv.Atoi(contractid)

	if err != nil {
		response.Message = "You must provide valid contract_id"
		response.Error = err.Error()
		response.Status = "failed"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	xlsFile, err := openFile(r)

	if err != nil {
		response.Message = "Error during opening file"
		response.Error = err.Error()
		response.Status = "failed"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var articles []ContractArticleResponseDTO

	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			response.Message = "Error during reading file"
			response.Error = err.Error()
			response.Status = "failed"
			_ = MarshalAndWriteJSON(w, response)
			w.WriteHeader(http.StatusInternalServerError)
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
				response.Message = "Error during reading column value"
				response.Error = err.Error()
				response.Status = "failed"
				_ = MarshalAndWriteJSON(w, response)
				w.WriteHeader(http.StatusInternalServerError)
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
						response.Message = "Error during converting neto price"
						response.Error = err.Error()
						response.Status = "failed"
						_ = MarshalAndWriteJSON(w, response)
						w.WriteHeader(http.StatusInternalServerError)
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
				response.Message = "Error during fetching articles"
				response.Error = err.Error()
				response.Status = "failed"
				_ = MarshalAndWriteJSON(w, response)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if len(res) == 0 {
				response.Message = "Artiklal \"" + title + "\" nije validan."
				response.Status = "failed"
				_ = MarshalAndWriteJSON(w, response)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			contractArticle, err := getProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
				ArticleID: &res[0].Id,
			})

			if err != nil {
				response.Message = "Error during checking article"
				response.Error = err.Error()
				response.Status = "failed"
				_ = MarshalAndWriteJSON(w, response)
				w.WriteHeader(http.StatusInternalServerError)
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
		response.Message = "You must provide valid public_procurement_id"
		response.Error = err.Error()
		response.Status = "failed"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	xlsFile, err := openFile(r)

	if err != nil {
		response.Message = "Error during opening file"
		response.Error = err.Error()
		response.Status = "failed"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
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
			response.Message = "Error during reading file"
			response.Error = err.Error()
			response.Status = "failed"
			_ = MarshalAndWriteJSON(w, response)
			w.WriteHeader(http.StatusInternalServerError)
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
				response.Message = "Error during reading column value"
				response.Error = err.Error()
				response.Status = "failed"
				_ = MarshalAndWriteJSON(w, response)
				w.WriteHeader(http.StatusInternalServerError)
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
						response.Message = "Error during converting neto price"
						response.Error = err.Error()
						response.Status = "failed"
						_ = MarshalAndWriteJSON(w, response)
						w.WriteHeader(http.StatusInternalServerError)
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
							response.Message = "Error during converting vat percentage"
							response.Error = err.Error()
							response.Status = "failed"
							_ = MarshalAndWriteJSON(w, response)
							w.WriteHeader(http.StatusInternalServerError)
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

func ReadArticlesSimpleProcurementHandler(w http.ResponseWriter, r *http.Request) {
	var response ProcurementArticleResponse

	procurementID := r.FormValue("public_procurement_id")

	publicProcurementID, err := strconv.Atoi(procurementID)

	if err != nil {
		response.Message = "You must provide valid public_procurement_id"
		response.Error = err.Error()
		response.Status = "failed"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	xlsFile, err := openFile(r)

	if err != nil {
		response.Message = "Error during opening file"
		response.Error = err.Error()
		response.Status = "failed"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
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
			response.Message = "Error during reading file"
			response.Error = err.Error()
			response.Status = "failed"
			_ = MarshalAndWriteJSON(w, response)
			w.WriteHeader(http.StatusInternalServerError)
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
				response.Message = "Error during reading column value"
				response.Error = err.Error()
				response.Status = "failed"
				_ = MarshalAndWriteJSON(w, response)
				w.WriteHeader(http.StatusInternalServerError)
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
						response.Message = "Error during converting neto price"
						response.Error = err.Error()
						response.Status = "failed"
						_ = MarshalAndWriteJSON(w, response)
						w.WriteHeader(http.StatusInternalServerError)
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
							response.Message = "Error during converting vat percentage"
							response.Error = err.Error()
							response.Status = "failed"
							_ = MarshalAndWriteJSON(w, response)
							w.WriteHeader(http.StatusInternalServerError)
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
						response.Message = "Error during converting amount"
						response.Error = err.Error()
						response.Status = "failed"
						_ = MarshalAndWriteJSON(w, response)
						w.WriteHeader(http.StatusInternalServerError)
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

func openFile(r *http.Request) (*excelize.File, error) {
	maxFileSize := int64(100 * 1024 * 1024) // file maximum 100 MB

	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		return nil, err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("", "uploaded-file-")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, err
	}

	xlsFile, err := excelize.OpenFile(tempFile.Name())

	if err != nil {
		return nil, err
	}

	return xlsFile, nil
}
