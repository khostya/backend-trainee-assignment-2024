package banner

import (
	"backend-trainee-assignment-2024/internal/httpserver/middleware/auth"
	"backend-trainee-assignment-2024/internal/model"
	"backend-trainee-assignment-2024/internal/usecase"
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"strconv"
)

type Router struct {
	banner usecase.Banner
}

func New(r chi.Router, useCase usecase.Banner) {
	banner := Router{banner: useCase}

	r.Use(auth.Required())
	r.Get("/user_banner", banner.getUserBanner)
	r.Get("/banner", banner.get)
	r.Group(func(r chi.Router) {
		r.Use(auth.AdminOnly())

		r.Post("/banner", banner.create)
		r.Patch("/banner/{id}", banner.update)
		r.Delete("/banner/{id}", banner.deleteById)
	})
}

type Banner struct {
	TagIds    []int  `json:"tag_ids"`
	FeatureId int    `json:"feature_id"`
	Content   string `json:"content" validate:"json"`
	IsActive  bool   `json:"is_active"`
}

type Create struct {
	TagIds    []int `json:"tag_ids"`
	FeatureId int   `json:"feature_id"`
	Content   any   `json:"content"`
	IsActive  bool  `json:"is_active"`
}

func (b Router) parseBanner(reader io.ReadCloser) (Banner, error) {
	request := new(Create)
	err := render.DecodeJSON(reader, request)
	if err != nil {
		return Banner{}, err
	}

	content := &bytes.Buffer{}
	enc := json.NewEncoder(content)
	if err := enc.Encode(request.Content); err != nil {
		return Banner{}, err
	}

	banner := Banner{
		IsActive:  request.IsActive,
		Content:   content.String(),
		TagIds:    request.TagIds,
		FeatureId: request.FeatureId,
	}
	return banner, nil
}

func (b Router) parseFilter(r *http.Request) model.Filter {
	featureId, featureErr := strconv.Atoi(r.URL.Query().Get("feature_id"))

	tagId, tagErr := strconv.Atoi(r.URL.Query().Get("tag_id"))

	filter := model.Filter{
		TagId:     sql.NullInt32{Valid: tagErr == nil, Int32: int32(tagId)},
		FeatureId: sql.NullInt32{Valid: featureErr == nil, Int32: int32(featureId)},
	}

	return filter
}

func (b Router) parsePage(r *http.Request) model.Page {
	limit, limitErr := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, offsetErr := strconv.Atoi(r.URL.Query().Get("offset"))

	page := model.Page{
		Limit:  sql.NullInt32{Int32: int32(limit), Valid: limitErr == nil},
		Offset: sql.NullInt32{Int32: int32(offset), Valid: offsetErr == nil},
	}

	return page
}

func (b Router) isAdmin(r *http.Request) bool {
	return r.Context().Value("token") == "admin"
}
