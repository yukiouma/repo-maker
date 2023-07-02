package repomaker

var RepoTemplate = `package {{.DestPackage}}

import (
	{{range $index, $import := .Imports}}
	{{$import}}
	{{end}}
)

{{range $index, $interface := .ImplementRender}}
type {{$interface.Name}} struct {
}

func New{{$interface.Name}}() domain.{{$interface.Name}} {
	return &{{$interface.Name}}{}
}

{{range $index, $method := $interface.Methods}}
{{$method}}

{{end}}
{{end}}

`
