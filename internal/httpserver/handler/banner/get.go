package banner

import (
	"backend-trainee-assignment-2024/internal/httpserver"
	"backend-trainee-assignment-2024/internal/model"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (b Router) getUserBanner(w http.ResponseWriter, r *http.Request) {
	filter, err := b.parseTagAndFeatureIds(r)
	if err != nil {
		httpserver.Error(http.StatusBadRequest, err, r, w)
		return
	}

	useLastRevision, err := strconv.ParseBool(chi.URLParam(r, "use_last_revision"))
	if err != nil {
		httpserver.Error(http.StatusBadRequest, err, r, w)
		return
	}

	banner, err := b.banner.GetUserBanner(r.Context(), filter, useLastRevision)
	if err != nil {
		httpserver.Error(http.StatusInternalServerError, err, r, w)
		return
	}

	httpserver.Response(http.StatusOK, banner.Content, r, w)
}

func (b Router) get(w http.ResponseWriter, r *http.Request) {
	filter, err := b.parseTagAndFeatureIds(r)
	if err != nil {
		httpserver.Error(http.StatusBadRequest, err, r, w)
		return
	}

	page, err := b.parsePage(r)
	if err != nil {
		httpserver.Error(http.StatusBadRequest, err, r, w)
		return
	}

	banners, err := b.banner.Get(r.Context(), filter, page)
	if err != nil {
		httpserver.Error(http.StatusInternalServerError, err, r, w)
		return
	}

	httpserver.Response(http.StatusOK, model.NewBanners(banners), r, w)
}
