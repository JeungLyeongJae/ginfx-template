package server

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"ginfx-template/pkg/version"
	"github.com/urfave/cli/v2"
	"math/rand"
	"os"
	"time"

	"go.uber.org/automaxprocs/maxprocs"
)

var (
	//go:embed banner.txt
	usage string
)

func init() {
	_, _ = maxprocs.Set()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	cli.VersionPrinter = func(c *cli.Context) {
		_ = json.NewEncoder(os.Stdout).Encode(map[string]string{
			"package":  version.Package,
			"version":  version.Version,
			"revision": version.Revision,
			"runtime":  version.GoVersion,
		})
	}
}

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "My App"
	app.Version = version.Version
	app.Usage = fmt.Sprintf(usage, version.Version)
	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "http端口",
			Value:   8080,
		},
	}
	app.Action = runServer
	_, _ = fmt.Fprintln(os.Stdout, app.Usage)
	return app
}
