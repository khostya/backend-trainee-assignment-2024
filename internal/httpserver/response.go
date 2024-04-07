package httpserver

import (
	"github.com/go-chi/render"
	"net/http"
)

type Err struct {
	Error string `json:"error"`
}

func Error(status int, err error, r *http.Request, w http.ResponseWriter) {
	render.Status(r, status)
	render.JSON(w, r, Err{Error: err.Error()})
}

func Response(status int, any any, r *http.Request, w http.ResponseWriter) {
	render.Status(r, status)
	render.JSON(w, r, any)
}
