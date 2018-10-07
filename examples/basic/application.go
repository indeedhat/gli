package main

import "fmt"

type Application struct {
    SubCommand SubCommand `gli:"sub,s,subcommand"`
    Flag       bool       `gli:"flag,f"`
    String     string     `gli:"string,s,S"`
}


func (app *Application) Run() int {
    fmt.Println("Main application command")
    fmt.Printf("Flag: %v\n", app.Flag)
    fmt.Printf("String: %v\n", app.String)

    return 0
}