package initializer

import (
	"github.com/pelletier/go-toml/v2"
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
	config := component.Config{
		Project: component.ConfigProject{
			Name: i.projectName,
		},
	}

	p := &component.Project{
		Config: config,
		Path:   i.projectPath,
	}

	if _, err := operation.CreateDirectory(p); err != nil {
		return err
	}

	if err := i.writeConfig(p, &config); err != nil {
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

func (i *fsInitializer) writeConfig(p *component.Project, cfg *component.Config) error {
	cfgContent, err := toml.Marshal(cfg)

	if err != nil {
		return err
	}

	return operation.WriteFile(p, cfgContent, "gown.toml")
}
