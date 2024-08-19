package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sandronister/socket-go/internal/entities"
)

type DeviceRepository struct {
	conn *pgxpool.Pool
}

func NewDeviceRepository(conn *pgxpool.Pool) *DeviceRepository {
	return &DeviceRepository{conn: conn}
}

func (d *DeviceRepository) SaveMessage(device entities.Ruptela) error {
	_, err := d.conn.Exec(context.Background(), "INSERT INTO go_parser.msgs2 (imei,time,timeunix,data,md5) VALUES ($1,$2,$3,$4,$5)", device.Md5Str)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceRepository) GetBlackList() ([]entities.BlackListDevice, error) {
	rows, err := d.conn.Query(context.Background(), "SELECT imei, blocked FROM go_parser.blacklist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entities.BlackListDevice
	for rows.Next() {
		var r entities.BlackListDevice
		err = rows.Scan(&r.Imei, &r.Blocked)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}
