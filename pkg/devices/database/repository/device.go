package repository

import (
	"github.com/go-pg/pg"
	"github.com/sandronister/socket-go/pkg/devices/types"
)

type DeviceRepository struct {
	conn *pg.DB
}

func NewDeviceRepository(conn *pg.DB) *DeviceRepository {
	return &DeviceRepository{conn: conn}
}

func (d *DeviceRepository) GetBlackList() ([]types.BlackListDevice, error) {
	return nil, nil
}
