package main

import "fmt"

type Application struct {
    SubCommand SubCommand `gli:"sub,s,subcommand" description:"does some sub command things"`
    Flag       bool       `gli:"flag,f"`
    String     string     `gli:"string,s,S"`
    Help       bool       `gli:"help,h,H"`
}


func (app *Application) Run() int {
    fmt.Println("Main application command")
    fmt.Printf("Flag: %v\n", app.Flag)
    fmt.Printf("String: %v\n", app.String)

    return 0
}