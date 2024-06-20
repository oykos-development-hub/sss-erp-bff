package repository

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
)

func (repo *MicroserviceRepository) GetMovementArticleList(filter dto.OveralSpendingFilter) ([]dto.ArticleReport, error) {
	res := &dto.ArticleReportMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.MovementReport, filter, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}
