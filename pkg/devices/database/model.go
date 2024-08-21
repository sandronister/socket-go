package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type IDatabase interface {
	GetPostGresConnection() *pgxpool.Pool
}
