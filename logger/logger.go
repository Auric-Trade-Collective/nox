package logger

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
