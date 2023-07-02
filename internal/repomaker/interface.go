package repomaker

type InterfaceDecl struct {
	Pkg    *Package
	Name   string
	Method []*Method
}
