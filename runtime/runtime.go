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
	Status  STATUS                 `json:"status"`
	Details map[string]interface{} `json:"details"`
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

type Env struct {
	Name  string
	Value string
}

type ProcessCommand struct {
	Envs     []Env
	Cmd      string
	Args     string
	Dir      string
	Addition map[string]string
}

type IRuntime interface {
	Init()
	Start(cmd ProcessCommand) int
	Stop()
	Idle() int
	Health() Status
	event.IEmitter
}

type Runtime struct {
	event.Emitter
	Meta map[string]interface{}
}
