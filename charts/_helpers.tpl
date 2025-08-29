{{/*
Expand the name of the chart.
*/}}
{{- define "aid.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "aid.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define `aid.chart` -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define `aid.labels` -}}
helm.sh/chart: "{{ include `aid.chart` . }}",
{{ include `aid.selectorLabels` . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: "{{ .Chart.AppVersion }}",
{{- end }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define `aid.selectorLabels` -}}
app.kubernetes.io/name: "{{ include `aid.name` . }}",
app.kubernetes.io/instance: "{{ .Release.Name }}",
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define `aid.serviceAccountName` -}}
{{- if .Values.serviceAccount.create }}
{{- default (include `aid.fullname` .) .Values.serviceAccount.name }}
{{- else }}
{{- default `default` .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Create the namespace based on the environment name
*/}}
{{- define `aid.namespace` -}}
{{- if eq .Values.environment `non-prod` }}
{{- printf `stage` }}
{{- else}}
{{- printf `prod` }}
{{- end -}}
{{- end }}