package webserver

import (
	"YendisFish/nox/logger"
	"YendisFish/nox/native"
	"net/http"
)

type Webserver struct {
	server *http.Server
	config *Config
}

func NewWebserver(config *Config) *Webserver {
	api, err := native.CreateApi(config.Nox.Api)
	if err != nil {
		logger.Panic(err.Error())
	}

	hand := &NoxHandler{ Root: config.Nox.Root, Api: api, DirView: nil }
	server := &Webserver{
		server: &http.Server{
			Addr: config.Nox.Addr,
			Handler: hand,
		},
		config: config,
	}

	return server
}

func (s *Webserver) Serve() {
	if !s.config.Nox.Tls.Enabled {
		err := s.server.ListenAndServe()
		if err != nil {
			logger.Panic(err.Error())
		}
	} else {
		err := s.server.ListenAndServeTLS(s.config.Nox.Tls.CertFile, 
		                                  s.config.Nox.Tls.KeyFile)
		
		if err != nil {
			logger.Panic(err.Error())
		}
	}
}
