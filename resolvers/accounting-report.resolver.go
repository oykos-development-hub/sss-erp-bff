package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"strconv"

	"github.com/graphql-go/graphql"
)

var OverallSpendingResolver = func(params graphql.ResolveParams) (interface{}, error) {
	year, yearOK := params.Args["year"].(string)
	officeID, officeIDOK := params.Args["office_id"].(int)
	search, searchOK := params.Args["search"].(string)
	exception, exceptionOK := params.Args["exception"].(bool)
	orgUnitID, orgUnitIDOK := params.Args["organization_unit_id"].(int)

	var filter dto.OveralSpendingFilter

	if officeIDOK && officeID > 0 {
		filter.OfficeID = &officeID
	}

	if yearOK && year != "" {
		filter.Year = &year
	}

	if searchOK && search != "" {
		filter.Title = &search
	}

	if exceptionOK {
		filter.Exception = &exception
	}

	if orgUnitIDOK && orgUnitID != 0 {
		filter.OrganizationUnitID = &orgUnitID
	}

	articles, err := getMovementArticleList(filter)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	var response []dto.ArticleReport

	if orgUnitIDOK && orgUnitID > 0 {
		organizationUnitID := strconv.Itoa(orgUnitID)
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
						response = append(response, article)
					}
				}
			}
		}
	} else {
		response = articles
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
