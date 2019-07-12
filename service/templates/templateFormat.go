package templates

import (
	"bytes"
	"text/template"

	"github.com/caioeverest/ingresso-watcher/service/errors"
)

func FormatMessage(templateType string, params map[string]string) (string, error) {
	content, ok := templateTypes[templateType]
	if !ok {
		return "", errors.TemplateNotFound
	}

	for _, param := range content.Params {
		if _, ok := params[param]; !ok {
			return "", errors.ParamNotFound(param)
		}
	}

	tmp := template.Must(template.New(templateType).Parse(content.Content))
	buf := &bytes.Buffer{}
	if err := tmp.Execute(buf, params); err != nil {
		panic(err)
	}

	return buf.String(), nil
}
