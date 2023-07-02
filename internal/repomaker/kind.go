package repomaker

type Kind struct {
	IsRef             bool
	PkgName, TypeName string
}

func (k *Kind) Render() (s string) {
	if k.IsRef {
		s += "*"
	}
	if len(k.PkgName) > 0 {
		s = s + k.PkgName + "."
	}
	s = s + k.TypeName
	return
}
