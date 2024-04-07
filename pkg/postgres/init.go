package postgres

import (
	"backend-trainee-assignment-2024/internal/entity"
	"github.com/uptrace/bun"
)

func initDB(db *bun.DB) {
	db.RegisterModel((*entity.Tag)(nil))
}
