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
    {{ if $value.Value }}value: {{ $value.Value }}{{ end }}
    {{ if $value.Editable }}editable: {{ $value.Editable }}{{ end }}
{{- end -}}
{{- end -}}
{{ if .Outputs }}
outputs:
{{- range $key, $value := .Outputs }}
  - info: {{ $value.Info }}
    {{ if $value.Value }}value: {{ $value.Value }}{{ end }}
{{- end -}}
{{- end -}}
{{ if .Vars }}
vars:
{{- range $key, $value := .Vars }}
  - dtype: {{ $value.Dtype }}
    name: {{ $value.Name }}
    {{ if $value.Value }}value: {{ $value.Value }}{{ end }}
{{- end -}}
{{- end }}
help: >
    {{ .Help }}
code:
{{- if .Code.Path }}
  path: {{ .Code.Path }}
{{- end }}