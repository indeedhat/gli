package util

import "reflect"

func GetKind(field reflect.Value) reflect.Kind {
    if IsSlice(field) {
        return field.Type().Elem().Kind()
    }

    return field.Type().Kind()
}

func CheckKind(field reflect.Value, kind reflect.Kind) bool {
    return GetKind(field) == kind
}