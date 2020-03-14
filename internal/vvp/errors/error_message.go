package vvperrors

import (
	"encoding/json"

	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-operator/internal/vvp/appmanager-api-client"
	platformApi "github.com/fintechstudios/ververica-platform-k8s-operator/internal/vvp/platform-api-client"
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

// GetVVPErrorMessage deconstructs errors coming from the Platform or AppManger APIs
// to get the real error message
func GetVVPErrorMessage(err error) string {
	if err == nil {
		return "UnknownError"
	}

	switch err := err.(type) {
	case appManagerApi.GenericSwaggerError:
		return bodyMessage(err, err.Body())
	case platformApi.GenericSwaggerError:
		return bodyMessage(err, err.Body())
	default:
		return err.Error()
	}
}
