#!/usr/bin/env bash

set -e

main() {
  TILLER_BIN=${1:-"tiller"}
  NAMESPACE=${TILLER_NAMESPACE:-"tiller"}
  # Even though we'll be running tiller locally, it still needs a namespace
  kubectl get namespace "$NAMESPACE" >/dev/null 2>&1 || kubectl create namespace "$NAMESPACE" >/dev/null 2>&1

  ${TILLER_BIN} -listen=localhost:44134 -storage=secret -alsologtostderr >/dev/null 2>&1
}

main "$@"
