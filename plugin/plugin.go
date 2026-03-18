package plugin

/*
#cgo LDFLAGS: -ldl
#include "../include/nox.h"
*/
import "C"
import "unsafe"

var Routines []Plugin
var Blockers []Plugin

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
	go triggerRoutines(tp)
}

func triggerRoutines(tp EventType) {

}
