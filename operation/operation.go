package operation

import (
	"fmt"
	"gown/component"
	"os"
	"path"
)

const filePermission = 0744
const dirPermission = 0755

// Creates a package with a single go file inside.
func CreatePackage(p *component.Project, packagePath ...string) (component.Package, error) {
	relPath := path.Join(packagePath...)
	fullPath := path.Join(p.Path, relPath)

	if err := os.Mkdir(fullPath, dirPermission); err != nil {
		return component.Package{}, err
	}

	packageName := packagePath[len(packagePath)-1]

	pkg := component.Package{
		Projct: p,
		Name:   packageName,
		Path:   relPath,
	}

	file, err := CreateSourceFile(p.Path, packageName, path.Join(relPath, packageName))

	if err != nil {
		return component.Package{}, err
	}

	pkg.Add(file)

	return pkg, nil
}

// Delete a package and all it's subpackages. It doesn't refactor the rest of the source to remove references.
func DeletePackage(p *component.Project, packagePath ...string) error {
	relPath := path.Join(packagePath...)
	fullPath := path.Join(p.Path, relPath)
	return os.RemoveAll(fullPath)
}

// Create a go source file.
func CreateSourceFile(projectPath string, packageName string, filePath string) (component.SourceFile, error) {
	fullPath := path.Join(projectPath, filePath) + ".go"
	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, filePermission)

	if err != nil {
		return component.SourceFile{}, err
	}

	fmt.Fprintf(file, "package %s\n\n", packageName)

	return component.SourceFile{}, nil
}

