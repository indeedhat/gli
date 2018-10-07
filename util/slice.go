package util

import "reflect"


// check the reflection value to see if a type is an array or slice
func IsSlice(kind reflect.Value) bool {
    return reflect.Slice == kind.Kind() ||
        reflect.Array == kind.Kind()
}