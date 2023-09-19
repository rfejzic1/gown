package main

import (
	"log"
	"os"

	"github.com/rfejzic1/gown"
)

func main() {
	projectPath := "./output"

	if len(os.Args) > 1 {
		projectPath = os.Args[1]
	}

	loader := gown.NewLoader(projectPath)
	p, err := loader.LoadProject()

	if err != nil {
		log.Fatal(err)
	}

	gown.PrintProjectStructure(p, os.Stdout)
}
