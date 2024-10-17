package gupdit

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
	"golang.org/x/mod/semver"
	"log"
)

type Gupdit struct {
	remote string
	strict bool
}

func New(remote string) *Gupdit {
	return &Gupdit{
		remote: remote,
		strict: true,
	}
}

func (g *Gupdit) SetStrict(strict bool) {
	g.strict = strict
}

func (g *Gupdit) IsLatest(version string) (bool, error) {
	tags, err := g.listTags()
	if err != nil {
		return false, err
	}

	if len(tags) == 0 {
		return false, nil
	}

	semver.Sort(tags)
	isLatest := semver.Compare(tags[len(tags)-1], version) == 0

	return isLatest, nil
}

func (g *Gupdit) listTags() ([]string, error) {

	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{g.remote},
	})

	refs, err := rem.List(&git.ListOptions{
		PeelingOption: git.AppendPeeled,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Filters the references list and only keeps tags
	var tags []string
	for _, ref := range refs {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		}
	}

	return tags, nil
}
