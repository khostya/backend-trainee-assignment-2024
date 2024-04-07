package memory

import (
	"backend-trainee-assignment-2024/internal/entity"
	"github.com/dgraph-io/ristretto"
	"time"
)

type Banner struct {
	cache *ristretto.Cache
	ttl   time.Duration
}

func NewBanner(cache *ristretto.Cache, ttl time.Duration) Banner {
	return Banner{cache: cache, ttl: ttl}
}

func (mem Banner) Get(id int) (*entity.Banner, error) {
	obj, ok := mem.cache.Get(id)
	if !ok {
		return nil, entity.ErrNotFound
	}
	return obj.(*entity.Banner), nil
}

func (mem Banner) Delete(id int) {
	mem.cache.Del(id)
}

func (mem Banner) Set(banner entity.Banner) {
	mem.cache.SetWithTTL(banner.Id, &banner, 1, mem.ttl)
}
