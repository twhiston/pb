########################################################
# PatchBlock: {{ .Name }}
# {{ .Username }} - {{ .DateNow }}
#
# Generated by pbgen https://github.com/twhiston/pbgen
########################################################
name: {{ .Name }}
category: {{ .Category }}
info: {{ .Info }}
{{ if .Inputs -}}
inputs:
{{- range $key, $value := .Inputs }}
  - info: {{ $value.Info }}
    value: {{ $value.Value }}
    editable: {{ $value.Editable }}
{{ end -}}
{{ end -}}
{{ if .Outputs -}}
outputs:
{{- range $key, $value := .Outputs }}
  - info: {{ $value.Info }}
    value: {{ $value.Value }}
{{ end -}}
{{ end -}}
{{ if .Vars -}}
vars:
{{- range $key, $value := .Vars }}
  - dtype: {{ $value.Dtype }}
    name: {{ $value.Name }}
{{ end -}}
{{ end -}}
help: {{ .Help }}
code:
{{- with .Code }}
{{- if .Path }}
  path: {{ .Path }}
{{- end }}
{{- if .Block }}
  block: {{ .Block }}
{{ end -}}
{{ end }}