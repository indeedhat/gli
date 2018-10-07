package main

import "fmt"

type SubCommand struct {
    Slice    []string `gli:"data,d"`
    Required int `gli:"!r"`
}


func (cmd *SubCommand) Run() int {
    fmt.Println("Sub command")
    fmt.Printf("Slice: %v\n", cmd.Slice)
    fmt.Printf("Required: %v\n", cmd.Required)

    return 0
}