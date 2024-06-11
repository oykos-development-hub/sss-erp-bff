package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateRevision(ctx context.Context, revision *structs.Revision) (*structs.Revision, error) {
	res := &dto.GetRevisionResponseMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.Revisions, revision, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetRevisors() ([]*structs.Revisor, error) {
	res := &dto.GetRevisors{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Revisors, nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateRevision(ctx context.Context, id int, revision *structs.Revision) (*structs.Revision, error) {
	res := &dto.GetRevisionResponseMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.Revisions+"/"+strconv.Itoa(id), revision, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteRevision(ctx context.Context, id int) error {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.Revisions+"/"+strconv.Itoa(id), nil, nil, header)

	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetRevisionByID(id int) (*structs.Revision, error) {
	res := &dto.GetRevisionResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Revisions+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetRevisionList(input *dto.GetRevisionsInput) (*dto.GetRevisionListResponseMS, error) {
	res := &dto.GetRevisionListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Revisions, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ------------------------

func (repo *MicroserviceRepository) GetRevisionPlanList(input *dto.GetPlansInput) (*dto.GetRevisionPlanResponseMS, error) {
	res := &dto.GetRevisionPlanResponseMS{}

	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.RevisionPlan, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetRevisionPlanByID(id int) (*dto.RevisionPlanItem, error) {
	res := &dto.GetPlanResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.RevisionPlan+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteRevisionPlan(ctx context.Context, id int) error {

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.RevisionPlan+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateRevisionPlan(ctx context.Context, plan *dto.RevisionPlanItem) (*dto.RevisionPlanItem, error) {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	res := &dto.GetPlanResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.RevisionPlan, plan, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateRevisionPlan(ctx context.Context, id int, plan *dto.RevisionPlanItem) (*dto.RevisionPlanItem, error) {
	res := &dto.GetPlanResponseMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.RevisionPlan+"/"+strconv.Itoa(id), plan, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// --------------------

func (repo *MicroserviceRepository) GetRevisionsList(input *dto.GetRevisionFilter) (*dto.GetRevisionsResponseMS, error) {
	res := &dto.GetRevisionsResponseMS{}

	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Revision, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetRevisionsByID(id int) (*structs.Revisions, error) {
	res := &dto.GetRevisionMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Revision+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteRevisions(ctx context.Context, id int) error {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.Revision+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateRevisions(ctx context.Context, plan *structs.Revisions) (*structs.Revisions, error) {
	res := &dto.GetRevisionMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.Revision, plan, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateRevisions(ctx context.Context, id int, plan *structs.Revisions) (*structs.Revisions, error) {
	res := &dto.GetRevisionMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.Revision+"/"+strconv.Itoa(id), plan, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// --------------

func (repo *MicroserviceRepository) GetRevisionTipsList(input *dto.GetRevisionTipFilter) (*dto.GetRevisionTipsResponseMS, error) {
	res := &dto.GetRevisionTipsResponseMS{}

	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.RevisionTips, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetRevisionTipByID(id int) (*structs.RevisionTips, error) {
	res := &dto.GetRevisionTipMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.RevisionTips+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteRevisionTips(ctx context.Context, id int) error {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.RevisionTips+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateRevisionTips(ctx context.Context, plan *structs.RevisionTips) (*structs.RevisionTips, error) {
	res := &dto.GetRevisionTipMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.RevisionTips, plan, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateRevisionTips(ctx context.Context, id int, plan *structs.RevisionTips) (*structs.RevisionTips, error) {
	res := &dto.GetRevisionTipMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.RevisionTips+"/"+strconv.Itoa(id), plan, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteRevisionRevisor(id int) error {

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.RevisionRevisors+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateRevisionRevisor(plan *dto.RevisionRevisor) error {
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.RevisionRevisors, plan, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetRevisionRevisorList(input *dto.RevisionRevisorFilter) ([]*dto.RevisionRevisor, error) {
	res := &dto.GetRevisionRevisorResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.RevisionRevisors, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteRevisionOrgUnit(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.RevisionOrgUnit+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateRevisionOrgUnit(plan *dto.RevisionOrgUnit) error {
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.RevisionOrgUnit, plan, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetRevisionOrgUnitList(input *dto.RevisionOrgUnitFilter) ([]*dto.RevisionOrgUnit, error) {
	res := &dto.GetRevisionOrgUnitResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.RevisionOrgUnit, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
