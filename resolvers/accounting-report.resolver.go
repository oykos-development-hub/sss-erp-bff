package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
)

var OverallSpendingResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data dto.OveralSpendingFilter

	dataBytes, _ := json.Marshal(params.Args["data"])
	_ = json.Unmarshal(dataBytes, &data)

	flagArticles := false

	if len(data.Articles) > 0 {
		flagArticles = true
	}

	articles, err := getMovementArticleList(data)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	var response []dto.ArticleReport

	if data.OrganizationUnitID != nil && *data.OrganizationUnitID > 0 {
		organizationUnitID := strconv.Itoa(*data.OrganizationUnitID)
		res, err := getOfficeDropdownSettings(&dto.GetOfficesOfOrganizationInput{
			Value: &organizationUnitID,
		})

		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, article := range articles {
			for _, office := range res.Data {
				if article.OfficeID == office.Id {
					found := false
					for i, existingArticle := range response {
						if article.Title == existingArticle.Title && article.Year == existingArticle.Year && article.Description == existingArticle.Description {
							response[i].Amount += article.Amount
							found = true
							break
						}
					}

					if !found {
						if !flagArticles {
							response = append(response, article)
						} else {
							for _, articleString := range data.Articles {
								parts := strings.SplitN(articleString, " ", 2)
								if len(parts) != 2 {
									continue
								}
								if strings.Contains(parts[0], article.Year) && strings.Contains(parts[1], article.Title) {
									response = append(response, article)
								}

							}
						}
					}
				}
			}
		}
	} else {
		if !flagArticles {
			response = articles
		} else {
			for _, article := range articles {
				for _, articleString := range data.Articles {
					parts := strings.SplitN(articleString, " ", 2)
					if len(parts) != 2 {
						continue
					}
					if strings.Contains(parts[0], article.Year) && strings.Contains(parts[1], article.Title) {
						response = append(response, article)
					}
				}
			}
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "You fetched articles!",
		Items:   response,
	}, nil
}

func getMovementArticleList(filter dto.OveralSpendingFilter) ([]dto.ArticleReport, error) {
	res := &dto.ArticleReportMS{}
	_, err := shared.MakeAPIRequest("GET", config.MOVEMENT_REPORT_ENDPOINT, filter, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
