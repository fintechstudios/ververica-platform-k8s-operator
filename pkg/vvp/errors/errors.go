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
	// Response errors

	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unathorized")
	ErrForbidden    = errors.New("forbidden")
	ErrConflict     = errors.New("conflict")
	ErrNotFound     = errors.New("not found")
	ErrUnknown      = errors.New("unknown error")

	// Authorized context errors

	ErrAuthContext = errors.New("couldn't get authorized context")
)

func bodyOrErrMessage(err error, body []byte) string {
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
		return bodyOrErrMessage(err, err.Body())
	case platformapi.GenericSwaggerError:
		return bodyOrErrMessage(err, err.Body())
	default:
		return err.Error()
	}
}

func IsResponseError(res *http.Response) bool {
	return res != nil && res.StatusCode >= 400
}

func errForStatusCode(code int, message string) error {
	switch code {
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

func FormatResponseError(res *http.Response, clientError error) error {
	message := getClientErrorMessage(clientError)
	return errForStatusCode(res.StatusCode, message)
}

func WrapAuthContextError(namespaceName string, err error) error {
	return fmt.Errorf("%w for namespace %v: %v", ErrAuthContext, namespaceName, err)
}
