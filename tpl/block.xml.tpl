<block name="{{ .Name }}">
    <category>{{ .Category }}</category>
    <info>{{ .Info }}</info>
    <data name="{{ .BlockNameLower }}">
    {{- range $key, $value := .Inputs }}
        <variable socket="in" {{ if eq $value.Editable "1" }}editable="{{ $value.Editable }}" {{ end }}{{ if ne $value.Value "0" }}value="{{ $value.Value }}" {{ end }}info="{{ $value.Info }}" />
    {{- end -}}
    {{- range $key, $value := .Outputs }}
        <variable socket="out" info="{{ $value.Info }}" />
    {{- end -}}
    {{- range $key, $value := .Vars }}
        <variable dtype="{{ $value.Dtype }}" name="{{ $value.Name }}" {{ if ne $value.Value "" }}value="{{ $value.Value }}" {{ end }}/>
    {{- end }}
    </data>
    <function name="f_{{ .BlockNameLower }}">
        <![CDATA[
{{- range $key, $value := .ParsedFunction -}}
{{ $value }}
{{ end }}	]]>
    </function>
    <help>
        <![CDATA[
{{ .Help }}
	]]>
    </help>
</block>