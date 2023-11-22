package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	lineTmpl string = `+-----------------+---------------------+-------------` +
		`--------+--------------------------------------------------------------+`

	tagTmpl = `{{printf "%-15s" .Tag}}`

	startAtTmpl = `{{ .StartAt.Format "2006-01-02 15:04:05" | printf "%-10s"}}`

	endAtTmpl = `{{ if .EndAt }}{{ .EndAt.Format "2006-01-02 15:04:05" | printf "%-10s"}}` +
		`{{ else }}{{ printf "%-10s" "0000-00-00 00:00:00" }}{{ end }}`

	urlTmpl = `{{printf "%-60s" .URL}}`
)

var tableTmpl string = fmt.Sprintf(`
%s{{range .}}
| %s | %s | %s | %s |
%s{{end}}
`, lineTmpl, tagTmpl, startAtTmpl, endAtTmpl, urlTmpl, lineTmpl,
)

type TaskView struct {
	tmpl string
	*template.Template
}

func NewTaskView() (*TaskView, error) {
	t, err := template.New("table").Parse(tableTmpl)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse template: %v", err)
	}
	return &TaskView{tmpl: tableTmpl, Template: t}, nil
}

func (t TaskView) Render(tasks *[]*Task) error {
	if len(*tasks) > 0 {
		for _, task := range *tasks {
			task.Localtime()
		}
		return t.Execute(os.Stdout, tasks)
	}
	return nil
}
