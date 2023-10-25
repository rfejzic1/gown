package loader

import "gown/component"

type Loader interface {
	Load() (*component.Project, error)
}
