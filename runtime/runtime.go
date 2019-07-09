package runtime

import "time"

func init() {}

type IRuntime interface {
	Start(arg string) int
	Stop()
	ResStart(arg string) int
	Idle() int
	Health() int
}

type Runtime struct {
	CreateTime time.Time
}
