package repomaker

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

type RepoMaker interface {
	Render() error
}

type repoMaker struct {
	input, output string
}

func NewRepoMaker(input, output string) (repo RepoMaker, err error) {
	r := new(repoMaker)
	input, err = convertToAbs(input)
	if err != nil {
		return
	}
	output, err = convertToAbs(output)
	if err != nil {
		return
	}
	r.input = input
	r.output = output
	repo = r
	return
}

func (r *repoMaker) Render() error {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports | packages.NeedFiles,
	}
	pkgs, err := packages.Load(cfg, r.input)
	if err != nil {
		return err
	}
	for _, p := range pkgs {
		for i, s := range p.Syntax {
			f := &File{
				Output: r.output,
				Name:   filepath.Base(p.GoFiles[i]),
				Pkg: &Package{
					Name: p.Name,
					Path: fmt.Sprintf(`"%s"`, p.PkgPath),
				},
				Imports:        make(map[string]*Package, 0),
				InterfaceDecls: make([]*InterfaceDecl, 0, 1),
			}
			f.InterfaceDecls = f.FindInterfaceDecls(s.Decls)
			f.FindImports(s.Imports)
			impl, _ := f.ToImpl()
			if s, err := impl.Render(); err != nil {
				return err
			} else if len(s) > 0 {
				newFile, err := os.Create(path.Join(r.output, f.ExportFileName()))
				if err != nil {
					return err
				} else {
					newFile.Write(s)
				}
			}
		}

	}
	return nil
}

func convertToAbs(path string) (s string, err error) {
	if filepath.IsAbs(path) {
		s = path
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	s = filepath.Join(pwd, path)
	return
}
