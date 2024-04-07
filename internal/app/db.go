package app

import (
	"backend-trainee-assignment-2024/config"
	"backend-trainee-assignment-2024/pkg/postgres"
)

func openDB(cfg config.PG) (*postgres.Postgres, error) {
	return postgres.New(cfg.URL,
		postgres.MaxOpenConnsDB(cfg.MaxOpenConns),
		postgres.ConnMaxIdleTimeDB(cfg.ConnMaxIdleTime),
		postgres.MaxIdleConnsDB(cfg.MaxIdleConns),
		postgres.ConnMaxLifetimeDB(cfg.ConnMaxLifetime),
	)
}
