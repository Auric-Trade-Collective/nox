package js

import (
	"github.com/dop251/goja"
)

type Js struct {
	Runtime *goja.Runtime
	Config *Config
}

func CreateRuntime() *Js {
	runtime := goja.New()

	var js = &Js{ Runtime: runtime, Config: nil }
	runtime.Set("createNox", func(call goja.FunctionCall) goja.Value {
		var configMap = make(map[string]interface{})
		err := runtime.ExportTo(call.Argument(0), &configMap)
		if err != nil {
			panic(err)
		}

		conf := &Config{
			Root: configMap["root"].(string),
			Ip: configMap["ip"].(string),
			Port: configMap["port"].(string),
		}
		js.Config = conf

		return goja.Undefined()
	})

	return js
}
