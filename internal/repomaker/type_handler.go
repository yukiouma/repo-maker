package repomaker

import (
	"go/ast"
	"reflect"
)

var goTypeSet = make(map[string]bool)

func init() {
	types := []interface{}{
		false,
		int(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		uintptr(0),
		float32(0),
		float64(0),
		complex64(0),
		complex128(0),
		"",
	}
	for _, t := range types {
		tp := reflect.TypeOf(t)
		goTypeSet[tp.Name()] = true
	}
	goTypeSet["error"] = true
	goTypeSet["any"] = true
}

func (f *File) identHandler(param *Param, t *ast.Ident) {
	if param == nil || param.Kind == nil {
		return
	}
	name := t.Name
	param.Kind.TypeName = name
	if !isBaseTypeInGo(name) {
		param.Kind.PkgName = f.Pkg.Name
	}
}

func (f *File) selectorExprHandler(param *Param, t *ast.SelectorExpr) {
	if param == nil || param.Kind == nil {
		return
	}
	param.Kind.TypeName = t.Sel.Name
	if t.X != nil {
		if x, ok := t.X.(*ast.Ident); ok {
			pkgName := x.Name
			param.Kind.PkgName = pkgName
			if _, ok = f.Imports[pkgName]; !ok {
				f.Imports[pkgName] = &Package{}
			}
		}
	}
}

func (f *File) starExprHandler(param *Param, t *ast.StarExpr) {
	if param == nil || param.Kind == nil {
		return
	}
	param.Kind.IsRef = true
	switch st := t.X.(type) {
	case *ast.Ident:
		param.Kind.TypeName = st.Name
		param.Kind.PkgName = f.Pkg.Name
	case *ast.SelectorExpr:
		f.selectorExprHandler(param, st)
	default:
	}
}

func (f *File) arrayTypeHandler(param *Param, t *ast.ArrayType) {
	if param == nil || param.Kind == nil {
		return
	}
	param.Kind.IsRepeat = true
	switch at := t.Elt.(type) {
	case *ast.Ident:
		param.Kind.TypeName = at.Name
		param.Kind.PkgName = f.Pkg.Name
	case *ast.SelectorExpr:
		f.selectorExprHandler(param, at)
	case *ast.StarExpr:
		f.starExprHandler(param, at)
	default:
	}
}

func isBaseTypeInGo(t string) bool {
	return goTypeSet[t]
}
