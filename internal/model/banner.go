package model

import (
	"backend-trainee-assignment-2024/internal/entity"
	"time"
)

type Banner struct {
	Id int `json:"id"`

	Tags      []int  `json:"tags"`
	FeatureId int    `json:"feature_id"`
	Content   string `json:"content"`
	IsActive  bool   `json:"is_active"`

	UpdatedAt time.Time `json:"updated_at" bun:",nullzero"`
	CreatedAt time.Time `json:"created_at" bun:",nullzero"`
}

func NewBanner(banner entity.Banner) Banner {
	tags := make([]int, 0, len(banner.Tags))
	for _, tag := range banner.Tags {
		tags = append(tags, tag.TagId)
	}
	return Banner{
		Id:        banner.Id,
		Tags:      tags,
		FeatureId: banner.FeatureId,
		Content:   banner.Content,
		IsActive:  banner.IsActive,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}
}

func NewBanners(banners []entity.Banner) []Banner {
	dtos := make([]Banner, 0, len(banners))

	for _, banner := range banners {
		dtos = append(dtos, NewBanner(banner))
	}

	return dtos
}
