package main

import (
	"fmt"
	"log"

	"golang.org/x/tools/go/packages"
)

// NOTE: This might be the way to go about reading
//       the project. I'll have access to all the pacakges
//       in the project and it's source files (via pkg.Syntax)

func main() {
	cfg := &packages.Config{
		Mode: packages.NeedFiles | packages.NeedSyntax,
	}

	pkgs, err := packages.Load(cfg, "./...")

	if err != nil {
		log.Fatal(err)
	}

	if packages.PrintErrors(pkgs) > 0 {
		log.Fatal()
	}

	for _, pkg := range pkgs {
		fmt.Println(pkg.ID, pkg.GoFiles)
	}
}
