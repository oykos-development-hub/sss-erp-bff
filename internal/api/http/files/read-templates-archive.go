package files

/*
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

		offices, err := h.Repo.GetOfficeDropdownSettings(&dto.GetOfficesOfOrganizationInput{})

		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		mapOfOffices := make(map[string]int)

		for _, obj := range offices.Data {
			mapOfOffices[obj.Title] = obj.ID
		}
		var fals bool
		organizationUnits, err := h.Repo.GetOrganizationUnits(
			&dto.GetOrganizationUnitsInput{IsParent: &fals},
		)

		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		mapOfOrganizationUnits := make(map[string]int)

		for _, obj := range organizationUnits.Data {
			mapOfOrganizationUnits[obj.Title] = obj.ID
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
			Type := "movable"
			res, _ := h.Repo.GetAllInventoryItem(dto.InventoryItemFilter{SourceOrganizationUnitID: &organizationUnitID, Type: &Type})

			total := res.Total

			for rows.Next() {
				rowindex++
				if rowindex <= total+9 {
					continue
				}

				if rowindex > total+9+90 {
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
					case 2:
						id, exists := mapOfOffices[value]
						if exists {
							article.Dispatch.OfficeID = id
							article.Article.OfficeID = article.Dispatch.OfficeID
							article.Dispatch.Type = "allocation"
						} else if value != "" {
							//index := strings.Index(value, "-")
							index := -1
							if index == -1 {
								id, exists = mapOfOrganizationUnits[value]
								if !exists && value != "" {
									newOffice := structs.SettingsDropdown{
										Value:  strconv.Itoa(organizationUnitID),
										Title:  value,
										Entity: config.OfficeTypes,
									}

									itemRes, err := h.Repo.CreateDropdownSettings(&newOffice)

									if err != nil {
										responseMessage := ValidationResponse{
											Column:  2,
											Row:     rowindex,
											Message: "Greska prilikom dodavanja kancelarije!",
										}
										response.Validation = append(response.Validation, responseMessage)
									}
									article.Dispatch.OfficeID = itemRes.ID
									article.Article.OfficeID = itemRes.ID
									article.Dispatch.Type = "allocation"
									mapOfOffices[itemRes.Title] = itemRes.ID
								} else {
									if err != nil {
										responseMessage := ValidationResponse{
											Column:  2,
											Row:     rowindex,
											Message: "Lokacija nije validna!",
										}
										response.Validation = append(response.Validation, responseMessage)
									}
								}
								article.Dispatch.IsAccepted = true
								article.Dispatch.SourceOrganizationUnitID = organizationUnitID
								article.Dispatch.TargetOrganizationUnitID = id
								article.Dispatch.Type = "revers"
							} else {
								organizationUnit := value[:index-1]
								office := value[index+2:]

								id, exists = mapOfOrganizationUnits[organizationUnit]
								if !exists && value != "" {
									responseMessage := ValidationResponse{
										Column:  2,
										Row:     rowindex,
										Message: "Lokacija nije validna!",
									}
									response.Validation = append(response.Validation, responseMessage)
								}

								article.ReversDispatch.IsAccepted = true
								article.ReversDispatch.SourceOrganizationUnitID = organizationUnitID
								article.ReversDispatch.TargetOrganizationUnitID = id
								article.ReversDispatch.Type = "revers"
								article.Article.TargetOrganizationUnitID = id

								newOffice := structs.SettingsDropdown{
									Value:  strconv.Itoa(id),
									Title:  office,
									Entity: config.OfficeTypes,
								}

								itemRes, err := h.Repo.CreateDropdownSettings(&newOffice)

								if err != nil {
									responseMessage := ValidationResponse{
										Column:  2,
										Row:     rowindex,
										Message: "Greska prilikom dodavanja kancelarije!",
									}
									response.Validation = append(response.Validation, responseMessage)
								}
								article.Dispatch.OfficeID = itemRes.ID
								article.Article.OfficeID = itemRes.ID
								article.Dispatch.Type = "allocation"
								mapOfOffices[itemRes.Title] = itemRes.ID
							}
						}
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
							response.Validation = append(response.Validation, responseMessage)
						} else {
							article.Article.GrossPrice = float64(price)
							article.FirstAmortization.GrossPriceDifference = float64(price)
						}
					case 14:
						DateOfAssessment, err := parseDate(value)

						if value != "" && err != nil {
							responseMessage := ValidationResponse{
								Column:  14,
								Row:     rowindex,
								Message: "Datum amortizacije nije validno unijet!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else if value != "" {
							dateOfAssessment := DateOfAssessment.Format("2006-01-02T15:04:05Z")
							article.Article.DateOfAssessment = &dateOfAssessment
							article.SecondAmortization.DateOfAssessment = &dateOfAssessment
						}
					case 15:
						grossPriceNew, err := strconv.ParseFloat(value, 32)

						if value != "" && err != nil {
							responseMessage := ValidationResponse{
								Column:  15,
								Row:     rowindex,
								Message: "Nova cijena nije validno unijeta!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else if grossPriceNew > 0 {
							article.SecondAmortization.GrossPriceDifference = float64(grossPriceNew)
							article.Article.GrossPrice = float64(grossPriceNew)
						}
					case 16:
						residualPrice, err := strconv.ParseFloat(value, 32)
						if value != "" && err != nil {
							responseMessage := ValidationResponse{
								Column:  16,
								Row:     rowindex,
								Message: "Rezidualna cijena nije validno unijeta!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else if value != "" {
							residualPricefloat64 := float64(residualPrice)
							article.SecondAmortization.ResidualPrice = &residualPricefloat64
						}
					case 23:
						estimatedDuration, err := strconv.Atoi(value)
						if value != "" && err != nil {
							responseMessage := ValidationResponse{
								Column:  24,
								Row:     rowindex,
								Message: "Vijek trajanja nije validno unijet!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else if estimatedDuration > 0 {
							article.SecondAmortization.EstimatedDuration = estimatedDuration
						}
					case 25:
						article.Article.Description = value
					case 26:
						if _, exists := mapOfClassTypes[value]; !exists && value != "" && value != "0" {
							responseMessage := ValidationResponse{
								Column:  26,
								Row:     rowindex,
								Message: "Klasa sredstva " + value + " nije validna.",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							article.Article.ClassTypeID = mapOfClassTypes[value]
						}
					case 28:
						price, err := strconv.ParseFloat(value, 32)
						if (value != "0" && value != "" && value != "#DIV/0!") && err != nil {
							responseMessage := ValidationResponse{
								Column:  28,
								Row:     rowindex,
								Message: "Cijena nije validno unijeta!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else if value != "" {
							pricefloat64 := float64(price)
							article.Article.AssessmentPrice = pricefloat64
						}
					case 29:
						dateOfPurchase, err := parseDate(value)

						if value != "" && err != nil {
							responseMessage := ValidationResponse{
								Column:  29,
								Row:     rowindex,
								Message: "Datum nabavke nije validno unijet!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							dateOfPurchaseString := dateOfPurchase.Format("2006-01-02T15:04:05Z")
							article.Article.DateOfPurchase = dateOfPurchaseString
							article.Article.DateOfAssessment = &dateOfPurchaseString
							article.FirstAmortization.DateOfAssessment = &dateOfPurchaseString
						}
					case 30:
						if id, exists := mapOfDeprecationTypes[value]; !exists && value != "" {
							responseMessage := ValidationResponse{
								Column:  30,
								Row:     rowindex,
								Message: "Amortizaciona grupa " + value + " nije validna.",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							article.Article.DepreciationTypeID = id
							article.FirstAmortization.DepreciationTypeID = id
							article.SecondAmortization.DepreciationTypeID = id
						}
					case 32:
						estimatedDuration, err := strconv.Atoi(value)
						if value != "" && err != nil {
							responseMessage := ValidationResponse{
								Column:  35,
								Row:     rowindex,
								Message: "Vijek trajanja nije validno unijet!",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else if estimatedDuration > 0 {
							article.FirstAmortization.EstimatedDuration = estimatedDuration
						}
					}
				}
				articles = append(articles, article)
			}
		}

		if len(response.Validation) == 0 {
			defaultTime := "0001-01-01T00:00:00Z"
			for _, article := range articles {
				if article.Article.Title == "" {
					continue
				}
				article.Article.OrganizationUnitID = organizationUnitID
				article.Article.Type = "movable"
				article.Article.Active = true

				if article.SecondAmortization.DateOfAssessment != nil && *article.SecondAmortization.DateOfAssessment != defaultTime {
					article.Article.DateOfAssessment = article.SecondAmortization.DateOfAssessment
				}

				if article.Dispatch.TargetOrganizationUnitID != 0 {
					article.Article.TargetOrganizationUnitID = article.Dispatch.TargetOrganizationUnitID
				}

				newArticle, err := h.Repo.CreateInventoryItem(&article.Article)
				if err != nil {
					handleError(w, err, http.StatusInternalServerError)
					return
				}

				if article.FirstAmortization.EstimatedDuration == 0 {
					article.FirstAmortization.EstimatedDuration = 10000
				}
				article.FirstAmortization.InventoryID = newArticle.ID
				article.FirstAmortization.Type = "financial"
				article.FirstAmortization.Active = true
				_, err = h.Repo.CreateAssessments(&article.FirstAmortization)
				if err != nil {
					handleError(w, err, http.StatusInternalServerError)
					return
				}

				if article.SecondAmortization.DateOfAssessment != nil {
					article.SecondAmortization.InventoryID = newArticle.ID
					article.SecondAmortization.Type = "financial"
					article.SecondAmortization.Active = true
					_, err = h.Repo.CreateAssessments(&article.SecondAmortization)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}
				}
				if article.ReversDispatch.TargetOrganizationUnitID != 0 {
					article.ReversDispatch.Date = article.Article.DateOfPurchase
					article.ReversDispatch.IsAccepted = true

					reversDispatch, err := h.Repo.CreateDispatchItem(&article.ReversDispatch)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}

					article.ReversDispatchItem.InventoryID = newArticle.ID
					article.ReversDispatchItem.DispatchID = reversDispatch.ID

					_, err = h.Repo.CreateDispatchItemItem(&article.ReversDispatchItem)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
					}
				}

				article.Dispatch.Date = article.Article.DateOfPurchase
				article.Dispatch.IsAccepted = true

				dispatch, err := h.Repo.CreateDispatchItem(&article.Dispatch)

				if err != nil {
					handleError(w, err, http.StatusInternalServerError)
					return
				}

				article.DispatchItem.InventoryID = newArticle.ID
				article.DispatchItem.DispatchID = dispatch.ID

				_, err = h.Repo.CreateDispatchItemItem(&article.DispatchItem)
				if err != nil {
					handleError(w, err, http.StatusInternalServerError)
				}
			}
		}

		response.Status = "success"
		response.Message = "The file was read successfully"

		_ = MarshalAndWriteJSON(w, response)
	}

	func (h *Handler) ImportExcelPS2(w http.ResponseWriter, r *http.Request) {
		var response ImportPS1Inventories

		xlsFile, err := openExcelFile(r)

		organizationUnitID := r.FormValue("organization_unit_id")

		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		offices, err := h.Repo.GetOfficeDropdownSettings(&dto.GetOfficesOfOrganizationInput{})

		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		mapOfOffices := make(map[string]int)

		for _, obj := range offices.Data {
			mapOfOffices[obj.Title] = obj.ID
		}

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

		sheetMap := xlsFile.GetSheetMap()
		//var articles []ImportInventoryArticles

		for _, sheetName := range sheetMap {

			if sheetName != "PS -2" {
				continue
			}

			rows, err := xlsFile.Rows(sheetName)
			if err != nil {
				handleError(w, err, http.StatusInternalServerError)
				return
			}

			rowindex := -1

			_, _ = h.Repo.GetAllInventoryItem(dto.InventoryItemFilter{OrganizationUnitID: &rowindex})

			for rows.Next() {
				rowindex++
				if rowindex < 9 {
					continue
				}

				if rowindex > 67 {
					break
				}
				cols := rows.Columns()
				if err != nil {
					handleError(w, err, http.StatusInternalServerError)
					return
				}
				var article ImportInventoryArticles
				outerLoop := true
				for cellIndex, cellValue := range cols {
					value := cellValue
					switch cellIndex {
					case 0:
						if value == "" {
							outerLoop = false
						}
					case 1:
						article.Article.InventoryNumber = value
					case 2:
						officeID, exists := mapOfOffices[value]
						if exists {
							article.Article.OfficeID = officeID
							article.Dispatch.OfficeID = officeID
						} else if value == "" {
							article.Article.Amount = 0
						} else {
							newOffice := structs.SettingsDropdown{
								Value:  organizationUnitID,
								Title:  value,
								Entity: config.OfficeTypes,
							}

							itemRes, err := h.Repo.CreateDropdownSettings(&newOffice)

							if err != nil {
								responseMessage := ValidationResponse{
									Column:  2,
									Row:     rowindex,
									Message: "Greska prilikom dodavanja kancelarije!",
								}
								response.Validation = append(response.Validation, responseMessage)
							}
							article.Dispatch.OfficeID = itemRes.ID
							article.Article.OfficeID = itemRes.ID
							article.Dispatch.Type = "allocation"
							mapOfOffices[itemRes.Title] = itemRes.ID
						}
					case 3:
						article.Article.Title = value
					case 14:
						article.Article.Description = value
					case 15:
						value = "15310000"
						if _, exists := mapOfClassTypes[value]; !exists && value != "" && value != "0" {
							responseMessage := ValidationResponse{
								Column:  26,
								Row:     rowindex,
								Message: "Klasa sredstva " + value + " nije validna.",
							}
							response.Validation = append(response.Validation, responseMessage)
						} else {
							article.Article.ClassTypeID = mapOfClassTypes[value]
						}
					case 18:
						dateOfPurchase, err := parseDate(value)

						if value != "" && err != nil {
							article.Article.DateOfPurchase = "1980-01-01T00:00:00Z"
							article.Article.DateOfAssessment = &article.Article.DateOfPurchase
							article.FirstAmortization.DateOfAssessment = &article.Article.DateOfPurchase
						} else {
							dateOfPurchaseString := dateOfPurchase.Format("2006-01-02T15:04:05Z")
							article.Article.DateOfPurchase = dateOfPurchaseString
							article.Article.DateOfAssessment = &dateOfPurchaseString
							article.FirstAmortization.DateOfAssessment = &dateOfPurchaseString
						}
					}
				}
				if outerLoop {
					if article.Article.Title == "" {
						continue
					}
					article.Article.OrganizationUnitID = 3
					article.Article.TargetOrganizationUnitID, _ = strconv.Atoi(organizationUnitID)
					article.Article.Type = "movable"
					article.Article.Active = true
					article.Article.DepreciationTypeID = 168

					newArticle, err := h.Repo.CreateInventoryItem(&article.Article)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}

					if article.FirstAmortization.EstimatedDuration == 0 {
						article.FirstAmortization.EstimatedDuration = 10
					}
					article.FirstAmortization.InventoryID = newArticle.ID
					article.FirstAmortization.Type = "financial"
					article.FirstAmortization.Active = true
					article.FirstAmortization.DateOfAssessment = article.Article.DateOfAssessment
					article.FirstAmortization.DepreciationTypeID = 166
					_, err = h.Repo.CreateAssessments(&article.FirstAmortization)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}

					article.ReversDispatch.TargetOrganizationUnitID, _ = strconv.Atoi(organizationUnitID)
					article.ReversDispatch.Date = article.Article.DateOfPurchase
					article.ReversDispatch.IsAccepted = true
					article.ReversDispatch.Type = "revers"

					reversDispatch, err := h.Repo.CreateDispatchItem(&article.ReversDispatch)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}

					article.ReversDispatchItem.InventoryID = newArticle.ID
					article.ReversDispatchItem.DispatchID = reversDispatch.ID

					_, err = h.Repo.CreateDispatchItemItem(&article.ReversDispatchItem)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
					}

					article.Dispatch.Date = article.Article.DateOfPurchase
					article.Dispatch.IsAccepted = true

					dispatch, err := h.Repo.CreateDispatchItem(params.Context, &article.Dispatch)

					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
						return
					}

					article.DispatchItem.InventoryID = newArticle.ID
					article.DispatchItem.DispatchID = dispatch.ID

					_, err = h.Repo.CreateDispatchItemItem(&article.DispatchItem)
					if err != nil {
						handleError(w, err, http.StatusInternalServerError)
					}
				}
			}
		}
		response.Status = "success"
		response.Message = "File was read successfuly"
		_ = MarshalAndWriteJSON(w, response)
	}
*/
