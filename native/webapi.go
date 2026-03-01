package native

/*
#cgo LDFLAGS: -ldl
#include "webapi.h"
#include <stdint.h>
*/
import "C"
import (
	"YendisFish/nox/logger"
	"io"
	"net/http"
	"os"
	"runtime/cgo"
	"unsafe"
)

//c exports

//export WriteFile
func WriteFile(w *C.HttpResponse, dat *C.NoxData) {
	gohandle := cgo.Handle(w.gohandle)
	wrt, ok := gohandle.Value().(http.ResponseWriter)
	if !ok {
		logger.Warn("Could not write to stream!")
		return
	}

	filename := C.GoString(dat.filename)

	f, fErr := os.Open(filename)
	if fErr != nil {
		logger.Warn(fErr.Error())
	}
	defer f.Close()

	buff := make([]byte, 512)
	_, err := f.Read(buff)
	if err != nil && err != io.EOF {
		logger.Warn(err.Error())
	}

	conType := http.DetectContentType(buff)

	f.Seek(0, 0)

	wrt.Header().Set("Content-Type", conType)
	wrt.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	io.Copy(wrt, f)
}

//export WriteCopy
func WriteCopy(w *C.HttpResponse, dat *C.NoxData) {
	gohandle := cgo.Handle(w.gohandle)
	wrt, ok := gohandle.Value().(http.ResponseWriter)
	if !ok {
		logger.Warn("Could not write to stream!")
		return
	}

	buff := unsafe.Slice((*byte)(dat.buff), dat.length)

	toWrite := make([]byte, dat.length)
	copy(toWrite, buff)

	wrt.Write(buff)
}

//export WriteMove
func WriteMove(w *C.HttpResponse, dat *C.NoxData) {
	gohandle := cgo.Handle(w.gohandle)
	wrt, ok := gohandle.Value().(http.ResponseWriter)
	if !ok {
		logger.Warn("Could not write to stream!")
		return
	}

	//IT IS VITAL THAT THIS FUNCTION CLEANS UP THE dat POINTER AFTER ITSELF!!!!!
	// right now it just copies over the buffer, should eventually do more than that
	buff := unsafe.Slice((*byte)(dat.buff), dat.length)
	wrt.Write(buff)

	C.FreeData(dat)
}

//export WriteText
func WriteText(w *C.HttpResponse, dat *C.char, length C.int) {
	gohandle := cgo.Handle(w.gohandle)
	wrt, ok := gohandle.Value().(http.ResponseWriter)
	if !ok {
		logger.Warn("Could not write to stream!")
		return
	}

	buff := C.GoBytes(unsafe.Pointer(dat), length)
	wrt.Write(buff)
}

//export CreateGet
func CreateGet(coll *C.NoxEndpointCollection, path *C.char, cb C.apiCallback) {
	C.CreateNoxEndpoint(coll, path, cb, 0)
}

//export CreatePost
func CreatePost(coll *C.NoxEndpointCollection, path *C.char, cb C.apiCallback) {
	C.CreateNoxEndpoint(coll, path, cb, 1)
}

//export CreatePut
func CreatePut(coll *C.NoxEndpointCollection, path *C.char, cb C.apiCallback) {
	C.CreateNoxEndpoint(coll, path, cb, 2)
}

//export CreateDelete
func CreateDelete(coll *C.NoxEndpointCollection, path *C.char, cb C.apiCallback) {
	C.CreateNoxEndpoint(coll, path, cb, 3)
}

type NoxApi struct {
	handle    *C.NoxEndpointCollection
	Endpoints map[string]map[string]unsafe.Pointer
	Auth      unsafe.Pointer //eventually there will be an option to load and use an auth function

	//may put more in here in order to store API information and stuff
}

func CreateApi(libpath string) (*NoxApi, error) {
	cstr := C.CString(libpath)
	endp := C.LoadApi(cstr)

	defer C.free(unsafe.Pointer(cstr))

	if endp == nil {
		logger.Panic("Lib does not exist! " + libpath)
	}

	nox := &NoxApi{
		handle:    endp,
		Endpoints: make(map[string]map[string]unsafe.Pointer),
	}

	endps := getNoxEndpointSlice(endp)

	for _, ep := range endps {
		var method string
		switch ep.method {
		case 0:
			method = http.MethodGet
		case 1:
			method = http.MethodPost
		case 2:
			method = http.MethodPut
		case 3:
			method = http.MethodDelete
		}

		path := C.GoString(ep.endpoint)
		end, ok := nox.Endpoints[path]
		if !ok {
			nox.Endpoints[path] = make(map[string]unsafe.Pointer)
			end = nox.Endpoints[path]
		}
		end[method] = unsafe.Pointer(ep.callback)
	}

	return nox, nil
}

func (api *NoxApi) ExecuteEndpoint(path string, resp http.ResponseWriter, req *http.Request) {
	goHandle := cgo.NewHandle(resp)
	ptr := C.uintptr_t(goHandle)
	defer goHandle.Delete()

	pthStr := C.CString(path)
	defer C.free(unsafe.Pointer(pthStr))

	method := C.CString(req.Method)
	defer C.free(unsafe.Pointer(method))

	cResp := &C.HttpResponse{
		gohandle: ptr,
	}
	cReq := &C.HttpRequest{
		endpoint: pthStr,
		method:   method,
	}

	C.InvokeApiCallback((*[0]byte)(api.Endpoints[path][req.Method]), cResp, cReq)
}

func (api *NoxApi) CloseApi() {
	C.CloseApi(api.handle)
}

func getNoxEndpointSlice(endps *C.NoxEndpointCollection) []C.NoxEndpoint {
	count := int(endps.endpointCount)
	if count <= 0 {
		logger.Warn("No endpoints registered!")
		return nil
	}

	endpoints := (*[1 << 28]C.NoxEndpoint)(unsafe.Pointer(endps.endpoints))[:count:count]
	return endpoints
}
