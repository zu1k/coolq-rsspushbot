package config

import (
	"bytes"
	"strings"
	"text/template"
)

var (
	PreviewText    int = 0
	SQLitePath     string
	UpdateInterval int  = 10
	ErrorThreshold uint = 100
	MessageTpl     *template.Template
)

const (
	defaultMessageTpl = `{{.SourceTitle}}

{{.ContentTitle}}

{{.RawLink}}
`
)

func init() {
	MessageTpl = template.Must(template.New("message").Parse(defaultMessageTpl))
}

type TplData struct {
	SourceTitle  string
	ContentTitle string
	RawLink      string
	PreviewText  string
}

func (t TplData) Render() (string, error) {
	var buf []byte
	wb := bytes.NewBuffer(buf)

	if err := MessageTpl.Execute(wb, t); err != nil {
		return "", err
	}

	return strings.TrimSpace(string(wb.Bytes())), nil
}
