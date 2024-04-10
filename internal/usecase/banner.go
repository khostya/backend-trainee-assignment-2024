package usecase

import (
	"backend-trainee-assignment-2024/internal/entity"
	"backend-trainee-assignment-2024/internal/model"
	"backend-trainee-assignment-2024/internal/usecase/repo/memory"
	"backend-trainee-assignment-2024/internal/usecase/repo/postgres"
	"context"
)

type Banner struct {
	pg  postgres.Banner
	mem memory.Banner
}

func NewBannerUseCase(pg postgres.Banner, memory memory.Banner) Banner {
	return Banner{pg, memory}
}

func (uc Banner) Create(ctx context.Context, banner entity.Banner) (int, error) {
	id, err := uc.pg.Create(ctx, banner)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (uc Banner) Update(ctx context.Context, banner entity.Banner) error {
	err := uc.pg.Update(ctx, banner)
	return err
}

func (uc Banner) DeleteById(ctx context.Context, id int) error {
	return uc.pg.DeleteById(ctx, id)
}

func (uc Banner) GetUserBanner(ctx context.Context, filter model.Filter, useLastRevision bool, isAdmin bool) (entity.Banner, error) {
	banner, err := uc.pg.GetForUser(ctx, filter, isAdmin)
	return banner, err
}

func (uc Banner) Get(ctx context.Context, filter model.Filter, page model.Page, isAdmin bool) ([]entity.Banner, error) {
	return uc.pg.Get(ctx, filter, page, isAdmin)
}
