package main

import (
	// "net/http"
	// "net/http/httputil"

	"YendisFish/nox/native"
	"YendisFish/nox/webserver"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Test struct {

	} `cmd:"" help:"Test"`
	Dll struct {
		Dir string `cmd:"--dir" help:"Test load a DLL file"`
	} `cmd:"" help:"Load DLLS"`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "dll":
		api, _ := native.CreateApi()
		_ = api
	case "test":
		serv := webserver.NewWebserver(":5432")
		serv.Serve()
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
