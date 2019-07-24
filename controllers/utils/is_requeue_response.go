package utils

import (ctrl "sigs.k8s.io/controller-runtime")

func IsRequeueResponse(res ctrl.Result, err error) bool {
	return err != nil || res.RequeueAfter > 0 || res.Requeue
}
