package runtime

import (
	event "github.com/swift9/ares-event"
	"time"
)

func init() {}

type IRuntime interface {
	Start(cmd string, args ...string) int
	Stop()
	ReStart(cmd string, args ...string) int
	Idle() int
	Health() int
	GetCreateTime() time.Time
	Init()
	event.IEventEmitter
}

type Runtime struct {
	event.EventEmitter
	CreateTime time.Time
}

func (runtime *Runtime) GetCreateTime() time.Time {
	return runtime.CreateTime
}

func (runtime *Runtime) Init() {
	runtime.CreateTime = time.Now()
}
