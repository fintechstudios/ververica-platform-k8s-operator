package vvperrors

import (
	"encoding/json"
	"errors"
	"fmt"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
	platformapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/platform-api"
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

func bodyMessage(err error, body []byte) string {
	if body == nil {
		return err.Error()
	}

	var i interface{}
	if jsonErr := json.Unmarshal(body, &i); jsonErr != nil {
		return err.Error()
	}
	m := i.(map[string]interface{})

	if m["message"] == nil {
		return err.Error()
	}

	return m["message"].(string)
}

// getClientErrorMessage deconstructs errors coming from the Platform or AppManger APIs
// to get the real error message
func getClientErrorMessage(err error) string {
	if err == nil {
		return "UnknownError"
	}

	switch err := err.(type) {
	case appmanagerapi.GenericSwaggerError:
		return bodyMessage(err, err.Body())
	case platformapi.GenericSwaggerError:
		return bodyMessage(err, err.Body())
	default:
		return err.Error()
	}
}

func IsResponseError(res *http.Response) bool {
	return res != nil && res.StatusCode >= 400
}

func FormatResponseError(res *http.Response, clientError error) error {
	message := getClientErrorMessage(clientError)

	switch res.StatusCode {
	case 400:
		return fmt.Errorf("%w: %v", ErrBadRequest, message)
	case 401:
		return fmt.Errorf("%w: %v", ErrUnauthorized, message)
	case 403:
		return fmt.Errorf("%w: %v", ErrForbidden, message)
	case 404:
		return fmt.Errorf("%w: %v", ErrNotFound, message)
	case 409:
		return fmt.Errorf("%w: %v", ErrConflict, message)
	default:
		return fmt.Errorf("%w: %v", ErrUnknown, message)
	}
}
