package initializer

import (
	"gown/component"
	"gown/operation"

	"github.com/go-git/go-git/v5"
	"github.com/pelletier/go-toml/v2"
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

	if err := i.writeStaticFiles(p); err != nil {
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

	if _, err := operation.CreatePackageMain(p, "cmd", "web"); err != nil {
		return err
	}

	if err := i.gitInit(p); err != nil {
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

func (i *fsInitializer) writeStaticFiles(p *component.Project) error {
	return renderFiles(p, func(fileName string, content []byte) error {
		return operation.WriteFile(p, content, fileName)
	})
}

func (i *fsInitializer) gitInit(p *component.Project) error {
	opt := &git.PlainInitOptions{}
	repo, err := git.PlainInitWithOptions(p.Path, opt)

	if err != nil {
		return err
	}

	w, err := repo.Worktree()

	if err != nil {
		return err
	}

	if err := w.AddGlob("*"); err != nil {
		return err
	}

	commitOpts := &git.CommitOptions{}
	_, err = w.Commit("Initial commit", commitOpts)

	if err != nil {
		return err
	}

	return nil
}
