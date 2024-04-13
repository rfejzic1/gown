package initializer

import (
	"gown/component"
	"gown/operation"
)

type fsInitializer struct {
	projectName string
	projectPath string
}

func NewFsInitializer(projectName string, projectPath string) fsInitializer {
	return fsInitializer{
		projectName,
		projectPath,
	}
}

func (i *fsInitializer) Initialize() error {
	p := &component.Project{
		Name: i.projectName,
		Path: i.projectPath,
	}

	if _, err := operation.CreateDirectory(p); err != nil {
		return err
	}

	if _, err := operation.CreatePackage(p, "app"); err != nil {
		return err
	}

	if _, err := operation.CreatePackage(p, "setup"); err != nil {
		return err
	}

	if _, err := operation.CreatePackage(p, "web"); err != nil {
		return err
	}

	if _, err := operation.CreateDirectory(p, "cmd"); err != nil {
		return err
	}

	if _, err := operation.CreatePackageMain(p, "cmd", i.projectName); err != nil {
		return err
	}

	return nil
}
