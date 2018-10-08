package gli

import (
    "github.com/indeedhat/gli/color"
    "fmt"
    "reflect"
    "sort"
    "strings"
)


const TAB     = "    "
const TABSIZE = 4

type Documenter struct {
    Expected []*ExpectedArg
    Subject  *ExpectedArg
}


func NewDocumenter(expected []*ExpectedArg) (doc *Documenter) {
    doc = &Documenter{}
    doc.Expected = expected

    return
}


func (doc *Documenter) Build() string {
    commands, args, options := "", "", ""
    for _, doc.Subject = range doc.Expected {
        if reflect.Struct == doc.Subject.ArgType.Kind() {
            if "" == commands {
                commands = doc.buildHeader("Commands")
            }
            commands += doc.buildCommandEntry()
        } else if 0 == len(doc.Subject.Keys) {
            if "" == args {
                args += doc.buildHeader("Arguments")
            }
            args += doc.buildPositionalEntry()
        } else {
            if "" == options {
                options += doc.buildHeader("Options")
            }
            options += doc.buildOptionEntry()
        }
    }

    return fmt.Sprintf(
        "%s%s%s",
        args,
        options,
        commands,
    )
}


func (doc *Documenter) buildPositionalEntry() string {
    // get arg name from struct field
    name := strings.ToLower(doc.Subject.FieldName)

    return doc.buildArgEntry(name)
}


func (doc *Documenter) buildOptionEntry() string {
    // build list of options
    var opts []string
    for i := 0; i < len(doc.Subject.Keys); i++ {
        if 1 == len(doc.Subject.Keys[i]) {
            opts = append(opts, "-" + doc.Subject.Keys[i])
        } else {
            opts = append([]string{"--" + doc.Subject.Keys[i]}, opts...)
        }
    }

    return doc.buildArgEntry(strings.Join(opts, ", "))
}


func (doc *Documenter) buildArgEntry(argspec string) string {
    // add default value
    def := ""
    if "" != doc.Subject.DefaultVal && reflect.Bool != doc.Subject.ArgType.Kind() {
        def = fmt.Sprintf("[=%s]", doc.Subject.DefaultVal)
    }

    // add required bit
    required := ""
    if doc.Subject.Required {
        required = fmt.Sprintf("%s!Required!", TAB)
    }

    // add the description
    description := ""
    if "" != doc.Subject.Description {
        description = fmt.Sprintf("%s%s%s\n", TAB, TAB, doc.Subject.Description)
    }

    // build the output
    return fmt.Sprintf(
        "%s%s%s%s\n%s\n",
        TAB,
        color.Wrap(argspec, color.White),
        color.Wrap(def, color.DarkGray),
        color.Wrap(required, color.Red),
        color.Wrap(description, color.LightGray),
    )
}


func (doc *Documenter) buildCommandEntry() string {
    var opts = doc.Subject.Keys
    sort.Slice(opts, func (i, j int) bool {
        return len(opts[i]) > len(opts[j])
    })

    return doc.buildArgEntry(
        fmt.Sprintf("%s [%s]", opts[0], strings.Join(opts[1:], ", ")),
    )
}


// create a header block
func (doc *Documenter) buildHeader(text string) (entry string) {
    return fmt.Sprintf("\n%s:\n\n", color.Wrap(text, color.White))
}