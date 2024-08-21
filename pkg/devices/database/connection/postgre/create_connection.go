package postgre

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func createConnection(dbUrl string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dbUrl)

	if err != nil {
		return nil, err
	}

	return db, nil

}
