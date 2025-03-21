package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"gravitum-test/api"
	"gravitum-test/config"
	"gravitum-test/repository/users"
	users2 "gravitum-test/service/users"
)

func main() {
	cfg := config.Read()
	db := newDatabase(cfg)
	err := goose.Up(db.DB, "migrations")
	if err != nil {
		panic(err)
	}

	usersRepository := users.NewRepository(db)
	usersService := users2.NewService(usersRepository)
	server := api.NewServer(cfg, usersService)
	server.ListenAndServe()
}

func newDatabase(cfg config.Settings) *sqlx.DB {
	pool, err := pgxpool.New(context.Background(), cfg.Database.PostgresConnection)
	if err != nil {
		panic(err)
	}

	return sqlx.NewDb(stdlib.OpenDBFromPool(pool), "pgx")
}
