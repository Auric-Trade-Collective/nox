package logger

/*
#cgo LDFLAGS: -ldl
*/
import "C"
import (
	"os"

	"github.com/YendisFish/temper"
)

var tty bool = (func() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}

	return (fi.Mode() & os.ModeCharDevice) != 0
})()

func Write(msg string) {
	temper.Info(msg)
}

func Error(msg string) {
	temper.Error(msg)
}

func Warn(msg string) {
	temper.Warn(msg)	
}

func Panic(msg string) {
	temper.Error(msg)
	os.Exit(-1)
}

func Color(code string, msg string) {
	
}

func rawOutput(outp string) {

}

/*

Ideally we want to use the already defined functions
IE: "Write()/Warn()/Error()/Panic()" in these, instead
of having their own behavior. This way any logging behavior
that we may later define in the above functions will also
manifest in the ABI.

*/

//export LogWrite
func LogWrite(namespace *C.char, msg *C.char) {
	goNamespace := C.GoString(namespace)
	goMsg := C.GoString(msg)

	out := "[" + goNamespace + "] " + goMsg
	Write(out)
}

//export LogWarn
func LogWarn(namespace *C.char, msg *C.char) {
	goNamespace := C.GoString(namespace)
	goMsg := C.GoString(msg)

	out := "[" + goNamespace + "] " + goMsg
	Warn(out)
}

//export LogError
func LogError(namespace *C.char, msg *C.char) {
	goNamespace := C.GoString(namespace)
	goMsg := C.GoString(msg)

	out := "[" + goNamespace + "] " + goMsg
	Error(out)
}


//export LogPanic
func LogPanic(namespace *C.char, msg *C.char) {
	goNamespace := C.GoString(namespace)
	goMsg := C.GoString(msg)

	out := "[" + goNamespace + "] " + goMsg
	Panic(out)
}
