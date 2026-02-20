package native

/*
#cgo LDFLAGS: -ldl
#include "webapi.h"
#include <stdint.h>
*/
import "C"
import (
	"fmt"
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
		fmt.Println("Could not write to stream!")
		return
	}

	filename := C.GoString(dat.filename)

	f, fErr := os.Open(filename)
	if fErr != nil {
		fmt.Println(fErr.Error())
	}
	defer f.Close()

	buff := make([]byte, 512)
	_, err := f.Read(buff)
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
	}

	conType := http.DetectContentType(buff)

	f.Seek(0, 0)

	wrt.Header().Set("Content-Type", conType)
	wrt.Header().Set("Content-Disposition", `attachment; filename="` + filename + `"`)
	io.Copy(wrt, f)
}

//export WriteCopy
func WriteCopy(w *C.HttpResponse, dat *C.NoxData) {
	gohandle := cgo.Handle(w.gohandle)
	wrt, ok := gohandle.Value().(http.ResponseWriter)
	if !ok {
		fmt.Println("Could not write to stream!")
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
		fmt.Println("Could not write to stream!")
		return
	}

	//IT IS VITAL THAT THIS FUNCTION CLEANS UP THE dat POINTER AFTER ITSELF!!!!!
	// right now it just copies over the buffer, should eventually do more than that
	buff := unsafe.Slice((*byte)(dat.buff), dat.length)
	wrt.Write(buff)

	C.FreeData(dat);
}

//export WriteText
func WriteText(w *C.HttpResponse, dat *C.char, length C.int) {
	gohandle := cgo.Handle(w.gohandle)
	wrt, ok := gohandle.Value().(http.ResponseWriter)
	if !ok {
		fmt.Println("Could not write to stream!")
		return
	}

	buff := C.GoBytes(unsafe.Pointer(dat), length)
	wrt.Write(buff)
}

type NoxApi struct {
	handle *C.NoxEndpointCollection
	Endpoints map[string]unsafe.Pointer
	Auth unsafe.Pointer //eventually there will be an option to load and use an auth function

	//may put more in here in order to store API information and stuff
}

func CreateApi(libpath string) (*NoxApi, error) {
	cstr := C.CString(libpath)
	endp := C.LoadApi(cstr)

	defer C.free(unsafe.Pointer(cstr))

	if endp == nil {
		panic("Lib does not exist!")
	}	

	nox := &NoxApi{
		handle: endp,
		Endpoints: make(map[string]unsafe.Pointer),
	}

	endps := getNoxEndpointSlice(endp)

	for _, ep := range endps {
		nox.Endpoints[C.GoString(ep.endpoint)] = unsafe.Pointer(ep.callback)
	}

	return nox, nil;
}

func (api *NoxApi) ExecuteEndpoint(path string, resp http.ResponseWriter, req *http.Request) {
	goHandle := cgo.NewHandle(resp)
	ptr := C.uintptr_t(goHandle)
	defer goHandle.Delete()

	cResp := &C.HttpResponse{
		gohandle: ptr,
	}
	cReq := &C.HttpRequest{}

	C.InvokeApiCallback((*[0]byte)(api.Endpoints[path]), cResp, cReq)
}

func (api *NoxApi) CloseApi() {
	C.CloseApi(api.handle);
}

func getNoxEndpointSlice(endps *C.NoxEndpointCollection) []C.NoxEndpoint {
	count := int(endps.endpointCount)
	if count <= 0 {
		fmt.Println("WARNING: No endpoints registered!")
		return nil
	}

	endpoints := (*[1 << 28]C.NoxEndpoint)(unsafe.Pointer(endps.endpoints))[:count:count]
	return endpoints
}
