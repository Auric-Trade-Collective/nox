package native

/*
#cgo LDFLAGS: -ldl
#include "webapi.h"
#include <stdint.h>
*/
import "C"
import (
	"YendisFish/nox/logger"
	"YendisFish/nox/pages"
	"io"
	"net/http"
	"os"
	"runtime/cgo"
	"strings"
	"time"
	"unsafe"
)

//c exports

//export TryGetCookie
func TryGetCookie(req *C.HttpRequest, key *C.char) *C.char {
	gohandle := cgo.Handle(req.gohandle)
	r, ok := gohandle.Value().(*http.Request)
	if !ok {
		logger.Warn("Could not get request body!")
	}

	goKey := C.GoString(key)

	cookie, err := r.Cookie(goKey)
	if err != nil {
		// return 0 and set nil to value
	}

	return C.CString(cookie.Value)
}

//export TrySetCookie
func TrySetCookie(w *C.HttpResponse, key *C.char, value *C.char, path *C.char, expires C.long, secure C.bool, httponly C.bool) {
	gohandle := cgo.Handle(w.gohandle)
	wrt, ok := gohandle.Value().(http.ResponseWriter)
	if !ok {
		logger.Warn("Could not write to stream!")
	}

	goKey := C.GoString(key)
	goValue := C.GoString(value)
	goPath := C.GoString(path)
	goExpires := int64(expires)
	goSecure := bool(secure)
	goHttpOnly := bool(httponly)

	http.SetCookie(wrt, &http.Cookie{
		Name: goKey,
		Value: goValue,
		Path: goPath,
		Expires: time.Unix(goExpires, 0),
		Secure: goSecure,
		HttpOnly: goHttpOnly,
	})
}

//export TryGetResponseHeader
func TryGetResponseHeader(w *C.HttpResponse, key *C.char, num C.size_t, out **C.char) C.int {
	gohandle := cgo.Handle(w.gohandle)
	wrt, ok := gohandle.Value().(http.ResponseWriter)
	if !ok {
		logger.Warn("Could not write to stream!")
		return 0
	}

	goKey := C.GoString(key)

	if vals := wrt.Header().Values(goKey); vals != nil {
		if len(vals) > int(num) {
			cStr := C.CString(vals[int(num)])

			*out = cStr
			return 1
		}
	}

	return 0
}

//export TrySetResponseHeader
func TrySetResponseHeader(w *C.HttpResponse, key *C.char, val *C.char, add C.int) C.int {
	gohandle := cgo.Handle(w.gohandle)
	wrt, ok := gohandle.Value().(http.ResponseWriter)
	if !ok {
		logger.Warn("Could not write to stream!")
		return 0
	}

	goKey := C.GoString(key)
	goVal := C.GoString(val)


	if int(add) == 1 {
		wrt.Header().Add(goKey, goVal)
	} else {
		wrt.Header().Set(goKey, goVal)
	}

	return 1
}

//export TryGetRequestHeader
func TryGetRequestHeader(req *C.HttpRequest, key *C.char, num C.size_t, out **C.char) C.int {
	gohandle := cgo.Handle(req.gohandle)
	r, ok := gohandle.Value().(*http.Request)
	if !ok {
		logger.Warn("Could not get request body!")
		return 0
	}

	goKey := C.GoString(key)

	if strings.EqualFold(goKey, "Host") && r.Host != "" {
		cStr := C.CString(r.Host)
		*out = cStr

		return 1
	}

	if vals := r.Header.Values(goKey); vals != nil {

		if len(vals) > int(num) {
			cStr := C.CString(vals[int(num)])

			*out = cStr
			return 1
		}
	}

	return 0
}

//export TrySetRequestHeader
func TrySetRequestHeader(req *C.HttpRequest, key *C.char, val *C.char, add C.int) C.int {
	gohandle := cgo.Handle(req.gohandle)
	r, ok := gohandle.Value().(*http.Request)
	if !ok {
		logger.Warn("Could not get request body!")
		return 0
	}

	goKey := C.GoString(key)
	goVal := C.GoString(val)

	if int(add) == 1 {
		r.Header.Add(goKey, goVal)
	} else {
		r.Header.Set(goKey, goVal)
	}
	
	return 1
}

//export ReadBody
func ReadBody(req *C.HttpRequest, buffer *C.uint8_t, numBytes C.size_t) C.size_t {
	gohandle := cgo.Handle(req.gohandle)
	r, ok := gohandle.Value().(*http.Request)
	if !ok {
		logger.Warn("Could not get request body!")
		return C.size_t(0)
	}

	n, err := r.Body.Read(unsafe.Slice((*byte)(unsafe.Pointer(buffer)), int(numBytes)))
	if err != nil && err != io.EOF {
		logger.Error(err.Error())
	}

	return C.size_t(n)
}

//export GetUri
func GetUri(req *C.HttpRequest, outLength *C.size_t) *C.char {
	gohandle := cgo.Handle(req.gohandle)
	r, ok := gohandle.Value().(*http.Request)
	if !ok {
		logger.Warn("Could not get request URL!")
		return nil
	}

	ret := r.URL.RequestURI()
	*outLength = C.size_t(len(ret))
	return C.CString(ret)
}

//export TryGetUriParam
func TryGetUriParam(req *C.HttpRequest, key *C.char, index C.size_t, out **C.char, outLen *C.size_t) C.int {
	gohandle := cgo.Handle(req.gohandle)
	r, ok := gohandle.Value().(*http.Request)
	if !ok {
		logger.Warn("Could not get request URL!")
		return 0
	}

	if val, ok := r.URL.Query()[C.GoString(key)]; ok {
		if len(val) > int(index) {
			cStr := C.CString(val[int(index)])
			*out = cStr
			*outLen = C.size_t(len(val[int(index)]))
			return 1
		}
	}

	return 0
}

//export GetUriParamCount
func GetUriParamCount(req *C.HttpRequest, val *C.char) C.size_t {
	gohandle := cgo.Handle(req.gohandle)
	r, ok := gohandle.Value().(*http.Request)
	if !ok {
		logger.Warn("Could not get request URL!")
		return C.size_t(0)
	}

	query := r.URL.Query()
	if value, ok := query[C.GoString(val)]; ok {
		return C.size_t(len(value))
	}

	return C.size_t(0)
}

// //export GetHeader
// func GetHeader(req *C.HttpRequest, header *C.char) *C.char {
// 	gohandle := cgo.Handle(req.gohandle)
// 	r, ok := gohandle.Value().(*http.Request)
// 	if !ok {
// 		logger.Warn("Could not get request URL!")
// 		return nil
// 	}
//
// 	return C.CString(r.Header.Get(C.GoString(header)))
// }

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

	
	conType := http.DetectContentType(buff)
	wrt.Header().Set("Content-Type", conType)

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

	conType := http.DetectContentType(buff)
	wrt.Header().Set("Content-Type", conType)

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
	wrt.Header().Set("Content-Type", "text/plain")
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
	Auth      *C.authCallback //eventually there will be an option to load and use an auth function

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
		Auth: endp.auth,
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

func (api *NoxApi) Authenticate(req *C.HttpRequest) bool {
	if api.Auth == nil {
		return true
	}

	val := int(C.InvokeAuth(*api.Auth, req))

	if val == 1 {
		return true
	}

	return false
}

func (api *NoxApi) ExecuteEndpoint(path string, resp http.ResponseWriter, req *http.Request) {
	goHandle := cgo.NewHandle(resp)
	ptr := C.uintptr_t(goHandle)
	defer goHandle.Delete()

	goHandle2 := cgo.NewHandle(req)
	ptr2 := C.uintptr_t(goHandle2)
	defer goHandle2.Delete()

	pthStr := C.CString(path)
	defer C.free(unsafe.Pointer(pthStr))

	method := C.CString(req.Method)
	defer C.free(unsafe.Pointer(method))

	remoteAddr := C.CString(req.RemoteAddr)
	defer C.free(unsafe.Pointer(remoteAddr))

	cResp := &C.HttpResponse{
		gohandle: ptr,
	}
	cReq := &C.HttpRequest{
		gohandle:   ptr2,
		endpoint:   pthStr,
		method:     method,
		remoteAddr: remoteAddr,
	}

	if api.Authenticate(cReq) {
		C.InvokeApiCallback((*[0]byte)(api.Endpoints[path][req.Method]), cResp, cReq)
	} else {
		resp.WriteHeader(http.StatusUnauthorized)
		resp.Write([]byte(pages.Pg401))
	}
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
