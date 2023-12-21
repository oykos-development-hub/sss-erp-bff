package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) SyncPermissions(roleID int, input []*structs.RolePermission) ([]structs.RolePermission, error) {
	res := &dto.GetInsertRolesPermissionListResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.ROLES+"/"+strconv.Itoa(roleID)+"/permissions/sync", input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetPermissionList(roleID int) ([]structs.Permissions, error) {
	res := &dto.GetPermissionListForRoleResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.ROLES+"/"+strconv.Itoa(roleID)+"/permissions", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
