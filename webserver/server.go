package webserver

import (
	"YendisFish/nox/logger"
	"YendisFish/nox/pages"
	"strconv"

	// "YendisFish/nox/webapi"
	"net/http"
	"os"
	"path/filepath"
)

type Webserver struct {
	server *http.Server
	config *Config
}

func NewWebserver(config *Config) *Webserver {
	var api *NoxApi = nil
	var err error
	if len(config.Nox.Api) > 0 {
		api, err = CreateApi(config.Nox.Api, config.Nox.AuthLocation)
		if err != nil {
			logger.Panic(err.Error())
		}
	}

	hand := &NoxHandler{Root: config.Nox.Root, Api: api}
	Handler = hand
	server := &Webserver{
		server: &http.Server{
			Addr:    config.Nox.Addr,
			Handler: hand,
		},
		config: config,
	}

	setupErrorPages(config.Nox.Root)

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

func setupErrorPages(root string) {
	fpath := filepath.Join(root, "/status/")
	for k, _ := range pages.Pages {
		num := strconv.Itoa(k)
		final := filepath.Join(fpath, num+".html")
		_, err := os.Stat(final)
		if err != nil {
			logger.Debug("Could not get information for custom error: " + final)
			continue
		}

		buff, err := os.ReadFile(final)
		if err != nil {
			logger.Panic("Could not read file " + final + " when setting custom errors")
		}

		pages.Pages[k] = string(buff)
	}
}
