package repomaker

import (
	"go/ast"
	"strings"
)

type File struct {
	Output         string
	Name           string
	Pkg            *Package
	Imports        map[string]*Package
	InterfaceDecls []*InterfaceDecl
}

// Generate the export file name, with ".impl". For example:
// input filename is "repo.go", then return "repo.impl.go"
func (f *File) ExportFileName() (s string) {
	origin := strings.Split(f.Name, ".")
	if size := len(origin); size > 0 {
		s = strings.Join(origin[:size-1], ".") + ".impl" + "." + origin[size-1]
	} else {
		s = f.Name + ".impl"
	}
	return
}

func (f *File) FindImports(in []*ast.ImportSpec) {
	for _, v := range in {
		path := v.Path.Value
		name := ""
		alias := false
		if v.Name != nil {
			name = v.Name.Name
			alias = true
		} else {
			p := strings.Split(path, "/")
			if size := len(p); size > 0 {
				name = strings.ReplaceAll(p[size-1], `"`, "")
			}
		}
		if _, ok := f.Imports[name]; ok {
			f.Imports[name].Path = path
			if alias {
				f.Imports[name].Name = name
			}
		}
	}
}

func (f *File) FindInterfaceDecls(decls []ast.Decl) (interfaces []*InterfaceDecl) {
	for _, d := range decls {
		decl, ok := d.(*ast.GenDecl)
		if !ok {
			continue
		}
		if !filterByComment(decl.Doc) {
			continue
		}
		specs := decl.Specs
		for _, s := range specs {
			spec, ok := s.(*ast.TypeSpec)
			if !ok {
				continue
			}
			typeInfo, ok := spec.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}
			i := &InterfaceDecl{
				Name:   spec.Name.String(),
				Method: f.FindMethods(typeInfo.Methods.List...),
			}
			interfaces = append(interfaces, i)
		}

	}
	return
}

func (f *File) FindMethods(ms ...*ast.Field) (methods []*Method) {
	for _, m := range ms {
		methods = append(methods, f.FindMethod(m))
	}
	return
}

func (f *File) FindMethod(m *ast.Field) (method *Method) {
	method = new(Method)
	if len(m.Names) == 0 {
		return
	}
	method.Name = m.Names[0].Name
	t, ok := m.Type.(*ast.FuncType)
	if !ok {
		return
	}

	method.Input = f.FindParameter(t.Params.List...)
	method.Output = f.FindParameter(t.Results.List...)

	return
}

func (f *File) FindParameter(fields ...*ast.Field) (params []*Param) {
	for _, field := range fields {
		param := new(Param)
		if len(field.Names) > 0 {
			param.Arg = field.Names[0].Name
		}
		switch t := field.Type.(type) {
		case *ast.Ident:
			param.Kind = &Kind{
				TypeName: t.Name,
			}
		case *ast.SelectorExpr:
			param.Kind = &Kind{
				TypeName: t.Sel.Name,
			}
			if t.X != nil {
				if x, ok := t.X.(*ast.Ident); ok {
					pkgName := x.Name
					param.Kind.PkgName = pkgName
					if _, ok = f.Imports[pkgName]; !ok {
						f.Imports[pkgName] = &Package{}
					}
				}
			}
		case *ast.StarExpr:
			param.Kind = &Kind{
				IsRef: true,
			}
			switch st := t.X.(type) {
			case *ast.Ident:
				param.Kind.TypeName = st.Name
				param.Kind.PkgName = f.Pkg.Name
			case *ast.SelectorExpr:
				param.Kind.TypeName = st.Sel.Name
				if st.X != nil {
					if x, ok := st.X.(*ast.Ident); ok {
						param.Kind.PkgName = x.Name
					}
				}
			default:
			}
		default:
		}
		params = append(params, param)
	}
	return
}

func (f *File) ToImpl() (impl *Implement, err error) {
	impl = new(Implement)
	ps := strings.Split(f.Output, "/")
	if size := len(ps); size > 0 {
		impl.DestPackage = ps[size-1]
	} else {
		impl.DestPackage = f.Output
	}
	impl.Imports = make([]string, 0, len(f.Imports)+1)
	for _, p := range f.Imports {
		impl.Imports = append(impl.Imports, p.Render())
	}
	impl.Imports = append(impl.Imports, f.Pkg.Render())
	impl.ImplementRender = make([]*ImplementRender, 0)
	for _, v := range f.InterfaceDecls {
		methods := make([]string, 0, len(v.Method))
		for _, m := range v.Method {
			methods = append(methods, m.Render(v.Name))
		}
		impl.ImplementRender = append(impl.ImplementRender, &ImplementRender{
			Name:    v.Name,
			Methods: methods,
		})
	}
	return
}

func filterByComment(cg *ast.CommentGroup) (result bool) {
	if cg == nil {
		return
	}
	for _, comment := range cg.List {
		if comment.Text == "// +repo.impl" {
			result = true
			break
		}
	}
	return
}
