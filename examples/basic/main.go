package main

import (
	"github.com/indeedhat/gli"
)

func main() {
	app := &Application{}

	cli := gli.NewApplication(app, "Example app")
	cli.Run()
	//
}
