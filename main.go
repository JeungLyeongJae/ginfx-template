package main

import (
	"fmt"
	"ginfx-template/cmd/server"
	"os"
)

// @title GinFx-Template
// @version 1.0.0
// @description 演示说明文档
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:9060
// @BasePath /
func main() {
	app := server.App()
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Model Builder: %s\n", err)
		os.Exit(1)
	}
}
