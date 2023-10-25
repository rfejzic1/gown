package loader

import "github.com/rfejzic1/gown/component"

type Loader interface {
	Load() (*component.Project, error)
}
