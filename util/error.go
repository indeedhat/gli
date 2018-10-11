package util


func PanicOnError(err error) {
    if nil != err {
        panic(err.Error())
    }
}