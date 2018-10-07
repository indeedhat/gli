package gli

import "reflect"

type ValidArg struct {
    FieldName string
    ArgType   reflect.Value
    Value     []string
}


// constructor for the valid argument struct
func newValidArg(arg *ExpectedArg, val string) (valid *ValidArg) {
    if "" == val {
        val = arg.DefaultVal
    }

    valid = &ValidArg{
        arg.FieldName,
        arg.ArgType,
        []string{ val },
    }

    return
}