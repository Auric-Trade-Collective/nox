package native

/*
#cgo LDFLAGS: -ldl
#include "webapi.h"
*/
import "C"

type NoxApi struct {
	handle C.NoxEndpointCollection
	Endpoints map[string]interface{}

	//may put more in here in order to store API information and stuff
}


type HttpRequest struct {

}

type HttpResponse struct {

}

func CreateApi() (*NoxApi, error) {
	return nil, nil;
}
