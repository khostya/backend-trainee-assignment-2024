package memory

import (
	"backend-trainee-assignment-2024/internal/entity"
	"backend-trainee-assignment-2024/internal/model"
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

func (mem Banner) Get(key any) (entity.Banner, error) {
	obj, ok := mem.cache.Get(key)
	if !ok {
		return entity.Banner{}, model.ErrNotFound
	}
	return *obj.(*entity.Banner), nil
}

func (mem Banner) Delete(id int) error {
	banner, err := mem.Get(id)
	if err != nil {
		return err
	}
	mem.cache.Del(banner.Id)
	for _, tag := range banner.Tags {
		key := Key{FeatureId: tag.FeatureId, TagId: tag.TagId}.String()
		mem.cache.Del(key)
	}
	mem.cache.Del(id)
	return nil
}

func (mem Banner) Set(banner entity.Banner) {
	for _, tag := range banner.Tags {
		key := Key{FeatureId: tag.FeatureId, TagId: tag.TagId}.String()
		mem.cache.SetWithTTL(key, &banner, 1, mem.ttl)
	}
	mem.cache.SetWithTTL(banner.Id, &banner, 1, mem.ttl)
}
