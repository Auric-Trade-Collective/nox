package webserver

import (
	"YendisFish/nox/js"
)

type Webserver struct {
	runtime *js.Js
}

func NewWebserver(conf *js.Config) *Webserver {
	return new(Webserver)
}

func (ws *Webserver) Start() {

}
