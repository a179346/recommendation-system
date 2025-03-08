package console

import (
	"fmt"
	"log"
	"os"
)

const reset = "\033[0m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"

const flag = log.Ldate | log.Ltime | log.Lshortfile

var infoLog = log.New(os.Stdout, green+"INFO "+reset, flag)
var warningLog = log.New(os.Stdout, yellow+"WARNING "+reset, flag)
var errorLog = log.New(os.Stderr, red+"ERROR "+reset, flag)

func Info(v ...any) {
	println(infoLog, v...)
}

func Infof(format string, v ...any) {
	printf(infoLog, format, v...)
}

func Warn(v ...any) {
	println(warningLog, v...)
}

func Warnf(format string, v ...any) {
	printf(warningLog, format, v...)
}

func Error(v ...any) {
	println(errorLog, v...)
}

func Errorf(format string, v ...any) {
	printf(errorLog, format, v...)
}

func println(logger *log.Logger, v ...any) {
	var b []byte
	_ = logger.Output(3, string(fmt.Appendln(b, v...)))
}

func printf(logger *log.Logger, format string, v ...any) {
	var b []byte
	_ = logger.Output(3, string(fmt.Appendf(b, format, v...)))
}
