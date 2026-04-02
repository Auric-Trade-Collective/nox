package webserver

import (
	"YendisFish/nox/logger"
	"YendisFish/nox/pages"
	"YendisFish/nox/webapi"
	"net/http"
	"os"
	"path/filepath"
)

// Looks silly but trust, it's much better
var SupportedHttpMethods = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"DELETE": true,
}

type NoxHandler struct {
	Root string
	Api  *webapi.NoxApi
	//eventually map the endpoints to some ABI functions
	DirView interface{}
}

func (h *NoxHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if h.Root == "" {
		if h.Api != nil {
			h.handleLogicReq(w, req)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(pages.Pg404))
		}

		return
	}
	
	reqPath := filepath.Join(h.Root, req.URL.Path)
	sanitized := filepath.Clean(reqPath) //force it to be within the root!
	reqInfo, statErr := os.Stat(sanitized)
	if statErr != nil {
		if !os.IsNotExist(statErr) {
			w.Write([]byte(pages.Pg500))
			return
		}

		if h.Root != "" {
			h.handleLogicReq(w, req)
		}
		return
	}

	if reqInfo.IsDir() {
		h.handleDirReq(w, req, sanitized)
		return
	}

	http.ServeFile(w, req, sanitized)
}

func (h *NoxHandler) handleLogicReq(w http.ResponseWriter, req *http.Request) {
	if h.Api == nil {
		return
	}

	if !SupportedHttpMethods[req.Method] {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//check for existing API endpoints, or anything else, call them and return output
	if _, ok := h.Api.Endpoints[req.URL.Path]; ok {
		h.Api.ExecuteEndpoint(req.URL.Path, w, req)
		logger.Write("[" + req.Method + "] " + req.URL.Path + " called by " + req.RemoteAddr)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(pages.Pg404))
	}
}

func (h *NoxHandler) handleDirReq(w http.ResponseWriter, req *http.Request, path string) {
	indexPath := filepath.Join(path, "index.html")
	_, indexErr := os.Stat(indexPath)
	if indexErr == nil {
		http.ServeFile(w, req, indexPath)
		return
	} else if !os.IsNotExist(indexErr) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(pages.Pg404))
		return
	}

	if h.DirView != nil {
		//generate dir view
		return
	}

	http.ServeFile(w, req, path)
}
