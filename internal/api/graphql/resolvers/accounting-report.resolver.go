package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"encoding/json"
	"strings"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) OverallSpendingResolver(params graphql.ResolveParams) (interface{}, error) {
	var data dto.OveralSpendingFilter

	dataBytes, _ := json.Marshal(params.Args["data"])
	_ = json.Unmarshal(dataBytes, &data)

	flagArticles := false

	if len(data.Articles) > 0 {
		flagArticles = true
	}

	articles, err := r.Repo.GetMovementArticleList(data)

	if err != nil {
		return errors.HandleAPPError(err)
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
