package command

import "gown/component"

type AddModule struct {
	Name string
}

func (c *AddModule) Execute(p *component.Project) error {
	// TODO: Implement add module command
	return nil
}
