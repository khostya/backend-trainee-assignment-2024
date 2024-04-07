package banner

import (
	"backend-trainee-assignment-2024/internal/entity"
	handlers "backend-trainee-assignment-2024/internal/httpserver"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (b Router) update(w http.ResponseWriter, r *http.Request) {
	request, err := b.parseBanner(r.Body)
	if err != nil {
		handlers.Error(http.StatusBadRequest, err, r, w)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		handlers.Error(http.StatusBadRequest, err, r, w)
		return
	}

	content := request.Content
	err = b.banner.Update(r.Context(), id, content)
	if errors.Is(err, entity.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
