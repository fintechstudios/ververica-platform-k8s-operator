#!/usr/bin/env bash

NAMESPACE=vvp
RELEASE_NAME=vvp

kubectl --namespace "${NAMESPACE}" port-forward service/"${RELEASE_NAME}"-ververica-platform 8081:80
