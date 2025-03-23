package cmd

import (
	"n-puzzle-solver/internal/cli"
)

func Execute() error {
	app := cli.NewApp()
	return app.Execute()
}
