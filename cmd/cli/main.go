package main

import (
	"log"
	"os"

	"github.com/rfejzic1/gown/cli"
)

func main() {
	cli := cli.New()

	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
