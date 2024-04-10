package banner

import (
	"backend-trainee-assignment-2024/internal/entity"
	"backend-trainee-assignment-2024/internal/httpserver"
	"backend-trainee-assignment-2024/internal/model"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (b Router) getUserBanner(w http.ResponseWriter, r *http.Request) {
	filter := b.parseFilter(r)
	if !filter.TagId.Valid || !filter.FeatureId.Valid {
		httpserver.Error(http.StatusBadRequest, errors.New("feature_id and tag_id is required"), r, w)
		return
	}

	useLastRevision, _ := strconv.ParseBool(chi.URLParam(r, "use_last_revision"))

	banner, err := b.banner.GetUserBanner(r.Context(), filter, useLastRevision)
	if errors.Is(err, entity.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		httpserver.Error(http.StatusInternalServerError, err, r, w)
		return
	}

	httpserver.Response(http.StatusOK, banner.Content, r, w)
}

func (b Router) get(w http.ResponseWriter, r *http.Request) {
	filter := b.parseFilter(r)

	page := b.parsePage(r)

	banners, err := b.banner.Get(r.Context(), filter, page)
	if err != nil {
		httpserver.Error(http.StatusInternalServerError, err, r, w)
		return
	}

	httpserver.Response(http.StatusOK, model.NewBanners(banners), r, w)
}
