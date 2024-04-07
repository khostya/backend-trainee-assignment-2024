package entity

import "errors"

var ErrNotFound = errors.New("not found")

type Error struct {
	Error string `json:"error"`
}
