package main

import (
	"fmt"
	"github.com/TheQueenIsDead/gupdit"
	"log"
)

var Version = "development"

func main() {

	g := gupdit.New("https://github.com/TheQueenIsDead/gupdit.git")
	latest, err := g.IsLatest(Version)
	if err != nil {
		panic(err)
	}
	if latest {
		fmt.Printf("%s is the latest\n", Version)
	} else {
		log.Panicf("%s is out date!\n", Version)
	}
}
