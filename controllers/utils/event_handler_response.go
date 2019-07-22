package utils

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"time"
)

func EventHandlerResponse(duration time.Duration, err error) (ctrl.Result, error) {
	res := ctrl.Result{}

	if duration > 0 {
		res.RequeueAfter = duration
	}

	return res, err
}
