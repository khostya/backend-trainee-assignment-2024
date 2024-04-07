package app

import (
	"backend-trainee-assignment-2024/internal/usecase"
	"backend-trainee-assignment-2024/internal/usecase/repo/memory"
	pg "backend-trainee-assignment-2024/internal/usecase/repo/postgres"
	"backend-trainee-assignment-2024/pkg/postgres"
	"github.com/dgraph-io/ristretto"
	"time"
)

func newDependencies(db *postgres.Postgres, cache *ristretto.Cache, ttl time.Duration) usecase.Dependencies {
	pgRepositories := pg.NewRepositories(db)
	memoryRepositories := memory.NewRepositories(cache, ttl)
	return usecase.Dependencies{
		Pg:     pgRepositories,
		Memory: memoryRepositories,
	}
}
