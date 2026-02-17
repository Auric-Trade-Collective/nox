package webserver

import (
	"YendisFish/nox/native"
	"net/http"
	"path/filepath"
)

type Webserver struct {
	server *http.Server
}

func NewWebserver(addr string, api *native.NoxApi) *Webserver {
	abs, err := filepath.Abs("./") //later should be read from config
	if err != nil {
		panic(err.Error())
	}

	hand := &NoxHandler{ Root: abs, Api: api, DirView: nil }
	server := &Webserver{
		server: &http.Server{
			Addr: addr,
			Handler: hand,
		},
	}

	return server
}

func (s *Webserver) Serve() {
	err := s.server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
