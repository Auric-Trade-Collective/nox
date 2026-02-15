package native

import (
	"errors"

	"github.com/ebitengine/purego"
)

type NoxApi struct {
	handle uintptr
	Endpoints map[string]NoxApiCallback

	//may put more in here in order to store API information and stuff
}

type NoxApiCallback struct {
	handle uintptr
	GoFunc func(*HttpResponse, *HttpRequest)
}

type HttpRequest struct {

}

type HttpResponse struct {

}

func CreateApi(libpath string) (*NoxApi, error) {
	lib, libErr := purego.Dlopen(libpath, purego.RTLD_DEFAULT)
	if libErr != nil {
		return nil, errors.New("Cannot open library")
	}

	api := &NoxApi{
		handle: lib,
		Endpoints: make(map[string]NoxApiCallback),
	}
	
	cRegister := purego.NewCallback(func(name string, ptr uintptr) {
		var goCallback func(*HttpResponse, *HttpRequest)
		purego.RegisterFunc(&goCallback, ptr)

		api.Endpoints[name] = NoxApiCallback{
			handle: ptr,
			GoFunc: goCallback,
		}
	})

	var cCreateApi func(uintptr)
	purego.RegisterLibFunc(&cCreateApi, api.handle, "createNox")
	cCreateApi(cRegister);

	return api, nil
}


func (api *NoxApi) CloseApi() {
	purego.Dlclose(api.handle)

	//honestly I would go around "nil-ifying" every function after this, since the lib gets freed
}

