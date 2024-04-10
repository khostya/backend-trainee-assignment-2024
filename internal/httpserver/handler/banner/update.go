package banner

import (
	"backend-trainee-assignment-2024/internal/entity"
	"backend-trainee-assignment-2024/internal/httpserver"
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (b Router) update(w http.ResponseWriter, r *http.Request) {
	request, err := b.parseBanner(r.Body)
	if err != nil {
		httpserver.Error(http.StatusBadRequest, err, r, w)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpserver.Error(http.StatusBadRequest, err, r, w)
		return
	}

	banner := entity.Banner{
		Id:        id,
		Content:   request.Content,
		IsActive:  sql.NullBool{Valid: true, Bool: request.IsActive},
		FeatureId: request.FeatureId,
	}
	for _, tagId := range request.TagIds {
		tag := entity.Tag{TagId: tagId, FeatureId: banner.FeatureId, BannerId: id}
		banner.Tags = append(banner.Tags, tag)
	}

	err = b.banner.Update(r.Context(), banner)
	if errors.Is(err, entity.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		httpserver.Error(http.StatusBadRequest, err, r, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
