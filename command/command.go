package command

import "gown/component"

type Command interface {
	Execute(*component.Project) error
	Undo(*component.Project) error
}
