package webserver

import (
	"YendisFish/nox/native"
	"fmt"
	"net/http"
)

type Webserver struct {
	server *http.Server
}

func NewWebserver(config *Config, api *native.NoxApi) *Webserver {
	hand := &NoxHandler{ Root: config.Nox.Root, Api: api, DirView: nil }
	server := &Webserver{
		server: &http.Server{
			Addr: config.Nox.Addr,
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

func (s *Webserver) Display() {
	fmt.Println(s.server.Addr)
}
