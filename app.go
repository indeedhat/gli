package gli

import (
    "fmt"
    "os"
    "reflect"
)

type App struct {
    Structure Command // the full structure of the application
    Subject   Command // the command that is currently being run
    Parser    Parser
}


// App constructor
func NewApplication(structure Command) (app *App) {
    app = &App{}
    app.Structure = structure
    app.Subject   = structure

    return
}


// parse input and run appropriate command
func (app *App) Run() {
    code := 0 // response/exit code
    args := os.Args[1:]

    app.Parser = NewParser(app, args)
    err    := app.Parser.Parse()

    // inflate struct with args if they could be parsed
    if nil == err {
        inflater := NewInflater(app.Subject, app.Parser.Valid)
        err = inflater.Run()
    }

    // Check flags and show help if required
    if app.ShowHelp(false) {
        os.Exit(0)
    }

    // run app if command if there is no error
    if nil == err {
        code = app.Subject.Run()
    }

    if nil != err || 0 != code {
        // display error if present
        if nil != err {
            os.Stderr.WriteString(err.Error())
        }

        // set default fail code
        if 0 == code {
            code = 1
        }

        // show help if auto helper is implemented
        if ah, ok := app.Subject.(AutoHelper); ok && ah.AutoHelp() {
            app.ShowHelp(true)
        }
    }

    os.Exit(code)
}


// select a command by name
// this will only search for the command within the current active command
func (app *App) SelectCommand(fieldName string) bool {
    if t, ok := reflect.TypeOf(app.Subject).Elem().FieldByName(fieldName); ok {
        cmd, ok := reflect.New(t.Type).Interface().(Command)
        if ok {
            app.Subject = cmd
            return true
        }
    }

    return false
}


func (app *App) ShowHelp(force bool) bool {
    if !force {
        for field, arg := range app.Parser.Valid {
            if "Help" != field { continue }

            if reflect.Bool == arg.ArgType.Kind() && "true" == arg.Value[len(arg.Value) - 1] {
                force = true
            }
        }

        if !force {
            return false
        }
    }

    if cmd, ok := app.Subject.(CustomHelp); ok {
        fmt.Print(cmd.Help(app.Parser.Expected))
    } else {
        doc := NewDocumenter(app.Parser.Expected)
        fmt.Print(doc.Build())
    }

    return true
}