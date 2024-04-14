package banner

import (
	"backend-trainee-assignment-2024/internal/httpserver"
	"backend-trainee-assignment-2024/internal/model"
	"errors"
	"net/http"
	"strconv"
)

func (b Router) getUserBanner(w http.ResponseWriter, r *http.Request) {
	isAdmin := b.isAdmin(r)
	filter := b.parseFilter(r)
	if !filter.TagId.Valid || !filter.FeatureId.Valid {
		httpserver.Error(http.StatusBadRequest, errors.New("feature_id and tag_id is required"), r, w)
		return
	}

	useLastRevision, _ := strconv.ParseBool(r.URL.Query().Get("use_last_revision"))

	banner, err := b.banner.GetUserBanner(r.Context(), filter, useLastRevision, isAdmin)
	if errors.Is(err, model.ErrNotFound) {
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
	isAdmin := b.isAdmin(r)
	filter := b.parseFilter(r)
	page := b.parsePage(r)

	banners, err := b.banner.Get(r.Context(), filter, page, isAdmin)
	if err != nil {
		httpserver.Error(http.StatusInternalServerError, err, r, w)
		return
	}

	httpserver.Response(http.StatusOK, model.NewBanners(banners), r, w)
}
