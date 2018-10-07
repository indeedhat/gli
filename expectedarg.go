package gli

import (
    "./util"
    "reflect"
    "strings"
)


type ExpectedArg struct {
    FieldName   string
    Keys        []string
    ArgType     reflect.Value
    DefaultVal  string
    Required    bool
    Description string
}


// constructor for an expected argument
func newExpectedArg(field reflect.StructField, val reflect.Value) *ExpectedArg {
    // skip fields without a gli tag
    gliTag := field.Tag.Get("gli")
    if "" == gliTag { return nil }

    // parser gli tag
    required  := "!" == gliTag[:1]
    offset, _ := util.IfElse(required, 1, 0).(int)

    // parse description tag
    description := field.Tag.Get("description")

    // parse default tag for all but bools, arrays and slices (they cannot have default values)
    defaultVal, _ := util.IfElse(
        reflect.Bool == val.Kind() || util.IsSlice(val),
        "",
        field.Tag.Get("default"),
    ).(string)

    // create arg for return
    return &ExpectedArg{
        field.Name,
        strings.Split(gliTag[offset:], ","),
        val,
        defaultVal,
        required,
        description,
    }
}


// basically this is in_array for the keys slice
func (arg *ExpectedArg) hasKey(key string) bool {
    for _, v := range arg.Keys {
        if key == v {
            return true
        }
    }

    return false
}