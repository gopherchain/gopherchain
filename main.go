package main

import (
	"os"

	"github.com/gopherchain/gopherchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
