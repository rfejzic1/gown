package operation

import (
	"fmt"
	"go/format"
	"gown/component"
	"os"
	"path/filepath"
)

const filePermission = 0744
const dirPermission = 0755

func CreateDirectory(p *component.Project, dirPath ...string) error {
	relPath := filepath.Join(dirPath...)
	fullPath := filepath.Join(p.Path, relPath)
	return os.Mkdir(fullPath, dirPermission)
}

// Creates a package with a single go file inside.
func CreatePackage(p *component.Project, packagePath ...string) (component.Package, error) {
	relPath := filepath.Join(packagePath...)
	fullPath := filepath.Join(p.Path, relPath)

	if err := os.Mkdir(fullPath, dirPermission); err != nil {
		return component.Package{}, err
	}

	packageName := packagePath[len(packagePath)-1]

	pkg := component.Package{
		Projct: p,
		Name:   packageName,
		Path:   relPath,
	}

	_, err := CreateSourceFile(p, packageName, filepath.Join(relPath, packageName))

	if err != nil {
		return component.Package{}, err
	}

	return pkg, nil
}

// Creates a package that acts as an entrypoint
func CreateEntrypoint(p *component.Project, packagePath ...string) (*component.Package, error) {
	relPath := filepath.Join(packagePath...)
	fullPath := filepath.Join(p.Path, relPath)

	if err := os.Mkdir(fullPath, dirPermission); err != nil {
		return nil, err
	}

	packageName := "main"

	pkg := &component.Package{
		Projct: p,
		Name:   packageName,
		Path:   relPath,
	}

	file, err := CreateSourceFile(p, packageName, filepath.Join(relPath, packageName))

	if err != nil {
		return nil, err
	}

	if err := AppendToSource(file, `import "fmt"; func main() { fmt.Println("Hello, World") }`); err != nil {
		return nil, err
	}

	return pkg, nil
}

// Delete a package and all it's subpackages. It doesn't refactor the rest of the source to remove references.
func DeletePackage(p *component.Project, packagePath ...string) error {
	relPath := filepath.Join(packagePath...)
	fullPath := filepath.Join(p.Path, relPath)
	return os.RemoveAll(fullPath)
}

// Create a go source file.
func CreateSourceFile(p *component.Project, packageName string, filePath string) (*component.SourceFile, error) {
	fullPath := filepath.Join(p.Path, filePath) + ".go"
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, filePermission)

	if err != nil {
		return nil, err
	}

	source := fmt.Sprintf("package %s\n\n", packageName)
	formatted, err := format.Source([]byte(source))

	if err != nil {
		return nil, err
	}

	fmt.Fprintf(file, string(formatted))

	return &component.SourceFile{
		Path: fullPath,
		File: file,
	}, nil
}

func AppendToSource(f *component.SourceFile, content string) error {
	formatted, err := format.Source([]byte(content))

	if err != nil {
		return err
	}

	_, err = f.File.Write(formatted)

	return err
}
