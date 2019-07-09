package runtime

import (
	event "github.com/swift9/ares-event"
	"time"
)

func init() {}

type IRuntime interface {
	Start(cmd string, args ...string) int
	Stop()
	ResStart(cmd string, args ...string) int
	Idle() int
	Health() int
}

type Runtime struct {
	event.Event
	CreateTime time.Time
}
