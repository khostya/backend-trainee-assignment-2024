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
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strconv"
)

type Router struct {
	banner    usecase.Banner
	validator *validator.Validate
}

func New(r chi.Router, useCase usecase.Banner) {
	validator := validator.New()

	banner := Router{banner: useCase, validator: validator}

	r.Use(auth.Required())
	r.Get("/user_banner/{id}", banner.getUserBanner)

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

func (b Router) parseBanner(reader io.ReadCloser) (Banner, error) {
	type Request struct {
		TagIds    []int `json:"tag_ids"`
		FeatureId int   `json:"feature_id"`
		Content   any   `json:"content" validate:"json"`
		IsActive  bool  `json:"is_active"`
	}

	request := new(Request)
	err := render.DecodeJSON(reader, request)
	if err != nil {
		return Banner{}, err
	}

	content := &bytes.Buffer{}
	enc := json.NewEncoder(content)
	if err := enc.Encode(content); err != nil {
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

func (b Router) parseTagAndFeatureIds(r *http.Request) (model.Filter, error) {
	featureId, err := strconv.Atoi(chi.URLParam(r, "feature_id"))
	if err != nil {
		return model.Filter{}, err
	}

	tagId, err := strconv.Atoi(chi.URLParam(r, "tag_id"))
	if err != nil {
		return model.Filter{}, err
	}

	filter := model.Filter{
		TagId:     sql.NullInt32{Valid: true, Int32: int32(tagId)},
		FeatureId: sql.NullInt32{Valid: true, Int32: int32(featureId)},
	}

	return filter, nil
}

func (b Router) parsePage(r *http.Request) (model.Page, error) {
	limit, err := strconv.Atoi(chi.URLParam(r, "limit"))
	if err != nil {
		return model.Page{}, err
	}

	offset, err := strconv.Atoi(chi.URLParam(r, "offset"))
	if err != nil {
		return model.Page{}, err
	}

	page := model.Page{
		Limit:  limit,
		Offset: offset,
	}

	return page, nil
}
