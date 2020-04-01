#!/usr/bin/env bash

export HELM_HOST=:44134
helm init --client-only

helm repo add ververica https://charts.ververica.com
helm repo add jetstack https://charts.jetstack.io

