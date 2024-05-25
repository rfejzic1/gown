package loader

import (
	"fmt"
	"gown/component"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"golang.org/x/tools/go/packages"
)

type fsLoader struct {
	projectPath string
	cfg         *packages.Config
}

func NewFsLoader(projectPath string) fsLoader {
	cfg := &packages.Config{
		Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedName | packages.NeedModule | packages.NeedCompiledGoFiles,
	}

	return fsLoader{
		projectPath: filepath.Clean(projectPath),
		cfg:         cfg,
	}
}

func (l *fsLoader) Load() (*component.Project, error) {
	cfg, err := l.loadConfig(l.projectPath)

	if err != nil {
		return nil, err
	}

	project := component.NewProject(l.projectPath, cfg)

	pkgs, err := packages.Load(l.cfg, "./...")

	if err != nil {
		return nil, err
	}

	if packages.PrintErrors(pkgs) > 0 {
		return nil, fmt.Errorf("errors occured while parsing packages")
	}

	for i := range pkgs {
		pkg := pkgs[i]

		packagePathParts := strings.Split(pkg.PkgPath, "/")
		packagePath := path.Join(packagePathParts[1:]...)

		if packageMatch(packagePath, "app") {
			fmt.Printf("handle 'app' package\n")
			project.Application.Sources = l.loadCompiledSources(pkg)
		} else if packageMatch(packagePath, "app/*") {
			fmt.Printf("handle '%s' package [module]\n", packagePath)

			module := component.Module{
				Name:    pkg.Name,
				Sources: l.loadCompiledSources(pkg),
			}

			project.Application.AddModule(module)
		} else if packageMatch(packagePath, "setup") {
			fmt.Printf("handle 'setup' package\n")

			project.Setup = &component.Setup{
				Sources: l.loadCompiledSources(pkg),
			}
		} else if packageMatch(packagePath, "web") {
			fmt.Printf("handle 'web' package\n")

			project.Web = &component.Web{
				Sources: l.loadCompiledSources(pkg),
			}
		} else if packageMatch(packagePath, "cmd/*") {
			fmt.Printf("handle '%s' package [command]\n", packagePath)

			cmd := component.Command{
				Sources: l.loadCompiledSources(pkg),
			}

			project.AddCommand(cmd)
		} else {
			fmt.Printf("unhadnled '%s' package\n", packagePath)
		}

		fmt.Println()
	}

	return project, nil
}

func packageMatch(packagePath string, pattern string) bool {
	match, err := path.Match(pattern, packagePath)

	if err != nil {
		log.Fatalf("invalid path match pattern: %s", err)
	}

	return match
}

func (l *fsLoader) loadCompiledSources(pkg *packages.Package) []component.SourceFile {
	sources := []component.SourceFile{}

	fmt.Printf("loading sources...\n")

	for i, pth := range pkg.CompiledGoFiles {
		rel, err := filepath.Rel(l.projectPath, pth)

		if err != nil {
			log.Printf("failed to resolve relative path for file: %s", pth)
			continue
		}

		source := component.SourceFile{
			Path: rel,
			File: pkg.Syntax[i],
		}
		sources = append(sources, source)
		fmt.Printf("loaded file: %s\n", source.Path)
	}

	return sources
}

func (l *fsLoader) loadConfig(projectPath string) (component.Config, error) {
	var cfg component.Config

	path := filepath.Join(projectPath, "gown.toml")

	content, err := os.ReadFile(path)

	if err != nil {
		return component.Config{}, err
	}

	if err := toml.Unmarshal(content, &cfg); err != nil {
		return component.Config{}, err
	}

	return cfg, nil
}
