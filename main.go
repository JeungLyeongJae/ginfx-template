package main

import (
	"fmt"
	"ginfx-template/cmd/server"
	_ "ginfx-template/docs"
	"os"
)

// @title GinFx-Template
// @version 1.0.0
// @description 演示说明文档
// @contact.name go-swagger帮助文档
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:9060
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
func main() {
	app := server.App()
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Model Builder: %s\n", err)
		os.Exit(1)
	}
}
