package main

import (
	"log"
	"os"

	"github.com/glispy/glispy"
)

func main() {
	g := glispy.New()
	if _, err := g.EvalReader(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
