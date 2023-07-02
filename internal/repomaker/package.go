package repomaker

type Package struct {
	Name, Path string
}

func (p *Package) Render() (s string) {
	if len(p.Name) > 0 {
		s = s + p.Name + " "
	}
	s += p.Path
	return
}
