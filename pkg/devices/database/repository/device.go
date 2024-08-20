package repository

import (
	"context"

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
	rows, err := d.conn.Query(context.Background(), "SELECT imei, blocked FROM go_parser.blacklist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []types.BlackListDevice
	for rows.Next() {
		var r types.BlackListDevice
		err = rows.Scan(&r.Imei, &r.Blocked)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}
