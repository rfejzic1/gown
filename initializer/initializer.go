package initializer

import "gown/component"

type Initializer interface {
	Initialize() (*component.Project, error)
}
