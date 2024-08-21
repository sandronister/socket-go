package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sandronister/socket-go/pkg/devices/types"
)

type DeviceRepository struct {
	conn *pgxpool.Pool
}

func NewDeviceRepository(conn *pgxpool.Pool) *DeviceRepository {
	return &DeviceRepository{conn: conn}
}

func (d *DeviceRepository) GetBlackList() ([]types.BlackListDevice, error) {
	return nil, nil
}
