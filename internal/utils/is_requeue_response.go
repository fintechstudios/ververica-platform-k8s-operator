package utils

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

// IsRequeueResponse determines whether or not a response will result in re-queueing the message
func IsRequeueResponse(res ctrl.Result, err error) bool {
	return err != nil || res.RequeueAfter > 0 || res.Requeue
}
