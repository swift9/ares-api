package runtime

import (
	event "github.com/swift9/ares-event"
	"time"
)

type IRuntime interface {
	Start(cmd string, args ...string) int
	Stop()
	Idle() int
	Health() int
	GetCreateTime() time.Time
	Init()
	event.IEventEmitter
}
