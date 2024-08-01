package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) SyncPermissions(roleID int, input []*structs.RolePermission) ([]structs.RolePermission, error) {
	res := &dto.GetInsertRolesPermissionListResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.Roles+"/"+strconv.Itoa(roleID)+"/permissions/sync", input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetPermissionList(roleID int) ([]structs.Permissions, error) {
	res := &dto.GetPermissionListForRoleResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Roles+"/"+strconv.Itoa(roleID)+"/permissions", nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetUsersByPermission(title config.PermissionPath) ([]structs.UserAccounts, error) {
	res := &dto.GetUserAccountListResponseMS{}
	input := structs.Permissions{
		Title: string(title),
	}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.GetUserByPermission, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}
