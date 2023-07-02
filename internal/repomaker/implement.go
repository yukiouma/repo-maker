package repomaker

import (
	"bytes"
	"go/format"
	"log"
	"text/template"
)

var (
	tmpl *template.Template
	err  error
)

func init() {
	tmpl, err = template.New("repo").Parse(RepoTemplate)
	if err != nil {
		log.Fatal(err)
	}
}

type Implement struct {
	DestPackage     string
	Imports         []string
	ImplementRender []*ImplementRender
}

type ImplementRender struct {
	Name    string
	Methods []string
}

func (i *Implement) Render() (s []byte, err error) {
	if len(i.ImplementRender) == 0 {
		return
	}
	sb := bytes.NewBuffer(s)
	err = tmpl.Execute(sb, i)
	s = sb.Bytes()
	if err != nil {
		return
	}
	s, err = format.Source(s)
	return
}
