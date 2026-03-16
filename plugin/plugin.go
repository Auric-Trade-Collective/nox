package plugin

/*
#cgo LDFLAGS: -ldl
#include "../include/nox.h"
*/
import "C"
import "unsafe"

var ActivePlugins []Plugin

type Plugin struct {
	events map[EventType]unsafe.Pointer
}

type EventType int
const (
	OnLog	EventType = iota
	OnGet
	OnPost
	OnPut
	OnDelete
	OnError
)

func TriggerEvent(tp EventType) {

}
