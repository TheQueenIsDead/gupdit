package gupdit

import (
	"errors"
	"fmt"
	git "github.com/go-git/go-git/v5"
)

type Gupdit struct {
	remote string
}

func New(remote string) *Gupdit {
	return &Gupdit{
		remote: remote,
	}
}

func (g *Gupdit) ListVersions() (error, []string) {

	r, err := git.PlainOpen(g.remote)
	if err != nil {
		return err, nil
	}

	fmt.Println(r.Tags())

	return errors.New("not yet implemented"), nil
}
