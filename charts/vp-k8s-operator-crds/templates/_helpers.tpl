{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "vp-k8s-operator-crds.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "vp-k8s-operator-crds.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "vp-k8s-operator-crds.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "vp-k8s-operator-crds.labels" -}}
app.kubernetes.io/name: {{ include "vp-k8s-operator-crds.name" . }}
helm.sh/chart: {{ include "vp-k8s-operator-crds.chart" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "vp-k8s-operator-crds.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "vp-k8s-operator-crds.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}


{{/*
Webhook cert injection
*/}}
{{- define "vp-k8s-operator-crds.webhookCert" -}}
{{- $ns := .Values.webhookCert.namespace | default .Release.Namespace -}}
{{ printf "%s/%s" $ns .Values.webhookCert.name }}
{{- end -}}


{{/*
Webhook caBundle defaults to newline (PEM-encoded) as a placeholder
*/}}
{{- define "vp-k8s-operator-crds.webhookCaBundle" -}}
{{ default "Cg==" .Values.webhookCert.caBundle }}
{{- end -}}
