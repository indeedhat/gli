package util

import (
    "fmt"
    "github.com/go-errors/errors"
)

func PanicOnError(err error) {
    if nil != err {
        cerr := errors.New(err)
        panic(fmt.Sprintf("%s\n%s", cerr.Error(), cerr.ErrorStack()))
    }
}