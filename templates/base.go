package templates

var BaseTpl = `
{{- define "base" }}
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};page-break-inside:avoid;"></div>
</div>

<script type="text/javascript">
    "use strict";
    var goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
    var option_{{ .ChartID | safeJS }} = {{ .JSONNotEscaped | safeJS }};
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});

    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
</script>
{{ end }}
`
