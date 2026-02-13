package webserver

import (
	"net/http"
	"os"
	"path/filepath"
)

type NoxHandler struct {
	Root string
	//eventually map the endpoints to some ABI functions
	DirView interface{}
}

func (h *NoxHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqPath := filepath.Join(h.Root, req.URL.Path)
	sanitized := filepath.Clean(reqPath) //force it to be within the root!
	reqInfo, statErr := os.Stat(sanitized)
	if statErr != nil {
		if !os.IsNotExist(statErr) {
			//return 500 internal error
		}
		
		h.handleLogicReq(w, req)
		return
	}

	if reqInfo.IsDir() {
		h.handleDirReq(w, req, sanitized)
		return
	}
	
	http.ServeFile(w, req, sanitized)
}

func (h *NoxHandler) handleLogicReq(w http.ResponseWriter, req *http.Request) {
	//check for existing API endpoints, or anything else, call them and return output
}

func (h *NoxHandler) handleDirReq(w http.ResponseWriter, req *http.Request, path string) {
	indexPath := filepath.Join(path, "index.html")
	_, indexErr := os.Stat(indexPath)
	if indexErr == nil {
		http.ServeFile(w, req, indexPath)
		return
	} else if !os.IsNotExist(indexErr) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if h.DirView != nil {
		//generate dir view
		return
	}

	http.ServeFile(w, req, path)
}
