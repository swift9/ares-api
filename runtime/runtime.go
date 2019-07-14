package runtime

import (
	event "github.com/swift9/ares-event"
	"time"
)

type IRuntime interface {
	Init()
	Start(cmd string, args ...string) int
	Stop()
	Idle() int
	Health() int
	GetCreateTime() time.Time
	event.IEventEmitter
}

type Runtime struct {
	event.EventEmitter
	CreateTime time.Time
}

func (r *Runtime) GetCreateTime() time.Time {
	return r.CreateTime
}
