package entity

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Banner struct {
	bun.BaseModel `bun:"table:banners.banners"`

	Id int `json:"id" bun:",pk,autoincrement"`

	Tags      []Tag  `json:"tags" bun:"rel:has-many,join:id=banner_id"`
	FeatureId int    `json:"feature_id"`
	Content   string `json:"content"`
	IsActive  bool   `json:"is_active"`

	UpdatedAt time.Time `json:"updated_at" bun:",nullzero"`
	CreatedAt time.Time `json:"created_at" bun:",nullzero"`
}

func (b *Banner) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		b.CreatedAt = time.Now()
		b.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		b.UpdatedAt = time.Now()
	}
	return nil
}
