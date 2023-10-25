package command

import "github.com/rfejzic1/gown/component"

type Command interface {
	Execute(*component.Project) error
}
