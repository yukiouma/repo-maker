package repomaker

import maker "github.com/yukiouma/repo-maker/internal/repomaker"

func MakeRepo(input, output string) error {
	m, err := maker.NewRepoMaker(input, output)
	if err != nil {
		return err
	}
	return m.Render()
}
