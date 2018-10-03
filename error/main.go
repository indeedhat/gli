package error

import (
	"fmt"
	"os"
)

type GliError struct {
	Message string
	Code    int
}


func Panic(err error, code int) {
	panic(GliError{
		Message: err.Error(),
		Code: code,
	})
}


func PassToStdOut(err error) {
	fmt.Println(err.Error())
}


func PassToStdErr(err error) {
	os.Stderr.WriteString(err.Error())
}
