package vvperrors

import (
	"errors"
	"net/http"
)

var (
	// TODO: make custom types or wrap w/ response body content
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unathorized")
	ErrForbidden    = errors.New("forbidden")
	ErrConflict     = errors.New("conflict")
	ErrNotFound     = errors.New("not found")
	ErrUnknown      = errors.New("unknown error")

	ErrAuthContext = errors.New("couldn't get authorized context")
)


func IsResponseError(res *http.Response) bool {
	return res != nil && res.StatusCode >= 400
}

func FormatResponseError(res *http.Response) error {
	switch res.StatusCode {
	case 400:
		return ErrBadRequest
	case 401:
		return ErrUnauthorized
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	case 409:
		return ErrConflict
	default:
		return ErrUnknown
	}
}