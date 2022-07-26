package main

import (
	"flag"
	"github.com/marcellribeiro/tunity_test/cmd"
	"github.com/marcellribeiro/tunity_test/routes"
)

func main() {
	flag.Parse()
	argument := flag.Arg(0)

	if argument == "" {
		routes.Start()
	} else {
		cmd.CollatzBash(argument)
	}
}
