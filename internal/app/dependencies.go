package app

import (
	"backend-trainee-assignment-2024/config"
	"backend-trainee-assignment-2024/internal/usecase"
	"backend-trainee-assignment-2024/internal/usecase/repo/memory"
	pg "backend-trainee-assignment-2024/internal/usecase/repo/postgres"
	"backend-trainee-assignment-2024/pkg/postgres"
	"github.com/dgraph-io/ristretto"
)

func newDependencies(db *postgres.Postgres, cache *ristretto.Cache, cfg config.CACHE) usecase.Dependencies {
	pgRepositories := pg.NewRepositories(db)
	memoryRepositories := memory.NewRepositories(cache, cfg.TTL)
	return usecase.Dependencies{
		Pg:     pgRepositories,
		Memory: memoryRepositories,
	}
}
