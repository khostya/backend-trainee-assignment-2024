package httpserver

import (
	"backend-trainee-assignment-2024/internal/model"
	"github.com/go-chi/render"
	"net/http"
)

func Error(status int, err error, r *http.Request, w http.ResponseWriter) {
	render.Status(r, status)
	render.JSON(w, r, model.Error{Error: err.Error()})
}

func Response(status int, any any, r *http.Request, w http.ResponseWriter) {
	render.Status(r, status)
	render.JSON(w, r, any)
}
