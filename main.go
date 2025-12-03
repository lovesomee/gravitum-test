package main
import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"go.uber.org/zap"

	"gravitum-test/api"
	"gravitum-test/config"
	"gravitum-test/logger"
	"gravitum-test/repository/users"
	users2 "gravitum-test/service/users"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logg, err := logger.New()
	if err != nil {
		panic(err)
	}
	defer func() { _ = logg.Sync() }()

	cfg := config.Read()

	db := newDatabase(ctx, cfg, logg)
	if err := goose.Up(db.DB, "migrations"); err != nil {
		logg.Fatal("failed to run database migrations", zap.Error(err))
	}

	usersRepository := users.NewRepository(db, logg)
	usersService := users2.NewService(usersRepository, logg)
	server := api.NewServer(cfg, logg, usersService)

	logg.Info("starting http server", zap.Int("port", cfg.Port))
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logg.Fatal("server failed", zap.Error(err))
	}
}

func newDatabase(ctx context.Context, cfg config.Settings, logg *zap.Logger) *sqlx.DB {
	pool, err := pgxpool.New(ctx, cfg.Database.PostgresConnection)
	if err != nil {
		logg.Fatal("failed to create postgres pool", zap.Error(err))
	}

	logg.Info("connected to database")
	return sqlx.NewDb(stdlib.OpenDBFromPool(pool), "pgx")
}
