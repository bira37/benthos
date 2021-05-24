package template

import (
	"bytes"
	"text/template"

	"github.com/Jeffail/benthos/v3/internal/docs"
)

// TODO: Put this somewhere else and use the new Go 1.16 file APIs.
var templateDocsTemplate = `---
title: Templating
description: Learn how Benthos templates work.
---

<!--
     THIS FILE IS AUTOGENERATED!

     To make changes please edit the contents of:
     internal/docs/template_docs.go
-->

EXPERIMENTAL: Templates are an experimental feature and therefore subject to change (or removal) outside of major version releases.

Templates are a way to define new Benthos components (similar to plugins) that are implemented by generating Benthos configs. This is useful when a common pattern of Benthos configuration is used but with varying parameters each time.

A template is defined in a YAML file that can be imported when Benthos runs using the flag ` + "`-t`" + `:

` + "```sh" + `
benthos -t "./templates/*.yaml" -c ./config.yaml
` + "```" + `

You can see examples of templates, including some that are included as part of the standard Benthos distribution, at [https://github.com/Jeffail/benthos/tree/master/template](https://github.com/Jeffail/benthos/tree/master/template).

## Fields

The schema of a template file is as follows:

{{range $i, $field := .Fields -}}
### ` + "`{{$field.Name}}`" + `

{{$field.Description}}
{{if $field.Interpolated -}}
This field supports [interpolation functions](/docs/configuration/interpolation#bloblang-queries).
{{end}}

{{if gt (len $field.Type) 0}}
Type: {{if $field.IsArray}}list of {{end}}{{if $field.IsMap}}map of {{end}}` + "`{{$field.Type}}`" + `  
{{end}}
{{if gt (len $field.AnnotatedOptions) 0}}
| Option | Summary |
|---|---|
{{range $j, $option := $field.AnnotatedOptions -}}` + "| `" + `{{index $option 0}}` + "` |" + ` {{index $option 1}} |
{{end}}
{{else if gt (len $field.Options) 0}}Options: {{range $j, $option := $field.Options -}}
{{if ne $j 0}}, {{end}}` + "`" + `{{$option}}` + "`" + `{{end}}.
{{end}}
{{if gt (len $field.Examples) 0 -}}
` + "```yaml" + `
# Examples

{{range $j, $example := $field.Examples -}}
{{if ne $j 0}}
{{end}}{{$example}}{{end -}}
` + "```" + `

{{end -}}
{{end -}}
`

type templateContext struct {
	Fields docs.FieldSpecs
}

// DocsMarkdown returns a markdown document for the templates documentation.
func DocsMarkdown() ([]byte, error) {
	var buf bytes.Buffer
	err := template.Must(template.New("templates").Parse(templateDocsTemplate)).Execute(&buf, templateContext{
		Fields: docs.FieldCommon("", "").WithChildren(ConfigSpec()...).FlattenChildrenForDocs(),
	})

	return buf.Bytes(), err
}
