package gli

import (
	"reflect"
)


type ValidArg struct {
	ArgType   reflect.Type
	Value     interface{}
}
