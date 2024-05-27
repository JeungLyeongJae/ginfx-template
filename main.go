package main

import (
	"fmt"
	"ginfx-template/cmd/server"
	"os"
)

func main() {
	app := server.App()
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Model Builder: %s\n", err)
		os.Exit(1)
	}
}
