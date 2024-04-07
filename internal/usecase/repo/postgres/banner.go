package postgres

import (
	"backend-trainee-assignment-2024/internal/entity"
	"backend-trainee-assignment-2024/internal/model"
	"backend-trainee-assignment-2024/pkg/postgres"
	"context"
	"database/sql"
	"errors"
	"github.com/uptrace/bun"
	"strconv"
)

type Banner struct {
	db *bun.DB
}

func NewBanner(pg *postgres.Postgres) Banner {
	return Banner{db: pg.DB}
}

func (r Banner) Create(ctx context.Context, banner entity.Banner) (int, error) {
	return banner.Id, r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(&banner).Exec(ctx)
		if err != nil {
			return err
		}

		for i := range banner.Tags {
			banner.Tags[i].BannerId = banner.Id
		}

		_, err = tx.NewInsert().Model(&banner.Tags).Exec(ctx)
		return err
	})
}

func (r Banner) Update(ctx context.Context, id int, content string) error {
	exec, err := r.db.NewUpdate().
		Model(&entity.Banner{}).
		Where("id = ?", id).
		Set("content = ?", content).
		Exec(ctx)
	affected, _ := exec.RowsAffected()
	if affected == 0 {
		return entity.ErrNotFound
	}
	return err
}

func (r Banner) DeleteById(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().Model(&entity.Banner{}).Where("id = ?", id).Exec(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.ErrNotFound
	}
	return err
}

func (r Banner) Get(ctx context.Context, filter model.Filter, page model.Page) ([]entity.Banner, error) {
	where := ""
	if filter.TagId.Valid {
		where = "tag_id = " + strconv.FormatInt(int64(filter.TagId.Int32), 10)
	}
	if filter.FeatureId.Valid {
		where += "and "
		where += "feature_id = " + strconv.FormatInt(int64(filter.FeatureId.Int32), 10)
	}

	bannersIds := r.db.NewSelect().
		Model(&entity.Tag{}).
		Column("banner_id").
		Where(where).
		Distinct()

	var banners []entity.Banner
	_, err := r.db.NewSelect().
		Model(&banners).
		Where("id in (?)", bannersIds).
		Limit(page.Limit).
		Offset(page.Offset).
		Relation("TagsIds").
		Exec(ctx)
	return banners, err
}

func (r Banner) GetById(ctx context.Context, id int) (entity.Banner, error) {
	var banner = new(entity.Banner)
	err := r.db.NewSelect().Model(banner).
		Where("id = ?", id).
		Relation("Tags").
		Scan(ctx)
	if err != nil {
		return entity.Banner{}, err
	}
	return *banner, err
}

func (r Banner) getTagsQuery(featureId int) *bun.SelectQuery {
	return r.db.NewSelect().
		Column("tag_id").
		Table("banners.tags").
		Where("feature_id = ?", featureId)
}

func (r Banner) GetForUser(ctx context.Context, filter model.Filter) (entity.Banner, error) {
	var (
		banner    = new(entity.Banner)
		tagsQuery = r.getTagsQuery(int(filter.FeatureId.Int32))
	)

	err := r.db.NewSelect().
		Model(banner).
		Where("feature_id = ? and ? in (?)", filter.FeatureId, filter.TagId, tagsQuery).
		Scan(ctx)
	if err != nil {
		return entity.Banner{}, err
	}
	return *banner, nil
}
