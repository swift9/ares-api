package runtime

import (
	event "github.com/swift9/ares-event"
	"time"
)

func init() {}

type IRuntime interface {
	Start(arg string) int
	Stop()
	ResStart(arg string) int
	Idle() int
	Health() int
}

type Runtime struct {
	event.Event
	CreateTime time.Time
}
