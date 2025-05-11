package museLogError

import (
	"fmt"
	"runtime"
)

var funcInfoFormat = "{%s:%d} [%s]"

func getFuncInfo(pc uintptr, file string, line int) string {
	f := runtime.FuncForPC(pc)
	if f == nil {
		return fmt.Sprintf(funcInfoFormat, file, line, "unknown")
	}
	return fmt.Sprintf(funcInfoFormat, file, line, f.Name())
}

const wrapFormat = "%s\n%w"

func wrap(err error, msg string) error {
	pc, file, line, ok := runtime.Caller(1)

	if !ok {
		return fmt.Errorf(wrapFormat, msg, err)
	}

	// {file:line} [funcName] msg
	stack := fmt.Sprintf("%s %s", getFuncInfo(pc, file, line), msg)
	return fmt.Errorf(wrapFormat, stack, err)
}

type util struct{}

func (e *util) Wrap(error error, msg ...string) error {
	var message string
	if len(msg) == 0 {
		message = ""
	} else {
		message = msg[0]
	}

	return wrap(error, message)
}

var Util = &util{}
