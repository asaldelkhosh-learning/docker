package process

import "github.com/hamed-yousefi/gowl"

type Process interface {
	Start() error
	Name() string
	PID() gowl.PID
}
