{{/*
内建数据库名称
*/}}
{{- define "bk-hcm.mariadbName" -}}
{{- if .Values.mariadb.fullnameOverride -}}
{{- .Values.mariadb.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default "mariadb" .Values.mariadb.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}

{{/*
数据库配置，处理了内建和外部数据库场景
*/}}
{{- define "bk-hcm.database" -}}
{{- if .Values.mariadb.enabled }}
host: {{ include "bk-hcm.mariadbName" . }}
port: {{ .Values.mariadb.primary.service.ports.mysql }}
database: {{ .Values.mariadb.auth.database }}
user: {{ .Values.mariadb.auth.username }}
password: {{ .Values.mariadb.auth.password }}
{{- else -}}
host: {{ .Values.externalDatabase.host }}
port: {{ .Values.externalDatabase.port }}
database: {{ .Values.externalDatabase.database }}
user: {{ .Values.externalDatabase.user }}
password: {{ .Values.externalDatabase.password }}
{{- end }}
{{- end -}}

{{- define "bk-hcm.databaseConfig" -}}
{{- $cfg := fromYaml (include "bk-hcm.database" .) -}}
resource:
  endpoints:
    - {{ $cfg.host }}:{{ $cfg.port }}
  database: {{ $cfg.database }}
  user: {{ $cfg.user }}
  password: {{ $cfg.password }}
  dialTimeoutSec:
  readTimeoutSec:
  writeTimeoutSec:
  maxIdleTimeoutMin:
  maxOpenConn:
  maxIdleConn:
  tls:
    insecureSkipVerify:
    certFile:
    keyFile:
    caFile:
    password:
maxSlowLogLatencyMS: 200
limiter:
  qps: 500
  burst: 500
{{- end -}}

{{- define "bk-hcm.database.host" -}}
{{- $cfg := fromYaml (include "bk-hcm.database" .) -}}
{{ $cfg.host }}
{{- end -}}

{{- define "bk-hcm.database.port" -}}
{{- $cfg := fromYaml (include "bk-hcm.database" .) -}}
{{ $cfg.port }}
{{- end -}}

{{- define "bk-hcm.database.user" -}}
{{- $cfg := fromYaml (include "bk-hcm.database" .) -}}
{{ $cfg.user }}
{{- end -}}

{{- define "bk-hcm.database.password" -}}
{{- $cfg := fromYaml (include "bk-hcm.database" .) -}}
{{ $cfg.password }}
{{- end -}}

{{- define "bk-hcm.database.database" -}}
{{- $cfg := fromYaml (include "bk-hcm.database" .) -}}
{{ $cfg.database }}
{{- end -}}
