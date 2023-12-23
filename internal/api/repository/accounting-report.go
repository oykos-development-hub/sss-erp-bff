package repository

import (
	"bff/internal/api/dto"
)

func (repo *MicroserviceRepository) GetMovementArticleList(filter dto.OveralSpendingFilter) ([]dto.ArticleReport, error) {
	res := &dto.ArticleReportMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.MovementReport, filter, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
