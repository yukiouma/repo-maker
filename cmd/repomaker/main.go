package main

import (
	"flag"
	"log"

	repomaker "github.com/yukiouma/repo-maker"
)

func main() {
	in := flag.String("in", ".", "path of repository interface definition")
	out := flag.String("out", ".", "path of implements")
	flag.Parse()
	if err := repomaker.MakeRepo(*in, *out); err != nil {
		log.Printf("failed to generate implements for repository because: %v", err)
	}
}
