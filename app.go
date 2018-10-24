package gli

import (
    "fmt"
    "github.com/indeedhat/gli/util"
    "os"
    "reflect"
)

type App struct {
    Structure Command // the full structure of the application
    Subject   ActiveCommand // the command that is currently being run
    Parser    Parser
    Debug     bool // dont recover in debug mode
}


// App constructor
func NewApplication(structure Command, description string) (app *App) {
    app = &App{}
    app.Structure = structure
    app.Subject   = newActiveCommand(structure, description)

    return
}


// parse input and run appropriate command
func (app *App) Run() {
    code := 0 // response/exit code

    if !app.Debug {
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
                if _, ok := app.Subject.Cmd.(Helper); ok {
                    app.ShowHelp(true)
                }
            }

            os.Exit(code)
        }()
    }

    // parse args
    app.Parser = NewParser(app, os.Args[1:])
    util.PanicOnError(
        app.Parser.Parse(),
    )

    // inflate the command struct with parsed args
    inf := NewInflater(app.Subject.Cmd, app.Parser.Valid)
    util.PanicOnError(
        inf.Run(),
    )

    // Check flags and show help if required
    if app.ShowHelp(false) {
        os.Exit(0)
    }

    code = app.Subject.Cmd.Run()
}


// select a command by name
// this will only search for the command within the current active command
func (app *App) SelectCommand(fieldName string) bool {
    if t, ok := reflect.TypeOf(app.Subject.Cmd).Elem().FieldByName(fieldName); ok {
        cmd, ok := reflect.New(t.Type).Interface().(Command)
        if ok {
            app.Subject = newActiveCommand(cmd, t.Tag.Get("description"))
            return true
        }
    }

    return false
}


func (app *App) ShowHelp(force bool) bool {
    if !force {
        if help, ok := app.Subject.Cmd.(Helper); !ok || !help.NeedHelp() {
            return false
        }
    }

    if cmd, ok := app.Subject.Cmd.(CustomHelp); ok {
        os.Stdout.WriteString(cmd.GenerateHelp(app.Parser.Expected, app.Subject.Description))
    } else {
        os.Stdout.WriteString(
            NewDocumenter(app.Parser.Expected).Build(app.Subject.Description),
        )
    }

    return true
}
