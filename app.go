package gli

import (
    "fmt"
    "github.com/indeedhat/gli/util"
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

    defer func() {
        msg := recover()
        // check for error
        if nil != msg || 0 != code {
            // set appropriate exit code
            code, _ = util.IfElse(0 == code, 1, code).(int)

            // show error message
            if nil != msg {
                fmt.Fprintln(os.Stdout, msg)
            }

            // show help
            if _, ok := app.Subject.(Helper); ok {
                app.ShowHelp(true)
            }
        }

        os.Exit(code)
    }()

    // parse args
    app.Parser = NewParser(app, os.Args[1:])
    util.PanicOnError(
        app.Parser.Parse(),
    )

    // inflate the command struct with parsed args
    inf := NewInflater(app.Subject, app.Parser.Valid)
    util.PanicOnError(
        inf.Run(),
    )

    // Check flags and show help if required
    if app.ShowHelp(false) {
        os.Exit(0)
    }

    code = app.Subject.Run()
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
        if help, ok := app.Subject.(Helper); !ok || !help.NeedHelp() {
            return false
        }
    }

    if cmd, ok := app.Subject.(CustomHelp); ok {
        os.Stdout.WriteString(cmd.GenerateHelp(app.Parser.Expected))
    } else {
        os.Stdout.WriteString(
            NewDocumenter(app.Parser.Expected).Build(),
        )
    }

    return true
}