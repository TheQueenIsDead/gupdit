package main

import (
	"fmt"
	"github.com/TheQueenIsDead/gupdit"
)

var Version = "development"

func main() {

	g := gupdit.New("https://github.com/TheQueenIsDead/gupdit.git")
	err, versions := g.ListVersions()
	fmt.Println(err, versions)
}
