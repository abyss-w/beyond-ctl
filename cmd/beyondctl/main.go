package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const flagConfig = "config"

var app = cli.App{
	Name: "beyondctl",
	Flags: []cli.Flag{
		&configFlag,
	},
	Commands: []*cli.Command{
		lsCmd,
		profileCmd,
	},
}

var configFlag = cli.StringFlag{
	Name:    flagConfig,
	Usage:   "Load config from `FILE`",
	Aliases: []string{"c"},
	EnvVars: []string{
		"BEYOND_CTL_CONFIG",
	},
	Value: fmt.Sprintf("%s/beyondctl/config.toml", userConfigDir()),
}

func main() {
	logger, _ := zap.NewDevelopment()

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal("beyondctl execute", zap.Error(err))
	}
}

func userConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("$HOME is not specified")
	}
	return configDir
}