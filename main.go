package main

import (
	// "net/http"
	// "net/http/httputil"

	"YendisFish/nox/js"
	"fmt"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Test struct {

	} `cmd:"" help:"Test"`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "test":
		jsc := js.CreateRuntime()
		_, err := jsc.Runtime.RunString("createNox({ root: \"test\", ip: \"\", port: \"100\" });")
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(jsc.Config.Root)
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
