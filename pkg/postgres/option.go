package postgres

import (
	"github.com/uptrace/bun"
	"time"
)

func MaxOpenConnsDB(max int) bun.DBOption {
	return func(db *bun.DB) {
		db.SetMaxOpenConns(max)
	}
}

func MaxIdleConnsDB(max int) bun.DBOption {
	return func(db *bun.DB) {
		db.SetMaxIdleConns(max)
	}
}

func ConnMaxLifetimeDB(max time.Duration) bun.DBOption {
	return func(db *bun.DB) {
		db.SetConnMaxLifetime(max)
	}
}

func ConnMaxIdleTimeDB(max time.Duration) bun.DBOption {
	return func(db *bun.DB) {
		db.SetConnMaxIdleTime(max)
	}
}
