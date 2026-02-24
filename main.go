package main

import (
	"YendisFish/nox/webserver"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	toml "github.com/pelletier/go-toml/v2"
)

var CLI struct {
	Spin struct {
		Dir string `cmd:"--dir" help:"Test load a DLL file"`
	} `cmd:"" help:"Spinup a nox server"`
}

func main() {
	ctx := kong.Parse(&CLI)

	switch ctx.Command() {
	case "spin":
		dir, err := filepath.Abs("./")
		if err != nil {
			panic(err.Error())
		}
		
		if CLI.Spin.Dir != "" {
			dir, _ = filepath.Abs(CLI.Spin.Dir)
		}

		buff, err := os.ReadFile(filepath.Join(dir, "nox.toml"))

		var conf webserver.Config
		toml.Unmarshal(buff, &conf)

		cDir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}

		os.Chdir(dir)
		conf.Nox.Root, _ = filepath.Abs(conf.Nox.Root)
		conf.Nox.Api, _ = filepath.Abs(conf.Nox.Api)
		conf.Nox.Tls.CertFile, _ = filepath.Abs(conf.Nox.Tls.CertFile)
		conf.Nox.Tls.KeyFile, _ = filepath.Abs(conf.Nox.Tls.KeyFile)
		os.Chdir(cDir)
		
		fmt.Println(conf.Nox.Root);

		serve := webserver.NewWebserver(&conf)
		serve.Serve()
	}
}
