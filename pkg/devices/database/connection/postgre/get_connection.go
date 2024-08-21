package postgre

import (
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetPostGresConnection() (*pgxpool.Pool, error) {
	if db == nil {
		conn, err := createConnection(os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}
