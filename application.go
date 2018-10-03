package gli

import (
  "reflect"
  "strings"
  "os"
  gliError "./error"
  "errors"
)


type Application struct {
  Structure Command  // The full struct of the pp
  Subject   Command  // the struct of the command that has been triggered
}


func (app *Application) Run() {
  // catch fatal errors
  defer app.appRecover()

  // select command
  skip := app.discoverCommand()

  // parse remaining args
  remainingArgs := os.Args[skip + 2:]
  if 0 < len(remainingArgs) {
    // do the parsy things
    parser := CreateParser(remainingArgs, app.Subject)
    parser.assignArguments()
  }

  // run command
  app.Subject.Run()
}


func (app *Application) HasCommand(name string) bool {
  r := reflect.ValueOf(app.Subject).Elem()
  c := r.FieldByName(strings.Title(name))

  return c.IsValid()
}


func (app *Application) SelectCommand(name string) {
  r := reflect.ValueOf(app.Subject).Elem()
  c := r.FieldByName(strings.Title(name))

  if c.IsValid() {
    cmd, ok := c.Interface().(Command)
    if ok {
      app.Subject = cmd
    }
  }
}


func (app *Application) ShowHelp() {
  // TODO: implement help dialog
}


func (app *Application) discoverCommand() (i int) {
  args := os.Args[1:]

  for i = 0; i < len(args); i++ {
    if !app.HasCommand(args[i]) {
      break
    }

    app.SelectCommand(args[i])
  }

  if nil == app.Subject {
    gliError.Panic(errors.New("Cound not find a command to run"), gliError.COMMAND_NOT_FOUND)
  }

  return i
}


func (app *Application) appRecover() {
  r := recover()
  if nil == r {
    return
  }

  e, ok := r.(gliError.GliError)
  if !ok {
    os.Stderr.WriteString("Unexpected error found on panic")
    os.Exit(gliError.UNKNOWN_ERROR)
  }

  os.Stderr.WriteString(e.Message)
  if help, ok := app.Subject.(AutoHelper); ok && help.AutoHelp() {
    app.ShowHelp()
  }

  os.Exit(e.Code)
}