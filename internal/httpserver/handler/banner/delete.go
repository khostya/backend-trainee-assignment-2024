package banner

import (
	"backend-trainee-assignment-2024/internal/httpserver"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (b Router) deleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bannerId, err := strconv.Atoi(id)
	if err != nil {
		httpserver.Error(http.StatusBadRequest, err, r, w)
		return
	}

	err = b.banner.DeleteById(r.Context(), bannerId)
	if err != nil {
		httpserver.Error(http.StatusInternalServerError, err, r, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
