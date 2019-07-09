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
	GetCreateTime() time.Time
	event.IEventEmitter
}

type Runtime struct {
	event.EventEmitter
	CreateTime time.Time
}

func (runtime *Runtime) getCreateTime() time.Time {
	return runtime.CreateTime
}
