package util


// return value based on the outcome of a condition
// this pretty much implements single line syntax for if/else statements in go
func IfElse(condition bool, is, isnt interface{}) interface{} {
    if condition {
        return is
    }

    return isnt
}


// simplify if statements requiring an xor
func Xor(first, second bool) bool {
    return first != second
}