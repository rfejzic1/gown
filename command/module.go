package command

import (
	"gown/component"
	"gown/operation"
)

type AddModule struct {
	Name string
}

func (c *AddModule) Execute(p *component.Project) error {
	_, err := operation.CreatePackage(p, "app", c.Name)

	if err != nil {
		return err
	}

	return nil
}

func (c *AddModule) Undo(p *component.Project) error {
	return operation.DeletePackage(p, "app", c.Name)
}
