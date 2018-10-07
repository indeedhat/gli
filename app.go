package gli

import (
    "fmt"
    "os"
    "reflect"
)

type App struct {
    Structure Command // the full structure of the application
    Subject   Command // the command that is currently being run
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
    args := os.Args[1:]

    parser := NewParser(app, args)
    err    := parser.Parse()

    if nil == err {
        inflater := NewInflater(app.Subject, parser.Valid)
        err = inflater.Run()
    }

    if nil != err {
        os.Stderr.WriteString(err.Error())

        if ah, ok := app.Subject.(AutoHelper); ok && ah.AutoHelp() {
            app.ShowHelp()
        }

        os.Exit(1)
    }

    code := app.Subject.Run()
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


func (app *App) ShowHelp() {
    fmt.Println("Showing auto help")
}