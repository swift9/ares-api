package runtime

import (
	event "github.com/swift9/ares-event"
)

type STATUS string

const (
	StatusUp   STATUS = "UP"
	StatusDown STATUS = "DOWN"
)

type Status struct {
	Status  STATUS
	Details map[string]interface{}
}

func NewStatusUp() Status {
	return Status{
		Status: StatusUp,
	}
}

func NewStatusDown() Status {
	return Status{
		Status: StatusDown,
	}
}

type IRuntime interface {
	Init()
	Start(cmd string, args ...string) int
	Stop()
	Idle() int
	Health() Status
	event.IEventEmitter
}

type Runtime struct {
	event.EventEmitter
	Meta map[string]interface{}
}
