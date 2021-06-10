package main

import (
	"github.com/asppj/g/cli"
	"os"
)

func main() {
	if err := cli.SubApp().Run(os.Args); err != nil {
		os.Exit(1)
	}
}
