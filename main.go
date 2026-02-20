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
	Test struct {

	} `cmd:"" help:"Test"`
	Dll struct {
		Dir string `cmd:"--dir" help:"Test load a DLL file"`
	} `cmd:"" help:"Load DLLS"`
	Spin struct {
		Dir string `cmd:"--dir" help:"Test load a DLL file"`
	} `cmd:"" help:"Spinup a nox server"`
}

func main() { ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "dll":
		// api, _ := native.CreateApi(CLI.Dll.Dir)
		// 
		// for k, _ := range api.Endpoints {
		// 	fmt.Println("Registered: " + k)
		//
		//
		// serve := webserver.NewWebserver(":5432", api)
		// 
		// serve.Serve()
		//
		// api.CloseApi()
	case "test":
		// serv := webserver.NewWebserver(":5432")
		// serv.Serve()

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

		conf.Nox.Root = filepath.Join(dir, conf.Nox.Root)
		
		fmt.Println(conf.Nox.Root);

		serve := webserver.NewWebserver(&conf, nil)
		serve.Serve()
	}
}

	// httputil.NewSingleHostReverseProxy()
	// 
	// handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 //        switch r.Host {
 //        case "y.example.com":
 //            proxyY.ServeHTTP(w, r)
 //        case "z.example.com":
 //            proxyZ.ServeHTTP(w, r)
 //        default:
 //            http.Error(w, "Not Authorized", http.StatusForbidden)
 //        }
 //    })
