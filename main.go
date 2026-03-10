package main

import (
	"YendisFish/nox/logger"
	"YendisFish/nox/webserver"
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	toml "github.com/pelletier/go-toml/v2"
)

var Version = "0.0.1"

var CLI struct {
	Version kong.VersionFlag `help:"Print version and exit"`

	Spin struct {
		Dir    string `help:"Test load a DLL file" default:"."`
		Config string `help:"Path to config file" default:"nox.toml"`
	} `cmd:"" help:"Spinup a nox server"`

	Install struct {} `help:"Install nox into your system and path" cmd:""`
	Init struct {} `help:"Create a nox.toml file" cmd:""`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("nox"),
		kong.Description("Nox webserver -- Version: "+Version),
		kong.Vars{
			"version": Version,
		},
		kong.UsageOnError(),
	)

	switch ctx.Command() {
	case "init": initNoxToml()
	case "install": installNox()
	case "spin":
		dir, err := filepath.Abs("./")
		if err != nil {
			panic(err.Error())
		}

		if CLI.Spin.Dir != "" {
			dir, _ = filepath.Abs(CLI.Spin.Dir)
		}

		var configPath string
		if filepath.IsAbs(CLI.Spin.Config) {
			configPath = CLI.Spin.Config
		} else {
			configPath = filepath.Join(dir, CLI.Spin.Config)
		}

		buff, err := os.ReadFile(configPath)
		if err != nil {
			logger.Panic("Failed to read config file: " + err.Error())
		}

		var conf webserver.Config
		if err := toml.Unmarshal(buff, &conf); err != nil {
			logger.Panic("Failed to parse config file: " + err.Error())
		}

		cDir, err := os.Getwd()
		if err != nil {
			logger.Panic(err.Error())
		}

		os.Chdir(dir)
		conf.Nox.Root, _ = filepath.Abs(conf.Nox.Root)
		conf.Nox.Api, _ = filepath.Abs(conf.Nox.Api)
		conf.Nox.Tls.CertFile, _ = filepath.Abs(conf.Nox.Tls.CertFile)
		conf.Nox.Tls.KeyFile, _ = filepath.Abs(conf.Nox.Tls.KeyFile)
		os.Chdir(cDir)

		logger.Write("Root is: " + conf.Nox.Root)

		serve := webserver.NewWebserver(&conf)
		serve.Serve()
	}
}

func installNox() {
	binPath, err := os.Executable()
	if err != nil {
		logger.Panic("Cannot find local noxfile to copy into /usr/bin")
	}

	buff, err := os.ReadFile(binPath)
	if err != nil {
		logger.Panic("Cannot copy " + binPath + " to /usr/bin")
	}

	err = os.WriteFile("/usr/bin/nox", buff, 0755)
	if err != nil {
		logger.Panic("Cannot copy " + binPath + " to /usr/bin")
	}
}

var tomlDat = `
[nox]
addr = ":5432"
root = "./web/"
api = "./abi/libapi.so"
`

func initNoxToml() {
	curDur, err := os.Getwd()
	if err != nil {
		logger.Panic("Failed to get current directory.")
	}

	path := filepath.Join(curDur, "nox.toml")

	fle, err := os.Create(path)
	if err != nil {
		logger.Panic("Could not create nox.toml")
	}

	_, err = fle.WriteString(tomlDat)
	if err != nil {
		logger.Panic("Could not write to nox.toml")
	}
}
