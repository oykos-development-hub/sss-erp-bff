package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"encoding/json"
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

	if !flagArticles {
		response = articles
	} else {
		for _, article := range articles {
			for _, articleString := range data.Articles {
				parts := strings.SplitN(articleString, " ", 2)

				if strings.Contains(parts[0], article.Year) && strings.Contains(parts[1], article.Title) {
					response = append(response, article)
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
