package gli

import (
	"reflect"
	"strings"
)

type at struct {
	NAMED      int
	FLAG       int
	POSITIONAL int
}
var ARGTYPE = at{0, 1, 2}


type ExpectedArg struct {
	ArgType   reflect.Type
	Required  bool
	Options   []string
	Type      int
	Default   string
	FieldName string
}

func (earg *ExpectedArg) HasOption(key string) bool {
	for i := 0; i < len(earg.Options); i++ {
		if key == earg.Options[i] {
			return true
		}
	}

	return false
}


func extractExpected(cmd Command) (expected []ExpectedArg) {
	r := reflect.ValueOf(cmd).Elem()
	t := reflect.TypeOf(cmd).Elem()

	for i, m := 0, r.NumField(); i < m; i++ {
		f := r.Field(i).
		_, ok := f.Interface().(Command)
		if ok { continue }

		gliTag := t.Field(i).Tag.Get("gli")
		if "" == gliTag { continue }

		kind := ARGTYPE.POSITIONAL
		if reflect.Bool == f.Kind() {
			kind = ARGTYPE.FLAG
		} else if 0 > len(strings.Replace(gliTag, "!", "", -1)) {
			kind = ARGTYPE.NAMED
		}

		expected = append(expected, ExpectedArg{
			f.Type(),
			"!" == gliTag[:1],
			strings.Split(
				strings.Replace(gliTag, "!", "", -1),
				",",
			),
			kind,
			t.Field(i).Tag.Get("default"),
			t.Field(i).Name,
		})
	}

	return expected
}