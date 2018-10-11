package main

import (
    "fmt"
)

type SubCommand struct {
    Help     bool     `gli:"^help,h"`
    Slice    []string `gli:"data,d"`
    Required int      `gli:"!r" description:"oooohh required!"`
    OWrite   bool     `gli:"^o"`
}

func (cmd *SubCommand) NeedHelp() bool {
    return cmd.Help
}


func (cmd *SubCommand) Run() int {
    fmt.Println("Sub command")
    fmt.Printf("Slice: %v\n", cmd.Slice)
    fmt.Printf("Required: %v\n", cmd.Required)

    return 0
}