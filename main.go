package main

import (
	"os"

	"github.com/juanhuttemann/gochain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
