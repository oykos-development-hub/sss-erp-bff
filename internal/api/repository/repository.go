package repository

import (
	"bff/config"
)

var _ MicroserviceRepositoryInterface = &MicroserviceRepository{}

type MicroserviceRepository struct {
	Config *config.Config
}

func NewMicroserviceRepository(cfg *config.Config) *MicroserviceRepository {
	return &MicroserviceRepository{
		Config: cfg,
	}
}
