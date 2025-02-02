package main

import (
	"os"

	"github.com/s-beats/sutil/cmd"
	"github.com/s-beats/sutil/logger"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "sutil"
	app.Usage = "A collection of useful commands for working with Slack."

	app.Commands = []*cli.Command{
		{
			Name:   "aggregate-messages",
			Action: cmd.AggregateMessages,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Error(err)
	}
}
