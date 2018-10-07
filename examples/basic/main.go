package main

import (
    gli "../.."
)

func main() {
    app := &Application{}

    cli := gli.NewApplication(app)
    cli.Run()
}