package repomaker

import (
	"fmt"
	"strings"
)

type Method struct {
	Name   string
	Input  []*Param
	Output []*Param
}

func (m *Method) Render(owner string) (s string) {
	inputs := make([]string, 0, len(m.Input))
	for _, v := range m.Input {
		inputs = append(inputs, v.Render())
	}
	outputs := make([]string, 0, len(m.Output))
	for _, v := range m.Output {
		outputs = append(outputs, v.Render())
	}
	s = fmt.Sprintf(
		`func (r *%s) %s(%s) (%s) {
	panic("Unimplemenet")
}`,
		owner,
		m.Name,
		strings.Join(inputs, ", "),
		strings.Join(outputs, ", "),
	)
	return
}

type Param struct {
	Arg  string
	Kind *Kind
}

func (p *Param) Render() (s string) {
	arg := p.Arg
	if len(arg) > 0 {
		s = s + p.Arg + " "
	}
	s = s + p.Kind.Render()
	return
}
