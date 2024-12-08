package database

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Connect() *bun.DB {
	dns := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	//pgconn := pgdriver.NewConnector(
	//	pgdriver.WithNetwork("tcp"),
	//	pgdriver.WithAddr("db:5432"),
	//	pgdriver.WithUser(os.Getenv("DB_USER")),
	//	pgdriver.WithPassword(os.Getenv("DB_PASSWORD")),
	//	pgdriver.WithDatabase(os.Getenv("DB_DATABASE")),
	//)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dns)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}
