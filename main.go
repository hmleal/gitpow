package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli/v2"
)

func RunCommand(command string, params ...string) ([]byte, error) {
	cmd := exec.Command(command, params...)

	result, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ShowStatus() bool {
	_, err := RunCommand("git", "rev-parse", "--is-inside-work-tree")

	return err == nil
}

func GitStatus() string {
	if !ShowStatus() {
		return ""
	}

	res, _ := RunCommand("git", "branch", "--show-current", "--no-color")

	return string(res)
}

func main() {
	app := &cli.App{
		Name:        "GitPow - (Git + Powershell)",
		Description: "GitPow is a program that integrates Git and Powershell by providing git status information and tab completion.",
		Version:     "0.0.1",
		Compiled:    time.Now(),
		Commands: []*cli.Command{
			{
				Name:    "status",
				Aliases: []string{"s"},
				Usage:   "This command is used to provide git status summary.",
				Action: func(ctx *cli.Context) error {
					fmt.Printf(" [%s]", GitStatus())

					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "This command is used to provide context sensitive command line completion.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "word",
						Usage:    "The value provided before completion is requested",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "commandline",
						Usage:    "The full command line for completion",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "position",
						Usage:    "The position of the cursor within the command line",
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					word := ctx.String("word")
					cmdline := ctx.String("commandline")
					position := ctx.String("position")

					fmt.Println("NotImplementedError", word, cmdline, position)

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
