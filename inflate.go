package gli

import (
    "github.com/indeedhat/util"
    "errors"
    "reflect"
    "strconv"
    "fmt"
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
            err = inf.inflateBool(val)
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            err = inf.inflateInt(val)
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
            err = inf.inflateUint(val)
        case reflect.Float32, reflect.Float64:
            err = inf.inflateFloat(val)
        case reflect.String:
            err = inf.inflateString(val)
        default:
            // do nothing
        }

        if nil != err { return }
    }

    return
}


func (inf *Inflater) inflateString(arg *ValidArg) error {
    if util.IsSlice(arg.ArgType) {
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        arg.ArgType.SetString(arg.Value[len(arg.Value) - 1])
    }

    return nil
}


func (inf *Inflater) inflateInt(arg *ValidArg) error {
    var err error
    var i int64

    if util.IsSlice(arg.ArgType) {
        var newvals []int64
        for _, v := range arg.Value {
            if i, err = strconv.ParseInt(v, 10, 64); nil == err {
                newvals = append(newvals, i)
            } else { break }
        }
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        i, err = strconv.ParseInt(arg.Value[len(arg.Value) - 1], 10, 64)
        if err == nil {
            arg.ArgType.SetInt(i)
        }
    }

    return generateError(err, arg, "int")
}


func (inf *Inflater) inflateUint(arg *ValidArg) error {
    var i uint64
    var err error

    if util.IsSlice(arg.ArgType) {
        var newvals []uint64
        for _, v := range arg.Value {
            if i, err = strconv.ParseUint(v, 10, 64); nil == err {
                newvals = append(newvals, i)
            } else { break }
        }
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        i, err = strconv.ParseUint(arg.Value[len(arg.Value) - 1], 10, 64)
        if err == nil {
            arg.ArgType.SetUint(i)
        }
    }

    return generateError(err, arg, "uint")
}


func (inf *Inflater) inflateFloat(arg *ValidArg) error {
    var i float64
    var err error

    if util.IsSlice(arg.ArgType) {
        var newvals []float64
        for _, v := range arg.Value {
            if i, err = strconv.ParseFloat(v, 64); nil == err {
                newvals = append(newvals, i)
            } else { break }
        }
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        i, err = strconv.ParseFloat(arg.Value[len(arg.Value) - 1], 64)
        if err == nil {
            arg.ArgType.SetFloat(i)
        }
    }

    return generateError(err, arg, "float")
}


func (inf *Inflater) inflateBool(arg *ValidArg) error {
    var i bool
    var err error

    if util.IsSlice(arg.ArgType) {
        var newvals []bool
        for _, v = range arg.Value {
            if i, err := strconv.ParseBool(v); nil == err {
                newvals = append(newvals, i)
            }
        }
        arg.ArgType.Set(reflect.AppendSlice(arg.ArgType, reflect.ValueOf(arg.Value)))
    } else {
        i, err = strconv.ParseBool(arg.Value[len(arg.Value) - 1])
        if err == nil {
            arg.ArgType.SetBool(i)
        }
    }

    return generateError(err, arg, "boolean")
}


func generateError(err error, arg *ValidArg, typ string) error {
    if nil == err {
        return nil
    }

    if POSITIONAL == arg.Type() {
        return errors.New(fmt.Sprintf("Arg given at position %d was not of type %s", arg.Offset, typ))
    } else {
        return errors.New(fmt.Sprintf("Arg given for %s was notof type %s", arg.Key, typ))
    }
}
