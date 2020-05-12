package main

import (
	"log"
	"os"

	"github.com/glispy/glispy"
	"github.com/hatchify/scribe"
)

func main() {
	out := scribe.New("Glispy CLI")
	out.Notification("Starting VM")
	g := glispy.New()
	out.Notification("Ready to read input :)\n\n\n")
	if _, err := g.EvalReader(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
