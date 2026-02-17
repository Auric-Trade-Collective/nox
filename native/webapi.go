package native

/*
#cgo LDFLAGS: -ldl
#include "webapi.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type NoxApi struct {
	handle *C.NoxEndpointCollection
	Endpoints map[string][0]byte

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
		Endpoints: make(map[string][0]byte),
	}

	endps := getNoxEndpointSlice(endp)

	for _, ep := range endps {
		nox.Endpoints[C.GoString(ep.endpoint)] = *ep.callback
	}

	return nox, nil;
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
