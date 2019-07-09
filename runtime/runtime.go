package runtime

import (
	event "github.com/swift9/ares-event"
	"time"
)

func init() {}

type IRuntime interface {
	Start(name string, args ...string) int
	Stop()
	ResStart(name string, args ...string) int
	Idle() int
	Health() int
}

type RuntimeEvent string

const (
	Log RuntimeEvent = "log"

	Fatal RuntimeEvent = "fatal"

	Exit RuntimeEvent = "exit"

	Error RuntimeEvent = "error"
)

type Runtime struct {
	event.Event
	CreateTime time.Time
}
