package gli

import "reflect"


const (
    POSITIONAL = iota
    NAMED
    FLAG
)


type ValidArg struct {
    FieldName string
    ArgType   reflect.Value
    Value     []string
    Offset    int
    Key       string
}


// constructor for the valid argument struct
func newValidArg(arg *ExpectedArg, val, key string, offset int) (valid *ValidArg) {
    if "" == val {
        val = arg.DefaultVal
    }

    valid = &ValidArg{
        arg.FieldName,
        arg.ArgType,
        []string{ val },
        offset,
        key,
    }

    return
}


func (varg *ValidArg) Type() int {
    if reflect.Bool == varg.ArgType.Kind() {
        return FLAG
    }

    if 0 >= varg.Offset {
        return POSITIONAL
    }

    return NAMED
}