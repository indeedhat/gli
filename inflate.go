package gli

import (
    "./util"
    "reflect"
    "strconv"
)

type Inflater struct {
    Values  map[string]*ValidArg
    Subject Command
}


// constructor
func NewInflater(subject Command, vals map[string]*ValidArg) (inflate *Inflater) {
    inflate = &Inflater{}
    inflate.Values  = vals
    inflate.Subject = subject

    return
}


func (inf *Inflater) Run() (err error) {
    for _, val := range inf.Values {
        tpe := val.ArgType.Type().Kind()
        if util.IsSlice(val.ArgType) {
            tpe = val.ArgType.Type().Elem().Kind()
        }

        switch tpe {
        case reflect.Bool:
            inf.inflateBool(val)
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            inf.inflateInt(val)
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
            inf.inflateUint(val)
        case reflect.Float32, reflect.Float64:
            inf.inflateFloat(val)
        case reflect.String:
            inf.inflateString(val)
        default:
            // do nothing
        }
    }

    return
}


func (inf *Inflater) inflateString(arg *ValidArg) {
    if util.IsSlice(arg.ArgType) {
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        arg.ArgType.SetString(arg.Value[len(arg.Value) - 1])
    }
}


func (inf *Inflater) inflateInt(arg *ValidArg) {
    if util.IsSlice(arg.ArgType) {
        var newvals []int64
        for _, v := range arg.Value {
            if i, err := strconv.ParseInt(v, 10, 64); nil == err {
                newvals = append(newvals, i)
            }
        }
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        i, err := strconv.ParseInt(arg.Value[len(arg.Value) - 1], 10, 64)
        if err == nil {
            arg.ArgType.SetInt(i)
        }
    }
}


func (inf *Inflater) inflateUint(arg *ValidArg) {
    if util.IsSlice(arg.ArgType) {
        var newvals []uint64
        for _, v := range arg.Value {
            if i, err := strconv.ParseUint(v, 10, 64); nil == err {
                newvals = append(newvals, i)
            }
        }
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        i, err := strconv.ParseUint(arg.Value[len(arg.Value) - 1], 10, 64)
        if err == nil {
            arg.ArgType.SetUint(i)
        }
    }
}


func (inf *Inflater) inflateFloat(arg *ValidArg) {
    if util.IsSlice(arg.ArgType) {
        var newvals []float64
        for _, v := range arg.Value {
            if i, err := strconv.ParseFloat(v, 64); nil == err {
                newvals = append(newvals, i)
            }
        }
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        i, err := strconv.ParseFloat(arg.Value[len(arg.Value) - 1], 64)
        if err == nil {
            arg.ArgType.SetFloat(i)
        }
    }
}


func (inf *Inflater) inflateBool(arg *ValidArg) {
    if util.IsSlice(arg.ArgType) {
        var newvals []bool
        for _, v := range arg.Value {
            if i, err := strconv.ParseBool(v); nil == err {
                newvals = append(newvals, i)
            }
        }
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        i, err := strconv.ParseBool(arg.Value[len(arg.Value) - 1])
        if err == nil {
            arg.ArgType.SetBool(i)
        }
    }
}