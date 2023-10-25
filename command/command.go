package command

import "gown/component"

type Command interface {
	Execute(*component.Project) error
}
