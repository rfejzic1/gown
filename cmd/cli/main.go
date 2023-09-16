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
	_, err := loader.LoadProject()

	if err != nil {
		log.Fatal(err)
	}
}
