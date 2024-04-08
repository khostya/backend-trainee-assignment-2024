package banner

import (
	"backend-trainee-assignment-2024/internal/entity"
	"backend-trainee-assignment-2024/internal/httpserver"
	"net/http"
)

func (b Router) create(w http.ResponseWriter, r *http.Request) {
	request, err := b.parseBanner(r.Body)
	if err != nil {
		httpserver.Error(http.StatusBadRequest, err, r, w)
		return
	}

	banner := entity.Banner{
		IsActive:  request.IsActive,
		FeatureId: request.FeatureId,
		Content:   request.Content,
	}

	for _, tagId := range request.TagIds {
		tag := entity.Tag{TagId: tagId, FeatureId: request.FeatureId}
		banner.Tags = append(banner.Tags, tag)
	}

	id, err := b.banner.Create(r.Context(), banner)
	if err != nil {
		httpserver.Error(http.StatusInternalServerError, err, r, w)
		return
	}

	type Response struct {
		BannerId int `json:"banner_id"`
	}

	httpserver.Response(http.StatusCreated, Response{BannerId: id}, r, w)
}
