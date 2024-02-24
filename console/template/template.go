package template

import (
	"io"
	"os"
)

type Template struct {
	html, style, script string
}

func NewTemplate() Template {
	t := Template{}
	return t
}

func (t Template) WithHTML(path string) Template {
	t.html = t.mustParseFile(path)
	return t
}

func (t Template) WithStyle(path string) Template {
	t.style = t.mustParseFile(path)
	return t
}

func (t Template) WithScript(path string) Template {
	t.script = t.mustParseFile(path)
	return t
}

func (t Template) mustParseFile(path string) string {
	b, _ := os.ReadFile(path)
	return string(b)
}

func (t Template) Execute(w io.Writer) error {
	var s string
	s += quoteContent("style", t.style)
	s += t.html
	s += quoteContent("script", t.script)

	_, err := w.Write([]byte(s))
	return err
}

func quoteContent(q string, s string) string {
	var str string
	str += "<" + q + ">"
	str += s
	str += "</" + q + ">"

	return str
}
