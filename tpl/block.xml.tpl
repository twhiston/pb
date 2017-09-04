<block name="{{ .name }}">
    <category>{{ .category }}</category>
    <info>{{ .info }}</info>
    <data name="{{ .blockNameLower }}">
    {{ range $key, $value := .inputs }}
        <variable socket="in" info="Signal input" />
        <variable socket="in" editable="1" value="100" info="Delay time in ms" />
    {{ end }}
    {{ range $key, $value := .outputs }}
        <variable socket="out" info="$value" />
    {{ end }}
    {{ range $key, $value := .vars }}
        <variable dtype="{{ $value }}{{ .type }}" name="{{ $value }}{{ .name }}" />
        //OR IF VALUE ADD value="0"
    {{ end }}
    </data>
    <function name="f_{{ .blockNameLower }}">
        <![CDATA[
    {{ .function }}
	]]>
    </function>
    <help>
        <![CDATA[
{{ .help }}
	]]>
    </help>
</block>